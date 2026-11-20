# Explain an HTTP client error

Remove credentials, cookies, and personal data before passing a response to your configured model.

~~~sh
printf 'HTTP/2 429\nRetry-After: 120\n' | fabric -p explain_http_error
~~~

## Example checks

| Input | Expected behavior |
| --- | --- |
| The example above | Identify rate limiting and interpret 120 as seconds. |
| HTTP 422 with a body saying the quantity field must be positive | Identify the field and suggest correcting its value. |
| HTTP 403 with an empty body | Acknowledge the unknown reason and request one relevant piece of evidence. |
| HTTP 200 with an unrelated error message | Ask for the failing response. |
