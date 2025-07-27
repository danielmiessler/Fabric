package fsdb

import (
	"testing"
)

func TestStorage_SaveAndLoad(t *testing.T) {
	dir := t.TempDir()
	storage := &StorageEntity{Dir: dir}
	name := "test"
	content := []byte("test content")
	if err := storage.Save(name, content); err != nil {
		t.Fatalf("failed to save content: %v", err)
	}
	loadedContent, err := storage.Load(name)
	if err != nil {
		t.Fatalf("failed to load content: %v", err)
	}
	if string(loadedContent) != string(content) {
		t.Errorf("expected %v, got %v", string(content), string(loadedContent))
	}
}

func TestStorage_Exists(t *testing.T) {
	dir := t.TempDir()
	storage := &StorageEntity{Dir: dir}
	name := "test"
	if storage.Exists(name) {
		t.Errorf("expected file to not exist")
	}
	if err := storage.Save(name, []byte("test content")); err != nil {
		t.Fatalf("failed to save content: %v", err)
	}
	if !storage.Exists(name) {
		t.Errorf("expected file to exist")
	}
}

func TestStorage_Delete(t *testing.T) {
	dir := t.TempDir()
	storage := &StorageEntity{Dir: dir}
	name := "test"
	if err := storage.Save(name, []byte("test content")); err != nil {
		t.Fatalf("failed to save content: %v", err)
	}
	if err := storage.Delete(name); err != nil {
		t.Fatalf("failed to delete content: %v", err)
	}
	if storage.Exists(name) {
		t.Errorf("expected file to be deleted")
	}
}
