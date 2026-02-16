package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// View renders the UI
func (m Model) View() string {
	if !m.ready {
		return "Initializing..."
	}

	var content strings.Builder

	// Render header with title
	content.WriteString(m.renderHeader())
	content.WriteString("\n")

	// Render tabs
	content.WriteString(m.renderTabs())
	content.WriteString("\n")

	// Render content based on active tab
	content.WriteString(m.renderContent())
	content.WriteString("\n")

	// Render footer with help
	content.WriteString(m.renderFooter())

	return content.String()
}

// renderHeader renders the application header
func (m Model) renderHeader() string {
	title := titleStyle.Render("🐉 DRAGON BALL CLI")
	return lipgloss.PlaceHorizontal(m.width, lipgloss.Center, title)
}

// renderTabs renders the tab navigation
func (m Model) renderTabs() string {
	var tabs []string

	// Characters tab
	if m.activeTab == CharactersTab {
		tabs = append(tabs, activeTabStyle.Render("Characters"))
	} else {
		tabs = append(tabs, inactiveTabStyle.Render("Characters"))
	}

	// Planets tab
	if m.activeTab == PlanetsTab {
		tabs = append(tabs, activeTabStyle.Render("Planets"))
	} else {
		tabs = append(tabs, inactiveTabStyle.Render("Planets"))
	}

	tabRow := lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
	return tabContainerStyle.Width(m.width - 2).Render(tabRow)
}

// renderContent renders the main content area based on active tab
func (m Model) renderContent() string {
	if m.activeTab == CharactersTab {
		return m.renderCharactersContent()
	}
	return m.renderPlanetsContent()
}

// renderCharactersContent renders the characters list
func (m Model) renderCharactersContent() string {
	switch m.charactersState {
	case StateLoading:
		return contentStyle.Render(
			fmt.Sprintf("%s Loading characters...", m.spinner.View()),
		)

	case StateError:
		errorMsg := errorStyle.Render(
			fmt.Sprintf("❌ Error loading characters:\n%s", m.charactersError),
		)
		return contentStyle.Render(errorMsg)

	case StateLoaded:
		if len(m.characters) == 0 {
			return contentStyle.Render("No characters found.")
		}

		// Render the list
		listView := m.charactersList.View()

		// Add pagination info
		pagination := m.renderPagination(
			m.charactersMeta.CurrentPage,
			m.charactersMeta.TotalPages,
			m.charactersMeta.TotalItems,
		)

		return lipgloss.JoinVertical(
			lipgloss.Left,
			listView,
			pagination,
		)

	default:
		return contentStyle.Render("Press 'r' to load characters")
	}
}

// renderPlanetsContent renders the planets list
func (m Model) renderPlanetsContent() string {
	switch m.planetsState {
	case StateLoading:
		return contentStyle.Render(
			fmt.Sprintf("%s Loading planets...", m.spinner.View()),
		)

	case StateError:
		errorMsg := errorStyle.Render(
			fmt.Sprintf("❌ Error loading planets:\n%s", m.planetsError),
		)
		return contentStyle.Render(errorMsg)

	case StateLoaded:
		if len(m.planets) == 0 {
			return contentStyle.Render("No planets found.")
		}

		// Render the list
		listView := m.planetsList.View()

		// Add pagination info
		pagination := m.renderPagination(
			m.planetsMeta.CurrentPage,
			m.planetsMeta.TotalPages,
			m.planetsMeta.TotalItems,
		)

		return lipgloss.JoinVertical(
			lipgloss.Left,
			listView,
			pagination,
		)

	default:
		return contentStyle.Render("Press 'r' to load planets")
	}
}

// renderPagination renders pagination information
func (m Model) renderPagination(current, total, totalItems int) string {
	if total == 0 {
		return ""
	}

	paginationText := fmt.Sprintf(
		"Page %d of %d • Total: %d items",
		current,
		total,
		totalItems,
	)

	return paginationStyle.Width(m.width - 4).Render(paginationText)
}

// renderFooter renders the footer with help text
func (m Model) renderFooter() string {
	var helpItems []string

	// Tab navigation
	helpItems = append(helpItems,
		fmt.Sprintf("%s %s", helpKeyStyle.Render("tab"), helpDescStyle.Render("switch tabs")),
	)

	// Number keys
	helpItems = append(helpItems,
		fmt.Sprintf("%s %s", helpKeyStyle.Render("1/2"), helpDescStyle.Render("jump to tab")),
	)

	// Pagination
	if m.activeTab == CharactersTab {
		if m.charactersMeta.TotalPages > 1 {
			helpItems = append(helpItems,
				fmt.Sprintf("%s %s", helpKeyStyle.Render("←/→"), helpDescStyle.Render("prev/next page")),
			)
		}
	} else {
		if m.planetsMeta.TotalPages > 1 {
			helpItems = append(helpItems,
				fmt.Sprintf("%s %s", helpKeyStyle.Render("←/→"), helpDescStyle.Render("prev/next page")),
			)
		}
	}

	// Refresh
	helpItems = append(helpItems,
		fmt.Sprintf("%s %s", helpKeyStyle.Render("r"), helpDescStyle.Render("refresh")),
	)

	// Quit
	helpItems = append(helpItems,
		fmt.Sprintf("%s %s", helpKeyStyle.Render("q"), helpDescStyle.Render("quit")),
	)

	helpText := lipgloss.JoinHorizontal(lipgloss.Left, helpItems...)
	return footerStyle.Width(m.width).Render(helpText)
}
