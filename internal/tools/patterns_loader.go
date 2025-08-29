package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

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
		Name:             label,
		SetupDescription: "Patterns - Downloads patterns [required]",
		EnvNamePrefix:    plugins.BuildEnvVariablePrefix(label),
		ConfigureCustom:  ret.configure,
	}

	ret.DefaultGitRepoUrl = ret.AddSetupQuestionCustom("Git Repo Url", true,
		"Enter the default Git repository URL for the patterns")
	ret.DefaultGitRepoUrl.Value = DefaultPatternsGitRepoUrl

	ret.DefaultFolder = ret.AddSetupQuestionCustom("Git Repo Patterns Folder", true,
		"Enter the default folder in the Git repository where patterns are stored")
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
	// Use a consistent temp folder name regardless of the source path structure
	tempDir, err := os.MkdirTemp("", "fabric-patterns-")
	if err != nil {
		return fmt.Errorf("failed to create temporary patterns folder: %w", err)
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
	fmt.Printf("Downloading patterns and Populating %s...\n", o.Patterns.Dir)
	fmt.Println()

	originalPath := o.DefaultFolder.Value
	if err = o.gitCloneAndCopy(); err != nil {
		return fmt.Errorf("failed to download patterns from git repository: %w", err)
	}

	// If the path was migrated during gitCloneAndCopy, we need to save the updated configuration
	if o.DefaultFolder.Value != originalPath {
		fmt.Printf("💾 Saving updated configuration (path changed from '%s' to '%s')...\n", originalPath, o.DefaultFolder.Value)
		// The configuration will be saved by the calling code after this returns successfully
	}

	if err = o.movePatterns(); err != nil {
		return fmt.Errorf("failed to move patterns to config directory: %w", err)
	}

	fmt.Printf("✅ Successfully downloaded and installed patterns to %s\n", o.Patterns.Dir)

	// Create the unique patterns file after patterns are successfully moved
	if err = o.createUniquePatternsFile(); err != nil {
		return fmt.Errorf("failed to create unique patterns file: %w", err)
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
		return fmt.Errorf("failed to access patterns directory '%s': %w", o.Patterns.Dir, err)
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
				fmt.Printf("Warning: failed to preserve custom pattern '%s': %v\n", currentPattern.Name(), copyErr)
			} else {
				fmt.Printf("Preserved custom pattern: %s\n", currentPattern.Name())
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
		err = fmt.Errorf("no patterns were successfully copied to %s", o.Patterns.Dir)
		return
	}

	//create an empty file to indicate that the patterns have been updated if not exists
	if _, err = os.Create(o.loadedFilePath); err != nil {
		return fmt.Errorf("failed to create loaded marker file '%s': %w", o.loadedFilePath, err)
	}

	err = os.RemoveAll(patternsDir)
	return
}

func (o *PatternsLoader) gitCloneAndCopy() (err error) {
	// Create temp folder if it doesn't exist
	if err = os.MkdirAll(filepath.Dir(o.tempPatternsFolder), os.ModePerm); err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}

	fmt.Printf("Cloning repository %s (path: %s)...\n", o.DefaultGitRepoUrl.Value, o.DefaultFolder.Value)

	// Try to fetch files with the current path
	err = githelper.FetchFilesFromRepo(githelper.FetchOptions{
		RepoURL:    o.DefaultGitRepoUrl.Value,
		PathPrefix: o.DefaultFolder.Value,
		DestDir:    o.tempPatternsFolder,
	})
	if err != nil {
		return fmt.Errorf("failed to download patterns from %s: %w", o.DefaultGitRepoUrl.Value, err)
	}

	// Check if patterns were downloaded
	if patternCount, checkErr := o.countPatternsInDirectory(o.tempPatternsFolder); checkErr != nil {
		return fmt.Errorf("failed to read temp patterns directory: %w", checkErr)
	} else if patternCount == 0 {
		// No patterns found with current path, try automatic migration
		if migrationErr := o.tryPathMigration(); migrationErr != nil {
			return fmt.Errorf("no patterns found in repository at path %s and migration failed: %w", o.DefaultFolder.Value, migrationErr)
		}
		// Migration successful, try downloading again
		return o.gitCloneAndCopy()
	} else {
		fmt.Printf("Downloaded %d patterns to temporary directory\n", patternCount)
	}

	return nil
}

// tryPathMigration attempts to migrate from old pattern paths to new restructured paths
func (o *PatternsLoader) tryPathMigration() (err error) {
	// Check if current path is the old "patterns" path
	if o.DefaultFolder.Value == "patterns" {
		fmt.Println("🔄 Detected old pattern path 'patterns', trying migration to 'data/patterns'...")

		// Try the new restructured path
		newPath := "data/patterns"
		testTempFolder := filepath.Join(os.TempDir(), "fabric-patterns-test")

		// Clean up any existing test temp folder
		if err := os.RemoveAll(testTempFolder); err != nil {
			fmt.Printf("Warning: failed to remove test temporary folder '%s': %v\n", testTempFolder, err)
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
				fmt.Printf("✅ Found %d patterns at new path '%s', updating configuration...\n", patternCount, newPath)

				// Update the configuration
				o.DefaultFolder.Value = newPath
				// Clean up the main temp folder and replace it with the test one
				os.RemoveAll(o.tempPatternsFolder)
				if renameErr := os.Rename(testTempFolder, o.tempPatternsFolder); renameErr != nil {
					// If rename fails, try copy
					if copyErr := copy.Copy(testTempFolder, o.tempPatternsFolder); copyErr != nil {
						return fmt.Errorf("failed to move test patterns to temp folder: %w", copyErr)
					}
					os.RemoveAll(testTempFolder)
				}

				return nil
			}
		}

		// Clean up test folder
		os.RemoveAll(testTempFolder)
	}

	return fmt.Errorf("unable to find patterns at current path '%s' or migrate to new structure", o.DefaultFolder.Value)
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
		return fmt.Errorf("failed to read patterns directory: %w", err)
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
			debuglog.Log("📂 Also included patterns from custom directory: %s\n", o.Patterns.CustomPatternsDir)
		} else {
			debuglog.Log("Warning: Could not read custom patterns directory %s: %v\n", o.Patterns.CustomPatternsDir, customErr)
		}
	}

	if len(patternNamesMap) == 0 {
		if o.Patterns.CustomPatternsDir != "" {
			return fmt.Errorf("no patterns found in directories %s and %s", o.Patterns.Dir, o.Patterns.CustomPatternsDir)
		}
		return fmt.Errorf("no patterns found in directory %s", o.Patterns.Dir)
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
		return fmt.Errorf("failed to write unique patterns file: %w", err)
	}

	fmt.Printf("📝 Created unique patterns file with %d patterns\n", len(patternNames))
	return nil
}
