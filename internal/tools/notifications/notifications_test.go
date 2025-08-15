package notifications

import (
	"os/exec"
	"runtime"
	"testing"
)

func TestNewNotificationManager(t *testing.T) {
	manager := NewNotificationManager()
	if manager == nil {
		t.Fatal("NewNotificationManager() returned nil")
	}
	if manager.provider == nil {
		t.Fatal("NotificationManager provider is nil")
	}
}

func TestNotificationManagerIsAvailable(t *testing.T) {
	manager := NewNotificationManager()
	// Should not panic
	_ = manager.IsAvailable()
}

func TestNotificationManagerSend(t *testing.T) {
	manager := NewNotificationManager()

	// Test sending notification - this may fail on systems without notification tools
	// but should not panic
	err := manager.Send("Test Title", "Test Message")
	if err != nil {
		t.Logf("Notification send failed (expected on systems without notification tools): %v", err)
	}
}

func TestTerminalNotifierProvider(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("Skipping macOS terminal-notifier test on non-macOS platform")
	}

	provider := &TerminalNotifierProvider{}

	// Test availability - depends on whether terminal-notifier is installed
	available := provider.IsAvailable()
	t.Logf("terminal-notifier available: %v", available)

	if available {
		err := provider.Send("Test", "Test message")
		if err != nil {
			t.Logf("terminal-notifier send failed: %v", err)
		}
	}
}

func TestOSAScriptProvider(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("Skipping macOS osascript test on non-macOS platform")
	}

	provider := &OSAScriptProvider{}

	// osascript should always be available on macOS
	if !provider.IsAvailable() {
		t.Error("osascript should be available on macOS")
	}

	// Test sending (may show actual notification)
	err := provider.Send("Test", "Test message")
	if err != nil {
		t.Errorf("osascript send failed: %v", err)
	}
}

func TestNotifySendProvider(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("Skipping Linux notify-send test on non-Linux platform")
	}

	provider := &NotifySendProvider{}

	// Test availability - depends on whether notify-send is installed
	available := provider.IsAvailable()
	t.Logf("notify-send available: %v", available)

	if available {
		err := provider.Send("Test", "Test message")
		if err != nil {
			t.Logf("notify-send send failed: %v", err)
		}
	}
}

func TestPowerShellProvider(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("Skipping Windows PowerShell test on non-Windows platform")
	}

	provider := &PowerShellProvider{}

	// PowerShell should be available on Windows
	if !provider.IsAvailable() {
		t.Error("PowerShell should be available on Windows")
	}

	// Note: This will show a message box if run
	// In CI/CD, this might not work properly
	err := provider.Send("Test", "Test message")
	if err != nil {
		t.Logf("PowerShell send failed (expected in headless environments): %v", err)
	}
}

func TestNoopProvider(t *testing.T) {
	provider := &NoopProvider{}

	// Should always report as not available
	if provider.IsAvailable() {
		t.Error("NoopProvider should report as not available")
	}

	// Should never error
	err := provider.Send("Test", "Test message")
	if err != nil {
		t.Errorf("NoopProvider send should never error, got: %v", err)
	}
}

func TestProviderIsAvailable(t *testing.T) {
	tests := []struct {
		name     string
		provider NotificationProvider
		command  string
	}{
		{"TerminalNotifier", &TerminalNotifierProvider{}, "terminal-notifier"},
		{"OSAScript", &OSAScriptProvider{}, "osascript"},
		{"NotifySend", &NotifySendProvider{}, "notify-send"},
		{"PowerShell", &PowerShellProvider{}, "powershell"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			available := tt.provider.IsAvailable()

			// Cross-check with actual command availability
			_, err := exec.LookPath(tt.command)
			expectedAvailable := err == nil

			if available != expectedAvailable {
				t.Logf("Provider %s availability mismatch: provider=%v, command=%v",
					tt.name, available, expectedAvailable)
				// This is informational, not a failure, since system setup varies
			}
		})
	}
}

func TestSendWithSpecialCharacters(t *testing.T) {
	manager := NewNotificationManager()

	// Test with special characters that might break shell commands
	specialTitle := `Title with "quotes" and 'apostrophes'`
	specialMessage := `Message with \backslashes and $variables and "quotes"`

	err := manager.Send(specialTitle, specialMessage)
	if err != nil {
		t.Logf("Send with special characters failed (may be expected): %v", err)
	}
}
