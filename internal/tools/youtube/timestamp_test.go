package youtube

import (
	"testing"
)

func TestParseTimestampToSeconds(t *testing.T) {
	tests := []struct {
		timestamp string
		expected  int
		shouldErr bool
	}{
		{"00:30", 30, false},
		{"01:30", 90, false},
		{"01:05:30", 3930, false}, // 1 hour 5 minutes 30 seconds
		{"10:00", 600, false},
		{"invalid", 0, true},
		{"1:2:3:4", 0, true}, // too many parts
	}

	for _, test := range tests {
		result, err := parseTimestampToSeconds(test.timestamp)

		if test.shouldErr {
			if err == nil {
				t.Errorf("Expected error for timestamp %s, but got none", test.timestamp)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for timestamp %s: %v", test.timestamp, err)
			}
			if result != test.expected {
				t.Errorf("For timestamp %s, expected %d seconds, got %d", test.timestamp, test.expected, result)
			}
		}
	}
}

func TestShouldIncludeRepeat(t *testing.T) {
	tests := []struct {
		lastTimestamp    string
		currentTimestamp string
		expected         bool
		description      string
	}{
		{"00:30", "01:30", true, "60 second gap should allow repeat"},
		{"00:30", "00:45", true, "15 second gap should allow repeat"},
		{"01:00", "01:10", true, "10 second gap should allow repeat (boundary case)"},
		{"01:00", "01:09", false, "9 second gap should not allow repeat"},
		{"00:30", "00:35", false, "5 second gap should not allow repeat"},
		{"invalid", "01:30", true, "invalid timestamp should err on side of inclusion"},
		{"01:30", "invalid", true, "invalid timestamp should err on side of inclusion"},
	}

	for _, test := range tests {
		result := shouldIncludeRepeat(test.lastTimestamp, test.currentTimestamp)
		if result != test.expected {
			t.Errorf("%s: expected %v, got %v", test.description, test.expected, result)
		}
	}
}

func TestParseFFmpegFrameTimes(t *testing.T) {
	logOutput := `
[Parsed_showinfo_1 @ 0x1] n:   0 pts:      0 pts_time:0
[Parsed_showinfo_1 @ 0x1] n:   1 pts:    500 pts_time:0.5
[Parsed_showinfo_1 @ 0x1] n:   2 pts:   2500 pts_time:2.5
`

	frameTimes := parseFFmpegFrameTimes(logOutput)
	if len(frameTimes) != 3 {
		t.Fatalf("expected 3 frame times, got %d", len(frameTimes))
	}
	if frameTimes[0] != 0 || frameTimes[1] != 0.5 || frameTimes[2] != 2.5 {
		t.Fatalf("unexpected frame times: %#v", frameTimes)
	}
}

func TestVisualCueRange_UsesActualFrameTimes(t *testing.T) {
	frameTimes := []float64{0, 0.5, 2.5}

	start0, end0 := visualCueRange(frameTimes, 0)
	if start0 != "00:00:00.000" || end0 != "00:00:00.499" {
		t.Fatalf("unexpected first cue range: %s --> %s", start0, end0)
	}

	start1, end1 := visualCueRange(frameTimes, 1)
	if start1 != "00:00:00.500" || end1 != "00:00:02.499" {
		t.Fatalf("unexpected second cue range: %s --> %s", start1, end1)
	}
}
