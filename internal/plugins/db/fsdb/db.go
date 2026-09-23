package fsdb

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/danielmiessler/fabric/internal/i18n"
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

	envMu sync.Mutex
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
		err = fmt.Errorf(i18n.T("db_error_loading_env_file"), err)
	}
	return
}

func (o *Db) IsEnvFileExists() (ret bool) {
	_, err := os.Stat(o.EnvFilePath)
	ret = !os.IsNotExist(err)
	return
}

func (o *Db) SaveEnv(content string) error {
	return o.WithEnvLock(func() error {
		if err := writeFileAtomic(o.EnvFilePath, []byte(content)); err != nil {
			return fmt.Errorf(i18n.T("db_error_updating_env_file"), err)
		}
		return nil
	})
}

func (o *Db) ReadEnvFile() (map[string]string, error) {
	env, err := godotenv.Read(o.EnvFilePath)
	if err != nil {
		if o.IsEnvFileExists() {
			return nil, fmt.Errorf(i18n.T("db_error_loading_env_file"), err)
		}
		return map[string]string{}, nil
	}
	return env, nil
}

func (o *Db) WithEnvLock(fn func() error) error {
	o.envMu.Lock()
	defer o.envMu.Unlock()

	lockFile, err := os.OpenFile(o.EnvFilePath+".lock", os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return fmt.Errorf(i18n.T("db_error_updating_env_file"), err)
	}
	defer lockFile.Close()
	if err := lockExclusive(lockFile); err != nil {
		return fmt.Errorf(i18n.T("db_error_updating_env_file"), err)
	}
	defer unlockExclusive(lockFile)

	return fn()
}

// UpdateEnvVars merges non-empty values into .env under a file lock.
// Comments and key order are not preserved.
func (o *Db) UpdateEnvVars(updates map[string]string) error {
	return o.WithEnvLock(func() error {
		return o.ApplyEnvUpdates(updates)
	})
}

// ApplyEnvUpdates writes non-empty updates atomically. Callers holding WithEnvLock use this.
func (o *Db) ApplyEnvUpdates(updates map[string]string) error {
	env, err := o.ReadEnvFile()
	if err != nil {
		return err
	}
	for key, value := range updates {
		if strings.TrimSpace(value) == "" {
			continue
		}
		env[key] = value
	}
	if err := writeEnvFileAtomic(o.EnvFilePath, env); err != nil {
		return fmt.Errorf(i18n.T("db_error_updating_env_file"), err)
	}
	return nil
}

func writeEnvFileAtomic(path string, env map[string]string) error {
	content, err := godotenv.Marshal(env)
	if err != nil {
		return err
	}
	return writeFileAtomic(path, []byte(content+"\n"))
}

func writeFileAtomic(path string, content []byte) error {
	tmp, err := os.CreateTemp(filepath.Dir(path), ".env.tmp-")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName)

	if err := tmp.Chmod(0600); err != nil {
		_ = tmp.Close()
		return err
	}
	_, err = tmp.Write(content)
	if err == nil {
		err = tmp.Sync()
	}
	if cerr := tmp.Close(); err == nil {
		err = cerr
	}
	if err != nil {
		return err
	}
	return os.Rename(tmpName, path)
}

func (o *Db) FilePath(fileName string) (ret string) {
	return filepath.Join(o.Dir, fileName)
}

type DirectoryChange struct {
	Dir       string
	Timestamp time.Time
}
