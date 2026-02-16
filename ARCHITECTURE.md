# Architecture Documentation

## Overview

KI is a Terminal User Interface (TUI) application built with Go and the Bubble Tea framework. It demonstrates best practices for building interactive CLI applications with asynchronous data fetching, state management, and responsive UI design.

## Technology Stack

- **Language**: Go 1.25+
- **TUI Framework**: Bubble Tea (The Elm Architecture for Go)
- **Styling**: Lipgloss (CSS-like styling for terminals)
- **Components**: Bubbles (reusable TUI components)
- **CLI**: Cobra (command-line interface framework)
- **API**: Dragon Ball API (RESTful HTTP API)

## The Elm Architecture (TEA)

Bubble Tea implements The Elm Architecture, which consists of three core concepts:

### 1. Model (State)

The `Model` struct contains all application state:

```go
type Model struct {
    // API client
    client *api.Client
    
    // UI state
    activeTab Tab
    width     int
    height    int
    ready     bool
    
    // Characters data
    characters      []api.Character
    charactersList  list.Model
    charactersState LoadingState
    charactersError string
    charactersMeta  api.Meta
    
    // Planets data
    planets      []api.Planet
    planetsList  list.Model
    planetsState LoadingState
    planetsError string
    planetsMeta  api.Meta
    
    // Loading spinner
    spinner spinner.Model
    
    // Pagination
    currentCharactersPage int
    currentPlanetsPage    int
    itemsPerPage          int
}
```

### 2. Update (State Transitions)

The `Update` function receives messages and returns an updated model:

```go
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        // Handle keyboard input
    case CharactersLoadedMsg:
        // Handle API response
    case tea.WindowSizeMsg:
        // Handle terminal resize
    }
    return m, cmd
}
```

### 3. View (Rendering)

The `View` function renders the current model state to a string:

```go
func (m Model) View() string {
    return lipgloss.JoinVertical(
        lipgloss.Left,
        m.renderHeader(),
        m.renderTabs(),
        m.renderContent(),
        m.renderFooter(),
    )
}
```

## Project Structure

```
ki/
├── cmd/                          # Command-line interface
│   ├── root.go                   # Root Cobra command
│   └── tui.go                    # TUI subcommand
│
├── internal/
│   ├── api/                      # API client layer
│   │   ├── client.go             # HTTP client implementation
│   │   └── types.go              # API data models
│   │
│   └── tui/                      # TUI implementation
│       ├── commands.go           # Tea commands (async operations)
│       ├── model.go              # Application state
│       ├── styles.go             # Lipgloss styles
│       ├── update.go             # Message handling
│       └── view.go               # UI rendering
│
├── main.go                       # Application entry point
├── go.mod                        # Go module definition
└── go.sum                        # Dependency checksums
```

## Module Responsibilities

### cmd/

**Responsibility**: Command-line interface setup and routing

- **root.go**: Defines the root command and global flags
- **tui.go**: Implements the `tui` subcommand that launches the TUI

### internal/api/

**Responsibility**: Communication with the Dragon Ball API

- **client.go**: 
  - HTTP client with timeout configuration
  - Methods for fetching characters and planets
  - Error handling and status code validation
  
- **types.go**:
  - Data models matching API response structure
  - JSON serialization tags
  - Pagination metadata structures

### internal/tui/

**Responsibility**: Terminal user interface implementation

- **model.go**:
  - Application state definition
  - Model initialization
  - List item implementations

- **update.go**:
  - Message handling logic
  - State transitions
  - Keyboard input processing
  - Command orchestration

- **view.go**:
  - UI rendering functions
  - Layout composition
  - Loading/error state displays

- **commands.go**:
  - Async command definitions
  - API call wrappers
  - Message type definitions

- **styles.go**:
  - Lipgloss style definitions
  - Color palette
  - Component styling

## Message Flow

### Application Startup

```
1. main.go
   └─> cmd.Execute()
       └─> tuiCmd.Run()
           └─> runTUI()
               └─> tea.NewProgram(model)
                   └─> model.Init()
                       ├─> spinner.Tick (recurring)
                       └─> LoadCharactersCmd (async)
```

### User Input Flow

```
User presses key
    ↓
tea.KeyMsg
    ↓
Update(tea.KeyMsg)
    ↓
Switch on key string
    ├─> "tab": Change activeTab
    ├─> "q": Return tea.Quit
    ├─> "r": Return LoadCharactersCmd
    └─> "n": Increment page, return LoadCharactersCmd
    ↓
Return (updatedModel, cmd)
    ↓
View(updatedModel)
    ↓
Render to terminal
```

### API Call Flow

```
1. User action triggers command
   └─> LoadCharactersCmd(client, page, limit)

2. Command executes in goroutine
   └─> client.GetCharacters(page, limit)
       ├─> Success: Return CharactersLoadedMsg{response}
       └─> Error: Return CharactersErrorMsg{err}

3. Message sent to Update()
   └─> Update(CharactersLoadedMsg)
       ├─> Set charactersState = StateLoaded
       ├─> Store characters data
       ├─> Update list items
       └─> Return (model, nil)

4. View() re-renders with new data
```

## State Management

### Loading States

The application uses a finite state machine for data loading:

```go
type LoadingState int

const (
    StateIdle    LoadingState = iota  // Initial state
    StateLoading                      // API call in progress
    StateLoaded                       // Data successfully loaded
    StateError                        // Error occurred
)
```

State transitions:

```
StateIdle ─────> StateLoading ─────> StateLoaded
                      │
                      └─────────────> StateError
                                           │
                      ┌────────────────────┘
                      ↓
                 StateLoading (retry)
```

### Tab Management

```go
type Tab int

const (
    CharactersTab Tab = iota
    PlanetsTab
)
```

Each tab maintains its own:
- Data array
- List component
- Loading state
- Error message
- Pagination state

This allows switching between tabs without re-fetching data.

## Concurrency Model

### Tea Commands

All asynchronous operations are wrapped in `tea.Cmd`:

```go
type Cmd func() Msg
```

Commands execute in goroutines and send messages back to the update loop:

```go
func LoadCharactersCmd(client *api.Client, page, limit int) tea.Cmd {
    return func() tea.Msg {
        // This runs in a goroutine
        resp, err := client.GetCharacters(page, limit)
        if err != nil {
            return CharactersErrorMsg{Err: err}
        }
        return CharactersLoadedMsg{Response: resp}
    }
}
```

### Thread Safety

- **Main Loop**: Single-threaded update loop (no race conditions)
- **API Calls**: Run in goroutines but send messages to main loop
- **No Shared State**: Each goroutine operates independently

## UI Layout

### Screen Structure

```
┌─────────────────────────────────────────────────────┐
│                🐉 DRAGON BALL CLI                   │
├─────────────────────────────────────────────────────┤
│ [Characters] [ Planets ]                            │
├─────────────────────────────────────────────────────┤
│                                                     │
│  • Goku              Saiyan • Z Fighter             │
│  • Vegeta            Saiyan • Z Fighter             │
│  • Piccolo           Namekian • Z Fighter           │
│  • Frieza            Unknown • Frieza Force         │
│  ...                                                │
│                                                     │
│           Page 1 of 5 • Total: 85 items            │
├─────────────────────────────────────────────────────┤
│ tab switch tabs │ 1/2 jump to tab │ ←/→ page │ q quit │
└─────────────────────────────────────────────────────┘
```

### Responsive Design

The UI adapts to terminal size:

```go
case tea.WindowSizeMsg:
    m.width = msg.Width
    m.height = msg.Height
    
    listWidth := msg.Width - 4
    listHeight := msg.Height - 10
    
    m.charactersList.SetSize(listWidth, listHeight)
    m.planetsList.SetSize(listWidth, listHeight)
```

## Error Handling Strategy

### API Layer

```go
func (c *Client) GetCharacters(page, limit int) (*CharactersResponse, error) {
    // Network error
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch characters: %w", err)
    }
    
    // HTTP error
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
    }
    
    // Decode error
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, fmt.Errorf("failed to decode response: %w", err)
    }
    
    return &result, nil
}
```

### TUI Layer

```go
case CharactersErrorMsg:
    m.charactersState = StateError
    m.charactersError = msg.Err.Error()
    // View() will render error message
```

### User Display

```go
errorMsg := errorStyle.Render(
    fmt.Sprintf("❌ Error loading characters:\n%s", m.charactersError),
)
```

## Styling System

### Color Palette

```go
var (
    primaryColor   = lipgloss.Color("#FF6B35") // Orange
    secondaryColor = lipgloss.Color("#F7931E") // Golden
    accentColor    = lipgloss.Color("#4ECDC4") // Cyan
    errorColor     = lipgloss.Color("#FF4757") // Red
    successColor   = lipgloss.Color("#2ED573") // Green
)
```

### Style Composition

Lipgloss uses a functional API for composing styles:

```go
activeTabStyle = lipgloss.NewStyle().
    Bold(true).
    Foreground(textColor).
    Background(primaryColor).
    Padding(0, 2).
    MarginRight(1)
```

### Rendering Pipeline

```
Data → Style → Render → String
```

Example:

```go
title := titleStyle.Render("🐉 DRAGON BALL CLI")
centered := lipgloss.PlaceHorizontal(width, lipgloss.Center, title)
```

## Performance Considerations

### Efficient Rendering

- Only re-render when model changes
- Use string builders for complex views
- Minimize allocations in hot paths

### API Optimization

- HTTP client reuse with connection pooling
- Configurable timeout (10 seconds)
- Pagination to limit response size

### Memory Management

- Fixed-size page limits (20 items default)
- No caching (stateless between sessions)
- Garbage collector handles cleanup

## Testing Strategy

### Unit Tests

```go
// internal/api/client_test.go
func TestGetCharacters(t *testing.T) {
    // Test successful response
    // Test error handling
    // Test pagination
}
```

### Integration Tests

```go
// internal/tui/model_test.go
func TestTabSwitching(t *testing.T) {
    m := NewModel()
    m, _ = m.Update(tea.KeyMsg{Type: tea.KeyTab})
    assert.Equal(t, PlanetsTab, m.activeTab)
}
```

### Manual Testing

Run the application and verify:
- Tab switching works
- Pagination works
- Error handling works
- UI is responsive

## Future Enhancements

### Potential Features

1. **Detail View**: Show full character/planet information
2. **Search**: Filter items by name
3. **Caching**: Store API responses locally
4. **Offline Mode**: Work without internet
5. **Favorites**: Bookmark items
6. **Export**: Save data to JSON/CSV
7. **Themes**: Customizable color schemes
8. **Config File**: Persistent settings

### Architecture Extensions

```
internal/
├── cache/         # Local data caching
├── config/        # Configuration management
├── export/        # Data export functionality
└── search/        # Search and filtering
```

## Dependencies

### Direct Dependencies

- `github.com/charmbracelet/bubbletea` - TUI framework
- `github.com/charmbracelet/lipgloss` - Styling
- `github.com/charmbracelet/bubbles` - UI components
- `github.com/spf13/cobra` - CLI framework

### Indirect Dependencies

- `github.com/sahilm/fuzzy` - Fuzzy search (used by list)
- `github.com/atotto/clipboard` - Clipboard support
- `golang.org/x/sys` - System calls
- `github.com/mattn/go-runewidth` - Unicode width calculation

## Build and Deployment

### Build Process

```bash
# Development build
go build -o ki

# Production build with optimizations
go build -ldflags="-s -w" -o ki

# Cross-compilation
GOOS=linux GOARCH=amd64 go build -o ki-linux
GOOS=windows GOARCH=amd64 go build -o ki-windows.exe
GOOS=darwin GOARCH=arm64 go build -o ki-darwin-arm64
```

### Distribution

- Single binary (no dependencies)
- ~15MB compiled size
- Runs on Linux, macOS, Windows

## References

- [Bubble Tea Documentation](https://github.com/charmbracelet/bubbletea)
- [Lipgloss Documentation](https://github.com/charmbracelet/lipgloss)
- [The Elm Architecture](https://guide.elm-lang.org/architecture/)
- [Dragon Ball API](https://dragonball-api.com)