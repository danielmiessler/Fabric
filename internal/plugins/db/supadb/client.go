package supadb

import (
	"context"
	"fmt"
	"os"

	supabase "github.com/supabase-community/supabase-go"
)

// Client wraps the Supabase SDK to expose typed helpers for Fabric
type Client struct {
	client *supabase.Client
}

// NewClientFromEnv instantiates the Supabase client when credentials are present.
func NewClientFromEnv() (*Client, error) {
	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_SERVICE_ROLE_KEY")
	if url == "" || key == "" {
		return nil, fmt.Errorf("supabase credentials missing: SUPABASE_URL or SUPABASE_SERVICE_ROLE_KEY not set")
	}

	client, err := supabase.NewClient(url, key, nil)
	if err != nil {
		return nil, err
	}

	return &Client{client: client}, nil
}

// Ping verifies the Supabase connection using a lightweight RPC.
func (c *Client) Ping(ctx context.Context) error {
	_ = ctx
	if c == nil || c.client == nil {
		return fmt.Errorf("supabase client not initialized")
	}

	// Perform a simple fetch with a limit of 1 row against sessions table.
	_, err := c.client.From("sessions").Select("id", "", false).Limit(1, "").ExecuteTo(&[]Session{})
	return err
}

// Sessions returns a typed repository for session CRUD operations.
func (c *Client) Sessions() *SessionRepository {
	return &SessionRepository{client: c.client}
}

// Patterns returns a typed repository for pattern CRUD operations.
func (c *Client) Patterns() *PatternRepository {
	return &PatternRepository{client: c.client}
}

// Notes returns a typed repository for note CRUD operations.
func (c *Client) Notes() *NoteRepository {
	return &NoteRepository{client: c.client}
}
