package fsdb

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/danielmiessler/fabric/internal/plugins/template"
	"github.com/danielmiessler/fabric/internal/util"
)

const inputSentinel = "__FABRIC_INPUT_SENTINEL_TOKEN__"

type PatternsEntity struct {
	*StorageEntity
	SystemPatternFile      string
	UniquePatternsFilePath string
	CustomPatternsDir      string
}

// Pattern represents a single pattern with its metadata
type Pattern struct {
	Name        string
	Description string
	Pattern     string
}

// GetApplyVariables main entry point for getting patterns from any source
func (o *PatternsEntity) GetApplyVariables(
	source string, variables map[string]string, input string) (pattern *Pattern, err error) {

	if pattern, err = o.loadPattern(source); err != nil {
		return
	}

	err = o.applyVariables(pattern, variables, input)
	return
}

// GetWithoutVariables returns a pattern with only the {{input}} placeholder processed
// and skips template variable replacement
func (o *PatternsEntity) GetWithoutVariables(source, input string) (pattern *Pattern, err error) {

	if pattern, err = o.loadPattern(source); err != nil {
		return
	}

	o.applyInput(pattern, input)
	return
}

func (o *PatternsEntity) loadPattern(source string) (pattern *Pattern, err error) {
	// Determine if this is a file path
	isFilePath := strings.HasPrefix(source, "\\") ||
		strings.HasPrefix(source, "/") ||
		strings.HasPrefix(source, "~") ||
		strings.HasPrefix(source, ".")

	if isFilePath {
		// Resolve the file path using GetAbsolutePath
		var absPath string
		if absPath, err = util.GetAbsolutePath(source); err != nil {
			return nil, fmt.Errorf("could not resolve file path: %v", err)
		}

		// Use the resolved absolute path to get the pattern
		pattern, _ = o.getFromFile(absPath)
	} else {
		// Otherwise, get the pattern from the database
		pattern, err = o.getFromDB(source)
	}

	return
}

func (o *PatternsEntity) ensureInput(pattern *Pattern) {
	if !strings.Contains(pattern.Pattern, "{{input}}") {
		if !strings.HasSuffix(pattern.Pattern, "\n") {
			pattern.Pattern += "\n"
		}
		pattern.Pattern += "{{input}}"
	}
}

func (o *PatternsEntity) applyInput(pattern *Pattern, input string) {
	o.ensureInput(pattern)
	pattern.Pattern = strings.ReplaceAll(pattern.Pattern, "{{input}}", input)
}

func (o *PatternsEntity) applyVariables(
	pattern *Pattern, variables map[string]string, input string) (err error) {

	o.ensureInput(pattern)

	// Temporarily replace {{input}} with a sentinel token to protect it
	// from recursive variable resolution
	withSentinel := strings.ReplaceAll(pattern.Pattern, "{{input}}", inputSentinel)

	// Process all other template variables in the pattern
	// At this point, our sentinel ensures {{input}} won't be affected
	var processed string
	if processed, err = template.ApplyTemplate(withSentinel, variables, ""); err != nil {
		return
	}

	// Finally, replace our sentinel with the actual user input
	// The input has already been processed for variables if InputHasVars was true
	pattern.Pattern = strings.ReplaceAll(processed, inputSentinel, input)
	return
}

// retrieves a pattern from the database by name
func (o *PatternsEntity) getFromDB(name string) (ret *Pattern, err error) {
	// First check custom patterns directory if it exists
	if o.CustomPatternsDir != "" {
		customPatternPath := filepath.Join(o.CustomPatternsDir, name, o.SystemPatternFile)
		if pattern, customErr := os.ReadFile(customPatternPath); customErr == nil {
			ret = &Pattern{
				Name:    name,
				Pattern: string(pattern),
			}
			return ret, nil
		}
	}

	// Fallback to main patterns directory
	patternPath := filepath.Join(o.Dir, name, o.SystemPatternFile)

	var pattern []byte
	if pattern, err = os.ReadFile(patternPath); err != nil {
		return
	}

	patternStr := string(pattern)
	ret = &Pattern{
		Name:    name,
		Pattern: patternStr,
	}
	return
}

func (o *PatternsEntity) PrintLatestPatterns(latestNumber int) (err error) {
	var contents []byte
	if contents, err = os.ReadFile(o.UniquePatternsFilePath); err != nil {
		err = fmt.Errorf("could not read unique patterns file. Please run --updatepatterns (%s)", err)
		return
	}
	uniquePatterns := strings.Split(string(contents), "\n")
	if latestNumber > len(uniquePatterns) {
		latestNumber = len(uniquePatterns)
	}

	for i := len(uniquePatterns) - 1; i > len(uniquePatterns)-latestNumber-1; i-- {
		fmt.Println(uniquePatterns[i])
	}
	return
}

// reads a pattern from a file path and returns it
func (o *PatternsEntity) getFromFile(pathStr string) (pattern *Pattern, err error) {
	// Handle home directory expansion
	if strings.HasPrefix(pathStr, "~/") {
		var homedir string
		if homedir, err = os.UserHomeDir(); err != nil {
			err = fmt.Errorf("could not get home directory: %v", err)
			return
		}
		pathStr = filepath.Join(homedir, pathStr[2:])
	}

	var content []byte
	if content, err = os.ReadFile(pathStr); err != nil {
		err = fmt.Errorf("could not read pattern file %s: %v", pathStr, err)
		return
	}
	pattern = &Pattern{
		Name:    pathStr,
		Pattern: string(content),
	}
	return
}

// GetNames overrides StorageEntity.GetNames to include custom patterns directory
func (o *PatternsEntity) GetNames() (ret []string, err error) {
	// Get names from main patterns directory
	mainNames, err := o.StorageEntity.GetNames()
	if err != nil {
		return nil, err
	}

	// Create a map to track unique pattern names (custom patterns override main ones)
	nameMap := make(map[string]bool)
	for _, name := range mainNames {
		nameMap[name] = true
	}

	// Get names from custom patterns directory if it exists
	if o.CustomPatternsDir != "" {
		// Create a temporary StorageEntity for the custom directory
		customStorage := &StorageEntity{
			Dir:           o.CustomPatternsDir,
			ItemIsDir:     o.StorageEntity.ItemIsDir,
			FileExtension: o.StorageEntity.FileExtension,
		}

		customNames, customErr := customStorage.GetNames()
		if customErr == nil {
			// Add custom patterns, they will override main patterns with same name
			for _, name := range customNames {
				nameMap[name] = true
			}
		}
		// Ignore errors from custom directory (it might not exist)
	}

	// Convert map keys back to slice
	ret = make([]string, 0, len(nameMap))
	for name := range nameMap {
		ret = append(ret, name)
	}

	// Sort the patterns alphabetically
	sort.Strings(ret)

	return ret, nil
}

// ListNames overrides StorageEntity.ListNames to use PatternsEntity.GetNames
func (o *PatternsEntity) ListNames(shellCompleteList bool) (err error) {
	var names []string
	if names, err = o.GetNames(); err != nil {
		return
	}

	if len(names) == 0 {
		if !shellCompleteList {
			fmt.Printf("\nNo %v\n", o.StorageEntity.Label)
		}
		return
	}

	for _, item := range names {
		fmt.Printf("%s\n", item)
	}
	return
}

// Get required for Storage interface
func (o *PatternsEntity) Get(name string) (*Pattern, error) {
	// Use GetPattern with no variables
	return o.GetApplyVariables(name, nil, "")
}
func (o *PatternsEntity) Save(name string, content []byte) (err error) {
	patternDir := filepath.Join(o.Dir, name)
	if err = os.MkdirAll(patternDir, os.ModePerm); err != nil {
		return fmt.Errorf("could not create pattern directory: %v", err)
	}
	patternPath := filepath.Join(patternDir, o.SystemPatternFile)
	if err = os.WriteFile(patternPath, content, 0644); err != nil {
		return fmt.Errorf("could not save pattern: %v", err)
	}
	return nil
}
