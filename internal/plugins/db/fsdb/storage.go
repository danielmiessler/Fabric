package fsdb

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/util"
)

// StorageEntity is the filesystem-backed db.Storage implementation.
// Each method that gets a name requires one that obeys
// ValidateStorageName. It rejects other names with
// *InvalidStorageNameError, which HTTP handlers map to 400.
type StorageEntity struct {
	Label         string
	Dir           string
	ItemIsDir     bool
	FileExtension string
}

func (o *StorageEntity) Configure() (err error) {
	if err = os.MkdirAll(o.Dir, os.ModePerm); err != nil {
		return
	}
	return
}

// GetNames finds all patterns in the patterns directory and enters the id, name, and pattern into a slice of Entry structs. it returns these entries or an error
func (o *StorageEntity) GetNames() (ret []string, err error) {
	// Resolve the directory path to an absolute path
	absDir, err := util.GetAbsolutePath(o.Dir)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("storage_error_resolve_directory"), err)
	}

	// Read the directory entries
	var entries []os.DirEntry
	if entries, err = os.ReadDir(absDir); err != nil {
		return nil, fmt.Errorf(i18n.T("storage_error_read_directory"), err)
	}

	for _, entry := range entries {
		entryPath := filepath.Join(absDir, entry.Name())

		// Get metadata for the entry, including symlink info
		fileInfo, err := os.Lstat(entryPath)
		if err != nil {
			return nil, fmt.Errorf(i18n.T("storage_error_stat_entry"), entryPath, err)
		}

		// Determine if the entry should be included
		if o.ItemIsDir {
			// Include directories or symlinks to directories
			if fileInfo.IsDir() || (fileInfo.Mode()&os.ModeSymlink != 0 && util.IsSymlinkToDir(entryPath)) {
				ret = append(ret, entry.Name())
			}
		} else {
			// Include files, optionally filtering by extension
			if !fileInfo.IsDir() {
				if o.FileExtension == "" || filepath.Ext(entry.Name()) == o.FileExtension {
					ret = append(ret, strings.TrimSuffix(entry.Name(), o.FileExtension))
				}
			}
		}
	}

	return ret, nil
}

func (o *StorageEntity) Delete(name string) (err error) {
	var path string
	if path, err = o.resolvedPath(name); err != nil {
		return
	}
	if err = os.RemoveAll(path); err != nil {
		err = fmt.Errorf(i18n.T("storage_error_delete"), name, err)
	}
	return
}

func (o *StorageEntity) Exists(name string) (ret bool) {
	path, err := o.resolvedPath(name)
	if err != nil {
		return false
	}
	_, err = os.Stat(path)
	ret = !os.IsNotExist(err)
	return
}

func (o *StorageEntity) Rename(oldName, newName string) (err error) {
	var oldPath, newPath string
	if oldPath, err = o.resolvedPath(oldName); err != nil {
		return
	}
	if newPath, err = o.resolvedPath(newName); err != nil {
		return
	}
	if err = os.Rename(oldPath, newPath); err != nil {
		err = fmt.Errorf(i18n.T("storage_error_rename"), oldName, newName, err)
	}
	return
}

func (o *StorageEntity) Save(name string, content []byte) (err error) {
	var path string
	if path, err = o.resolvedPath(name); err != nil {
		return
	}
	if err = os.WriteFile(path, content, 0644); err != nil {
		err = fmt.Errorf(i18n.T("storage_error_save"), name, err)
	}
	return
}

func (o *StorageEntity) Load(name string) (ret []byte, err error) {
	var path string
	if path, err = o.resolvedPath(name); err != nil {
		return
	}
	if ret, err = os.ReadFile(path); err != nil {
		err = fmt.Errorf(i18n.T("storage_error_load"), name, err)
	}
	return
}

func (o *StorageEntity) ListNames(shellCompleteList bool) (err error) {
	var names []string
	if names, err = o.GetNames(); err != nil {
		return
	}

	if len(names) == 0 {
		if !shellCompleteList {
			fmt.Printf("%s\n", fmt.Sprintf(i18n.T("no_items_found"), o.Label))
		}
		return
	}

	for _, item := range names {
		fmt.Printf("%s\n", item)
	}
	return
}

func (o *StorageEntity) BuildFilePath(fileName string) (ret string) {
	ret = filepath.Join(o.Dir, fileName)
	return
}

func (o *StorageEntity) buildFileName(name string) string {
	return fmt.Sprintf("%s%v", name, o.FileExtension)
}

// InvalidStorageNameError reports a name that storage-name validation
// rejected. HTTP handlers map it to 400 Bad Request. All other storage
// errors stay 500 errors.
type InvalidStorageNameError struct {
	Name    string
	Message string // optional: the default is the storage_invalid_name translation
}

func (e *InvalidStorageNameError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return fmt.Sprintf(i18n.T("storage_invalid_name"), e.Name)
}

// windowsReservedNames are DOS device names. On Windows, these names
// identify devices, not files. The match ignores case and all text after
// the first dot, because Windows maps "CON.tar.gz" to the CON device.
var windowsReservedNames = map[string]bool{
	"CON": true, "PRN": true, "AUX": true, "NUL": true,
	"COM1": true, "COM2": true, "COM3": true, "COM4": true, "COM5": true,
	"COM6": true, "COM7": true, "COM8": true, "COM9": true,
	"LPT1": true, "LPT2": true, "LPT3": true, "LPT4": true, "LPT5": true,
	"LPT6": true, "LPT7": true, "LPT8": true, "LPT9": true,
}

// ValidateStorageName rejects an empty name, ".", "..", and each name
// that is not a single path element. It also rejects names that are
// dangerous only on Windows: names with ":" (an NTFS alternate data
// stream suffix), names with a dot or space at the end, and reserved
// DOS device names. Windows removes a dot or space at the end, and the
// shortened name then collides with an existing entry. The policy is
// the same on each platform, and an entry made on one system stays
// valid on the other systems. A Unix
// entry that already has a name against these rules shows in GetNames
// but is not accessible. To repair it, rename its file or directory on
// disk. Call this function before you join a name to a storage
// directory.
func ValidateStorageName(name string) error {
	if name == "" || name == "." || name == ".." {
		return &InvalidStorageNameError{Name: name}
	}
	if strings.ContainsAny(name, `/\:`) {
		return &InvalidStorageNameError{Name: name}
	}
	if name != strings.TrimRight(name, ". ") {
		return &InvalidStorageNameError{Name: name}
	}
	base := strings.ToUpper(name)
	if i := strings.IndexByte(base, '.'); i >= 0 {
		base = base[:i]
	}
	if windowsReservedNames[base] {
		return &InvalidStorageNameError{Name: name}
	}
	return nil
}

// symlinkContained rejects an entry at path if the entry resolves,
// through symlinks, to a target outside absDir. A missing entry passes,
// because the lexical check in resolvedPath already keeps the path that
// a write will make in the directory. The two inputs must be absolute
// paths. If they are not, you cannot compare the resolved forms. The
// check does not fully prevent local races. If a hostile local writer
// enters the threat model, move to os.Root.
func symlinkContained(absDir, path, name string) error {
	target, err := filepath.EvalSymlinks(path)
	if err != nil {
		if os.IsNotExist(err) {
			if _, lerr := os.Lstat(path); os.IsNotExist(lerr) {
				return nil
			}
			// This is a dangling symlink. A write through it makes the
			// outside target.
			return &InvalidStorageNameError{Name: name}
		}
		return err
	}
	resolvedDir, err := filepath.EvalSymlinks(absDir)
	if err != nil {
		return err
	}
	rel, err := filepath.Rel(resolvedDir, target)
	if err != nil || !filepath.IsLocal(rel) {
		return &InvalidStorageNameError{Name: name}
	}
	return nil
}

// resolvedPath keeps name in the entity directory. It validates the
// name, checks containment again after absolute resolution, and
// rejects a symlinked entry that resolves out of the directory.
// Symlinks that stay in the directory are permitted. A storage
// directory that is a symlink is also permitted.
func (o *StorageEntity) resolvedPath(name string) (string, error) {
	if err := ValidateStorageName(name); err != nil {
		return "", err
	}
	absDir, err := filepath.Abs(o.Dir)
	if err != nil {
		return "", err
	}
	absFull, err := filepath.Abs(filepath.Join(o.Dir, o.buildFileName(name)))
	if err != nil {
		return "", err
	}
	rel, err := filepath.Rel(absDir, absFull)
	if err != nil || !filepath.IsLocal(rel) {
		return "", &InvalidStorageNameError{Name: name}
	}
	if err := symlinkContained(absDir, absFull, name); err != nil {
		return "", err
	}
	return absFull, nil
}

func (o *StorageEntity) SaveAsJson(name string, item any) (err error) {
	var jsonString []byte
	if jsonString, err = json.Marshal(item); err == nil {
		err = o.Save(name, jsonString)
	} else {
		err = fmt.Errorf(i18n.T("storage_error_marshal"), name, err)
	}

	return err
}

func (o *StorageEntity) LoadAsJson(name string, item any) (err error) {
	var content []byte
	if content, err = o.Load(name); err != nil {
		return
	}

	if err = json.Unmarshal(content, &item); err != nil {
		err = fmt.Errorf(i18n.T("storage_error_unmarshal"), name, err)
	}
	return
}
