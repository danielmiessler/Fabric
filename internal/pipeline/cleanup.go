package pipeline

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	cleanupRunDirEnv   = "FABRIC_PIPELINE_CLEANUP_RUN_DIR"
	cleanupDelayMsEnv  = "FABRIC_PIPELINE_CLEANUP_DELAY_MS"
	defaultCleanupWait = 5 * time.Second
)

type RunState struct {
	RunID       string     `json:"run_id"`
	Status      string     `json:"status"`
	PID         int        `json:"pid"`
	StartedAt   time.Time  `json:"started_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
	Pipeline    string     `json:"pipeline"`
	RunDir      string     `json:"run_dir"`
	SourceMode  SourceMode `json:"source_mode"`
}

func CleanupRunDirFromEnv() (bool, error) {
	runDir := os.Getenv(cleanupRunDirEnv)
	if runDir == "" {
		return false, nil
	}

	delay := defaultCleanupWait
	if rawDelay := os.Getenv(cleanupDelayMsEnv); rawDelay != "" {
		var millis int
		if _, err := fmt.Sscanf(rawDelay, "%d", &millis); err != nil {
			return true, fmt.Errorf("parse %s: %w", cleanupDelayMsEnv, err)
		}
		delay = time.Duration(millis) * time.Millisecond
	}

	time.Sleep(delay)
	if err := deleteRunDir(runDir); err != nil {
		return true, err
	}
	return true, nil
}

func startCleanupHelper(runDir string, delay time.Duration) error {
	executable, err := os.Executable()
	if err != nil {
		return fmt.Errorf("resolve current executable for cleanup helper: %w", err)
	}

	cmd := exec.Command(executable)
	cmd.Env = append(os.Environ(),
		cleanupRunDirEnv+"="+runDir,
		cleanupDelayMsEnv+"="+fmt.Sprintf("%d", delay.Milliseconds()),
	)
	cmd.Stdout = nil
	cmd.Stderr = nil
	cmd.Stdin = nil
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("start cleanup helper: %w", err)
	}
	return nil
}

func cleanupExpiredRuns(runRoot string, now time.Time) error {
	info, err := os.Stat(runRoot)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("stat run root %s: %w", runRoot, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("run root %s is not a directory", runRoot)
	}

	entries, err := os.ReadDir(runRoot)
	if err != nil {
		return fmt.Errorf("read run root %s: %w", runRoot, err)
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		runDir := filepath.Join(runRoot, entry.Name())
		state, err := readRunState(filepath.Join(runDir, "run.json"))
		if err != nil {
			continue
		}
		if state.ExpiresAt != nil && !state.ExpiresAt.After(now) {
			if err := deleteRunDir(runDir); err != nil {
				return err
			}
		}
	}

	return removeRunRootIfEmpty(runRoot)
}

func readRunState(path string) (*RunState, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var state RunState
	if err := json.Unmarshal(content, &state); err != nil {
		return nil, err
	}
	return &state, nil
}

func deleteRunDir(runDir string) error {
	if err := os.RemoveAll(runDir); err != nil {
		return fmt.Errorf("remove run dir %s: %w", runDir, err)
	}
	return removeRunRootIfEmpty(filepath.Dir(runDir))
}

func removeRunRootIfEmpty(runRoot string) error {
	entries, err := os.ReadDir(runRoot)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("read run root %s: %w", runRoot, err)
	}
	if len(entries) > 0 {
		return nil
	}
	if err := os.Remove(runRoot); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("remove empty run root %s: %w", runRoot, err)
	}
	return nil
}
