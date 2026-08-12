package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/danielmiessler/fabric/internal/tools/githelper"

	"github.com/otiai10/copy"
)

const DefaultPatternsGitRepoUrl = "https://github.com/danielmiessler/fabric.git"
const DefaultPatternsGitRepoFolder = "data/patterns"

func NewPatternsLoader(patterns *fsdb.PatternsEntity) (ret *PatternsLoader) {
	label := "Patterns Loader"
	ret = &PatternsLoader{
		Patterns:       patterns,
		loadedFilePath: patterns.BuildFilePath("loaded"),
	}

	ret.PluginBase = &plugins.PluginBase{
		Name:             i18n.T("patterns_loader_label"),
		SetupDescription: i18n.T("patterns_setup_description") + " " + i18n.T("required_marker"),
		EnvNamePrefix:    plugins.BuildEnvVariablePrefix(label),
		ConfigureCustom:  ret.configure,
	}

	ret.DefaultGitRepoUrl = ret.AddSetupQuestionWithEnvName("Git Repo Url", true,
		i18n.T("patterns_git_repo_url_question"))
	ret.DefaultGitRepoUrl.Value = DefaultPatternsGitRepoUrl

	ret.DefaultFolder = ret.AddSetupQuestionWithEnvName("Git Repo Patterns Folder", true,
		i18n.T("patterns_git_repo_folder_question"))
	ret.DefaultFolder.Value = DefaultPatternsGitRepoFolder

	return
}

type PatternsLoader struct {
	*plugins.PluginBase
	Patterns *fsdb.PatternsEntity

	DefaultGitRepoUrl *plugins.SetupQuestion
	DefaultFolder     *plugins.SetupQuestion

	loadedFilePath string

	pathPatternsPrefix string
	tempPatternsFolder string
}

func (o *PatternsLoader) configure() (err error) {
	o.pathPatternsPrefix = fmt.Sprintf("%v/", o.DefaultFolder.Value)
	return
}

// ensureTempPatternsFolder lazily creates the scratch folder used by the pattern
// download path. Only PopulateDB and the helpers it calls need it, so it must not
// be created in configure(): configure() runs on every plugin registry
// construction, which leaked one empty temp directory per fabric invocation.
func (o *PatternsLoader) ensureTempPatternsFolder() (err error) {
	if o.tempPatternsFolder != "" {
		return
	}

	// Use a consistent temp folder name regardless of the source path structure
	var tempDir string
	if tempDir, err = os.MkdirTemp("", "fabric-patterns-"); err != nil {
		return fmt.Errorf(i18n.T("patterns_failed_create_temp_folder"), err)
	}
	o.tempPatternsFolder = tempDir

	return
}

func (o *PatternsLoader) IsConfigured() (ret bool) {
	ret = o.PluginBase.IsConfigured()
	if ret {
		if _, err := os.Stat(o.loadedFilePath); os.IsNotExist(err) {
			ret = false
		}
	}
	return
}

func (o *PatternsLoader) Setup() (err error) {
	if err = o.PluginBase.Setup(); err != nil {
		return
	}

	if err = o.PopulateDB(); err != nil {
		return
	}
	return
}

// PopulateDB downloads patterns from the internet and populates the patterns folder
func (o *PatternsLoader) PopulateDB() (err error) {
	if err = o.ensureTempPatternsFolder(); err != nil {
		return
	}
	// movePatterns removes the folder on success; this also covers the error paths
	// and leaves the loader reusable if PopulateDB is called again.
	defer func() {
		os.RemoveAll(o.tempPatternsFolder)
		o.tempPatternsFolder = ""
	}()

	fmt.Printf(i18n.T("patterns_downloading"), o.Patterns.Dir)
	fmt.Println()
	fmt.Println()

	originalPath := o.DefaultFolder.Value
	if err = o.gitCloneAndCopy(); err != nil {
		return fmt.Errorf(i18n.T("patterns_failed_download_from_git"), err)
	}

	// If the path was migrated during gitCloneAndCopy, we need to save the updated configuration
	if o.DefaultFolder.Value != originalPath {
		fmt.Printf(i18n.T("patterns_saving_updated_configuration"), originalPath, o.DefaultFolder.Value)
		// The configuration will be saved by the calling code after this returns successfully
	}

	if err = o.movePatterns(); err != nil {
		return fmt.Errorf(i18n.T("patterns_failed_move_patterns"), err)
	}

	fmt.Printf(i18n.T("patterns_download_success"), o.Patterns.Dir)

	// Create the unique patterns file after patterns are successfully moved
	if err = o.createUniquePatternsFile(); err != nil {
		return fmt.Errorf(i18n.T("patterns_failed_unique_file"), err)
	}

	return
}

// PersistPatterns copies custom patterns to the updated patterns directory
func (o *PatternsLoader) PersistPatterns() (err error) {
	// Check if patterns directory exists, if not, nothing to persist
	if _, err = os.Stat(o.Patterns.Dir); err != nil {
		if os.IsNotExist(err) {
			// No existing patterns directory, nothing to persist
			return nil
		}
		// Return unexpected errors (e.g., permission issues)
		return fmt.Errorf(i18n.T("patterns_failed_access_directory"), o.Patterns.Dir, err)
	}

	var currentPatterns []os.DirEntry
	if currentPatterns, err = os.ReadDir(o.Patterns.Dir); err != nil {
		return
	}

	newPatternsFolder := o.tempPatternsFolder
	var newPatterns []os.DirEntry
	if newPatterns, err = os.ReadDir(newPatternsFolder); err != nil {
		return
	}

	// Create a map of new patterns for faster lookup
	newPatternNames := make(map[string]bool)
	for _, newPattern := range newPatterns {
		if newPattern.IsDir() {
			newPatternNames[newPattern.Name()] = true
		}
	}

	// Copy custom patterns that don't exist in the new download
	for _, currentPattern := range currentPatterns {
		if currentPattern.IsDir() && !newPatternNames[currentPattern.Name()] {
			// This is a custom pattern, preserve it
			src := filepath.Join(o.Patterns.Dir, currentPattern.Name())
			dst := filepath.Join(newPatternsFolder, currentPattern.Name())
			if copyErr := copy.Copy(src, dst); copyErr != nil {
				fmt.Printf(i18n.T("patterns_preserve_warning"), currentPattern.Name(), copyErr)
			} else {
				fmt.Printf(i18n.T("patterns_preserved_custom_pattern"), currentPattern.Name())
			}
		}
	}
	return nil
}

// movePatterns copies the new patterns into the config directory
func (o *PatternsLoader) movePatterns() (err error) {
	if err = os.MkdirAll(o.Patterns.Dir, os.ModePerm); err != nil {
		return
	}

	patternsDir := o.tempPatternsFolder
	if err = o.PersistPatterns(); err != nil {
		return
	}

	if err = copy.Copy(patternsDir, o.Patterns.Dir); err != nil { // copies the patterns to the config directory
		return
	}

	// Verify that patterns were actually copied before creating the loaded marker
	var entries []os.DirEntry
	if entries, err = os.ReadDir(o.Patterns.Dir); err != nil {
		return
	}

	// Count actual pattern directories (exclude the loaded file itself)
	patternCount := 0
	for _, entry := range entries {
		if entry.IsDir() {
			patternCount++
		}
	}

	if patternCount == 0 {
		err = fmt.Errorf(i18n.T("patterns_no_patterns_copied"), o.Patterns.Dir)
		return
	}

	//create an empty file to indicate that the patterns have been updated if not exists
	if _, err = os.Create(o.loadedFilePath); err != nil {
		return fmt.Errorf(i18n.T("patterns_failed_loaded_marker"), o.loadedFilePath, err)
	}

	err = os.RemoveAll(patternsDir)
	return
}

func (o *PatternsLoader) gitCloneAndCopy() (err error) {
	// Create temp folder if it doesn't exist
	if err = os.MkdirAll(filepath.Dir(o.tempPatternsFolder), os.ModePerm); err != nil {
		return fmt.Errorf(i18n.T("patterns_failed_create_temp_dir"), err)
	}

	fmt.Printf(i18n.T("patterns_cloning_repository"), o.DefaultGitRepoUrl.Value, o.DefaultFolder.Value)

	// Try to fetch files with the current path
	err = githelper.FetchFilesFromRepo(githelper.FetchOptions{
		RepoURL:    o.DefaultGitRepoUrl.Value,
		PathPrefix: o.DefaultFolder.Value,
		DestDir:    o.tempPatternsFolder,
	})
	if err != nil {
		return fmt.Errorf(i18n.T("patterns_failed_download_from_repo"), o.DefaultGitRepoUrl.Value, err)
	}

	// Check if patterns were downloaded
	if patternCount, checkErr := o.countPatternsInDirectory(o.tempPatternsFolder); checkErr != nil {
		return fmt.Errorf(i18n.T("patterns_failed_read_temp_directory"), checkErr)
	} else if patternCount == 0 {
		// No patterns found with current path, try automatic migration
		if migrationErr := o.tryPathMigration(); migrationErr != nil {
			return fmt.Errorf(i18n.T("patterns_no_patterns_migration_failed"), o.DefaultFolder.Value, migrationErr)
		}
		// Migration successful, try downloading again
		return o.gitCloneAndCopy()
	} else {
		fmt.Printf(i18n.T("patterns_downloaded_temp"), patternCount)
	}

	return nil
}

// tryPathMigration attempts to migrate from old pattern paths to new restructured paths
func (o *PatternsLoader) tryPathMigration() (err error) {
	// Check if current path is the old "patterns" path
	if o.DefaultFolder.Value == "patterns" {
		fmt.Println(i18n.T("patterns_detected_old_path"))

		// Try the new restructured path
		newPath := "data/patterns"
		testTempFolder := filepath.Join(os.TempDir(), "fabric-patterns-test")

		// Clean up any existing test temp folder
		if err := os.RemoveAll(testTempFolder); err != nil {
			fmt.Printf(i18n.T("patterns_warning_remove_test_folder"), testTempFolder, err)
		}

		// Test if the new path works
		testErr := githelper.FetchFilesFromRepo(githelper.FetchOptions{
			RepoURL:    o.DefaultGitRepoUrl.Value,
			PathPrefix: newPath,
			DestDir:    testTempFolder,
		})

		if testErr == nil {
			// Check if patterns exist in the new path
			if patternCount, countErr := o.countPatternsInDirectory(testTempFolder); countErr == nil && patternCount > 0 {
				fmt.Printf(i18n.T("patterns_found_new_path"), patternCount, newPath)

				// Update the configuration
				o.DefaultFolder.Value = newPath
				// Clean up the main temp folder and replace it with the test one
				os.RemoveAll(o.tempPatternsFolder)
				if renameErr := os.Rename(testTempFolder, o.tempPatternsFolder); renameErr != nil {
					// If rename fails, try copy
					if copyErr := copy.Copy(testTempFolder, o.tempPatternsFolder); copyErr != nil {
						return fmt.Errorf(i18n.T("patterns_failed_move_test_patterns"), copyErr)
					}
					os.RemoveAll(testTempFolder)
				}

				return nil
			}
		}

		// Clean up test folder
		os.RemoveAll(testTempFolder)
	}

	return fmt.Errorf(i18n.T("patterns_unable_to_find_or_migrate"), o.DefaultFolder.Value)
}

// countPatternsInDirectory counts the number of pattern directories in a given directory
func (o *PatternsLoader) countPatternsInDirectory(dir string) (int, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0, err
	}

	patternCount := 0
	for _, entry := range entries {
		if entry.IsDir() {
			patternCount++
		}
	}

	return patternCount, nil
}

// createUniquePatternsFile creates the unique_patterns.txt file with all pattern names
func (o *PatternsLoader) createUniquePatternsFile() (err error) {
	// Read patterns from the main patterns directory
	entries, err := os.ReadDir(o.Patterns.Dir)
	if err != nil {
		return fmt.Errorf(i18n.T("patterns_failed_read_directory"), err)
	}

	patternNamesMap := make(map[string]bool) // Use map to avoid duplicates

	// Add patterns from main directory
	for _, entry := range entries {
		if entry.IsDir() {
			patternNamesMap[entry.Name()] = true
		}
	}

	// Add patterns from custom patterns directory if it exists
	if o.Patterns.CustomPatternsDir != "" {
		if customEntries, customErr := os.ReadDir(o.Patterns.CustomPatternsDir); customErr == nil {
			for _, entry := range customEntries {
				if entry.IsDir() {
					patternNamesMap[entry.Name()] = true
				}
			}
			debuglog.Log(i18n.T("patterns_debug_included_custom_directory"), o.Patterns.CustomPatternsDir)
		} else {
			debuglog.Log(i18n.T("patterns_warning_custom_directory"), o.Patterns.CustomPatternsDir, customErr)
		}
	}

	if len(patternNamesMap) == 0 {
		if o.Patterns.CustomPatternsDir != "" {
			return fmt.Errorf(i18n.T("patterns_no_patterns_found_in_directories"), o.Patterns.Dir, o.Patterns.CustomPatternsDir)
		}
		return fmt.Errorf(i18n.T("patterns_no_patterns_found_in_directory"), o.Patterns.Dir)
	}

	// Convert map to sorted slice
	var patternNames []string
	for name := range patternNamesMap {
		patternNames = append(patternNames, name)
	}

	// Sort patterns alphabetically for consistent output
	sort.Strings(patternNames)

	// Join pattern names with newlines
	content := strings.Join(patternNames, "\n") + "\n"
	if err = os.WriteFile(o.Patterns.UniquePatternsFilePath, []byte(content), 0644); err != nil {
		return fmt.Errorf(i18n.T("patterns_failed_write_unique_file"), err)
	}

	fmt.Printf(i18n.T("patterns_unique_file_created"), len(patternNames))
	return nil
}
