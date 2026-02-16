package tui

import (
	"github.com/bilalbaraz/ki/internal/api"
	tea "github.com/charmbracelet/bubbletea"
)

// Message types for API responses

// CharactersLoadedMsg is sent when characters are successfully loaded
type CharactersLoadedMsg struct {
	Response *api.CharactersResponse
}

// CharactersErrorMsg is sent when loading characters fails
type CharactersErrorMsg struct {
	Err error
}

// PlanetsLoadedMsg is sent when planets are successfully loaded
type PlanetsLoadedMsg struct {
	Response *api.PlanetsResponse
}

// PlanetsErrorMsg is sent when loading planets fails
type PlanetsErrorMsg struct {
	Err error
}

// Commands for async operations

// LoadCharactersCmd loads characters from the API
func LoadCharactersCmd(client *api.Client, page, limit int) tea.Cmd {
	return func() tea.Msg {
		resp, err := client.GetCharacters(page, limit)
		if err != nil {
			return CharactersErrorMsg{Err: err}
		}
		return CharactersLoadedMsg{Response: resp}
	}
}

// LoadPlanetsCmd loads planets from the API
func LoadPlanetsCmd(client *api.Client, page, limit int) tea.Cmd {
	return func() tea.Msg {
		resp, err := client.GetPlanets(page, limit)
		if err != nil {
			return PlanetsErrorMsg{Err: err}
		}
		return PlanetsLoadedMsg{Response: resp}
	}
}
