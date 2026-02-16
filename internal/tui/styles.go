package tui

import "github.com/charmbracelet/lipgloss"

var (
	// Color palette
	primaryColor   = lipgloss.Color("#FF6B35") // Orange (Dragon Ball energy)
	secondaryColor = lipgloss.Color("#F7931E") // Golden
	accentColor    = lipgloss.Color("#4ECDC4") // Cyan
	textColor      = lipgloss.Color("#FFFFFF") // White
	mutedColor     = lipgloss.Color("#95A3B3") // Gray
	bgColor        = lipgloss.Color("#1A1A2E") // Dark blue
	errorColor     = lipgloss.Color("#FF4757") // Red
	successColor   = lipgloss.Color("#2ED573") // Green

	// Base style
	baseStyle = lipgloss.NewStyle().
			Foreground(textColor)

	// Title style
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(primaryColor).
			Background(bgColor).
			Padding(0, 1).
			MarginBottom(1)

	// Tab styles
	activeTabStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(textColor).
			Background(primaryColor).
			Padding(0, 2).
			MarginRight(1)

	inactiveTabStyle = lipgloss.NewStyle().
				Foreground(mutedColor).
				Background(lipgloss.Color("#2D2D44")).
				Padding(0, 2).
				MarginRight(1)

	tabContainerStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(primaryColor).
				BorderBottom(true).
				MarginBottom(1)

	// Content styles
	contentStyle = lipgloss.NewStyle().
			Padding(1, 2).
			MarginTop(1)

	// List item styles
	selectedItemStyle = lipgloss.NewStyle().
				Foreground(primaryColor).
				Bold(true).
				PaddingLeft(2)

	normalItemStyle = lipgloss.NewStyle().
			Foreground(textColor).
			PaddingLeft(2)

	itemTitleStyle = lipgloss.NewStyle().
			Foreground(secondaryColor).
			Bold(true)

	itemDescStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			MarginLeft(4)

	// Status styles
	loadingStyle = lipgloss.NewStyle().
			Foreground(accentColor).
			Bold(true).
			Padding(1, 2)

	errorStyle = lipgloss.NewStyle().
			Foreground(errorColor).
			Bold(true).
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(errorColor)

	// Footer style
	footerStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Background(lipgloss.Color("#2D2D44")).
			Padding(0, 1).
			MarginTop(1)

	helpKeyStyle = lipgloss.NewStyle().
			Foreground(accentColor).
			Bold(true)

	helpDescStyle = lipgloss.NewStyle().
			Foreground(mutedColor)

	// Info styles
	infoLabelStyle = lipgloss.NewStyle().
			Foreground(accentColor).
			Bold(true)

	infoValueStyle = lipgloss.NewStyle().
			Foreground(textColor)

	// Pagination style
	paginationStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Align(lipgloss.Center).
			MarginTop(1)

	// Border styles
	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(primaryColor).
			Padding(1, 2)

	// Character specific styles
	characterNameStyle = lipgloss.NewStyle().
				Foreground(primaryColor).
				Bold(true)

	characterRaceStyle = lipgloss.NewStyle().
				Foreground(secondaryColor)

	characterKiStyle = lipgloss.NewStyle().
				Foreground(accentColor).
				Bold(true)

	// Planet specific styles
	planetNameStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true)

	planetDestroyedStyle = lipgloss.NewStyle().
				Foreground(errorColor).
				Bold(true)

	planetActiveStyle = lipgloss.NewStyle().
				Foreground(successColor).
				Bold(true)
)
