package core

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
	"github.com/danielmiessler/fabric/internal/plugins/ai/anthropic"
	"github.com/danielmiessler/fabric/internal/plugins/ai/azure"
	"github.com/danielmiessler/fabric/internal/plugins/ai/bedrock"
	"github.com/danielmiessler/fabric/internal/plugins/ai/dryrun"
	"github.com/danielmiessler/fabric/internal/plugins/ai/exolab"
	"github.com/danielmiessler/fabric/internal/plugins/ai/gemini"
	"github.com/danielmiessler/fabric/internal/plugins/ai/lmstudio"
	"github.com/danielmiessler/fabric/internal/plugins/ai/ollama"
	"github.com/danielmiessler/fabric/internal/plugins/ai/openai"
	"github.com/danielmiessler/fabric/internal/plugins/ai/openai_compatible"
	"github.com/danielmiessler/fabric/internal/plugins/ai/perplexity"
	"github.com/danielmiessler/fabric/internal/plugins/strategy"

	"github.com/samber/lo"

	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/danielmiessler/fabric/internal/plugins/ai"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/danielmiessler/fabric/internal/plugins/template"
	"github.com/danielmiessler/fabric/internal/tools"
	"github.com/danielmiessler/fabric/internal/tools/custom_patterns"
	"github.com/danielmiessler/fabric/internal/tools/jina"
	"github.com/danielmiessler/fabric/internal/tools/lang"
	"github.com/danielmiessler/fabric/internal/tools/youtube"
	"github.com/danielmiessler/fabric/internal/util"
)

// hasAWSCredentials checks if Bedrock is properly configured by ensuring both
// AWS credentials and BEDROCK_AWS_REGION are present. This prevents the Bedrock
// client from being initialized when AWS credentials exist for other purposes.
func hasAWSCredentials() bool {
	// First check if BEDROCK_AWS_REGION is set - this is required for Bedrock
	if os.Getenv("BEDROCK_AWS_REGION") == "" {
		return false
	}

	// Then check if AWS credentials are available
	if os.Getenv("AWS_PROFILE") != "" ||
		os.Getenv("AWS_ROLE_SESSION_NAME") != "" ||
		(os.Getenv("AWS_ACCESS_KEY_ID") != "" && os.Getenv("AWS_SECRET_ACCESS_KEY") != "") {

		return true
	}

	credFile := os.Getenv("AWS_SHARED_CREDENTIALS_FILE")
	if credFile == "" {
		if home, err := os.UserHomeDir(); err == nil {
			credFile = filepath.Join(home, ".aws", "credentials")
		}
	}
	if credFile != "" {
		if _, err := os.Stat(credFile); err == nil {
			return true
		}
	}
	return false
}

func NewPluginRegistry(db *fsdb.Db) (ret *PluginRegistry, err error) {
	ret = &PluginRegistry{
		Db:             db,
		VendorManager:  ai.NewVendorsManager(),
		VendorsAll:     ai.NewVendorsManager(),
		PatternsLoader: tools.NewPatternsLoader(db.Patterns),
		CustomPatterns: custom_patterns.NewCustomPatterns(),
		YouTube:        youtube.NewYouTube(),
		Language:       lang.NewLanguage(),
		Jina:           jina.NewClient(),
		Strategies:     strategy.NewStrategiesManager(),
	}

	var homedir string
	if homedir, err = os.UserHomeDir(); err != nil {
		return
	}
	ret.TemplateExtensions = template.NewExtensionManager(filepath.Join(homedir, ".config/fabric"))

	ret.Defaults = tools.NeeDefaults(ret.GetModels)

	// Create a vendors slice to hold all vendors (order doesn't matter initially)
	vendors := []ai.Vendor{}

	// Add non-OpenAI compatible clients
	vendors = append(vendors,
		openai.NewClient(),
		ollama.NewClient(),
		azure.NewClient(),
		gemini.NewClient(),
		anthropic.NewClient(),
		lmstudio.NewClient(),
		exolab.NewClient(),
		perplexity.NewClient(), // Added Perplexity client
	)

	if hasAWSCredentials() {
		vendors = append(vendors, bedrock.NewClient())
	}

	// Add all OpenAI-compatible providers
	for providerName := range openai_compatible.ProviderMap {
		provider, _ := openai_compatible.GetProviderByName(providerName)
		vendors = append(vendors, openai_compatible.NewClient(provider))
	}

	// Sort vendors by name for consistent ordering (case-insensitive)
	sort.Slice(vendors, func(i, j int) bool {
		return strings.ToLower(vendors[i].GetName()) < strings.ToLower(vendors[j].GetName())
	})

	// Add all sorted vendors to VendorsAll
	ret.VendorsAll.AddVendors(vendors...)
	_ = ret.Configure()

	return
}

func (o *PluginRegistry) ListVendors(out io.Writer) error {
	vendors := lo.Map(o.VendorsAll.Vendors, func(vendor ai.Vendor, _ int) string {
		return vendor.GetName()
	})
	fmt.Fprintf(out, "%s\n\n", i18n.T("available_vendors_header"))
	for _, vendor := range vendors {
		fmt.Fprintf(out, "%s\n", vendor)
	}
	return nil
}

type PluginRegistry struct {
	Db *fsdb.Db

	VendorManager      *ai.VendorsManager
	VendorsAll         *ai.VendorsManager
	Defaults           *tools.Defaults
	PatternsLoader     *tools.PatternsLoader
	CustomPatterns     *custom_patterns.CustomPatterns
	YouTube            *youtube.YouTube
	Language           *lang.Language
	Jina               *jina.Client
	TemplateExtensions *template.ExtensionManager
	Strategies         *strategy.StrategiesManager
}

func (o *PluginRegistry) SaveEnvFile() (err error) {
	// Now create the .env with all configured VendorsController info
	var envFileContent bytes.Buffer

	o.Defaults.Settings.FillEnvFileContent(&envFileContent)
	o.PatternsLoader.SetupFillEnvFileContent(&envFileContent)
	o.CustomPatterns.SetupFillEnvFileContent(&envFileContent)
	o.Strategies.SetupFillEnvFileContent(&envFileContent)

	for _, vendor := range o.VendorManager.Vendors {
		vendor.SetupFillEnvFileContent(&envFileContent)
	}

	o.YouTube.SetupFillEnvFileContent(&envFileContent)
	o.Jina.SetupFillEnvFileContent(&envFileContent)
	o.Language.SetupFillEnvFileContent(&envFileContent)

	err = o.Db.SaveEnv(envFileContent.String())
	return
}

func (o *PluginRegistry) Setup() (err error) {
	setupQuestion := plugins.NewSetupQuestion("Enter the number of the plugin to setup")
	groupsPlugins := util.NewGroupsItemsSelector("Available plugins (please configure all required plugins):",
		func(plugin plugins.Plugin) string {
			var configuredLabel string
			if plugin.IsConfigured() {
				configuredLabel = " (configured)"
			} else {
				configuredLabel = ""
			}
			return fmt.Sprintf("%v%v", plugin.GetSetupDescription(), configuredLabel)
		})

	groupsPlugins.AddGroupItems("AI Vendors [at least one, required]", lo.Map(o.VendorsAll.Vendors,
		func(vendor ai.Vendor, _ int) plugins.Plugin {
			return vendor
		})...)

	groupsPlugins.AddGroupItems("Tools", o.CustomPatterns, o.Defaults, o.Jina, o.Language, o.PatternsLoader, o.Strategies, o.YouTube)

	for {
		groupsPlugins.Print(false)

		if answerErr := setupQuestion.Ask("Plugin Number"); answerErr != nil {
			break
		}

		if setupQuestion.Value == "" {
			break
		}
		number, parseErr := strconv.Atoi(setupQuestion.Value)
		setupQuestion.Value = ""

		if parseErr == nil {
			var plugin plugins.Plugin
			if _, plugin, err = groupsPlugins.GetGroupAndItemByItemNumber(number); err != nil {
				return
			}

			if pluginSetupErr := plugin.Setup(); pluginSetupErr != nil {
				println(pluginSetupErr.Error())
			} else {
				if err = o.SaveEnvFile(); err != nil {
					break
				}
			}

			if _, ok := o.VendorManager.VendorsByName[plugin.GetName()]; !ok {
				var vendor ai.Vendor
				if vendor, ok = plugin.(ai.Vendor); ok {
					o.VendorManager.AddVendors(vendor)
				}
			}
		} else {
			break
		}
	}

	err = o.SaveEnvFile()

	return
}

func (o *PluginRegistry) SetupVendor(vendorName string) (err error) {
	if err = o.VendorsAll.SetupVendor(vendorName, o.VendorManager.VendorsByName); err != nil {
		return
	}
	err = o.SaveEnvFile()
	return
}

func (o *PluginRegistry) ConfigureVendors() {
	o.VendorManager.Clear()
	for _, vendor := range o.VendorsAll.Vendors {
		if vendorErr := vendor.Configure(); vendorErr == nil && vendor.IsConfigured() {
			o.VendorManager.AddVendors(vendor)
		}
	}
}

func (o *PluginRegistry) GetModels() (ret *ai.VendorsModels, err error) {
	o.ConfigureVendors()
	ret, err = o.VendorManager.GetModels()
	return
}

// Configure buildClient VendorsController based on the environment variables
func (o *PluginRegistry) Configure() (err error) {
	o.ConfigureVendors()
	_ = o.Defaults.Configure()
	if err := o.CustomPatterns.Configure(); err != nil {
		return fmt.Errorf("error configuring CustomPatterns: %w", err)
	}
	_ = o.PatternsLoader.Configure()

	// Refresh the database custom patterns directory after custom patterns plugin is configured
	customPatternsDir := os.Getenv("CUSTOM_PATTERNS_DIRECTORY")
	if customPatternsDir != "" {
		// Expand home directory if needed
		if strings.HasPrefix(customPatternsDir, "~/") {
			if homeDir, err := os.UserHomeDir(); err == nil {
				customPatternsDir = filepath.Join(homeDir, customPatternsDir[2:])
			}
		}
		o.Db.Patterns.CustomPatternsDir = customPatternsDir
		o.PatternsLoader.Patterns.CustomPatternsDir = customPatternsDir
	}

	//YouTube and Jina are not mandatory, so ignore not configured error
	_ = o.YouTube.Configure()
	_ = o.Jina.Configure()
	_ = o.Language.Configure()
	return
}

func (o *PluginRegistry) GetChatter(model string, modelContextLength int, vendorName string, strategy string, stream bool, dryRun bool) (ret *Chatter, err error) {
	ret = &Chatter{
		db:     o.Db,
		Stream: stream,
		DryRun: dryRun,
	}

	defaultModel := o.Defaults.Model.Value
	defaultModelContextLength, err := strconv.Atoi(o.Defaults.ModelContextLength.Value)
	defaultVendor := o.Defaults.Vendor.Value
	vendorManager := o.VendorManager

	if err != nil {
		defaultModelContextLength = 0
		err = nil
	}

	ret.modelContextLength = modelContextLength
	if ret.modelContextLength == 0 {
		ret.modelContextLength = defaultModelContextLength
	}

	if dryRun {
		ret.vendor = dryrun.NewClient()
		ret.model = model
		if ret.model == "" {
			ret.model = defaultModel
		}
	} else if model == "" {
		if vendorName != "" {
			ret.vendor = vendorManager.FindByName(vendorName)
		} else {
			ret.vendor = vendorManager.FindByName(defaultVendor)
		}
		ret.model = defaultModel
	} else {
		var models *ai.VendorsModels
		if models, err = vendorManager.GetModels(); err != nil {
			return
		}
		if vendorName != "" {
			// ensure vendor exists and provides model
			ret.vendor = vendorManager.FindByName(vendorName)
			availableVendors := models.FindGroupsByItem(model)
			if ret.vendor == nil || !lo.Contains(availableVendors, vendorName) {
				err = fmt.Errorf("model %s not available for vendor %s", model, vendorName)
				return
			}
		} else {
			availableVendors := models.FindGroupsByItem(model)
			if len(availableVendors) > 1 {
				debuglog.Log("Warning: multiple vendors provide model %s: %s. Using %s. Specify --vendor to select a vendor.\n", model, strings.Join(availableVendors, ", "), availableVendors[0])
			}
			ret.vendor = vendorManager.FindByName(models.FindGroupsByItemFirst(model))
		}
		ret.model = model
	}

	if ret.vendor == nil {
		var errMsg string
		if defaultModel == "" || defaultVendor == "" {
			errMsg = "Please run, fabric --setup, and select default model and vendor."
		} else {
			errMsg = "could not find vendor."
		}
		err = fmt.Errorf(
			" Requested Model = %s\n Default Model = %s\n Default Vendor = %s.\n\n%s",
			model, defaultModel, defaultVendor, errMsg)
		return
	}
	ret.strategy = strategy
	return
}
