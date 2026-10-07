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
	// dirCount starts at 1 to include rootDir.
	dirCount := 1
	fileCount := 0

	rootItem := FileItem{
		Type:     "directory",
		Name:     rootDir,
		Contents: []FileItem{},
	}

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip every path that contains ".git". This also matches .gitignore and .github.
		if strings.Contains(path, ".git") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

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

		// Read only regular files. Skip symlinks, devices and sockets, so
		// that a symlink cannot add a file from outside the project.
		if info.IsDir() {
			dirCount++
		} else if info.Mode().IsRegular() {
			fileCount++

			content, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("error reading file %s: %v", path, err)
			}

			addFileToDirectory(&rootItem, relPath, string(content), rootDir)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	var data []any
	data = append(data, rootItem)

	reportItem := map[string]any{
		"type":        "report",
		"directories": dirCount,
		"files":       fileCount,
	}
	data = append(data, reportItem)

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

	rootItem := FileItem{
		Type:     "directory",
		Name:     ".",
		Contents: []FileItem{},
	}

	for _, filePath := range files {
		info, err := os.Stat(filePath)
		if err != nil {
			return nil, fmt.Errorf("error accessing file %s: %v", filePath, err)
		}
		if info.IsDir() {
			continue
		}

		dir := filepath.Dir(filePath)
		if dir != "." {
			dirSet[dir] = true
		}

		fileCount++

		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("error reading file %s: %v", filePath, err)
		}

		cleanPath := filepath.Clean(filePath)
		if strings.HasPrefix(cleanPath, "./") {
			cleanPath = cleanPath[2:]
		}

		addFileToDirectory(&rootItem, cleanPath, string(content), ".")
	}

	var data []any
	data = append(data, rootItem)

	reportItem := map[string]any{
		"type":        "report",
		"directories": len(dirSet) + 1,
		"files":       fileCount,
	}
	data = append(data, reportItem)

	instructionsItem := map[string]any{
		"type":    "instructions",
		"name":    "code_change_instructions",
		"details": instructions,
	}
	data = append(data, instructionsItem)

	return json.MarshalIndent(data, "", "  ")
}

// addFileToDirectory adds a file under root at path. It creates missing intermediate directories.
func addFileToDirectory(root *FileItem, path, content, rootDir string) {
	parts := strings.Split(path, string(filepath.Separator))

	if len(parts) == 1 {
		root.Contents = append(root.Contents, FileItem{
			Type:    "file",
			Name:    parts[0],
			Content: content,
		})
		return
	}

	current := root
	for i := 0; i < len(parts)-1; i++ {
		dirName := parts[i]
		found := false

		for j, item := range current.Contents {
			if item.Type == "directory" && item.Name == dirName {
				current = &current.Contents[j]
				found = true
				break
			}
		}

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

	current.Contents = append(current.Contents, FileItem{
		Type:    "file",
		Name:    parts[len(parts)-1],
		Content: content,
	})
}
