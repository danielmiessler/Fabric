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
	fmt.Printf(i18n.T("patterns_downloading"), o.Patterns.Dir)
	fmt.Println()
	fmt.Println()

	// Create the temp folder here, not in configure(). A run that does not
	// download patterns must not leave an empty directory (issue #2190).
	var tempDir string
	if tempDir, err = os.MkdirTemp("", "fabric-patterns-"); err != nil {
		return fmt.Errorf(i18n.T("patterns_failed_create_temp_folder"), err)
	}
	o.tempPatternsFolder = tempDir
	defer os.RemoveAll(tempDir)

	originalPath := o.DefaultFolder.Value
	if err = o.gitCloneAndCopy(); err != nil {
		return fmt.Errorf(i18n.T("patterns_failed_download_from_git"), err)
	}

	// The caller saves the env file after this returns, so a migrated path persists.
	if o.DefaultFolder.Value != originalPath {
		fmt.Printf(i18n.T("patterns_saving_updated_configuration"), originalPath, o.DefaultFolder.Value)
	}

	if err = o.movePatterns(); err != nil {
		return fmt.Errorf(i18n.T("patterns_failed_move_patterns"), err)
	}

	fmt.Printf(i18n.T("patterns_download_success"), o.Patterns.Dir)

	if err = o.createUniquePatternsFile(); err != nil {
		return fmt.Errorf(i18n.T("patterns_failed_unique_file"), err)
	}

	return
}

// PersistPatterns copies custom patterns to the updated patterns directory
func (o *PatternsLoader) PersistPatterns() (err error) {
	if _, err = os.Stat(o.Patterns.Dir); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
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

	newPatternNames := make(map[string]bool)
	for _, newPattern := range newPatterns {
		if newPattern.IsDir() {
			newPatternNames[newPattern.Name()] = true
		}
	}

	for _, currentPattern := range currentPatterns {
		if currentPattern.IsDir() && !newPatternNames[currentPattern.Name()] {
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

	if err = copy.Copy(patternsDir, o.Patterns.Dir); err != nil {
		return
	}

	// Do not write the loaded marker when the copy produced no pattern directories.
	var entries []os.DirEntry
	if entries, err = os.ReadDir(o.Patterns.Dir); err != nil {
		return
	}

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

	// IsConfigured checks this marker file.
	if _, err = os.Create(o.loadedFilePath); err != nil {
		return fmt.Errorf(i18n.T("patterns_failed_loaded_marker"), o.loadedFilePath, err)
	}

	err = os.RemoveAll(patternsDir)
	return
}

func (o *PatternsLoader) gitCloneAndCopy() (err error) {
	if err = os.MkdirAll(filepath.Dir(o.tempPatternsFolder), os.ModePerm); err != nil {
		return fmt.Errorf(i18n.T("patterns_failed_create_temp_dir"), err)
	}

	fmt.Printf(i18n.T("patterns_cloning_repository"), o.DefaultGitRepoUrl.Value, o.DefaultFolder.Value)

	err = githelper.FetchFilesFromRepo(githelper.FetchOptions{
		RepoURL:    o.DefaultGitRepoUrl.Value,
		PathPrefix: o.DefaultFolder.Value,
		DestDir:    o.tempPatternsFolder,
	})
	if err != nil {
		return fmt.Errorf(i18n.T("patterns_failed_download_from_repo"), o.DefaultGitRepoUrl.Value, err)
	}

	if patternCount, checkErr := o.countPatternsInDirectory(o.tempPatternsFolder); checkErr != nil {
		return fmt.Errorf(i18n.T("patterns_failed_read_temp_directory"), checkErr)
	} else if patternCount == 0 {
		if migrationErr := o.tryPathMigration(); migrationErr != nil {
			return fmt.Errorf(i18n.T("patterns_no_patterns_migration_failed"), o.DefaultFolder.Value, migrationErr)
		}
		// Retry with the migrated path. tryPathMigration fails on a second call, so this recurses once.
		return o.gitCloneAndCopy()
	} else {
		fmt.Printf(i18n.T("patterns_downloaded_temp"), patternCount)
	}

	return nil
}

// tryPathMigration changes DefaultFolder from the old "patterns" path to "data/patterns" when that path has patterns.
func (o *PatternsLoader) tryPathMigration() (err error) {
	if o.DefaultFolder.Value == "patterns" {
		fmt.Println(i18n.T("patterns_detected_old_path"))

		newPath := "data/patterns"
		testTempFolder := filepath.Join(os.TempDir(), "fabric-patterns-test")

		if err := os.RemoveAll(testTempFolder); err != nil {
			fmt.Printf(i18n.T("patterns_warning_remove_test_folder"), testTempFolder, err)
		}

		testErr := githelper.FetchFilesFromRepo(githelper.FetchOptions{
			RepoURL:    o.DefaultGitRepoUrl.Value,
			PathPrefix: newPath,
			DestDir:    testTempFolder,
		})

		if testErr == nil {
			if patternCount, countErr := o.countPatternsInDirectory(testTempFolder); countErr == nil && patternCount > 0 {
				fmt.Printf(i18n.T("patterns_found_new_path"), patternCount, newPath)

				o.DefaultFolder.Value = newPath
				os.RemoveAll(o.tempPatternsFolder)
				if renameErr := os.Rename(testTempFolder, o.tempPatternsFolder); renameErr != nil {
					if copyErr := copy.Copy(testTempFolder, o.tempPatternsFolder); copyErr != nil {
						return fmt.Errorf(i18n.T("patterns_failed_move_test_patterns"), copyErr)
					}
					os.RemoveAll(testTempFolder)
				}

				return nil
			}
		}

		os.RemoveAll(testTempFolder)
	}

	return fmt.Errorf(i18n.T("patterns_unable_to_find_or_migrate"), o.DefaultFolder.Value)
}

// countPatternsInDirectory returns the number of subdirectories in dir.
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

// createUniquePatternsFile writes the sorted names from the main and custom directories to unique_patterns.txt.
func (o *PatternsLoader) createUniquePatternsFile() (err error) {
	entries, err := os.ReadDir(o.Patterns.Dir)
	if err != nil {
		return fmt.Errorf(i18n.T("patterns_failed_read_directory"), err)
	}

	patternNamesMap := make(map[string]bool)

	for _, entry := range entries {
		if entry.IsDir() {
			patternNamesMap[entry.Name()] = true
		}
	}

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

	var patternNames []string
	for name := range patternNamesMap {
		patternNames = append(patternNames, name)
	}

	sort.Strings(patternNames)

	content := strings.Join(patternNames, "\n") + "\n"
	if err = os.WriteFile(o.Patterns.UniquePatternsFilePath, []byte(content), 0644); err != nil {
		return fmt.Errorf(i18n.T("patterns_failed_write_unique_file"), err)
	}

	fmt.Printf(i18n.T("patterns_unique_file_created"), len(patternNames))
	return nil
}
