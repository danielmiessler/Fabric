package supadb

import (
	"context"

	"github.com/google/uuid"
	postgrest "github.com/supabase-community/postgrest-go"
	supabase "github.com/supabase-community/supabase-go"
)

type SessionRepository struct {
	client *supabase.Client
}

func (r *SessionRepository) List(ctx context.Context, limit uint) ([]Session, error) {
	_ = ctx
	var result []Session
	query := r.client.From("sessions").Select("*", "", false).Order("updated_at", &postgrest.OrderOpts{Ascending: false})
	if limit > 0 {
		query = query.Limit(int(limit), "")
	}
	_, err := query.ExecuteTo(&result)
	return result, err
}

func (r *SessionRepository) Get(ctx context.Context, id uuid.UUID) (*Session, error) {
	_ = ctx
	var result []Session
	_, err := r.client.From("sessions").Select("*", "", false).Eq("id", id.String()).Limit(1, "").ExecuteTo(&result)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	return &result[0], nil
}

func (r *SessionRepository) Insert(ctx context.Context, payload map[string]any) (*Session, error) {
	_ = ctx
	var result []Session
	_, err := r.client.From("sessions").Insert(payload, false, "", "representation", "").ExecuteTo(&result)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	return &result[0], nil
}

func (r *SessionRepository) Update(ctx context.Context, id uuid.UUID, payload map[string]any) (*Session, error) {
	_ = ctx
	var result []Session
	_, err := r.client.From("sessions").Update(payload, "representation", "").Eq("id", id.String()).Limit(1, "").ExecuteTo(&result)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	return &result[0], nil
}

type PatternRepository struct {
	client *supabase.Client
}

func (r *PatternRepository) List(ctx context.Context) ([]Pattern, error) {
	_ = ctx
	var result []Pattern
	_, err := r.client.From("patterns").Select("*", "", false).Order("name", &postgrest.OrderOpts{Ascending: true}).ExecuteTo(&result)
	return result, err
}

func (r *PatternRepository) GetByID(ctx context.Context, id uuid.UUID) (*Pattern, error) {
	_ = ctx
	var result []Pattern
	_, err := r.client.From("patterns").Select("*", "", false).Eq("id", id.String()).Limit(1, "").ExecuteTo(&result)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	return &result[0], nil
}

func (r *PatternRepository) Create(ctx context.Context, payload map[string]any) (*Pattern, error) {
	_ = ctx
	var result []Pattern
	_, err := r.client.From("patterns").Insert(payload, false, "", "representation", "").ExecuteTo(&result)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	return &result[0], nil
}

func (r *PatternRepository) Upsert(ctx context.Context, payload map[string]any) (*Pattern, error) {
	_ = ctx
	var result []Pattern
	_, err := r.client.From("patterns").Upsert(payload, "name", "representation", "").ExecuteTo(&result)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	return &result[0], nil
}

func (r *PatternRepository) DeleteByName(ctx context.Context, name string) error {
	_ = ctx
	_, _, err := r.client.From("patterns").Delete("", "").Eq("name", name).Execute()
	return err
}

func (r *PatternRepository) UpdateByID(ctx context.Context, id uuid.UUID, payload map[string]any) (*Pattern, error) {
	_ = ctx
	var result []Pattern
	_, err := r.client.From("patterns").Update(payload, "representation", "").Eq("id", id.String()).Limit(1, "").ExecuteTo(&result)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	return &result[0], nil
}

func (r *PatternRepository) DeleteByID(ctx context.Context, id uuid.UUID) error {
	_ = ctx
	_, _, err := r.client.From("patterns").Delete("", "").Eq("id", id.String()).Execute()
	return err
}

type NoteRepository struct {
	client *supabase.Client
}

func (r *NoteRepository) ListBySession(ctx context.Context, sessionID uuid.UUID) ([]Note, error) {
	_ = ctx
	var result []Note
	query := r.client.From("notes").Select("*", "", false).Order("updated_at", &postgrest.OrderOpts{Ascending: false})
	if sessionID != uuid.Nil {
		query = query.Eq("session_id", sessionID.String())
	}
	_, err := query.ExecuteTo(&result)
	return result, err
}

func (r *NoteRepository) Upsert(ctx context.Context, payload map[string]any) (*Note, error) {
	_ = ctx
	var result []Note
	_, err := r.client.From("notes").Upsert(payload, "id", "representation", "").ExecuteTo(&result)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	return &result[0], nil
}

func (r *NoteRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_ = ctx
	_, _, err := r.client.From("notes").Delete("", "").Eq("id", id.String()).Execute()
	return err
}
