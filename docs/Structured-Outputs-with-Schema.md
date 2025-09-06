# Structured Outputs with `--schema` Flag

Fabric now supports structured outputs using the new `--schema` flag. This feature allows you to define the expected JSON structure for the AI's response, ensuring more consistent and parseable results.

## How to Use

To enable structured outputs, use the `--schema` flag when running Fabric, providing the name of your desired JSON schema file located in the `schemas` folder of your .config directory (with or without the `.json` extension; or full path to the file).

Example:
```bash
fabric --pattern my_pattern --schema my_schema
```

### Including Schema in `system.md` Files

You can reference the JSON schema within your `system.md` pattern files using the `{{schema}}` template variable. This variable will be automatically replaced with the stringified content of the specified JSON schema, allowing the AI to understand the required output format.

Example `data/patterns/my_pattern/system.md`:
```markdown
    # IDENTITY and PURPOSE
    You are an AI assistant that generates structured data.

    # OUTPUT
    Please provide a valid JSON structured output object that conforms to the following schema:

    {{schema}}

    # EXAMPLE
    ```json
    {
    "name": "Example Item",
    "quantity": 10,
    "available": true
    }
    ```
```

### Schema Location

JSON schema files must be saved in the `Schemas` folder located within your Fabric configuration directory.

### Vendor Compatibility

The `--schema` flag has been tested and confirmed to work with the following AI vendors:

- OpenAI
- Anthropic
- Gemini
- Perplexity

Please note that its compatibility with other vendors has not yet been thoroughly tested (Bedrock, LLMStudio). Azure has not been implemented.

## Provider-Specific Behavior

### Perplexity Citations

Perplexity automatically provides web search citations with its responses. The schema plugin handles these citations intelligently:

- **JSON Outputs**: When using schemas, citations are embedded within the JSON structure:
  - If your schema includes a `citations` field (array of objects), citations are formatted to match the schema properties (e.g., `id`, `title`, `url`)
  - If your schema doesn't include citations, a `citations` field is automatically added to the output
  - Citations are formatted as objects with sequential IDs starting from 1

- **Text Outputs**: For non-JSON responses, citations are appended as markdown-formatted sources at the end of the response

Example schema with citations:
```json
{
  "type": "object",
  "properties": {
    "title": {"type": "string"},
    "summary": {"type": "string"},
    "citations": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "id": {"type": "string"},
          "title": {"type": "string"},
          "url": {"type": "string", "format": "uri"}
        }
      }
    }
  }
}
```

## Associated Patterns

Two new patterns are specifically designed to work with structured outputs and JSON schemas:

### `patterns/create_json_schema`

This pattern acts as a JSON Schema generator. Given a detailed user description of an output or goal, it generates a complete, standard-compliant JSON Schema that accurately models all relevant properties, data types, and constraints described. The output is strictly a valid JSON Schema format. When possible, it includes a name and description representing the generated JSON schema.

Example: `fabric "Create a json schema for full_name (required), company_name (optiona), address (optional) and telephone (required)" --pattern create_json_schema > ~/.config/fabric/Schemas/contact_details.json`

### `patterns/extract_data`

This pattern functions as a JSON Schema extraction engine. Its task is to read free-form user input and output a single JSON object that validates against a provided JSON Schema. It ensures the output conforms to all constraints defined in the schema (types, required, enum, format, pattern, etc.), uses defaults where values are missing, and never invents facts, only extracting what is present in the input.

Example: `cat <source doc> | fabric -p extract_data --schema contact_details > my_contact.json`
