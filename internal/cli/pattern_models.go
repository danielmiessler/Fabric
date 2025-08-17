package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// getPatternModelFile returns the path to the pattern models mapping file
func getPatternModelFile() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not determine user home directory: %w", err)
	}
	return filepath.Join(home, ".config", "fabric", "pattern_models.yaml"), nil
}

// loadPatternModelMapping loads the pattern->model mapping from disk. It returns
// an empty map if the file does not exist.
func loadPatternModelMapping() (map[string]string, error) {
	path, err := getPatternModelFile()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return map[string]string{}, nil
		}
		return nil, err
	}
	mapping := make(map[string]string)
	if err := yaml.Unmarshal(data, &mapping); err != nil {
		return nil, err
	}
	return mapping, nil
}

// setPatternModel updates the mapping file with the provided pattern and model.
func setPatternModel(pattern, model string) error {
	path, err := getPatternModelFile()
	if err != nil {
		return err
	}
	mapping, err := loadPatternModelMapping()
	if err != nil {
		return err
	}
	if mapping == nil {
		mapping = make(map[string]string)
	}
	mapping[pattern] = model
	data, err := yaml.Marshal(mapping)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// unsetPatternModel removes a pattern from the mapping file.
func unsetPatternModel(pattern string) error {
	path, err := getPatternModelFile()
	if err != nil {
		return err
	}
	mapping, err := loadPatternModelMapping()
	if err != nil {
		return err
	}
	delete(mapping, pattern)
	if len(mapping) == 0 {
		// Remove the mapping file if empty
		if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
			return err
		}
		return nil
	}
	data, err := yaml.Marshal(mapping)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// listPatternModels prints all pattern->model mappings to stdout.
func listPatternModels() error {
	mapping, err := loadPatternModelMapping()
	if err != nil {
		return err
	}
	if len(mapping) == 0 {
		fmt.Println("no pattern model mappings found")
		return nil
	}
	for pattern, model := range mapping {
		fmt.Printf("%s: %s\n", pattern, model)
	}
	return nil
}
