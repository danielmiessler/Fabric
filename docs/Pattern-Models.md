# Pattern Model Mapping

Fabric can automatically select a model for a pattern based on a user-defined mapping file.

## Configuration

Create `~/.config/fabric/pattern_models.yaml` with entries mapping pattern names to models:

```yaml
summarize: openai/gpt-4o-mini
ai: anthropic/claude-3-opus
```

The key is the pattern name. The value is the model identifier. If the value includes a vendor prefix (e.g. `openai/`), Fabric sets both the vendor and model accordingly.

When you run a pattern without specifying `--model`, Fabric consults this file to determine the model to use.

## Helper Command

You can manage this mapping from the command line:

```bash
fabric pattern-model list
fabric pattern-model set <pattern> <model>
fabric pattern-model unset <pattern>
```

`set` updates the mapping file, creating it if necessary. `unset` removes a mapping. `list` shows all current mappings.
