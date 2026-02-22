package log

import "testing"

func TestLevelFromInt(t *testing.T) {
	tests := []struct {
		in   int
		want Level
	}{
		{in: -1, want: Off},
		{in: 0, want: Off},
		{in: 1, want: Basic},
		{in: 2, want: Detailed},
		{in: 3, want: Trace},
		{in: 4, want: Wire},
		{in: 9, want: Wire},
	}

	for _, tc := range tests {
		if got := LevelFromInt(tc.in); got != tc.want {
			t.Fatalf("LevelFromInt(%d) = %v, want %v", tc.in, got, tc.want)
		}
	}
}
