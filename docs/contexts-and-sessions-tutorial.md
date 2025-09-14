# Contexts and Sessions in Fabric

Fabric uses **contexts** and **sessions** to manage conversation state and reusable prompt data. This guide focuses on how to use them from the CLI and REST API.

## What is a Context?

A context is named text that Fabric injects at the beginning of a conversation. Contexts live on disk under `~/.config/fabric/contexts`; each file name is the context name, and its contents are included as a system message.

Command-line helpers:

- `--context <name>` select a context
- `--listcontexts` list available contexts
- `--printcontext <name>` show the contents
- `--wipecontext <name>` delete it

## What is a Session?

A session tracks the message history of a conversation. When you specify a session name, Fabric loads any existing messages, appends new ones, and saves back to disk. Sessions are stored as JSON under `~/.config/fabric/sessions`.

Command-line helpers:

- `--session <name>` attach to a session
- `--listsessions` list stored sessions
- `--printsession <name>` print a session
- `--wipesession <name>` delete it

## Everyday Use Cases

Contexts and sessions serve different everyday needs:

- **Context** – Reuse prompt text such as preferred style, domain knowledge, or instructions for the assistant.
- **Session** – Maintain ongoing conversation history so Fabric remembers earlier exchanges.

Example workflow:

1. Create a context file manually in `~/.config/fabric/contexts/writer` with your writing guidelines.
2. Start a session while chatting to build on previous answers (`fabric --session mychat`). Sessions are automatically created if they don't exist.

## How Contexts and Sessions Interact

When Fabric handles a chat request, it loads any named context, combines it with pattern text, and adds the result as a system message before sending the conversation history to the model. The assistant's reply is appended to the session so future calls continue from the same state.

## REST API Endpoints

The REST server exposes CRUD endpoints for managing contexts and sessions:

- `/contexts/:name` – get or save a context
- `/contexts/names` – list available contexts
- `/sessions/:name` – get or save a session
- `/sessions/names` – list available sessions

## Summary

Contexts provide reusable system-level instructions, while sessions maintain conversation history. Together they allow Fabric to build rich, stateful interactions with language models.

## For Developers

### Loading Contexts from Disk

```go
// internal/plugins/db/fsdb/contexts.go
func (o *ContextsEntity) Get(name string) (*Context, error) {
    content, err := o.Load(name)
    if err != nil {
        return nil, err
    }
    return &Context{Name: name, Content: string(content)}, nil
}
```

### Handling Sessions

```go
// internal/plugins/db/fsdb/sessions.go
type Session struct {
    Name     string
    Messages []*chat.ChatCompletionMessage
}

func (o *SessionsEntity) Get(name string) (*Session, error) {
    session := &Session{Name: name}
    if o.Exists(name) {
        err = o.LoadAsJson(name, &session.Messages)
    } else {
        fmt.Printf("Creating new session: %s\n", name)
    }
    return session, err
}
```

### Building a Session

```go
// internal/core/chatter.go
if request.ContextName != "" {
    ctx, err := o.db.Contexts.Get(request.ContextName)
    if err != nil {
        return nil, fmt.Errorf("could not find context %s: %v", request.ContextName, err)
    }
    contextContent = ctx.Content
}

systemMessage := strings.TrimSpace(contextContent) + strings.TrimSpace(patternContent)
if systemMessage != "" {
    session.Append(&chat.ChatCompletionMessage{Role: chat.ChatMessageRoleSystem, Content: systemMessage})
}
```
