package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FileItem represents a file in the project
type FileItem struct {
	Type     string     `json:"type"`
	Name     string     `json:"name"`
	Content  string     `json:"content,omitempty"`
	Contents []FileItem `json:"contents,omitempty"`
}

// ProjectData represents the entire project structure with instructions
type ProjectData struct {
	Files        []FileItem `json:"files"`
	Instructions struct {
		Type    string `json:"type"`
		Name    string `json:"name"`
		Details string `json:"details"`
	} `json:"instructions"`
	Report struct {
		Type        string `json:"type"`
		Directories int    `json:"directories"`
		Files       int    `json:"files"`
	} `json:"report"`
}

// ScanDirectory scans a directory and returns a JSON representation of its structure
func ScanDirectory(rootDir string, maxDepth int, instructions string, ignoreList []string) ([]byte, error) {
	// Count totals for report
	dirCount := 1
	fileCount := 0

	// Create root directory item
	rootItem := FileItem{
		Type:     "directory",
		Name:     rootDir,
		Contents: []FileItem{},
	}

	// Walk through the directory
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip .git directory
		if strings.Contains(path, ".git") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Check if path matches any ignore pattern
		relPath, err := filepath.Rel(rootDir, path)
		if err != nil {
			return err
		}

		for _, pattern := range ignoreList {
			if strings.Contains(relPath, pattern) {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}

		if relPath == "." {
			return nil
		}

		depth := len(strings.Split(relPath, string(filepath.Separator)))
		if depth > maxDepth {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Create directory structure
		if info.IsDir() {
			dirCount++
		} else {
			fileCount++

			// Read file content
			content, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("error reading file %s: %v", path, err)
			}

			// Add file to appropriate parent directory
			addFileToDirectory(&rootItem, relPath, string(content), rootDir)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Create final data structure
	var data []any
	data = append(data, rootItem)

	// Add report
	reportItem := map[string]any{
		"type":        "report",
		"directories": dirCount,
		"files":       fileCount,
	}
	data = append(data, reportItem)

	// Add instructions
	instructionsItem := map[string]any{
		"type":    "instructions",
		"name":    "code_change_instructions",
		"details": instructions,
	}
	data = append(data, instructionsItem)

	return json.MarshalIndent(data, "", "  ")
}

// ScanFiles scans specific files and returns a JSON representation
func ScanFiles(files []string, instructions string) ([]byte, error) {
	fileCount := 0
	dirSet := make(map[string]bool)

	// Create root directory item
	rootItem := FileItem{
		Type:     "directory",
		Name:     ".",
		Contents: []FileItem{},
	}

	for _, filePath := range files {
		// Skip directories
		info, err := os.Stat(filePath)
		if err != nil {
			return nil, fmt.Errorf("error accessing file %s: %v", filePath, err)
		}
		if info.IsDir() {
			continue
		}

		// Track unique directories
		dir := filepath.Dir(filePath)
		if dir != "." {
			dirSet[dir] = true
		}

		fileCount++

		// Read file content
		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("error reading file %s: %v", filePath, err)
		}

		// Clean path for consistent handling
		cleanPath := filepath.Clean(filePath)
		if strings.HasPrefix(cleanPath, "./") {
			cleanPath = cleanPath[2:]
		}

		// Add file to the structure
		addFileToDirectory(&rootItem, cleanPath, string(content), ".")
	}

	// Create final data structure
	var data []any
	data = append(data, rootItem)

	// Add report
	reportItem := map[string]any{
		"type":        "report",
		"directories": len(dirSet) + 1,
		"files":       fileCount,
	}
	data = append(data, reportItem)

	// Add instructions
	instructionsItem := map[string]any{
		"type":    "instructions",
		"name":    "code_change_instructions",
		"details": instructions,
	}
	data = append(data, instructionsItem)

	return json.MarshalIndent(data, "", "  ")
}

// addFileToDirectory adds a file to the correct directory in the structure
func addFileToDirectory(root *FileItem, path, content, rootDir string) {
	parts := strings.Split(path, string(filepath.Separator))

	// If this is a file at the root level
	if len(parts) == 1 {
		root.Contents = append(root.Contents, FileItem{
			Type:    "file",
			Name:    parts[0],
			Content: content,
		})
		return
	}

	// Otherwise, find or create the directory path
	current := root
	for i := 0; i < len(parts)-1; i++ {
		dirName := parts[i]
		found := false

		// Look for existing directory
		for j, item := range current.Contents {
			if item.Type == "directory" && item.Name == dirName {
				current = &current.Contents[j]
				found = true
				break
			}
		}

		// Create directory if not found
		if !found {
			newDir := FileItem{
				Type:     "directory",
				Name:     dirName,
				Contents: []FileItem{},
			}
			current.Contents = append(current.Contents, newDir)
			current = &current.Contents[len(current.Contents)-1]
		}
	}

	// Add the file to the current directory
	current.Contents = append(current.Contents, FileItem{
		Type:    "file",
		Name:    parts[len(parts)-1],
		Content: content,
	})
}
