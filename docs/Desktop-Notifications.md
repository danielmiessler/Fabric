# Desktop Notifications

Fabric supports desktop notifications to alert you when commands complete, which is especially useful for long-running tasks or when you're multitasking.

## Quick Start

Enable notifications with the `--notification` flag:

```bash
fabric --pattern summarize --notification < article.txt
```

## Configuration

### Command Line Options

- `--notification`: Enable desktop notifications when command completes
- `--notification-command`: Use a custom notification command instead of built-in notifications

### YAML Configuration

Add notification settings to your `~/.config/fabric/config.yaml`:

```yaml
# Enable notifications by default
notification: true

# Optional: Custom notification command
notificationCommand: 'notify-send --urgency=normal "$1" "$2"'
```

## Platform Support

### macOS

- **Default**: Uses `osascript` (built into macOS)
- **Enhanced**: Install `terminal-notifier` for better notifications:

  ```bash
  brew install terminal-notifier
  ```

### Linux

- **Requirement**: Install `notify-send`:

  ```bash
  # Ubuntu/Debian
  sudo apt install libnotify-bin

  # Fedora
  sudo dnf install libnotify
  ```

### Windows

- **Default**: Uses PowerShell message boxes (built-in)

## Custom Notification Commands

The `--notification-command` flag allows you to use custom notification scripts or commands. The command receives the title as `$1` and message as `$2` as shell positional arguments.

**Security Note**: The title and message content are properly escaped to prevent command injection attacks from AI-generated output containing shell metacharacters.

### Examples

**macOS with custom sound:**

```bash
fabric --pattern analyze_claims --notification-command 'osascript -e "display notification \"$2\" with title \"$1\" sound name \"Ping\""' < document.txt
```

**Linux with urgency levels:**

```bash
fabric --pattern extract_wisdom --notification-command 'notify-send --urgency=critical "$1" "$2"' < video-transcript.txt
```

**Custom script:**

```bash
fabric --pattern summarize --notification-command '/path/to/my-notification-script.sh "$1" "$2"' < report.pdf
```

**Testing your custom command:**

```bash
# Test that $1 and $2 are passed correctly
fabric --pattern raw_query --notification-command 'echo "Title: $1, Message: $2"' "test input"
```

## Notification Content

Notifications include:

- **Title**: "Fabric Command Complete" or "Fabric: [pattern] Complete"
- **Message**: Brief summary of the output (first 100 characters)

For long outputs, the message is truncated with "..." to fit notification display limits.

## Use Cases

### Long-Running Tasks

```bash
# Process large document with notifications
fabric --pattern analyze_paper --notification < research-paper.pdf

# Extract wisdom from long video with alerts
fabric -y "https://youtube.com/watch?v=..." --pattern extract_wisdom --notification
```

### Background Processing

```bash
# Process multiple files and get notified when each completes
for file in *.txt; do
    fabric --pattern summarize --notification < "$file" &
done
```

### Integration with Other Tools

```bash
# Combine with other commands
curl -s "https://api.example.com/data" | \
    fabric --pattern analyze_data --notification --output results.md
```

## Troubleshooting

### No Notifications Appearing

1. **Check system notifications are enabled** for Terminal/your shell
2. **Verify notification tools are installed**:
   - macOS: `which osascript` (should exist)
   - Linux: `which notify-send`
   - Windows: `where.exe powershell`

3. **Test with simple command**:

   ```bash
   echo "test" | fabric --pattern raw_query --notification --dry-run
   ```

### Notification Permission Issues

On some systems, you may need to grant notification permissions to your terminal application:

- **macOS**: System Preferences → Security & Privacy → Privacy → Notifications → Enable for Terminal
- **Linux**: Depends on desktop environment; usually automatic
- **Windows**: Usually works by default

### Custom Commands Not Working

- Ensure your custom notification command is executable
- Test the command manually with sample arguments
- Check that all required dependencies are installed

## Advanced Configuration

### Environment-Specific Settings

Create different configuration files for different environments:

```bash
# Work computer (quieter notifications)
fabric --config ~/.config/fabric/work-config.yaml --notification

# Personal computer (with sound)
fabric --config ~/.config/fabric/personal-config.yaml --notification
```

### Integration with Task Management

```bash
# Custom script that also logs to task management system
notificationCommand: '/usr/local/bin/fabric-notify-and-log.sh "$1" "$2"'
```

## Examples

See `docs/notification-config.yaml` for a complete configuration example with various notification command options.
