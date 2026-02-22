package strategy

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/danielmiessler/fabric/internal/tools/githelper"
)

const DefaultStrategiesGitRepoUrl = "https://github.com/danielmiessler/fabric.git"
const DefaultStrategiesGitRepoFolder = "data/strategies"

func NewStrategiesManager() (sm *StrategiesManager) {
	label := "Prompt Strategies"
	strategies, err := LoadAllFiles()
	if err != nil {
		strategies = make(map[string]Strategy) // empty map
	}
	sm = &StrategiesManager{
		Strategies: strategies,
	}
	sm.PluginBase = &plugins.PluginBase{
		Name:             i18n.T("strategies_label"),
		SetupDescription: i18n.T("strategies_setup_description") + " " + i18n.T("required_marker"),
		EnvNamePrefix:    plugins.BuildEnvVariablePrefix(label),
		ConfigureCustom:  sm.configure,
	}

	sm.DefaultGitRepoUrl = sm.AddSetupQuestionWithEnvName("Git Repo Url", true,
		i18n.T("strategies_git_repo_url_question"))
	sm.DefaultGitRepoUrl.Value = DefaultStrategiesGitRepoUrl

	sm.DefaultFolder = sm.AddSetupQuestionWithEnvName("Git Repo Strategies Folder", true,
		i18n.T("strategies_git_repo_folder_question"))
	sm.DefaultFolder.Value = DefaultStrategiesGitRepoFolder

	return
}

type StrategiesManager struct {
	*plugins.PluginBase
	Strategies map[string]Strategy

	DefaultGitRepoUrl *plugins.SetupQuestion
	DefaultFolder     *plugins.SetupQuestion
}

type Strategy struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Prompt      string `json:"prompt"`
}

func LoadAllFiles() (strategies map[string]Strategy, err error) {
	strategies = make(map[string]Strategy)
	strategyDir, err := getStrategyDir()
	if err != nil {
		return
	}
	filepath.WalkDir(strategyDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && path != strategyDir {
			return filepath.SkipDir
		}

		if filepath.Ext(path) == ".json" {
			strategyName := strings.TrimSuffix(filepath.Base(path), ".json")
			strategy, err := LoadStrategy(strategyName)
			if err != nil {
				return err
			}
			strategies[strategy.Name] = *strategy
		}
		return nil
	})
	return

}

func (sm *StrategiesManager) IsConfigured() (ret bool) {
	ret = sm.PluginBase.IsConfigured()
	if ret {
		if len(sm.Strategies) == 0 {
			ret = false
		}
	}
	return
}

func (sm *StrategiesManager) Setup() (err error) {
	if err = sm.PluginBase.Setup(); err != nil {
		return
	}
	if err = sm.PopulateDB(); err != nil {
		return
	}
	// Reload strategies after downloading so IsConfigured() reflects the new state
	sm.Strategies, _ = LoadAllFiles()
	return
}

// PopulateDB downloads strategies from the internet and populates the strategies folder
func (sm *StrategiesManager) PopulateDB() (err error) {
	strategyDir, _ := getStrategyDir()
	fmt.Printf(i18n.T("strategies_downloading"), strategyDir)
	fmt.Println()
	fmt.Println()
	if err = sm.gitCloneAndCopy(); err != nil {
		return
	}
	fmt.Printf(i18n.T("strategies_download_success"), strategyDir)
	return
}

func (sm *StrategiesManager) gitCloneAndCopy() (err error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		err = fmt.Errorf(i18n.T("strategies_home_dir_error"), err)
		return
	}
	strategyDir := filepath.Join(homeDir, ".config", "fabric", "strategies")

	// Create the directory if it doesn't exist
	if err = os.MkdirAll(strategyDir, os.ModePerm); err != nil {
		return fmt.Errorf(i18n.T("strategies_failed_create_directory"), err)
	}

	fmt.Printf(i18n.T("strategies_cloning_repository"), sm.DefaultGitRepoUrl.Value, sm.DefaultFolder.Value)

	// Use the helper to fetch files
	err = githelper.FetchFilesFromRepo(githelper.FetchOptions{
		RepoURL:         sm.DefaultGitRepoUrl.Value,
		PathPrefix:      sm.DefaultFolder.Value,
		DestDir:         strategyDir,
		SingleDirectory: true,
	})
	if err != nil {
		return fmt.Errorf(i18n.T("strategies_failed_download"), err)
	}

	// Count downloaded strategies
	entries, readErr := os.ReadDir(strategyDir)
	if readErr == nil {
		strategyCount := 0
		for _, entry := range entries {
			if !entry.IsDir() && filepath.Ext(entry.Name()) == ".json" {
				strategyCount++
			}
		}
		fmt.Printf(i18n.T("strategies_downloaded_count"), strategyCount)
	}

	return nil
}

func (sm *StrategiesManager) configure() (err error) {
	sm.Strategies, err = LoadAllFiles()
	return
}

// getStrategyDir returns the path to the strategies directory
func getStrategyDir() (ret string, err error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		err = fmt.Errorf(i18n.T("strategies_home_dir_fallback"), err)
		ret = filepath.Join(".", "data/strategies")
		return
	}
	return filepath.Join(homeDir, ".config", "fabric", "strategies"), nil
}

// LoadStrategy loads a strategy from the given name
func LoadStrategy(filename string) (*Strategy, error) {
	if filename == "" {
		return nil, nil
	}

	// Get the strategy directory path
	strategyDir, err := getStrategyDir()
	if err != nil {
		return nil, err
	}

	// First try with .json extension
	strategyPath := filepath.Join(strategyDir, filename+".json")
	if _, err := os.Stat(strategyPath); os.IsNotExist(err) {
		// Try without extension
		strategyPath = filepath.Join(strategyDir, filename)
		if _, err := os.Stat(strategyPath); os.IsNotExist(err) {
			return nil, fmt.Errorf(i18n.T("strategy_not_found"), filename)
		}
	}

	data, err := os.ReadFile(strategyPath)
	if err != nil {
		return nil, err
	}

	var strategy Strategy
	if err := json.Unmarshal(data, &strategy); err != nil {
		return nil, err
	}
	strategy.Name = strings.TrimSuffix(filepath.Base(strategyPath), ".json")

	return &strategy, nil
}

// ListStrategies prints available strategies
func (sm *StrategiesManager) ListStrategies(shellCompleteList bool) error {
	if len(sm.Strategies) == 0 {
		return errors.New(i18n.T("strategies_none_found"))
	}
	if !shellCompleteList {
		fmt.Print(i18n.T("strategies_available_header"), "\n\n")
	}
	// Get all strategy names for sorting
	names := []string{}
	for name := range sm.Strategies {
		names = append(names, name)
	}

	// Sort the strategy names alphabetically
	sort.Strings(names)

	// Find the longest name to align descriptions
	maxNameLength := 0
	for _, name := range names {
		if len(name) > maxNameLength {
			maxNameLength = len(name)
		}
	}

	// Print each strategy with its description aligned
	formatString := "%-" + fmt.Sprintf("%d", maxNameLength+2) + "s %s\n"
	for _, name := range names {
		strategy := sm.Strategies[name]
		if shellCompleteList {
			fmt.Printf("%s\n", strategy.Name)
		} else {
			fmt.Printf(formatString, strategy.Name, strategy.Description)
		}
	}

	return nil
}
