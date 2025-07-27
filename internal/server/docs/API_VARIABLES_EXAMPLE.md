# REST API Pattern Variables Example

This example demonstrates how to use pattern variables in REST API calls to the `/chat` endpoint.

## Example: Using the `translate` pattern with variables

### Request

```json
{
  "prompts": [
    {
      "userInput": "Hello my name is Kayvan",
      "patternName": "translate",
      "model": "gpt-4o",
      "vendor": "openai",
      "contextName": "",
      "strategyName": "",
      "variables": {
        "lang_code": "fr"
      }
    }
  ],
  "language": "en",
  "temperature": 0.7,
  "topP": 0.9,
  "frequencyPenalty": 0.0,
  "presencePenalty": 0.0
}
```

### Pattern Content

The `translate` pattern contains:

```markdown
You are an expert translator... translate them as accurately and perfectly as possible into the language specified by its language code {{lang_code}}...

...

- Translate the document as accurately as possible keeping a 1:1 copy of the original text translated to {{lang_code}}.

{{input}}
```

### How it works

1. The pattern is loaded from `patterns/translate/system.md`
2. The `{{lang_code}}` variable is replaced with `"fr"` from the variables map
3. The `{{input}}` placeholder is replaced with `"Hello my name is Kayvan"`
4. The resulting processed pattern is sent to the AI model

### Expected Result

The AI would receive a prompt asking it to translate "Hello my name is Kayvan" to French (fr), and would respond with something like "Bonjour, je m'appelle Kayvan".

## Testing with curl

```bash
curl -X POST http://localhost:8080/api/chat \
  -H "Content-Type: application/json" \
  -d '{
    "prompts": [
      {
        "userInput": "Hello my name is Kayvan",
        "patternName": "translate",
        "model": "gpt-4o",
        "vendor": "openai",
        "variables": {
          "lang_code": "fr"
        }
      }
    ],
    "temperature": 0.7
  }'
```

## Multiple Variables Example

For patterns that use multiple variables:

```json
{
  "prompts": [
    {
      "userInput": "Analyze this business model",
      "patternName": "custom_analysis",
      "model": "gpt-4o",
      "variables": {
        "role": "expert consultant",
        "experience": "15",
        "focus_areas": "revenue, scalability, market fit",
        "output_format": "bullet points"
      }
    }
  ]
}
```

## Implementation Details

- Variables are passed in the `variables` field as a key-value map
- Variables are processed using Go's template system
- The `{{input}}` variable is automatically handled and should not be included in the variables map
- Variables support the same features as CLI variables (plugins, extensions, etc.)
