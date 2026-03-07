package pipeline

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

type Loader struct {
	BuiltInDir string
	UserDir    string
}

func NewDefaultLoader() (*Loader, error) {
	userDir, err := defaultUserPipelineDir()
	if err != nil {
		return nil, err
	}

	return &Loader{
		BuiltInDir: defaultBuiltInPipelineDir(),
		UserDir:    userDir,
	}, nil
}

func defaultBuiltInPipelineDir() string {
	if fromEnv := os.Getenv("FABRIC_BUILTIN_PIPELINES_DIR"); fromEnv != "" {
		return fromEnv
	}

	_, thisFile, _, ok := runtime.Caller(0)
	if ok {
		return filepath.Clean(filepath.Join(filepath.Dir(thisFile), "..", "..", "data", "pipelines"))
	}

	return filepath.Join("data", "pipelines")
}

func defaultUserPipelineDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("determine user home directory: %w", err)
	}
	return filepath.Join(homeDir, ".config", "fabric", "pipelines"), nil
}

func (l *Loader) List() ([]DiscoveryEntry, error) {
	builtIns, err := l.scanDir(l.BuiltInDir, DefinitionSourceBuiltIn)
	if err != nil {
		return nil, err
	}

	users, err := l.scanDir(l.UserDir, DefinitionSourceUser)
	if err != nil {
		return nil, err
	}

	builtInByName := make(map[string]DiscoveryEntry, len(builtIns))
	for _, entry := range builtIns {
		builtInByName[entry.Name] = entry
	}

	merged := make(map[string]DiscoveryEntry, len(builtIns)+len(users))
	for _, entry := range builtIns {
		merged[entry.Name] = entry
	}
	for _, entry := range users {
		if _, exists := builtInByName[entry.Name]; exists {
			entry.OverridesBuiltIn = true
		}
		merged[entry.Name] = entry
	}

	names := make([]string, 0, len(merged))
	for name := range merged {
		names = append(names, name)
	}
	sort.Strings(names)

	result := make([]DiscoveryEntry, 0, len(names))
	for _, name := range names {
		result = append(result, merged[name])
	}

	return result, nil
}

func (l *Loader) LoadNamed(name string) (*Pipeline, error) {
	entries, err := l.List()
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.Name == name {
			return l.LoadFile(entry.Path, entry.DefinitionSource)
		}
	}

	return nil, fmt.Errorf("pipeline %q not found", name)
}

func (l *Loader) LoadFile(path string, source DefinitionSource) (*Pipeline, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read pipeline file %s: %w", path, err)
	}

	var pipe Pipeline
	decoder := yaml.NewDecoder(bytes.NewReader(content))
	decoder.KnownFields(true)
	if err := decoder.Decode(&pipe); err != nil {
		return nil, fmt.Errorf("parse pipeline file %s: %w", path, err)
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("resolve pipeline path %s: %w", path, err)
	}

	fileName := filepath.Base(absPath)
	fileStem := strings.TrimSuffix(fileName, filepath.Ext(fileName))

	pipe.FilePath = absPath
	pipe.FileName = fileName
	pipe.FileStem = fileStem
	pipe.DefinitionSource = source

	return &pipe, nil
}

func (l *Loader) LoadFilePath(path string) (*Pipeline, error) {
	return l.LoadFile(path, DefinitionSourceUser)
}

func (l *Loader) scanDir(dir string, source DefinitionSource) ([]DiscoveryEntry, error) {
	if dir == "" {
		return nil, nil
	}

	info, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("stat pipeline directory %s: %w", dir, err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("pipeline directory %s is not a directory", dir)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read pipeline directory %s: %w", dir, err)
	}

	var result []DiscoveryEntry
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if filepath.Ext(entry.Name()) != ".yaml" {
			continue
		}

		result = append(result, DiscoveryEntry{
			Name:             strings.TrimSuffix(entry.Name(), ".yaml"),
			Path:             filepath.Join(dir, entry.Name()),
			DefinitionSource: source,
		})
	}

	sort.Slice(result, func(i, j int) bool { return result[i].Name < result[j].Name })
	return result, nil
}
