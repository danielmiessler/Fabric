package fsdb

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func NewDb(dir string) (db *Db) {

	db = &Db{Dir: dir}

	db.EnvFilePath = db.FilePath(".env")

	db.Patterns = &PatternsEntity{
		StorageEntity:          &StorageEntity{Label: "Patterns", Dir: db.FilePath("patterns"), ItemIsDir: true},
		SystemPatternFile:      "system.md",
		UniquePatternsFilePath: db.FilePath("unique_patterns.txt"),
		CustomPatternsDir:      "", // Will be set after loading .env file
	}

	db.Sessions = &SessionsEntity{
		&StorageEntity{Label: "Sessions", Dir: db.FilePath("sessions"), FileExtension: ".json"}}

	db.Contexts = &ContextsEntity{
		&StorageEntity{Label: "Contexts", Dir: db.FilePath("contexts")}}

	return
}

type Db struct {
	Dir string

	Patterns *PatternsEntity
	Sessions *SessionsEntity
	Contexts *ContextsEntity

	EnvFilePath string
}

func (o *Db) Configure() (err error) {
	if err = os.MkdirAll(o.Dir, os.ModePerm); err != nil {
		return
	}

	if err = o.LoadEnvFile(); err != nil {
		return
	}

	// Set custom patterns directory after loading .env file
	customPatternsDir := os.Getenv("CUSTOM_PATTERNS_DIRECTORY")
	if customPatternsDir != "" {
		// Expand home directory if needed
		if strings.HasPrefix(customPatternsDir, "~/") {
			if homeDir, err := os.UserHomeDir(); err == nil {
				customPatternsDir = filepath.Join(homeDir, customPatternsDir[2:])
			}
		}
		o.Patterns.CustomPatternsDir = customPatternsDir
	}

	if err = o.Patterns.Configure(); err != nil {
		return
	}

	if err = o.Sessions.Configure(); err != nil {
		return
	}

	if err = o.Contexts.Configure(); err != nil {
		return
	}

	return
}

func (o *Db) LoadEnvFile() (err error) {
	if err = godotenv.Load(o.EnvFilePath); err != nil {
		err = fmt.Errorf("error loading .env file: %s", err)
	}
	return
}

func (o *Db) IsEnvFileExists() (ret bool) {
	_, err := os.Stat(o.EnvFilePath)
	ret = !os.IsNotExist(err)
	return
}

func (o *Db) SaveEnv(content string) (err error) {
	err = os.WriteFile(o.EnvFilePath, []byte(content), 0644)
	return
}

func (o *Db) FilePath(fileName string) (ret string) {
	return filepath.Join(o.Dir, fileName)
}

type DirectoryChange struct {
	Dir       string
	Timestamp time.Time
}
