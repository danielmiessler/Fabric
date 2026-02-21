package fsdb

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/plugins/template"
	"github.com/danielmiessler/fabric/internal/util"
)

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
			return nil, fmt.Errorf(i18n.T("patterns_error_resolve_file_path"), err)
		}

		// Use the resolved absolute path to get the pattern
		if pattern, err = o.getFromFile(absPath); err != nil {
			return nil, fmt.Errorf(i18n.T("patterns_error_load_from_file"), absPath, err)
		}
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
	withSentinel := strings.ReplaceAll(pattern.Pattern, "{{input}}", template.InputSentinel)

	// Process all other template variables in the pattern
	// Pass the actual input so extension calls can use {{input}} within their value parameter
	var processed string
	if processed, err = template.ApplyTemplate(withSentinel, variables, input); err != nil {
		return
	}

	// Finally, replace our sentinel with the actual user input
	// The input has already been processed for variables if InputHasVars was true
	pattern.Pattern = strings.ReplaceAll(processed, template.InputSentinel, input)
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
		// Check if the patterns directory is empty to provide helpful error message
		if os.IsNotExist(err) {
			var entries []os.DirEntry
			entries, _ = os.ReadDir(o.Dir)
			if len(entries) == 0 || (len(entries) == 1 && entries[0].Name() == "loaded") {
				// Patterns directory is empty or only has 'loaded' file
				return nil, fmt.Errorf(i18n.T("pattern_not_found_no_patterns"), name)
			}
		}
		return nil, fmt.Errorf(i18n.T("pattern_not_found_list_available"), name)
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
		err = fmt.Errorf(i18n.T("patterns_error_read_unique_file"), err)
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
			err = fmt.Errorf(i18n.T("patterns_error_get_home_directory"), err)
			return
		}
		pathStr = filepath.Join(homedir, pathStr[2:])
	}

	var content []byte
	if content, err = os.ReadFile(pathStr); err != nil {
		err = fmt.Errorf(i18n.T("patterns_error_read_pattern_file"), pathStr, err)
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
		return fmt.Errorf(i18n.T("patterns_error_create_directory"), err)
	}
	patternPath := filepath.Join(patternDir, o.SystemPatternFile)
	if err = os.WriteFile(patternPath, content, 0644); err != nil {
		return fmt.Errorf(i18n.T("patterns_error_save_pattern"), err)
	}
	return nil
}
