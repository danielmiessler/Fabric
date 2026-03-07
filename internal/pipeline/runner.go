package pipeline

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/core"
)

type Runner struct {
	Stdout   io.Writer
	Stderr   io.Writer
	registry *core.PluginRegistry
}

func NewRunner(stdout, stderr io.Writer, registry *core.PluginRegistry) *Runner {
	return &Runner{Stdout: stdout, Stderr: stderr, registry: registry}
}

type RunOptions struct {
	InvocationDir  string
	CleanupDelay   time.Duration
	DisableCleanup bool
}

type RunResult struct {
	RunID       string
	RunDir      string
	FinalOutput string
}

type StageExecutionResult struct {
	Stdout       string
	FilesWritten []string
}

func (r *Runner) Run(ctx context.Context, p *Pipeline, source RunSource, opts RunOptions) (*RunResult, error) {
	if err := Preflight(ctx, p, PreflightOptions{Registry: r.registry}); err != nil {
		return nil, err
	}
	if err := validateAcceptedSource(p, source.Mode); err != nil {
		return nil, err
	}
	if opts.InvocationDir == "" {
		dir, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("determine working directory: %w", err)
		}
		opts.InvocationDir = dir
	}
	if opts.CleanupDelay <= 0 {
		opts.CleanupDelay = 5 * time.Second
	}

	runRoot := filepath.Join(opts.InvocationDir, ".pipeline")
	if !opts.DisableCleanup {
		if err := cleanupExpiredRuns(runRoot, time.Now().UTC()); err != nil {
			return nil, err
		}
	}

	now := time.Now().UTC()
	runID := fmt.Sprintf("%s-%09d", now.Format("20060102T150405Z"), now.Nanosecond())
	runDir := filepath.Join(runRoot, runID)
	if err := os.MkdirAll(runDir, 0o755); err != nil {
		return nil, fmt.Errorf("create run directory %s: %w", runDir, err)
	}

	manifest := &RunManifest{
		RunID:        runID,
		PipelineName: p.Name,
		PipelineFile: p.FilePath,
		Status:       "running",
		StartedAt:    now,
		Source: RunSourceManifest{
			Mode:      source.Mode,
			Reference: source.Reference,
		},
		Stages: make([]RunStageManifest, len(p.Stages)),
	}
	runState := &RunState{
		RunID:      runID,
		Status:     "running",
		PID:        os.Getpid(),
		StartedAt:  now,
		UpdatedAt:  now,
		Pipeline:   p.Name,
		RunDir:     runDir,
		SourceMode: source.Mode,
	}

	for i, stage := range p.Stages {
		manifest.Stages[i] = RunStageManifest{
			ID:       stage.ID,
			Role:     effectiveStageRole(stage),
			Executor: stage.Executor,
			Status:   "pending",
		}
	}

	if err := writeJSON(filepath.Join(runDir, "run_manifest.json"), manifest); err != nil {
		return nil, err
	}
	if err := writeJSON(filepath.Join(runDir, "run.json"), runState); err != nil {
		return nil, err
	}
	if err := writeJSON(filepath.Join(runDir, "source_manifest.json"), SourceManifest{
		Mode:         source.Mode,
		Reference:    source.Reference,
		PayloadBytes: len(source.Payload),
	}); err != nil {
		return nil, err
	}

	stagePayloads := make(map[string]string, len(p.Stages))
	stageArtifacts := make(map[string]map[string]string, len(p.Stages))
	finalOutput := ""
	finalOutputStageID := ""
	finalOutputEmitted := false
	lastValidateStageIndex := findLastValidateStageIndex(p.Stages)
	validationSatisfied := lastValidateStageIndex == -1
	noValidateWarningEmitted := false

	for i := range p.Stages {
		stage := p.Stages[i]
		stageStart := time.Now().UTC()
		manifest.Stages[i].Status = "running"
		manifest.Stages[i].StartedAt = &stageStart
		runState.UpdatedAt = stageStart
		_ = writeJSON(filepath.Join(runDir, "run_manifest.json"), manifest)
		_ = writeJSON(filepath.Join(runDir, "run.json"), runState)
		fmt.Fprintf(r.Stderr, "[%d/%d] %s ........ RUNNING\n", i+1, len(p.Stages), stage.ID)

		inputPayload, err := resolveStageInput(stage, i, source.Payload, p.Stages, stagePayloads, stageArtifacts)
		if err != nil {
			if r.shouldEmitFinalOutputOnFailure(stage, finalOutput, finalOutputEmitted, validationSatisfied) {
				if lastValidateStageIndex == -1 {
					noValidateWarningEmitted = r.emitNoValidateWarning(runDir, manifest, p.Name, noValidateWarningEmitted)
				}
				r.emitFinalOutput(finalOutput)
				finalOutputEmitted = true
			}
			return nil, r.failRun(runDir, manifest, runState, i, opts.CleanupDelay, opts.DisableCleanup, finalOutputStageID, finalOutput, err)
		}

		runtimeCtx := StageRuntimeContext{
			Pipeline:       p,
			Stage:          stage,
			Source:         source,
			InputPayload:   inputPayload,
			InvocationDir:  opts.InvocationDir,
			RunDir:         runDir,
			RunID:          runID,
			StageArtifacts: stageArtifacts,
			StagePayloads:  stagePayloads,
			Manifest:       manifest,
			FinalOutput:    finalOutput,
		}
		execResult, err := r.executeStage(ctx, stage, runtimeCtx)
		if err != nil {
			if r.shouldEmitFinalOutputOnFailure(stage, finalOutput, finalOutputEmitted, validationSatisfied) {
				if lastValidateStageIndex == -1 {
					noValidateWarningEmitted = r.emitNoValidateWarning(runDir, manifest, p.Name, noValidateWarningEmitted)
				}
				r.emitFinalOutput(finalOutput)
				finalOutputEmitted = true
			}
			return nil, r.failRun(runDir, manifest, runState, i, opts.CleanupDelay, opts.DisableCleanup, finalOutputStageID, finalOutput, err)
		}

		artifactPaths, writtenFiles, primaryPayload, err := resolveStageOutputs(stage, runDir, execResult.Stdout, execResult.FilesWritten)
		if err != nil {
			if r.shouldEmitFinalOutputOnFailure(stage, finalOutput, finalOutputEmitted, validationSatisfied) {
				if lastValidateStageIndex == -1 {
					noValidateWarningEmitted = r.emitNoValidateWarning(runDir, manifest, p.Name, noValidateWarningEmitted)
				}
				r.emitFinalOutput(finalOutput)
				finalOutputEmitted = true
			}
			return nil, r.failRun(runDir, manifest, runState, i, opts.CleanupDelay, opts.DisableCleanup, finalOutputStageID, finalOutput, err)
		}

		stageArtifacts[stage.ID] = artifactPaths
		stagePayloads[stage.ID] = primaryPayload
		manifest.Stages[i].Files = displayPathsForRun(runDir, writtenFiles)
		if stage.FinalOutput {
			finalOutput = primaryPayload
			finalOutputStageID = stage.ID
			manifest.FinalOutput = &FinalOutputReport{
				StageID: stage.ID,
				Bytes:   len(primaryPayload),
			}
		}
		if effectiveStageRole(stage) == StageRoleValidate && i == lastValidateStageIndex {
			validationSatisfied = true
		}

		stageEnd := time.Now().UTC()
		manifest.Stages[i].Status = "passed"
		manifest.Stages[i].FinishedAt = &stageEnd
		runState.UpdatedAt = stageEnd
		_ = writeJSON(filepath.Join(runDir, "run_manifest.json"), manifest)
		_ = writeJSON(filepath.Join(runDir, "run.json"), runState)
		fmt.Fprintf(r.Stderr, "[%d/%d] %s ........ PASS\n", i+1, len(p.Stages), stage.ID)
		if len(manifest.Stages[i].Files) > 0 {
			fmt.Fprintf(r.Stderr, "           files: %s\n", strings.Join(manifest.Stages[i].Files, ", "))
		}
	}

	finishedAt := time.Now().UTC()
	manifest.Status = "passed"
	manifest.FinishedAt = &finishedAt
	runState.Status = "completed"
	runState.UpdatedAt = finishedAt
	runState.CompletedAt = &finishedAt
	expiresAt := finishedAt.Add(opts.CleanupDelay)
	runState.ExpiresAt = &expiresAt

	if lastValidateStageIndex == -1 && finalOutput != "" {
		noValidateWarningEmitted = r.emitNoValidateWarning(runDir, manifest, p.Name, noValidateWarningEmitted)
	}

	if err := writeJSON(filepath.Join(runDir, "run_manifest.json"), manifest); err != nil {
		return nil, err
	}
	if err := writeJSON(filepath.Join(runDir, "run.json"), runState); err != nil {
		return nil, err
	}

	if finalOutput != "" {
		r.emitFinalOutput(finalOutput)
	}
	r.emitRunSummary(manifest, runDir)
	if !opts.DisableCleanup {
		if err := startCleanupHelper(runDir, opts.CleanupDelay); err != nil {
			return nil, err
		}
	}

	return &RunResult{
		RunID:       runID,
		RunDir:      runDir,
		FinalOutput: finalOutput,
	}, nil
}

func (r *Runner) executeStage(ctx context.Context, stage Stage, runtimeCtx StageRuntimeContext) (*StageExecutionResult, error) {
	switch stage.Executor {
	case ExecutorBuiltin:
		return r.executeBuiltinStage(ctx, stage, runtimeCtx)
	case ExecutorCommand:
		return r.executeCommandStage(ctx, stage, runtimeCtx)
	case ExecutorFabricPattern:
		return r.executePatternStage(ctx, stage, runtimeCtx)
	default:
		return nil, fmt.Errorf("unsupported executor %q", stage.Executor)
	}
}

func resolveStageInput(stage Stage, index int, sourcePayload string, stages []Stage, stagePayloads map[string]string, stageArtifacts map[string]map[string]string) (string, error) {
	if stage.Input == nil {
		if index == 0 {
			return sourcePayload, nil
		}
		prevStage := stages[index-1]
		return stagePayloads[prevStage.ID], nil
	}

	switch stage.Input.From {
	case StageInputSource:
		return sourcePayload, nil
	case StageInputPrevious:
		if index == 0 {
			return sourcePayload, nil
		}
		prevStage := stages[index-1]
		return stagePayloads[prevStage.ID], nil
	case StageInputArtifact:
		stageArtifactsForStage := stageArtifacts[stage.Input.Stage]
		if stageArtifactsForStage == nil {
			return "", fmt.Errorf("stage %q artifacts are not available", stage.Input.Stage)
		}
		artifactPath := stageArtifactsForStage[stage.Input.Artifact]
		if artifactPath == "" {
			return "", fmt.Errorf("stage %q artifact %q is not available", stage.Input.Stage, stage.Input.Artifact)
		}
		content, err := os.ReadFile(artifactPath)
		if err != nil {
			return "", fmt.Errorf("read artifact %s: %w", artifactPath, err)
		}
		return string(content), nil
	default:
		return "", fmt.Errorf("unsupported input.from %q", stage.Input.From)
	}
}

func buildArtifactMap(stage Stage, runDir string) map[string]string {
	result := make(map[string]string, len(stage.Artifacts))
	for _, artifact := range stage.Artifacts {
		result[artifact.Name] = filepath.Join(runDir, artifact.Path)
	}
	return result
}

func resolveStageOutputs(stage Stage, runDir string, stdout string, execWrittenFiles []string) (map[string]string, []string, string, error) {
	artifactPaths := buildArtifactMap(stage, runDir)
	writtenFiles := append([]string{}, execWrittenFiles...)
	for _, artifact := range stage.Artifacts {
		path := artifactPaths[artifact.Name]
		if stage.PrimaryOutput != nil && stage.PrimaryOutput.From == PrimaryOutputArtifact && stage.PrimaryOutput.Artifact == artifact.Name {
			if _, statErr := os.Stat(path); os.IsNotExist(statErr) && stdout != "" {
				if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
					return nil, nil, "", fmt.Errorf("create artifact directory %s: %w", filepath.Dir(path), err)
				}
				if err := os.WriteFile(path, []byte(stdout), 0o644); err != nil {
					return nil, nil, "", fmt.Errorf("write primary artifact %s: %w", path, err)
				}
				writtenFiles = append(writtenFiles, path)
			}
		}

		_, statErr := os.Stat(path)
		if artifact.IsRequired() && statErr != nil {
			if os.IsNotExist(statErr) {
				return nil, nil, "", fmt.Errorf("required artifact %q was not produced at %s", artifact.Name, path)
			}
			return nil, nil, "", fmt.Errorf("stat artifact %s: %w", path, statErr)
		}
		if os.IsNotExist(statErr) {
			delete(artifactPaths, artifact.Name)
			continue
		}
		writtenFiles = append(writtenFiles, path)
	}

	primaryPayload := stdout
	if stage.PrimaryOutput != nil && stage.PrimaryOutput.From == PrimaryOutputArtifact {
		path := artifactPaths[stage.PrimaryOutput.Artifact]
		if path == "" {
			return nil, nil, "", fmt.Errorf("primary artifact %q was not produced", stage.PrimaryOutput.Artifact)
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return nil, nil, "", fmt.Errorf("read primary artifact %s: %w", path, err)
		}
		primaryPayload = string(content)
	}

	return artifactPaths, dedupeStrings(writtenFiles), primaryPayload, nil
}

func (r *Runner) failRun(runDir string, manifest *RunManifest, runState *RunState, stageIndex int, cleanupDelay time.Duration, disableCleanup bool, finalOutputStageID string, finalOutput string, err error) error {
	stageEnd := time.Now().UTC()
	manifest.Stages[stageIndex].Status = "failed"
	manifest.Stages[stageIndex].FinishedAt = &stageEnd
	manifest.Stages[stageIndex].Error = err.Error()
	manifest.Status = "failed"
	manifest.FinishedAt = &stageEnd
	if finalOutputStageID != "" && finalOutput != "" && manifest.FinalOutput == nil {
		manifest.FinalOutput = &FinalOutputReport{
			StageID: finalOutputStageID,
			Bytes:   len(finalOutput),
		}
	}
	runState.Status = "failed"
	runState.UpdatedAt = stageEnd
	runState.CompletedAt = &stageEnd
	expiresAt := stageEnd.Add(cleanupDelay)
	runState.ExpiresAt = &expiresAt

	_ = writeJSON(filepath.Join(runDir, "run_manifest.json"), manifest)
	_ = writeJSON(filepath.Join(runDir, "run.json"), runState)
	fmt.Fprintf(r.Stderr, "[%d/%d] %s ........ FAIL\n", stageIndex+1, len(manifest.Stages), manifest.Stages[stageIndex].ID)
	r.emitRunSummary(manifest, runDir)
	if !disableCleanup {
		_ = startCleanupHelper(runDir, cleanupDelay)
	}
	return err
}

func (r *Runner) emitFinalOutput(output string) {
	if output == "" {
		return
	}
	fmt.Fprintln(r.Stdout, strings.TrimRight(output, "\n"))
}

func (r *Runner) shouldEmitFinalOutputOnFailure(stage Stage, finalOutput string, finalOutputEmitted bool, validationSatisfied bool) bool {
	if finalOutput == "" || finalOutputEmitted {
		return false
	}
	if effectiveStageRole(stage) != StageRolePublish {
		return false
	}
	return validationSatisfied
}

func (r *Runner) emitNoValidateWarning(runDir string, manifest *RunManifest, pipelineName string, alreadyEmitted bool) bool {
	if alreadyEmitted {
		return true
	}
	warning := fmt.Sprintf("warning: pipeline %s has no validate stage", pipelineName)
	fmt.Fprintln(r.Stderr, warning)
	manifest.Warnings = append(manifest.Warnings, warning)
	_ = writeJSON(filepath.Join(runDir, "run_manifest.json"), manifest)
	return true
}

func (r *Runner) emitRunSummary(manifest *RunManifest, runDir string) {
	finalBytes := 0
	finalStageID := ""
	if manifest.FinalOutput != nil {
		finalBytes = manifest.FinalOutput.Bytes
		finalStageID = manifest.FinalOutput.StageID
	}
	fmt.Fprintf(r.Stderr, "run summary: status=%s run_id=%s run_dir=%s final_stage=%s final_bytes=%d\n", manifest.Status, manifest.RunID, runDir, finalStageID, finalBytes)
}

func displayPathsForRun(runDir string, paths []string) []string {
	display := make([]string, 0, len(paths))
	for _, path := range dedupeStrings(paths) {
		if path == "" {
			continue
		}
		if rel, err := filepath.Rel(runDir, path); err == nil && !strings.HasPrefix(rel, "..") {
			display = append(display, rel)
			continue
		}
		display = append(display, path)
	}
	return display
}

func dedupeStrings(values []string) []string {
	seen := make(map[string]struct{}, len(values))
	result := make([]string, 0, len(values))
	for _, value := range values {
		if value == "" {
			continue
		}
		if _, exists := seen[value]; exists {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}

func validateAcceptedSource(p *Pipeline, mode SourceMode) error {
	if len(p.Accepts) == 0 {
		return nil
	}
	for _, allowed := range p.Accepts {
		if allowed == mode {
			return nil
		}
	}
	return fmt.Errorf("pipeline %q does not accept source mode %q", p.Name, mode)
}

func writeJSON(path string, value any) error {
	content, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal %s: %w", path, err)
	}
	if err := os.WriteFile(path, content, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}
