# Schema Plugin Architecture

This package provides a centralized, interface-based approach to handling structured outputs across different AI providers in Fabric.

## Phase 1: Core Infrastructure (COMPLETE)

### What's Been Built

1. **Core Interfaces** (`schema.go`)
   - `SchemaPlugin`: Main interface for schema handling
   - `ResponseParser`: Interface for provider-specific response parsing
   - `ProviderRequirements`: Configuration for each provider's needs

2. **Schema Manager** (`manager.go`)
   - Orchestrates schema transformation and response parsing
   - Provides validation capabilities
   - Central integration point for CLI

3. **Default Implementation** (`implementation.go`)
   - Concrete implementation of `SchemaPlugin`
   - Provider-specific schema transformations:
     - **Anthropic**: Tool-based approach with `get_structured_output`
     - **OpenAI**: Response format with `json_schema`
     - **Gemini**: Response schema in generation config
     - **Ollama**: Format parameter with JSON schema
   - Schema validation using `gojsonschema`

4. **Provider Parsers** (`parsers.go`)
   - Placeholder parsers for each provider
   - Ready for Phase 2 implementation
   - Basic structure for handling different response formats

5. **Tests** (`manager_test.go`)
   - Unit tests for core functionality
   - Schema validation tests
   - Provider support verification

### Architecture Benefits

- **Centralized Logic**: All schema handling in one place
- **Provider Agnostic**: Easy to add new providers
- **Minimal Provider Changes**: Providers only need `GetProviderName()` method
- **Extensible**: Plugin-based architecture allows customization
- **Testable**: Each component can be tested independently

## Next Steps (Phase 2)

1. **Minimal Provider Changes**
   - Add `GetProviderName()` to each provider
   - Add `SendRaw()` and `SendStreamRaw()` methods
   - Keep existing `HandleSchema()` methods temporarily

2. **Parser Implementation**
   - Implement actual parsing logic for Anthropic and OpenAI
   - Handle real response types from each provider
   - Implement streaming support

3. **CLI Integration**
   - Integrate schema manager into CLI flow
   - Replace direct provider `HandleSchema()` calls
   - Add comprehensive error handling

## Usage Example (Future)

```go
// In CLI
schemaManager := schema.NewManager()

// Before sending to provider
err := schemaManager.HandleSchemaTransformation("anthropic", opts)
if err != nil {
    return err
}

// After receiving from provider
parsedResponse, err := schemaManager.HandleResponseParsing("anthropic", rawResponse, opts)
if err != nil {
    return err
}

// Validate final output
err = schemaManager.ValidateOutput(parsedResponse, opts.SchemaContent)
if err != nil {
    return err
}
```

## Files Overview

- `schema.go` - Core interfaces and types
- `manager.go` - Schema manager orchestration
- `implementation.go` - Default plugin implementation
- `parsers.go` - Provider-specific response parsers (stubs)
- `manager_test.go` - Unit tests
- `README.md` - This documentation