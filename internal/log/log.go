package log

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// Level represents the debug verbosity.
type Level int

const (
	// Off disables all debug output.
	Off Level = iota
	// Basic provides minimal debugging information.
	Basic
	// Detailed provides more verbose debugging.
	Detailed
	// Trace is the most verbose level.
	Trace
)

var (
	mu     sync.RWMutex
	level  Level     = Off
	output io.Writer = os.Stderr
)

// SetLevel sets the global debug level.
func SetLevel(l Level) {
	mu.Lock()
	level = l
	mu.Unlock()
}

// LevelFromInt converts an int to a Level.
func LevelFromInt(i int) Level {
	switch {
	case i <= 0:
		return Off
	case i == 1:
		return Basic
	case i == 2:
		return Detailed
	case i >= 3:
		return Trace
	default:
		return Off
	}
}

// Debug writes a debug message if the global level permits.
func Debug(l Level, format string, a ...interface{}) {
	mu.RLock()
	current := level
	w := output
	mu.RUnlock()
	if current >= l {
		fmt.Fprintf(w, "DEBUG: "+format, a...)
	}
}

// Log writes a message unconditionally to stderr.
// This is for important messages that should always be shown regardless of debug level.
func Log(format string, a ...interface{}) {
	mu.RLock()
	w := output
	mu.RUnlock()
	fmt.Fprintf(w, format, a...)
}

// SetOutput allows overriding the output destination for debug logs.
func SetOutput(w io.Writer) {
	mu.Lock()
	output = w
	mu.Unlock()
}
