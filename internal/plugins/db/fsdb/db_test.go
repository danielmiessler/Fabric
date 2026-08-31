package fsdb

import (
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"testing"

	"github.com/joho/godotenv"
)

func TestDb_Configure(t *testing.T) {
	dir := t.TempDir()
	db := NewDb(dir)
	err := db.Configure()
	if err == nil {
		t.Fatalf("db is configured, but must not be at empty dir: %v", dir)
	}
	if db.IsEnvFileExists() {
		t.Fatalf("db file exists, but must not be at empty dir: %v", dir)
	}

	err = db.SaveEnv("")
	if err != nil {
		t.Fatalf("db can't save env for empty conf.: %v", err)
	}

	err = db.Configure()
	if err != nil {
		t.Fatalf("db is not configured, but shall be after save: %v", err)
	}
}

func TestDb_LoadEnvFile(t *testing.T) {
	dir := t.TempDir()
	db := NewDb(dir)
	content := "KEY=VALUE\n"
	err := os.WriteFile(db.EnvFilePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to write .env file: %v", err)
	}
	err = db.LoadEnvFile()
	if err != nil {
		t.Errorf("failed to load .env file: %v", err)
	}
}

func TestDb_SaveEnv(t *testing.T) {
	dir := t.TempDir()
	db := NewDb(dir)
	content := "KEY=VALUE\n"
	err := db.SaveEnv(content)
	if err != nil {
		t.Errorf("failed to save .env file: %v", err)
	}
	if _, err := os.Stat(db.EnvFilePath); os.IsNotExist(err) {
		t.Errorf("expected .env file to be saved")
	}
	assertEnvMode(t, db.EnvFilePath)
}

func TestDb_UpdateEnvVars(t *testing.T) {
	dir := t.TempDir()
	db := NewDb(dir)
	if err := db.SaveEnv("KEEP=old\nCODEX_REFRESH_TOKEN=stale\n"); err != nil {
		t.Fatalf("SaveEnv() error = %v", err)
	}

	if err := db.UpdateEnvVars(map[string]string{
		"CODEX_REFRESH_TOKEN": "rotated",
		"CODEX_ACCESS_TOKEN":  "fresh",
	}); err != nil {
		t.Fatalf("UpdateEnvVars() error = %v", err)
	}

	parsed, err := godotenv.Read(db.EnvFilePath)
	if err != nil {
		t.Fatalf("godotenv.Read() error = %v", err)
	}
	if parsed["KEEP"] != "old" {
		t.Fatalf("KEEP = %q, want old", parsed["KEEP"])
	}
	if parsed["CODEX_REFRESH_TOKEN"] != "rotated" {
		t.Fatalf("CODEX_REFRESH_TOKEN = %q, want rotated", parsed["CODEX_REFRESH_TOKEN"])
	}
	if parsed["CODEX_ACCESS_TOKEN"] != "fresh" {
		t.Fatalf("CODEX_ACCESS_TOKEN = %q, want fresh", parsed["CODEX_ACCESS_TOKEN"])
	}
	assertEnvMode(t, db.EnvFilePath)
}

func TestDb_UpdateEnvVars_MissingFile(t *testing.T) {
	dir := t.TempDir()
	db := NewDb(dir)

	if err := db.UpdateEnvVars(map[string]string{"CODEX_ACCESS_TOKEN": "fresh"}); err != nil {
		t.Fatalf("UpdateEnvVars() error = %v", err)
	}

	parsed, err := godotenv.Read(db.EnvFilePath)
	if err != nil {
		t.Fatalf("godotenv.Read() error = %v", err)
	}
	if parsed["CODEX_ACCESS_TOKEN"] != "fresh" {
		t.Fatalf("CODEX_ACCESS_TOKEN = %q, want fresh", parsed["CODEX_ACCESS_TOKEN"])
	}
	assertEnvMode(t, db.EnvFilePath)
}

func TestDb_UpdateEnvVars_CorruptFile(t *testing.T) {
	dir := t.TempDir()
	db := NewDb(dir)
	if err := os.Mkdir(db.EnvFilePath, 0700); err != nil {
		t.Fatalf("Mkdir() error = %v", err)
	}

	err := db.UpdateEnvVars(map[string]string{"CODEX_ACCESS_TOKEN": "fresh"})
	if err == nil {
		t.Fatal("UpdateEnvVars() error = nil, want corrupt-file error")
	}
}

func TestDb_UpdateEnvVars_SkipEmpty(t *testing.T) {
	dir := t.TempDir()
	db := NewDb(dir)
	if err := db.SaveEnv("CODEX_REFRESH_TOKEN=live\nKEEP=old\n"); err != nil {
		t.Fatalf("SaveEnv() error = %v", err)
	}

	if err := db.UpdateEnvVars(map[string]string{
		"CODEX_REFRESH_TOKEN": "",
		"CODEX_ACCESS_TOKEN":  "fresh",
	}); err != nil {
		t.Fatalf("UpdateEnvVars() error = %v", err)
	}

	parsed, err := godotenv.Read(db.EnvFilePath)
	if err != nil {
		t.Fatalf("godotenv.Read() error = %v", err)
	}
	if parsed["CODEX_REFRESH_TOKEN"] != "live" {
		t.Fatalf("CODEX_REFRESH_TOKEN = %q, want live", parsed["CODEX_REFRESH_TOKEN"])
	}
	if parsed["CODEX_ACCESS_TOKEN"] != "fresh" {
		t.Fatalf("CODEX_ACCESS_TOKEN = %q, want fresh", parsed["CODEX_ACCESS_TOKEN"])
	}
	if parsed["KEEP"] != "old" {
		t.Fatalf("KEEP = %q, want old", parsed["KEEP"])
	}
}

func TestDb_UpdateEnvVars_Concurrent(t *testing.T) {
	dir := t.TempDir()
	db := NewDb(dir)
	if err := db.SaveEnv("KEEP=old\n"); err != nil {
		t.Fatalf("SaveEnv() error = %v", err)
	}

	var wg sync.WaitGroup
	errs := make(chan error, 2)
	wg.Add(2)
	go func() {
		defer wg.Done()
		errs <- db.UpdateEnvVars(map[string]string{"CODEX_ACCESS_TOKEN": "one"})
	}()
	go func() {
		defer wg.Done()
		errs <- db.UpdateEnvVars(map[string]string{"CODEX_REFRESH_TOKEN": "two"})
	}()
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil {
			t.Fatalf("UpdateEnvVars() error = %v", err)
		}
	}

	parsed, err := godotenv.Read(db.EnvFilePath)
	if err != nil {
		t.Fatalf("godotenv.Read() error = %v", err)
	}
	if parsed["KEEP"] != "old" {
		t.Fatalf("KEEP = %q, want old", parsed["KEEP"])
	}
	if parsed["CODEX_ACCESS_TOKEN"] != "one" {
		t.Fatalf("CODEX_ACCESS_TOKEN = %q, want one", parsed["CODEX_ACCESS_TOKEN"])
	}
	if parsed["CODEX_REFRESH_TOKEN"] != "two" {
		t.Fatalf("CODEX_REFRESH_TOKEN = %q, want two", parsed["CODEX_REFRESH_TOKEN"])
	}
}

func assertEnvMode(t *testing.T, path string) {
	t.Helper()
	if runtime.GOOS == "windows" {
		return
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("Stat(%s) error = %v", path, err)
	}
	if perm := info.Mode().Perm(); perm != 0600 {
		t.Fatalf("%s mode = %o, want 0600", filepath.Base(path), perm)
	}
}
