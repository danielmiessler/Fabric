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

	// NoSystemPlugins makes GetApplyVariables run only the text and
	// datetime plugins. The REST server sets it, because a client can save
	// a pattern. Then the pattern cannot read the server environment or
	// files, fetch a URL, or run an extension.
	NoSystemPlugins bool
}

// Pattern represents a single pattern with its metadata
type Pattern struct {
	Name        string
	Description string
	Pattern     string
	// InputUsed is true when the pattern had an {{input}} placeholder,
	// so Pattern contains the input.
	InputUsed bool `json:"-"`
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

// GetRaw returns a pattern from storage without applying variable processing.
func (o *PatternsEntity) GetRaw(name string) (*Pattern, error) {
	return o.getFromDB(name)
}

// LooksLikePatternFilePath reports whether loadPattern uses source as a
// filesystem path. HTTP handlers must reject these names to keep the
// CLI file-path feature out of the REST API.
func LooksLikePatternFilePath(source string) bool {
	return strings.HasPrefix(source, "\\") ||
		strings.HasPrefix(source, "/") ||
		strings.HasPrefix(source, "~") ||
		strings.HasPrefix(source, ".")
}

func (o *PatternsEntity) loadPattern(source string) (pattern *Pattern, err error) {
	if LooksLikePatternFilePath(source) {
		var absPath string
		if absPath, err = util.GetAbsolutePath(source); err != nil {
			return nil, fmt.Errorf(i18n.T("patterns_error_resolve_file_path"), err)
		}

		if pattern, err = o.getFromFile(absPath); err != nil {
			return nil, fmt.Errorf(i18n.T("patterns_error_load_from_file"), absPath, err)
		}
	} else {
		pattern, err = o.getFromDB(source)
	}

	return
}

func (o *PatternsEntity) applyInput(pattern *Pattern, input string) {
	pattern.InputUsed = strings.Contains(pattern.Pattern, "{{input}}")
	pattern.Pattern = strings.ReplaceAll(pattern.Pattern, "{{input}}", input)
}

func (o *PatternsEntity) applyVariables(
	pattern *Pattern, variables map[string]string, input string) (err error) {

	pattern.InputUsed = strings.Contains(pattern.Pattern, "{{input}}")

	// Replace {{input}} with a sentinel so that ApplyTemplate does not
	// expand variables in the user input.
	withSentinel := strings.ReplaceAll(pattern.Pattern, "{{input}}", template.InputSentinel)

	// Pass input so that an extension call can use {{input}} in its
	// value parameter.
	apply := template.ApplyTemplate
	if o.NoSystemPlugins {
		apply = template.ApplyTemplateNoSystemPlugins
	}
	var processed string
	if processed, err = apply(withSentinel, variables, input); err != nil {
		return
	}

	// With InputHasVars, the caller already applied variables to the
	// input.
	pattern.Pattern = strings.ReplaceAll(processed, template.InputSentinel, input)
	return
}

func (o *PatternsEntity) getFromDB(name string) (ret *Pattern, err error) {
	if ValidateStorageName(name) != nil {
		// The typed error lets an HTTP route without a pre-validation
		// guard map this rejection to 400, not 500.
		return nil, &InvalidStorageNameError{
			Name:    name,
			Message: fmt.Sprintf(i18n.T("pattern_invalid_name"), name),
		}
	}

	// A custom pattern overrides the main pattern with the same name.
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

	patternPath := filepath.Join(o.Dir, name, o.SystemPatternFile)

	var pattern []byte
	if pattern, err = os.ReadFile(patternPath); err != nil {
		if os.IsNotExist(err) {
			var entries []os.DirEntry
			entries, _ = os.ReadDir(o.Dir)
			if len(entries) == 0 || (len(entries) == 1 && entries[0].Name() == "loaded") {
				// The "loaded" marker file from the patterns loader does
				// not count as a pattern.
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

// PrintPattern prints the raw contents of the named pattern to the terminal.
// It checks the custom patterns directory first, then falls back to the main directory.
func (o *PatternsEntity) PrintPattern(name string) (err error) {
	var pattern *Pattern
	if pattern, err = o.GetRaw(name); err != nil {
		return
	}
	fmt.Print(pattern.Pattern)
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

func (o *PatternsEntity) getFromFile(pathStr string) (pattern *Pattern, err error) {
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
	mainNames, err := o.StorageEntity.GetNames()
	if err != nil {
		return nil, err
	}

	nameMap := make(map[string]bool)
	for _, name := range mainNames {
		nameMap[name] = true
	}

	if o.CustomPatternsDir != "" {
		customStorage := &StorageEntity{
			Dir:           o.CustomPatternsDir,
			ItemIsDir:     o.StorageEntity.ItemIsDir,
			FileExtension: o.StorageEntity.FileExtension,
		}

		customNames, customErr := customStorage.GetNames()
		// A missing custom directory is not an error.
		if customErr == nil {
			for _, name := range customNames {
				nameMap[name] = true
			}
		}
	}

	ret = make([]string, 0, len(nameMap))
	for name := range nameMap {
		ret = append(ret, name)
	}

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
	return o.GetApplyVariables(name, nil, "")
}
func (o *PatternsEntity) Save(name string, content []byte) (err error) {
	// Do not store a name that loadPattern uses as a file path, for
	// example ".foo" or "~bar". For such a name, GetApplyVariables reads
	// from the filesystem, not from the database.
	if LooksLikePatternFilePath(name) {
		return &InvalidStorageNameError{
			Name:    name,
			Message: fmt.Sprintf(i18n.T("pattern_invalid_name"), name),
		}
	}
	var patternDir string
	if patternDir, err = o.resolvedPath(name); err != nil {
		return
	}
	if err = os.MkdirAll(patternDir, os.ModePerm); err != nil {
		return fmt.Errorf(i18n.T("patterns_error_create_directory"), err)
	}
	patternPath := filepath.Join(patternDir, o.SystemPatternFile)
	// The pattern file can be a symlink that already exists. Do not
	// write through a symlink that goes out of the pattern directory.
	if err = symlinkContained(patternDir, patternPath, name); err != nil {
		return err
	}
	if err = os.WriteFile(patternPath, content, 0644); err != nil {
		return fmt.Errorf(i18n.T("patterns_error_save_pattern"), err)
	}
	return nil
}

// Rename applies the file-path guard from Save to the destination name.
// Without the guard, the inherited StorageEntity.Rename accepts ".foo"
// or "~foo", and loadPattern then reads these names from the
// filesystem. A path-like source stays permitted, which lets you rename
// a legacy entry to a valid name.
func (o *PatternsEntity) Rename(oldName, newName string) error {
	if LooksLikePatternFilePath(newName) {
		return &InvalidStorageNameError{
			Name:    newName,
			Message: fmt.Sprintf(i18n.T("pattern_invalid_name"), newName),
		}
	}
	return o.StorageEntity.Rename(oldName, newName)
}
