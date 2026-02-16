# 🚀 Quick Start Guide

Get up and running with KI - Dragon Ball CLI in under 2 minutes!

## Prerequisites

- Go 1.25 or higher
- A terminal emulator (iTerm2, Alacritty, Windows Terminal recommended)

## Installation

### Option 1: Clone and Build (Recommended)

```bash
# Clone the repository
git clone https://github.com/bilalbaraz/ki.git
cd ki

# Download dependencies
go mod download

# Build the application
go build -o ki

# Run the TUI
./ki tui
```

### Option 2: Quick Run (No Build)

```bash
# Clone the repository
git clone https://github.com/bilalbaraz/ki.git
cd ki

# Run directly
go run main.go tui
```

### Option 3: Install Globally

```bash
go install github.com/bilalbaraz/ki@latest

# Run from anywhere
ki tui
```

## First Launch

When you first launch the TUI:

1. **Characters tab** loads automatically
2. Wait a few seconds for the API data to load
3. You'll see a list of Dragon Ball characters!

## Essential Controls

```
Tab          → Switch between tabs
1 or 2       → Jump to specific tab
↑ ↓          → Navigate list
← →          → Previous/Next page
r            → Refresh
q            → Quit
```

## What You'll See

```
┌─────────────────────────────────────────────────┐
│           🐉 DRAGON BALL CLI                    │
├─────────────────────────────────────────────────┤
│ [Characters] [ Planets ]                        │
├─────────────────────────────────────────────────┤
│                                                 │
│  • Goku              Saiyan • Z Fighter         │
│  • Vegeta            Saiyan • Z Fighter         │
│  • Piccolo           Namekian • Z Fighter       │
│  • Frieza            Unknown • Frieza Force     │
│  • Gohan             Saiyan • Z Fighter         │
│                                                 │
│         Page 1 of 5 • Total: 85 items          │
├─────────────────────────────────────────────────┤
│ tab switch │ ←/→ page │ r refresh │ q quit     │
└─────────────────────────────────────────────────┘
```

## Exploring the App

### Characters Tab

Press `1` or ensure you're on the Characters tab:
- Browse all Dragon Ball characters
- See their race and affiliation
- Navigate with arrow keys
- Page through results with `←` and `→`

### Planets Tab

Press `2` or `Tab` to switch:
- View all Dragon Ball planets
- See if they're destroyed or active
- Same navigation as Characters

### Pagination

- **Next page**: Press `n` or `→`
- **Previous page**: Press `p` or `←`
- **Page info**: Displayed at the bottom of the list

### Refreshing Data

Press `r` to reload the current tab's data from the API.

## Troubleshooting

### "Failed to fetch characters"

**Problem**: No internet or API is down

**Solution**:
```bash
# Test your connection
curl https://dragonball-api.com/api/characters

# Check DNS
ping dragonball-api.com
```

### UI looks weird/garbled

**Problem**: Terminal doesn't support colors or Unicode

**Solution**:
```bash
# Set proper terminal
export TERM=xterm-256color

# Use a modern terminal emulator
# Recommended: iTerm2 (macOS), Windows Terminal (Windows), Alacritty (Linux)
```

### Permission denied

**Problem**: Binary not executable

**Solution**:
```bash
chmod +x ki
./ki tui
```

## Next Steps

Now that you're up and running:

1. **Read the full docs**: Check out [README.md](README.md)
2. **Learn the architecture**: See [ARCHITECTURE.md](ARCHITECTURE.md)
3. **Contribute**: Read [CONTRIBUTING.md](CONTRIBUTING.md)
4. **Report issues**: Use GitHub Issues

## Getting Help

- **Documentation**: See README.md
- **Technical details**: See ARCHITECTURE.md
- **Common issues**: See TROUBLESHOOTING.md
- **GitHub Issues**: https://github.com/bilalbaraz/ki/issues

## Keyboard Shortcuts Reference

| Key | Action |
|-----|--------|
| `Tab` | Switch to next tab |
| `Shift+Tab` | Switch to previous tab |
| `1` | Jump to Characters tab |
| `2` | Jump to Planets tab |
| `↑` or `k` | Move up in list |
| `↓` or `j` | Move down in list |
| `←` or `p` | Previous page |
| `→` or `n` | Next page |
| `r` | Refresh current tab |
| `q` or `Ctrl+C` | Quit application |

## API Information

This app uses the [Dragon Ball API](https://dragonball-api.com):
- Base URL: `https://dragonball-api.com/api`
- Rate limits may apply
- Free to use, no API key required

## System Requirements

- **OS**: macOS, Linux, or Windows
- **Go**: 1.25 or higher
- **Terminal**: Any with 256 color support
- **Network**: Internet connection required
- **Disk**: ~15MB for compiled binary

## Building for Production

```bash
# Build with optimizations
go build -ldflags="-s -w" -o ki

# Cross-compile for different platforms
GOOS=linux GOARCH=amd64 go build -o ki-linux
GOOS=windows GOARCH=amd64 go build -o ki-windows.exe
GOOS=darwin GOARCH=arm64 go build -o ki-darwin-arm64
```

## Uninstalling

```bash
# Remove binary
rm ki

# Remove source (if cloned)
cd .. && rm -rf ki

# Remove globally installed version
rm $(which ki)
```

---

**Enjoy exploring the Dragon Ball universe!** 🐉

**Power Level: Over 9000!** 💪