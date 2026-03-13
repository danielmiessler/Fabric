package cli

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"
)

type builtBinary struct {
	once sync.Once
	path string
	err  error
	log  string
}

type commandResult struct {
	stdout   string
	stderr   string
	exitCode int
	args     []string
}

type binaryHarness struct {
	repoRoot string
	homeDir  string
	fabric   string
}

var binaryCache sync.Map

func newBinaryHarness(t *testing.T, patternNames ...string) *binaryHarness {
	t.Helper()

	if len(patternNames) == 0 {
		patternNames = []string{"summarize", "create_coding_feature"}
	}

	homeDir := filepath.Join(t.TempDir(), "home")
	configDir := filepath.Join(homeDir, ".config", "fabric")
	patternsDir := filepath.Join(configDir, "patterns")

	if err := os.MkdirAll(patternsDir, 0o755); err != nil {
		t.Fatalf("create patterns dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(configDir, ".env"), []byte{}, 0o644); err != nil {
		t.Fatalf("create env file: %v", err)
	}
	for _, patternName := range patternNames {
		copyPatternFixture(t, repoRoot(t), patternsDir, patternName)
	}
	if err := os.WriteFile(filepath.Join(patternsDir, "loaded"), []byte{}, 0o644); err != nil {
		t.Fatalf("create loaded marker: %v", err)
	}

	return &binaryHarness{
		repoRoot: repoRoot(t),
		homeDir:  homeDir,
		fabric:   buildGoBinary(t, "cmd/fabric"),
	}
}

func repoRoot(t *testing.T) string {
	t.Helper()
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("unable to determine repo root")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(filename), "..", ".."))
}

func buildGoBinary(t *testing.T, relativeDir string) string {
	t.Helper()

	entryAny, _ := binaryCache.LoadOrStore(relativeDir, &builtBinary{})
	entry := entryAny.(*builtBinary)

	entry.once.Do(func() {
		outputDir, err := os.MkdirTemp("", "fabric-binary-*")
		if err != nil {
			entry.err = err
			return
		}

		binaryName := filepath.Base(relativeDir)
		if runtime.GOOS == "windows" {
			binaryName += ".exe"
		}
		entry.path = filepath.Join(outputDir, binaryName)

		cmd := exec.Command("go", "build", "-o", entry.path, "./"+filepath.ToSlash(relativeDir))
		cmd.Dir = repoRoot(t)

		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		entry.err = cmd.Run()
		entry.log = stderr.String()
	})

	if entry.err != nil {
		t.Fatalf("build %s: %v\n%s", relativeDir, entry.err, entry.log)
	}

	return entry.path
}

func copyPatternFixture(t *testing.T, root, patternsDir, patternName string) {
	t.Helper()

	src := filepath.Join(root, "data", "patterns", patternName)
	dst := filepath.Join(patternsDir, patternName)
	if err := copyDir(src, dst); err != nil {
		t.Fatalf("copy pattern %s: %v", patternName, err)
	}
}

func copyDir(src, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", src)
	}
	if err := os.MkdirAll(dst, info.Mode()); err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())
		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
			continue
		}
		if err := copyFile(srcPath, dstPath); err != nil {
			return err
		}
	}

	return nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	info, err := in.Stat()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}

	out, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}

	return out.Close()
}

func (h *binaryHarness) baseEnv(extra ...string) []string {
	env := append([]string{}, os.Environ()...)
	env = append(env,
		"HOME="+h.homeDir,
		"USERPROFILE="+h.homeDir,
		"HOMEDRIVE=",
		"HOMEPATH=",
	)
	env = append(env, extra...)
	return env
}

func (h *binaryHarness) runFabric(t *testing.T, stdin, cwd string, extraEnv []string, args ...string) commandResult {
	t.Helper()
	return runCommand(t, h.fabric, stdin, cwdOrDefault(cwd, h.repoRoot), h.baseEnv(extraEnv...), args...)
}

func runCommand(t *testing.T, binaryPath, stdin, cwd string, env []string, args ...string) commandResult {
	t.Helper()

	cmd := exec.Command(binaryPath, args...)
	cmd.Dir = cwd
	cmd.Env = env
	if stdin != "" {
		cmd.Stdin = strings.NewReader(stdin)
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	result := commandResult{
		stdout:   stdout.String(),
		stderr:   stderr.String(),
		exitCode: 0,
		args:     append([]string{}, args...),
	}
	if err == nil {
		return result
	}

	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		result.exitCode = exitErr.ExitCode()
		return result
	}
	t.Fatalf("run %s %v: %v", binaryPath, args, err)
	return result
}

func cwdOrDefault(cwd, fallback string) string {
	if cwd != "" {
		return cwd
	}
	return fallback
}
