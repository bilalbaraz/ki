# Troubleshooting Guide

This guide covers common issues you might encounter when using or developing KI - Dragon Ball CLI.

## Table of Contents

- [Installation Issues](#installation-issues)
- [Runtime Issues](#runtime-issues)
- [API Issues](#api-issues)
- [UI/Display Issues](#ui-display-issues)
- [Performance Issues](#performance-issues)
- [Development Issues](#development-issues)

## Installation Issues

### Go Version Compatibility

**Problem**: Build fails with version errors

```
go: module requires Go 1.25 or later
```

**Solution**: 
```bash
# Check your Go version
go version

# Install or update Go from https://go.dev/dl/
# Or use a version manager like gvm
```

### Missing Dependencies

**Problem**: Missing module errors

```
module github.com/charmbracelet/bubbletea: no matching versions
```

**Solution**:
```bash
# Clean and reinstall dependencies
go clean -modcache
go mod download
go mod tidy
```

### Build Errors

**Problem**: Compilation fails

```
# error: package github.com/bilalbaraz/ki/internal/tui is not in GOROOT
```

**Solution**:
```bash
# Ensure you're in the project directory
cd ki

# Clean build cache
go clean -cache

# Rebuild
go build -o ki
```

## Runtime Issues

### Application Won't Start

**Problem**: Binary executes but TUI doesn't appear

**Solution**:
```bash
# Check if binary has execute permissions
chmod +x ki

# Run with verbose output
./ki tui 2>&1 | tee error.log

# Check for conflicting processes
ps aux | grep ki
```

### Immediate Crash on Launch

**Problem**: Application exits immediately

**Possible Causes**:
1. Terminal not supported
2. Missing terminal capabilities
3. Insufficient permissions

**Solution**:
```bash
# Try with different TERM value
TERM=xterm-256color ./ki tui

# Check terminal capabilities
echo $TERM
infocmp $TERM

# Run in a different terminal emulator
# Try: iTerm2, Alacritty, Windows Terminal, GNOME Terminal
```

### Keyboard Input Not Working

**Problem**: Keys don't respond

**Solution**:
```bash
# Verify terminal settings
stty -a

# Reset terminal
reset

# Check for key binding conflicts
# Some terminals may capture certain key combinations
```

## API Issues

### "Failed to fetch characters" Error

**Problem**: Cannot load data from API

```
❌ Error loading characters:
failed to fetch characters: Get "https://dragonball-api.com/api/characters": dial tcp: lookup dragonball-api.com: no such host
```

**Possible Causes**:
1. No internet connection
2. DNS resolution failure
3. Firewall blocking requests
4. API is down

**Solutions**:

1. **Check internet connection**:
```bash
ping google.com
ping dragonball-api.com
```

2. **Check DNS**:
```bash
nslookup dragonball-api.com
dig dragonball-api.com
```

3. **Test API directly**:
```bash
curl https://dragonball-api.com/api/characters
```

4. **Check firewall**:
```bash
# macOS
sudo /usr/libexec/ApplicationFirewall/socketfilterfw --listapps

# Linux (iptables)
sudo iptables -L

# Windows
netsh advfirewall show allprofiles
```

5. **Use proxy if needed**:
```bash
export HTTP_PROXY=http://proxy.example.com:8080
export HTTPS_PROXY=https://proxy.example.com:8080
./ki tui
```

### Timeout Errors

**Problem**: Requests time out

```
❌ Error loading characters:
context deadline exceeded
```

**Solution**:

The default timeout is 10 seconds. If your connection is slow:

1. **Check connection speed**:
```bash
curl -o /dev/null -s -w "Time: %{time_total}s\n" https://dragonball-api.com/api/characters
```

2. **Increase timeout** (requires code modification):

Edit `internal/api/client.go`:
```go
httpClient: &http.Client{
    Timeout: 30 * time.Second,  // Increased from 10s
},
```

### Rate Limiting

**Problem**: API returns 429 Too Many Requests

**Solution**:
- Wait a few minutes before retrying
- The Dragon Ball API has rate limits
- Consider implementing request caching (future feature)

### SSL/TLS Errors

**Problem**: Certificate verification fails

```
x509: certificate signed by unknown authority
```

**Solution**:
```bash
# Update system certificates (not recommended to disable verification)

# macOS
brew install ca-certificates

# Linux (Ubuntu/Debian)
sudo apt-get install ca-certificates
sudo update-ca-certificates

# Linux (RHEL/CentOS)
sudo yum install ca-certificates
sudo update-ca-trust
```

## UI/Display Issues

### Garbled Text or Weird Characters

**Problem**: UI displays incorrectly with strange symbols

**Cause**: Terminal doesn't support Unicode or colors properly

**Solution**:
```bash
# Set proper locale
export LANG=en_US.UTF-8
export LC_ALL=en_US.UTF-8

# Use terminal with UTF-8 support
# Recommended: iTerm2, Alacritty, Windows Terminal

# Check terminal encoding
locale
```

### No Colors

**Problem**: UI is black and white

**Solution**:
```bash
# Check color support
echo $TERM

# Set to 256 color terminal
export TERM=xterm-256color

# Test colors
for i in {0..255}; do printf "\x1b[38;5;${i}mcolor${i} "; done; echo
```

### Layout Issues / Text Overflow

**Problem**: Content doesn't fit properly

**Solution**:
1. Resize terminal to at least 80x24
2. Use fullscreen mode
3. Verify terminal reports size correctly:
```bash
echo "Columns: $COLUMNS, Rows: $LINES"
stty size
```

### Emoji Not Displaying

**Problem**: 🐉 appears as boxes or question marks

**Solution**:
- Use a font that supports emoji (Noto Color Emoji, Apple Color Emoji)
- Update your terminal emulator
- On Linux, install `fonts-noto-color-emoji`

## Performance Issues

### Slow Startup

**Problem**: Application takes long to start

**Possible Causes**:
1. Slow API response
2. DNS resolution delay
3. Large data set

**Solution**:
```bash
# Check DNS resolution time
time nslookup dragonball-api.com

# Use faster DNS (Google DNS)
# Add to /etc/resolv.conf or network settings
nameserver 8.8.8.8
nameserver 8.8.4.4
```

### Sluggish UI

**Problem**: UI feels laggy or unresponsive

**Solution**:
1. Reduce items per page (modify `itemsPerPage` in model.go)
2. Use a faster terminal emulator
3. Check system resources:
```bash
# Check CPU/memory usage
top
htop

# Check if system is under load
uptime
```

### Memory Usage

**Problem**: High memory consumption

**Solution**:
```bash
# Monitor memory usage
ps aux | grep ki

# Profile memory (during development)
go build -o ki
GODEBUG=gctrace=1 ./ki tui
```

## Development Issues

### Changes Not Reflected

**Problem**: Code changes don't appear when running

**Solution**:
```bash
# Ensure you rebuild after changes
go build -o ki

# Or use go run
go run main.go tui

# Clean build cache if needed
go clean -cache
go build -o ki
```

### Test Failures

**Problem**: Tests fail unexpectedly

```
FAIL: TestGetCharacters (5.00s)
    client_test.go:25: timeout waiting for response
```

**Solution**:
```bash
# Increase test timeout
go test -timeout 30s ./...

# Run specific test
go test -v -run TestGetCharacters ./internal/api/

# Skip network tests if API is down
go test -short ./...
```

### Import Cycle Errors

**Problem**: Circular import detected

```
import cycle not allowed
```

**Solution**:
- Restructure packages to avoid circular dependencies
- Move shared types to a common package
- Use interfaces to break dependency cycles

### Debugger Not Working with TUI

**Problem**: Can't use debugger with terminal UI

**Solution**:

Use logging instead:
```go
import (
    "log"
    "os"
)

func init() {
    f, _ := os.OpenFile("debug.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    log.SetOutput(f)
}

// In your code
log.Printf("Model state: %+v", m)
```

Then tail the log:
```bash
tail -f debug.log
```

### Module Version Conflicts

**Problem**: Incompatible module versions

```
go: github.com/charmbracelet/bubbletea@v1.3.10: parsing go.mod:
    module declares its path as: github.com/charmbracelet/bubbletea/v2
```

**Solution**:
```bash
# Update all dependencies
go get -u ./...
go mod tidy

# Or specific package
go get github.com/charmbracelet/bubbletea@latest
```

## Platform-Specific Issues

### macOS

**Problem**: "ki" is damaged and can't be opened

**Solution**:
```bash
# Remove quarantine attribute
xattr -d com.apple.quarantine ./ki

# Or allow in Security & Privacy settings
```

### Linux

**Problem**: Permission denied when running binary

**Solution**:
```bash
chmod +x ki
./ki tui
```

### Windows

**Problem**: Terminal colors not working

**Solution**:
- Use Windows Terminal (recommended)
- Enable ANSI colors in Command Prompt:
```cmd
reg add HKCU\Console /v VirtualTerminalLevel /t REG_DWORD /d 1
```

**Problem**: Emoji not displaying

**Solution**:
- Update Windows to version 1903 or later
- Use Windows Terminal
- Install a font with emoji support (Cascadia Code)

## Getting More Help

If you're still experiencing issues:

1. **Check existing issues**: https://github.com/bilalbaraz/ki/issues
2. **Create a bug report**: Include:
   - Operating system and version
   - Go version (`go version`)
   - Terminal emulator
   - Full error message
   - Steps to reproduce
3. **Enable debug mode** (if available):
```bash
export KI_DEBUG=1
./ki tui 2>&1 | tee debug.log
```

## Diagnostic Commands

Run these to gather system information for bug reports:

```bash
echo "=== System Info ==="
uname -a
echo ""

echo "=== Go Version ==="
go version
echo ""

echo "=== Terminal Info ==="
echo "TERM: $TERM"
echo "COLORTERM: $COLORTERM"
echo "Dimensions: $COLUMNS x $LINES"
echo ""

echo "=== Locale ==="
locale
echo ""

echo "=== Network Test ==="
curl -I https://dragonball-api.com/api/characters
echo ""

echo "=== DNS Test ==="
nslookup dragonball-api.com
```

Save output:
```bash
./diagnose.sh > diagnostic_report.txt 2>&1
```

## Known Issues

### Issue #1: Unicode Width Calculation

Some characters may not align properly in certain terminals.

**Workaround**: Use a terminal with proper Unicode support (iTerm2, Alacritty)

### Issue #2: Mouse Support Conflicts

Mouse scrolling may conflict with terminal emulator scrollback.

**Workaround**: Use keyboard navigation instead

## Preventive Measures

To avoid common issues:

1. **Keep dependencies updated**:
```bash
go get -u ./...
go mod tidy
```

2. **Use recommended terminals**:
   - macOS: iTerm2, Alacritty
   - Linux: GNOME Terminal, Konsole, Alacritty
   - Windows: Windows Terminal

3. **Verify environment**:
```bash
# Check Go installation
which go
go version

# Check GOPATH
go env GOPATH
```

4. **Regular testing**:
```bash
# Before committing changes
go test ./...
go build -o ki
./ki tui
```

---

If this guide doesn't solve your problem, please open an issue on GitHub with detailed information about your environment and the error you're experiencing.

**Power Level: Over 9000!** 💪