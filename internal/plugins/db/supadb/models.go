package supadb

import (
	"time"

	"github.com/google/uuid"
)

// Session represents a conversational session stored in Supabase.
type Session struct {
	ID          uuid.UUID      `json:"id"`
	UserID      *uuid.UUID     `json:"user_id"`
	Title       string         `json:"title"`
	Description *string        `json:"description"`
	Metadata    map[string]any `json:"metadata"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

// Message represents a message exchanged within a session.
type Message struct {
	ID        uuid.UUID      `json:"id"`
	SessionID uuid.UUID      `json:"session_id"`
	Role      string         `json:"role"`
	Content   string         `json:"content"`
	Metadata  map[string]any `json:"metadata"`
	CreatedAt time.Time      `json:"created_at"`
}

// Pattern persists prompt templates.
type Pattern struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Body        string    `json:"body"`
	Tags        []string  `json:"tags"`
	IsSystem    bool      `json:"is_system"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Note captures knowledge snippets from a session.
type Note struct {
	ID        uuid.UUID  `json:"id"`
	SessionID *uuid.UUID `json:"session_id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	Tags      []string   `json:"tags"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
