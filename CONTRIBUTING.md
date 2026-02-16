# Contributing to KI - Dragon Ball CLI

Thank you for your interest in contributing to KI! This document provides guidelines and instructions for contributing to the project.

## 📋 Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Project Structure](#project-structure)
- [How to Contribute](#how-to-contribute)
- [Coding Standards](#coding-standards)
- [Testing Guidelines](#testing-guidelines)
- [Commit Messages](#commit-messages)
- [Pull Request Process](#pull-request-process)
- [Reporting Bugs](#reporting-bugs)
- [Suggesting Enhancements](#suggesting-enhancements)

## Code of Conduct

This project adheres to a code of conduct that all contributors are expected to follow:

- Be respectful and inclusive
- Welcome newcomers and help them learn
- Focus on what is best for the community
- Show empathy towards other community members

## Getting Started

### Prerequisites

- Go 1.25 or higher
- Git
- A terminal with 256 color support
- Basic understanding of Go and TUI concepts

### Fork and Clone

1. Fork the repository on GitHub
2. Clone your fork locally:

```bash
git clone https://github.com/bilalbaraz/ki.git
cd ki
```

3. Add the upstream repository:

```bash
git remote add upstream https://github.com/bilalbaraz/ki.git
```

4. Create a new branch for your feature:

```bash
git checkout -b feature/your-feature-name
```

## Development Setup

### Install Dependencies

```bash
go mod download
```

### Build the Project

```bash
go build -o ki
```

### Run the Application

```bash
./ki tui
```

### Run Tests

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run with race detector
go test -race ./...

# Run specific package
go test ./internal/api/...
```

### Format Code

```bash
# Format all Go files
go fmt ./...

# Run linter (requires golangci-lint)
golangci-lint run
```

## Project Structure

```
ki/
├── cmd/              # CLI commands
├── internal/
│   ├── api/         # API client
│   └── tui/         # TUI implementation
├── main.go          # Entry point
└── README.md        # Documentation
```

See [ARCHITECTURE.md](ARCHITECTURE.md) for detailed architecture documentation.

## How to Contribute

### Types of Contributions

We welcome several types of contributions:

1. **Bug Fixes**: Fix issues in existing functionality
2. **Features**: Add new features or enhance existing ones
3. **Documentation**: Improve or add documentation
4. **Tests**: Add or improve test coverage
5. **Performance**: Optimize code for better performance
6. **UI/UX**: Improve the user interface or experience

### Good First Issues

Look for issues labeled `good first issue` - these are great for newcomers!

### Feature Requests

Check existing issues before creating a new one. If your idea is new:

1. Open an issue describing the feature
2. Explain why it would be useful
3. Provide examples or mockups if applicable
4. Wait for maintainer feedback before implementing

## Coding Standards

### Go Style Guide

Follow the official [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments):

- Use `gofmt` for formatting
- Follow Go naming conventions
- Keep functions small and focused
- Use meaningful variable names
- Add comments for exported functions

### Project-Specific Guidelines

#### API Layer (`internal/api/`)

```go
// Good: Clear error messages with context
func (c *Client) GetCharacters(page, limit int) (*CharactersResponse, error) {
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch characters: %w", err)
    }
    // ...
}
```

#### TUI Layer (`internal/tui/`)

```go
// Good: Descriptive message types
type CharactersLoadedMsg struct {
    Response *api.CharactersResponse
}

// Good: Clear state management
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case CharactersLoadedMsg:
        m.charactersState = StateLoaded
        m.characters = msg.Response.Items
        return m, nil
    }
    // ...
}
```

#### Styles (`internal/tui/styles.go`)

```go
// Good: Descriptive style names with semantic meaning
var characterNameStyle = lipgloss.NewStyle().
    Foreground(primaryColor).
    Bold(true)

// Avoid: Generic names
var redBoldStyle = lipgloss.NewStyle().
    Foreground(lipgloss.Color("#FF0000")).
    Bold(true)
```

### File Organization

- Keep files under 500 lines
- Group related functionality
- Use clear, descriptive file names
- Add package comments

### Comments

```go
// Good: Explains WHY, not WHAT
// LoadCharactersCmd fetches characters asynchronously to keep the UI responsive
func LoadCharactersCmd(client *api.Client, page, limit int) tea.Cmd {
    // ...
}

// Avoid: Redundant comments
// GetCharacters gets characters
func GetCharacters() {}
```

## Testing Guidelines

### Unit Tests

Write unit tests for:
- API client methods
- Data transformations
- Helper functions

```go
func TestGetCharacters(t *testing.T) {
    // Setup
    client := api.NewClient()
    
    // Execute
    resp, err := client.GetCharacters(1, 10)
    
    // Assert
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
    if len(resp.Items) == 0 {
        t.Error("expected items, got none")
    }
}
```

### Integration Tests

Test message flow and state transitions:

```go
func TestTabSwitching(t *testing.T) {
    m := tui.NewModel()
    
    // Switch to planets tab
    m, _ = m.Update(tea.KeyMsg{Type: tea.KeyTab})
    
    if m.ActiveTab() != tui.PlanetsTab {
        t.Error("expected PlanetsTab")
    }
}
```

### Test Coverage

- Aim for >80% coverage on critical paths
- Focus on edge cases and error handling
- Don't test external dependencies (mock them)

## Commit Messages

### Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Types

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

### Examples

```
feat(tui): add search functionality to character list

Implement fuzzy search that filters characters as the user types.
The search is case-insensitive and matches against character names.

Closes #42
```

```
fix(api): handle timeout errors gracefully

Previously, timeout errors would crash the application.
Now they are caught and displayed to the user with a retry option.

Fixes #56
```

### Best Practices

- Keep subject line under 50 characters
- Capitalize the subject line
- Use imperative mood ("add" not "added")
- Separate subject from body with blank line
- Wrap body at 72 characters
- Reference issues and PRs

## Pull Request Process

### Before Submitting

1. ✅ Code is formatted (`go fmt ./...`)
2. ✅ Tests pass (`go test ./...`)
3. ✅ No linter errors (`golangci-lint run`)
4. ✅ Documentation is updated
5. ✅ Commits follow commit message guidelines

### Submitting a PR

1. Push your branch to your fork
2. Open a Pull Request against `main`
3. Fill out the PR template completely
4. Link related issues
5. Request review from maintainers

### PR Template

```markdown
## Description
Brief description of what this PR does.

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Testing
Describe how you tested your changes.

## Screenshots (if applicable)
Add screenshots of UI changes.

## Checklist
- [ ] Code follows project style guidelines
- [ ] Self-review completed
- [ ] Comments added for complex code
- [ ] Documentation updated
- [ ] Tests added/updated
- [ ] All tests pass
```

### Review Process

- At least one maintainer must approve
- All CI checks must pass
- Conflicts must be resolved
- Changes may be requested

### After Approval

- Maintainer will merge your PR
- Your branch will be deleted
- You'll be added to contributors list!

## Reporting Bugs

### Before Reporting

1. Check existing issues
2. Verify you're on the latest version
3. Try to reproduce the bug

### Bug Report Template

```markdown
**Describe the bug**
A clear description of what the bug is.

**To Reproduce**
Steps to reproduce:
1. Go to '...'
2. Click on '...'
3. See error

**Expected behavior**
What you expected to happen.

**Screenshots**
If applicable, add screenshots.

**Environment:**
- OS: [e.g., macOS 13.0]
- Terminal: [e.g., iTerm2]
- Go version: [e.g., 1.25]
- KI version: [e.g., 1.0.0]

**Additional context**
Any other relevant information.
```

## Suggesting Enhancements

### Enhancement Template

```markdown
**Is your feature related to a problem?**
A clear description of the problem.

**Describe the solution**
How you'd like the feature to work.

**Describe alternatives**
Other solutions you've considered.

**Additional context**
Mockups, examples, or references.
```

## Development Tips

### Debugging TUI Applications

Use log files since stdout is captured:

```go
import "log"

func init() {
    f, _ := os.OpenFile("debug.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    log.SetOutput(f)
}

// In your code
log.Printf("Debug: %+v", model)
```

### Testing UI Locally

```bash
# Test with different terminal sizes
# Resize your terminal and verify layout

# Test with different color support
TERM=xterm ./ki tui    # 256 colors
TERM=linux ./ki tui    # Basic colors
```

### Common Issues

**"Module not found"**
```bash
go mod tidy
go mod download
```

**"Build failed"**
```bash
# Clean build cache
go clean -cache
go build -o ki
```

**"Tests hanging"**
```bash
# Use timeout
go test -timeout 30s ./...
```

## Getting Help

- **Questions**: Open a discussion on GitHub
- **Bugs**: Open an issue with bug report template
- **Features**: Open an issue with enhancement template
- **Security**: Email bilalbaraz@windowslive.com (do not open public issue)

## Recognition

Contributors are recognized in:
- GitHub contributors page
- Release notes
- CONTRIBUTORS.md file

Thank you for contributing to KI! 🐉

---

**Over 9000!** 💪
