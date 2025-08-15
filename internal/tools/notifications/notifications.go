package notifications

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// NotificationProvider interface for different notification backends
type NotificationProvider interface {
	Send(title, message string) error
	IsAvailable() bool
}

// NotificationManager handles cross-platform notifications
type NotificationManager struct {
	provider NotificationProvider
}

// NewNotificationManager creates a new notification manager with the best available provider
func NewNotificationManager() *NotificationManager {
	var provider NotificationProvider

	switch runtime.GOOS {
	case "darwin":
		// Try terminal-notifier first, then fall back to osascript
		provider = &TerminalNotifierProvider{}
		if !provider.IsAvailable() {
			provider = &OSAScriptProvider{}
		}
	case "linux":
		provider = &NotifySendProvider{}
	case "windows":
		provider = &PowerShellProvider{}
	default:
		provider = &NoopProvider{}
	}

	return &NotificationManager{provider: provider}
}

// Send sends a notification using the configured provider
func (nm *NotificationManager) Send(title, message string) error {
	if nm.provider == nil {
		return fmt.Errorf("no notification provider available")
	}
	return nm.provider.Send(title, message)
}

// IsAvailable checks if notifications are available
func (nm *NotificationManager) IsAvailable() bool {
	return nm.provider != nil && nm.provider.IsAvailable()
}

// macOS terminal-notifier implementation
type TerminalNotifierProvider struct{}

func (t *TerminalNotifierProvider) Send(title, message string) error {
	cmd := exec.Command("terminal-notifier", "-title", title, "-message", message, "-sound", "Glass")
	return cmd.Run()
}

func (t *TerminalNotifierProvider) IsAvailable() bool {
	_, err := exec.LookPath("terminal-notifier")
	return err == nil
}

// macOS osascript implementation
type OSAScriptProvider struct{}

func (o *OSAScriptProvider) Send(title, message string) error {
	// SECURITY: Use separate arguments instead of string interpolation to prevent AppleScript injection
	script := `display notification (system attribute "FABRIC_MESSAGE") with title (system attribute "FABRIC_TITLE") sound name "Glass"`
	cmd := exec.Command("osascript", "-e", script)

	// Set environment variables for the AppleScript to read safely
	cmd.Env = append(os.Environ(), "FABRIC_TITLE="+title, "FABRIC_MESSAGE="+message)
	return cmd.Run()
}

func (o *OSAScriptProvider) IsAvailable() bool {
	_, err := exec.LookPath("osascript")
	return err == nil
}

// Linux notify-send implementation
type NotifySendProvider struct{}

func (n *NotifySendProvider) Send(title, message string) error {
	cmd := exec.Command("notify-send", title, message)
	return cmd.Run()
}

func (n *NotifySendProvider) IsAvailable() bool {
	_, err := exec.LookPath("notify-send")
	return err == nil
}

// Windows PowerShell implementation
type PowerShellProvider struct{}

func (p *PowerShellProvider) Send(title, message string) error {
	// SECURITY: Use environment variables to avoid PowerShell injection attacks
	script := `Add-Type -AssemblyName System.Windows.Forms; [System.Windows.Forms.MessageBox]::Show($env:FABRIC_MESSAGE, $env:FABRIC_TITLE)`
	cmd := exec.Command("powershell", "-Command", script)

	// Set environment variables for PowerShell to read safely
	cmd.Env = append(os.Environ(), "FABRIC_TITLE="+title, "FABRIC_MESSAGE="+message)
	return cmd.Run()
}

func (p *PowerShellProvider) IsAvailable() bool {
	_, err := exec.LookPath("powershell")
	return err == nil
}

// NoopProvider for unsupported platforms
type NoopProvider struct{}

func (n *NoopProvider) Send(title, message string) error {
	// Silent no-op for unsupported platforms
	return nil
}

func (n *NoopProvider) IsAvailable() bool {
	return false
}
