package template

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestDateTimePlugin(t *testing.T) {
	plugin := &DateTimePlugin{}
	now := time.Now()

	tests := []struct {
		name      string
		operation string
		value     string
		validate  func(string) error
		wantErr   bool
	}{
		{
			name:      "now returns RFC3339",
			operation: "now",
			validate: func(got string) error {
				if _, err := time.Parse(time.RFC3339, got); err != nil {
					return err
				}
				return nil
			},
		},
		{
			name:      "time returns HH:MM:SS",
			operation: "time",
			validate: func(got string) error {
				if _, err := time.Parse("15:04:05", got); err != nil {
					return err
				}
				return nil
			},
		},
		{
			name:      "unix returns timestamp",
			operation: "unix",
			validate: func(got string) error {
				if _, err := strconv.ParseInt(got, 10, 64); err != nil {
					return err
				}
				return nil
			},
		},
		{
			name:      "today returns YYYY-MM-DD",
			operation: "today",
			validate: func(got string) error {
				if _, err := time.Parse("2006-01-02", got); err != nil {
					return err
				}
				return nil
			},
		},
		{
			name:      "full returns long date",
			operation: "full",
			validate: func(got string) error {
				if !strings.Contains(got, now.Month().String()) {
					return fmt.Errorf("full date missing month name")
				}
				return nil
			},
		},
		{
			name:      "relative positive hours",
			operation: "rel",
			value:     "2h",
			validate: func(got string) error {
				t, err := time.Parse(time.RFC3339, got)
				if err != nil {
					return err
				}
				expected := now.Add(2 * time.Hour)
				if t.Hour() != expected.Hour() {
					return fmt.Errorf("expected hour %d, got %d", expected.Hour(), t.Hour())
				}
				return nil
			},
		},
		{
			name:      "relative negative days",
			operation: "rel",
			value:     "-2d",
			validate: func(got string) error {
				t, err := time.Parse("2006-01-02", got)
				if err != nil {
					return err
				}
				expected := now.AddDate(0, 0, -2)
				if t.Day() != expected.Day() {
					return fmt.Errorf("expected day %d, got %d", expected.Day(), t.Day())
				}
				return nil
			},
		},
		// Error cases
		{
			name:      "invalid operation",
			operation: "invalid",
			wantErr:   true,
		},
		{
			name:      "empty relative value",
			operation: "rel",
			value:     "",
			wantErr:   true,
		},
		{
			name:      "invalid relative format",
			operation: "rel",
			value:     "2x",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := plugin.Apply(tt.operation, tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("DateTimePlugin.Apply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && tt.validate != nil {
				if err := tt.validate(got); err != nil {
					t.Errorf("DateTimePlugin.Apply() validation failed: %v", err)
				}
			}
		})
	}
}
