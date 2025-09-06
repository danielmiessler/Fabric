# For Perplexity

Perplexity always provides citation output, so if you want citations in the schema, you'll need to include `citations` in your schema definition:

```json
  "citations": {
      "type": "array",
      "description": "Array of citations referenced in the article",
      "items": {
        "type": "object",
        "properties": {
          "id": {
            "type": "string",
            "description": "Unique identifier for the citation"
          },
          "title": {
            "type": "string",
            "description": "Title of the cited work"
          },
          "url": {
            "type": "string",
            "format": "uri",
            "description": "URL to the cited work (if available)"
          }
        }
     }
   }
```