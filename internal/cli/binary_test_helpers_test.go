package cli

import (
	"bytes"
	"context"
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

// preBuildBinaries builds all binaries used by tests upfront and in parallel.
// Called from TestMain so the cost is paid once before any test runs.
func preBuildBinaries() error {
	targets := []string{
		"cmd/fabric",
		"cmd/to_pdf",
		"cmd/code2context",
		"cmd/generate_changelog",
	}

	root := preBuildRepoRoot()
	errs := make([]error, len(targets))
	var wg sync.WaitGroup
	for i, target := range targets {
		wg.Add(1)
		go func(idx int, relativeDir string) {
			defer wg.Done()
			errs[idx] = doBuildGoBinary(root, relativeDir)
		}(i, target)
	}
	wg.Wait()

	for i, err := range errs {
		if err != nil {
			return fmt.Errorf("pre-build %s: %w", targets[i], err)
		}
	}
	return nil
}

// preBuildRepoRoot determines the repo root without a *testing.T.
func preBuildRepoRoot() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("unable to determine repo root")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(filename), "..", ".."))
}

// doBuildGoBinary performs the actual build, caching the result in binaryCache.
func doBuildGoBinary(root, relativeDir string) error {
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
		cmd.Dir = root

		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		entry.err = cmd.Run()
		entry.log = stderr.String()
	})

	return entry.err
}

func TestMain(m *testing.M) {
	if err := preBuildBinaries(); err != nil {
		fmt.Fprintf(os.Stderr, "pre-build failed: %v\n", err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}

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

	// Try the cache first (populated by TestMain).
	if entryAny, ok := binaryCache.Load(relativeDir); ok {
		entry := entryAny.(*builtBinary)
		if entry.err != nil {
			t.Fatalf("build %s: %v\n%s", relativeDir, entry.err, entry.log)
		}
		return entry.path
	}

	// Fallback: build on demand if not pre-built (e.g. new binary added to a test).
	if err := doBuildGoBinary(repoRoot(t), relativeDir); err != nil {
		t.Fatalf("build %s: %v", relativeDir, err)
	}
	entryAny, _ := binaryCache.Load(relativeDir)
	return entryAny.(*builtBinary).path
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
	return h.runFabricContext(context.Background(), t, stdin, cwd, extraEnv, args...)
}

func (h *binaryHarness) runFabricContext(
	ctx context.Context, t *testing.T, stdin, cwd string, extraEnv []string, args ...string,
) commandResult {
	t.Helper()
	return runCommandContext(ctx, t, h.fabric, stdin, cwdOrDefault(cwd, h.repoRoot), h.baseEnv(extraEnv...), args...)
}

func runCommand(t *testing.T, binaryPath, stdin, cwd string, env []string, args ...string) commandResult {
	t.Helper()
	return runCommandContext(context.Background(), t, binaryPath, stdin, cwd, env, args...)
}

func runCommandContext(
	ctx context.Context, t *testing.T, binaryPath, stdin, cwd string, env []string, args ...string,
) commandResult {
	t.Helper()

	cmd := exec.CommandContext(ctx, binaryPath, args...)
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
	if ctxErr := ctx.Err(); ctxErr != nil {
		if errors.Is(ctxErr, context.DeadlineExceeded) {
			t.Fatalf("run %s %v: timed out: %v", binaryPath, args, ctxErr)
		}
		t.Fatalf("run %s %v: canceled: %v", binaryPath, args, ctxErr)
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
