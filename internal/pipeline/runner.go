package pipeline

import (
	"context"
	"encoding/json"
	"errors"
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

	stageExecutionOverride func(context.Context, Stage, StageRuntimeContext) (*StageExecutionResult, bool, error)
	jsonEvents             bool
}

func NewRunner(stdout, stderr io.Writer, registry *core.PluginRegistry) *Runner {
	return &Runner{Stdout: stdout, Stderr: stderr, registry: registry}
}

type RunOptions struct {
	InvocationDir  string
	CleanupDelay   time.Duration
	DisableCleanup bool
	FromStage      string
	ToStage        string
	OnlyStage      string
	JSONEvents     bool
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

type publishManifestRequest struct {
	stageIndex int
	path       string
}

func (r *Runner) Run(ctx context.Context, p *Pipeline, source RunSource, opts RunOptions) (*RunResult, error) {
	r.jsonEvents = opts.JSONEvents
	defer func() {
		r.jsonEvents = false
	}()

	if err := Preflight(ctx, p, PreflightOptions{
		Registry: r.registry,
		SkipPatternValidation: func(stage Stage) bool {
			return r.stageExecutionOverride != nil && stage.Executor == ExecutorFabricPattern
		},
	}); err != nil {
		return nil, err
	}
	if err := validateAcceptedSource(p, source.Mode); err != nil {
		return nil, err
	}

	selectedStageIndices, selectedStageSet, err := SelectStageIndices(p.Stages, opts.FromStage, opts.ToStage, opts.OnlyStage)
	if err != nil {
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
	result := &RunResult{
		RunID:  runID,
		RunDir: runDir,
	}

	for i, stage := range p.Stages {
		status := "pending"
		if _, selected := selectedStageSet[i]; !selected {
			status = "skipped"
		}
		manifest.Stages[i] = RunStageManifest{
			ID:       stage.ID,
			Role:     effectiveStageRole(stage),
			Executor: stage.Executor,
			Status:   status,
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
	publishManifestRequests := make([]publishManifestRequest, 0)
	finalOutput := ""
	finalOutputStageID := ""
	finalOutputEmitted := false
	lastValidateStageIndex := findLastValidateStageIndexForSelection(p.Stages, selectedStageSet)
	validationSatisfied := lastValidateStageIndex == -1
	noValidateWarningEmitted := false
	selectedStageIDs := stageIDsFromIndices(p.Stages, selectedStageIndices)
	r.emitJSONEvent("run_started", map[string]any{
		"run_id":             runID,
		"pipeline":           p.Name,
		"selected_stage_ids": selectedStageIDs,
		"selected_count":     len(selectedStageIndices),
	})

	for selectedIndex, stageIndex := range selectedStageIndices {
		stage := p.Stages[stageIndex]
		stageOrdinal := selectedIndex + 1
		stageTotal := len(selectedStageIndices)
		stageStart := time.Now().UTC()
		manifest.Stages[stageIndex].Status = "running"
		manifest.Stages[stageIndex].StartedAt = &stageStart
		runState.UpdatedAt = stageStart
		if err := persistRunState(runDir, manifest, runState); err != nil {
			return result, r.failRun(runDir, manifest, runState, stageIndex, stageOrdinal, stageTotal, opts.CleanupDelay, opts.DisableCleanup, finalOutputStageID, finalOutput, err, publishManifestRequests, source)
		}
		r.emitJSONEvent("stage_started", map[string]any{
			"run_id":       runID,
			"stage_id":     stage.ID,
			"stage_number": stageOrdinal,
			"stage_total":  stageTotal,
			"role":         effectiveStageRole(stage),
			"executor":     stage.Executor,
		})
		if !r.jsonEvents {
			fmt.Fprintf(r.Stderr, "[%d/%d] %s ........ RUNNING\n", stageOrdinal, stageTotal, stage.ID)
		}

		inputPayload, err := resolveStageInput(stage, stageIndex, source.Payload, p.Stages, stagePayloads, stageArtifacts)
		if err != nil {
			emitFinalOutputOnFailure := r.shouldEmitFinalOutputOnFailure(stage, finalOutput, finalOutputEmitted, validationSatisfied)
			if emitFinalOutputOnFailure {
				if lastValidateStageIndex == -1 {
					var warningErr error
					noValidateWarningEmitted, warningErr = r.emitNoValidateWarning(runDir, manifest, p.Name, noValidateWarningEmitted)
					if warningErr != nil {
						return result, r.failRun(runDir, manifest, runState, stageIndex, stageOrdinal, stageTotal, opts.CleanupDelay, opts.DisableCleanup, finalOutputStageID, finalOutput, errors.Join(err, warningErr), publishManifestRequests, source)
					}
				}
				result.FinalOutput = finalOutput
				r.emitFinalOutput(finalOutput)
				finalOutputEmitted = true
			} else {
				result.FinalOutput = ""
			}
			return result, r.failRun(runDir, manifest, runState, stageIndex, stageOrdinal, stageTotal, opts.CleanupDelay, opts.DisableCleanup, finalOutputStageID, finalOutput, err, publishManifestRequests, source)
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
			emitFinalOutputOnFailure := r.shouldEmitFinalOutputOnFailure(stage, finalOutput, finalOutputEmitted, validationSatisfied)
			if emitFinalOutputOnFailure {
				if lastValidateStageIndex == -1 {
					var warningErr error
					noValidateWarningEmitted, warningErr = r.emitNoValidateWarning(runDir, manifest, p.Name, noValidateWarningEmitted)
					if warningErr != nil {
						return result, r.failRun(runDir, manifest, runState, stageIndex, stageOrdinal, stageTotal, opts.CleanupDelay, opts.DisableCleanup, finalOutputStageID, finalOutput, errors.Join(err, warningErr), publishManifestRequests, source)
					}
				}
				result.FinalOutput = finalOutput
				r.emitFinalOutput(finalOutput)
				finalOutputEmitted = true
			} else {
				result.FinalOutput = ""
			}
			return result, r.failRun(runDir, manifest, runState, stageIndex, stageOrdinal, stageTotal, opts.CleanupDelay, opts.DisableCleanup, finalOutputStageID, finalOutput, err, publishManifestRequests, source)
		}

		artifactPaths, writtenFiles, primaryPayload, err := resolveStageOutputs(stage, runDir, execResult.Stdout, execResult.FilesWritten)
		if err != nil {
			emitFinalOutputOnFailure := r.shouldEmitFinalOutputOnFailure(stage, finalOutput, finalOutputEmitted, validationSatisfied)
			if emitFinalOutputOnFailure {
				if lastValidateStageIndex == -1 {
					var warningErr error
					noValidateWarningEmitted, warningErr = r.emitNoValidateWarning(runDir, manifest, p.Name, noValidateWarningEmitted)
					if warningErr != nil {
						return result, r.failRun(runDir, manifest, runState, stageIndex, stageOrdinal, stageTotal, opts.CleanupDelay, opts.DisableCleanup, finalOutputStageID, finalOutput, errors.Join(err, warningErr), publishManifestRequests, source)
					}
				}
				result.FinalOutput = finalOutput
				r.emitFinalOutput(finalOutput)
				finalOutputEmitted = true
			} else {
				result.FinalOutput = ""
			}
			return result, r.failRun(runDir, manifest, runState, stageIndex, stageOrdinal, stageTotal, opts.CleanupDelay, opts.DisableCleanup, finalOutputStageID, finalOutput, err, publishManifestRequests, source)
		}

		stageArtifacts[stage.ID] = artifactPaths
		stagePayloads[stage.ID] = primaryPayload
		manifest.Stages[stageIndex].Files = displayPathsForRun(runDir, writtenFiles)
		if stage.FinalOutput {
			finalOutput = primaryPayload
			result.FinalOutput = finalOutput
			finalOutputStageID = stage.ID
			manifest.FinalOutput = &FinalOutputReport{
				StageID: stage.ID,
				Bytes:   len(primaryPayload),
			}
		}
		if effectiveStageRole(stage) == StageRoleValidate && stageIndex == lastValidateStageIndex {
			validationSatisfied = true
		}
		if stage.Executor == ExecutorBuiltin && stage.Builtin != nil && stage.Builtin.Name == "write_publish_manifest" {
			publishManifestRequests = append(publishManifestRequests, publishManifestRequest{
				stageIndex: stageIndex,
				path:       builtinOutputPath(stage, runDir, "publish_manifest", "publish_manifest.json"),
			})
		}

		stageEnd := time.Now().UTC()
		manifest.Stages[stageIndex].Status = "passed"
		manifest.Stages[stageIndex].FinishedAt = &stageEnd
		runState.UpdatedAt = stageEnd
		if err := persistRunState(runDir, manifest, runState); err != nil {
			return result, r.failRun(runDir, manifest, runState, stageIndex, stageOrdinal, stageTotal, opts.CleanupDelay, opts.DisableCleanup, finalOutputStageID, finalOutput, err, publishManifestRequests, source)
		}
		r.emitJSONEvent("stage_passed", map[string]any{
			"run_id":       runID,
			"stage_id":     stage.ID,
			"stage_number": stageOrdinal,
			"stage_total":  stageTotal,
			"files":        manifest.Stages[stageIndex].Files,
		})
		if !r.jsonEvents {
			fmt.Fprintf(r.Stderr, "[%d/%d] %s ........ PASS\n", stageOrdinal, stageTotal, stage.ID)
			if len(manifest.Stages[stageIndex].Files) > 0 {
				fmt.Fprintf(r.Stderr, "           files: %s\n", strings.Join(manifest.Stages[stageIndex].Files, ", "))
			}
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
		var err error
		noValidateWarningEmitted, err = r.emitNoValidateWarning(runDir, manifest, p.Name, noValidateWarningEmitted)
		if err != nil {
			return result, err
		}
	}
	if err := writePublishManifests(runDir, manifest, source, finalOutput, publishManifestRequests); err != nil {
		return result, err
	}

	if err := persistRunState(runDir, manifest, runState); err != nil {
		return result, err
	}

	if finalOutput != "" {
		r.emitFinalOutput(finalOutput)
	}
	r.emitRunSummary(manifest, runDir)
	if !opts.DisableCleanup {
		if err := startCleanupHelper(runDir, opts.CleanupDelay); err != nil {
			return result, err
		}
	}

	return result, nil
}

func (r *Runner) executeStage(ctx context.Context, stage Stage, runtimeCtx StageRuntimeContext) (*StageExecutionResult, error) {
	if r.stageExecutionOverride != nil {
		if result, handled, err := r.stageExecutionOverride(ctx, stage, runtimeCtx); handled {
			return result, err
		}
	}

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
		payload, exists := stagePayloads[prevStage.ID]
		if !exists {
			return "", fmt.Errorf("stage %q requires previous stage %q output, but it is not available", stage.ID, prevStage.ID)
		}
		return payload, nil
	}

	switch stage.Input.From {
	case StageInputSource:
		return sourcePayload, nil
	case StageInputPrevious:
		if index == 0 {
			return sourcePayload, nil
		}
		prevStage := stages[index-1]
		payload, exists := stagePayloads[prevStage.ID]
		if !exists {
			return "", fmt.Errorf("stage %q requires previous stage %q output, but it is not available", stage.ID, prevStage.ID)
		}
		return payload, nil
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

func (r *Runner) failRun(runDir string, manifest *RunManifest, runState *RunState, stageIndex int, stageOrdinal int, stageTotal int, cleanupDelay time.Duration, disableCleanup bool, finalOutputStageID string, finalOutput string, err error, publishManifestRequests []publishManifestRequest, source RunSource) error {
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

	persistErr := writePublishManifests(runDir, manifest, source, finalOutput, publishManifestRequests)
	if stateErr := persistRunState(runDir, manifest, runState); stateErr != nil {
		persistErr = errors.Join(persistErr, stateErr)
	}
	r.emitJSONEvent("stage_failed", map[string]any{
		"run_id":       manifest.RunID,
		"stage_id":     manifest.Stages[stageIndex].ID,
		"stage_number": stageOrdinal,
		"stage_total":  stageTotal,
		"error":        err.Error(),
	})
	if !r.jsonEvents {
		fmt.Fprintf(r.Stderr, "[%d/%d] %s ........ FAIL\n", stageOrdinal, stageTotal, manifest.Stages[stageIndex].ID)
	}
	r.emitRunSummary(manifest, runDir)
	if !disableCleanup {
		if cleanupErr := startCleanupHelper(runDir, cleanupDelay); cleanupErr != nil {
			persistErr = errors.Join(persistErr, cleanupErr)
		}
	}
	return errors.Join(err, persistErr)
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

func (r *Runner) emitNoValidateWarning(runDir string, manifest *RunManifest, pipelineName string, alreadyEmitted bool) (bool, error) {
	if alreadyEmitted {
		return true, nil
	}
	warning := fmt.Sprintf("warning: pipeline %s has no validate stage", pipelineName)
	r.emitJSONEvent("warning", map[string]any{
		"run_id":  manifest.RunID,
		"message": warning,
	})
	if !r.jsonEvents {
		fmt.Fprintln(r.Stderr, warning)
	}
	manifest.Warnings = append(manifest.Warnings, warning)
	if err := persistRunManifest(runDir, manifest); err != nil {
		return true, err
	}
	return true, nil
}

func (r *Runner) emitRunSummary(manifest *RunManifest, runDir string) {
	finalBytes := 0
	finalStageID := ""
	if manifest.FinalOutput != nil {
		finalBytes = manifest.FinalOutput.Bytes
		finalStageID = manifest.FinalOutput.StageID
	}
	r.emitJSONEvent("run_summary", map[string]any{
		"run_id":      manifest.RunID,
		"status":      manifest.Status,
		"run_dir":     runDir,
		"final_stage": finalStageID,
		"final_bytes": finalBytes,
	})
	if !r.jsonEvents {
		fmt.Fprintf(r.Stderr, "run summary: status=%s run_id=%s run_dir=%s final_stage=%s final_bytes=%d\n", manifest.Status, manifest.RunID, runDir, finalStageID, finalBytes)
	}
}

func (r *Runner) emitJSONEvent(eventType string, payload map[string]any) {
	if !r.jsonEvents {
		return
	}
	event := map[string]any{
		"type": eventType,
		"time": time.Now().UTC().Format(time.RFC3339Nano),
	}
	for key, value := range payload {
		event[key] = value
	}
	encoded, err := json.Marshal(event)
	if err != nil {
		return
	}
	fmt.Fprintln(r.Stderr, string(encoded))
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

func SelectStageIndices(stages []Stage, fromStage string, toStage string, onlyStage string) ([]int, map[int]struct{}, error) {
	if len(stages) == 0 {
		return nil, nil, fmt.Errorf("pipeline has no stages")
	}
	if onlyStage != "" && (fromStage != "" || toStage != "") {
		return nil, nil, fmt.Errorf("--only-stage cannot be combined with --from-stage or --to-stage")
	}

	indexForID := make(map[string]int, len(stages))
	for i, stage := range stages {
		indexForID[stage.ID] = i
	}

	if onlyStage != "" {
		stageIndex, ok := indexForID[onlyStage]
		if !ok {
			return nil, nil, fmt.Errorf("stage %q not found", onlyStage)
		}
		set := map[int]struct{}{stageIndex: struct{}{}}
		return []int{stageIndex}, set, nil
	}

	start := 0
	end := len(stages) - 1
	if fromStage != "" {
		stageIndex, ok := indexForID[fromStage]
		if !ok {
			return nil, nil, fmt.Errorf("stage %q not found", fromStage)
		}
		start = stageIndex
	}
	if toStage != "" {
		stageIndex, ok := indexForID[toStage]
		if !ok {
			return nil, nil, fmt.Errorf("stage %q not found", toStage)
		}
		end = stageIndex
	}
	if start > end {
		return nil, nil, fmt.Errorf("--from-stage %q appears after --to-stage %q", fromStage, toStage)
	}

	indices := make([]int, 0, end-start+1)
	set := make(map[int]struct{}, end-start+1)
	for i := start; i <= end; i++ {
		indices = append(indices, i)
		set[i] = struct{}{}
	}
	return indices, set, nil
}

func stageIDsFromIndices(stages []Stage, indices []int) []string {
	ids := make([]string, 0, len(indices))
	for _, idx := range indices {
		ids = append(ids, stages[idx].ID)
	}
	return ids
}

func findLastValidateStageIndexForSelection(stages []Stage, selectedSet map[int]struct{}) int {
	last := -1
	for i, stage := range stages {
		if _, selected := selectedSet[i]; !selected {
			continue
		}
		if effectiveStageRole(stage) == StageRoleValidate {
			last = i
		}
	}
	return last
}

func persistRunManifest(runDir string, manifest *RunManifest) error {
	return writeJSON(filepath.Join(runDir, "run_manifest.json"), manifest)
}

func persistRunState(runDir string, manifest *RunManifest, runState *RunState) error {
	if err := persistRunManifest(runDir, manifest); err != nil {
		return err
	}
	return writeJSON(filepath.Join(runDir, "run.json"), runState)
}

func writePublishManifests(runDir string, manifest *RunManifest, source RunSource, finalOutput string, requests []publishManifestRequest) error {
	if len(requests) == 0 {
		return nil
	}
	filesByStage := make(map[int][]string, len(requests))
	for _, request := range requests {
		filesByStage[request.stageIndex] = append(filesByStage[request.stageIndex], request.path)
	}
	for stageIndex, files := range filesByStage {
		manifest.Stages[stageIndex].Files = displayPathsForRun(runDir, append(files, manifest.Stages[stageIndex].Files...))
	}
	for _, request := range requests {
		if err := os.MkdirAll(filepath.Dir(request.path), 0o755); err != nil {
			return fmt.Errorf("create publish manifest directory: %w", err)
		}
		payload := map[string]any{
			"run_manifest": manifest,
			"source": map[string]any{
				"mode":      source.Mode,
				"reference": source.Reference,
			},
			"final_output": finalOutput,
		}
		if err := writeJSON(request.path, payload); err != nil {
			return err
		}
	}
	return nil
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
