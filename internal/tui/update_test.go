package tui

import (
	"errors"
	"testing"

	"github.com/bilalbaraz/ki/internal/api"
	tea "github.com/charmbracelet/bubbletea"
)

func TestInit(t *testing.T) {
	m := NewModel()
	cmd := m.Init()

	if cmd == nil {
		t.Fatal("expected non-nil command from Init")
	}
}

func TestUpdate_WindowSizeMsg(t *testing.T) {
	m := NewModel()

	newModel, cmd := m.Update(tea.WindowSizeMsg{Width: 100, Height: 50})

	updatedModel := newModel.(Model)

	if updatedModel.width != 100 {
		t.Errorf("expected width 100, got %d", updatedModel.width)
	}

	if updatedModel.height != 50 {
		t.Errorf("expected height 50, got %d", updatedModel.height)
	}

	if !updatedModel.ready {
		t.Error("expected ready to be true after WindowSizeMsg")
	}

	if cmd != nil {
		t.Error("expected nil command after WindowSizeMsg")
	}
}

func TestUpdate_QuitKeys(t *testing.T) {
	tests := []struct {
		name string
		key  string
	}{
		{"ctrl+c", "ctrl+c"},
		{"q", "q"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewModel()

			_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(tt.key)})

			if cmd == nil {
				t.Fatal("expected non-nil command for quit")
			}

			if cmd() != tea.Quit() {
				t.Error("expected tea.Quit command")
			}
		})
	}
}

func TestUpdate_TabSwitching(t *testing.T) {
	m := NewModel()
	m.ready = true

	// Initially on CharactersTab
	if m.activeTab != CharactersTab {
		t.Fatal("expected initial tab to be CharactersTab")
	}

	// Switch to PlanetsTab with tab key
	newModel, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("tab")})
	m = newModel.(Model)

	if m.activeTab != PlanetsTab {
		t.Errorf("expected PlanetsTab after tab, got %d", m.activeTab)
	}

	if cmd == nil {
		t.Error("expected command to load planets")
	}

	// Mark planets as loaded
	m.planetsState = StateLoaded

	// Switch back to CharactersTab
	newModel, cmd = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("tab")})
	m = newModel.(Model)

	if m.activeTab != CharactersTab {
		t.Errorf("expected CharactersTab after second tab, got %d", m.activeTab)
	}

	if cmd != nil {
		t.Error("expected nil command when switching back to already loaded tab")
	}
}

func TestUpdate_NumberKeyNavigation(t *testing.T) {
	m := NewModel()
	m.ready = true

	// Press "1" to go to CharactersTab
	newModel, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("1")})
	m = newModel.(Model)

	if m.activeTab != CharactersTab {
		t.Errorf("expected CharactersTab after pressing 1, got %d", m.activeTab)
	}

	// Press "2" to go to PlanetsTab
	newModel, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("2")})
	m = newModel.(Model)

	if m.activeTab != PlanetsTab {
		t.Errorf("expected PlanetsTab after pressing 2, got %d", m.activeTab)
	}

	if cmd == nil {
		t.Error("expected command to load planets")
	}

	if m.planetsState != StateLoading {
		t.Error("expected planetsState to be StateLoading")
	}
}

func TestUpdate_Refresh(t *testing.T) {
	m := NewModel()
	m.ready = true
	m.activeTab = CharactersTab
	m.charactersState = StateLoaded

	// Refresh characters tab
	newModel, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("r")})
	m = newModel.(Model)

	if m.charactersState != StateLoading {
		t.Error("expected charactersState to be StateLoading after refresh")
	}

	if cmd == nil {
		t.Error("expected command to reload characters")
	}

	// Switch to planets tab and refresh
	m.activeTab = PlanetsTab
	m.planetsState = StateLoaded

	newModel, cmd = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("r")})
	m = newModel.(Model)

	if m.planetsState != StateLoading {
		t.Error("expected planetsState to be StateLoading after refresh")
	}

	if cmd == nil {
		t.Error("expected command to reload planets")
	}
}

func TestUpdate_NextPage_Characters(t *testing.T) {
	m := NewModel()
	m.ready = true
	m.activeTab = CharactersTab
	m.currentCharactersPage = 1
	m.charactersMeta = api.Meta{
		CurrentPage: 1,
		TotalPages:  5,
	}

	// Next page with "n" key
	newModel, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("n")})
	m = newModel.(Model)

	if m.currentCharactersPage != 2 {
		t.Errorf("expected page 2, got %d", m.currentCharactersPage)
	}

	if m.charactersState != StateLoading {
		t.Error("expected StateLoading after next page")
	}

	if cmd == nil {
		t.Error("expected command to load next page")
	}

	// Next page with "right" key
	m.charactersMeta.CurrentPage = 2
	newModel, cmd = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("right")})
	m = newModel.(Model)

	if m.currentCharactersPage != 3 {
		t.Errorf("expected page 3, got %d", m.currentCharactersPage)
	}

	// Try next page when at last page
	m.currentCharactersPage = 5
	m.charactersMeta.CurrentPage = 5
	newModel, cmd = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("n")})
	m = newModel.(Model)

	if m.currentCharactersPage != 5 {
		t.Errorf("expected to stay on page 5, got %d", m.currentCharactersPage)
	}

	if cmd != nil {
		t.Error("expected nil command when at last page")
	}
}

func TestUpdate_PreviousPage_Characters(t *testing.T) {
	m := NewModel()
	m.ready = true
	m.activeTab = CharactersTab
	m.currentCharactersPage = 3
	m.charactersMeta = api.Meta{
		CurrentPage: 3,
		TotalPages:  5,
	}

	// Previous page with "p" key
	newModel, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("p")})
	m = newModel.(Model)

	if m.currentCharactersPage != 2 {
		t.Errorf("expected page 2, got %d", m.currentCharactersPage)
	}

	if m.charactersState != StateLoading {
		t.Error("expected StateLoading after previous page")
	}

	if cmd == nil {
		t.Error("expected command to load previous page")
	}

	// Previous page with "left" key
	m.charactersMeta.CurrentPage = 2
	newModel, cmd = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("left")})
	m = newModel.(Model)

	if m.currentCharactersPage != 1 {
		t.Errorf("expected page 1, got %d", m.currentCharactersPage)
	}

	// Try previous page when at first page
	m.currentCharactersPage = 1
	newModel, cmd = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("p")})
	m = newModel.(Model)

	if m.currentCharactersPage != 1 {
		t.Errorf("expected to stay on page 1, got %d", m.currentCharactersPage)
	}

	if cmd != nil {
		t.Error("expected nil command when at first page")
	}
}

func TestUpdate_NextPage_Planets(t *testing.T) {
	m := NewModel()
	m.ready = true
	m.activeTab = PlanetsTab
	m.currentPlanetsPage = 1
	m.planetsMeta = api.Meta{
		CurrentPage: 1,
		TotalPages:  3,
	}

	// Next page
	newModel, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("n")})
	m = newModel.(Model)

	if m.currentPlanetsPage != 2 {
		t.Errorf("expected page 2, got %d", m.currentPlanetsPage)
	}

	if m.planetsState != StateLoading {
		t.Error("expected StateLoading after next page")
	}

	if cmd == nil {
		t.Error("expected command to load next page")
	}
}

func TestUpdate_PreviousPage_Planets(t *testing.T) {
	m := NewModel()
	m.ready = true
	m.activeTab = PlanetsTab
	m.currentPlanetsPage = 2
	m.planetsMeta = api.Meta{
		CurrentPage: 2,
		TotalPages:  3,
	}

	// Previous page
	newModel, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("p")})
	m = newModel.(Model)

	if m.currentPlanetsPage != 1 {
		t.Errorf("expected page 1, got %d", m.currentPlanetsPage)
	}

	if m.planetsState != StateLoading {
		t.Error("expected StateLoading after previous page")
	}

	if cmd == nil {
		t.Error("expected command to load previous page")
	}
}

func TestUpdate_CharactersLoadedMsg(t *testing.T) {
	m := NewModel()
	m.ready = true

	response := &api.CharactersResponse{
		Items: []api.Character{
			{ID: 1, Name: "Goku", Race: "Saiyan"},
			{ID: 2, Name: "Vegeta", Race: "Saiyan"},
		},
		Meta: api.Meta{
			TotalItems:   2,
			CurrentPage:  1,
			ItemsPerPage: 10,
		},
	}

	msg := CharactersLoadedMsg{Response: response}

	newModel, cmd := m.Update(msg)
	m = newModel.(Model)

	if m.charactersState != StateLoaded {
		t.Errorf("expected StateLoaded, got %d", m.charactersState)
	}

	if len(m.characters) != 2 {
		t.Errorf("expected 2 characters, got %d", len(m.characters))
	}

	if m.characters[0].Name != "Goku" {
		t.Errorf("expected first character to be Goku, got %s", m.characters[0].Name)
	}

	if m.charactersMeta.TotalItems != 2 {
		t.Errorf("expected TotalItems 2, got %d", m.charactersMeta.TotalItems)
	}

	if cmd != nil {
		t.Error("expected nil command after loading characters")
	}
}

func TestUpdate_CharactersErrorMsg(t *testing.T) {
	m := NewModel()
	m.ready = true

	msg := CharactersErrorMsg{Err: errors.New("network error")}

	newModel, cmd := m.Update(msg)
	m = newModel.(Model)

	if m.charactersState != StateError {
		t.Errorf("expected StateError, got %d", m.charactersState)
	}

	if m.charactersError != "network error" {
		t.Errorf("expected error message 'network error', got '%s'", m.charactersError)
	}

	if cmd != nil {
		t.Error("expected nil command after error")
	}
}

func TestUpdate_PlanetsLoadedMsg(t *testing.T) {
	m := NewModel()
	m.ready = true

	response := &api.PlanetsResponse{
		Items: []api.Planet{
			{ID: 1, Name: "Earth", IsDestroyed: false},
			{ID: 2, Name: "Namek", IsDestroyed: false},
		},
		Meta: api.Meta{
			TotalItems:   2,
			CurrentPage:  1,
			ItemsPerPage: 10,
		},
	}

	msg := PlanetsLoadedMsg{Response: response}

	newModel, cmd := m.Update(msg)
	m = newModel.(Model)

	if m.planetsState != StateLoaded {
		t.Errorf("expected StateLoaded, got %d", m.planetsState)
	}

	if len(m.planets) != 2 {
		t.Errorf("expected 2 planets, got %d", len(m.planets))
	}

	if m.planets[0].Name != "Earth" {
		t.Errorf("expected first planet to be Earth, got %s", m.planets[0].Name)
	}

	if m.planetsMeta.TotalItems != 2 {
		t.Errorf("expected TotalItems 2, got %d", m.planetsMeta.TotalItems)
	}

	if cmd != nil {
		t.Error("expected nil command after loading planets")
	}
}

func TestUpdate_PlanetsErrorMsg(t *testing.T) {
	m := NewModel()
	m.ready = true

	msg := PlanetsErrorMsg{Err: errors.New("api error")}

	newModel, cmd := m.Update(msg)
	m = newModel.(Model)

	if m.planetsState != StateError {
		t.Errorf("expected StateError, got %d", m.planetsState)
	}

	if m.planetsError != "api error" {
		t.Errorf("expected error message 'api error', got '%s'", m.planetsError)
	}

	if cmd != nil {
		t.Error("expected nil command after error")
	}
}

func TestUpdate_ShiftTab(t *testing.T) {
	m := NewModel()
	m.ready = true
	m.activeTab = CharactersTab

	newModel, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("shift+tab")})
	m = newModel.(Model)

	if m.activeTab != PlanetsTab {
		t.Errorf("expected PlanetsTab after shift+tab, got %d", m.activeTab)
	}
}

func TestUpdate_UnknownKey(t *testing.T) {
	m := NewModel()
	m.ready = true

	newModel, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("x")})

	if newModel == nil {
		t.Fatal("expected non-nil model")
	}

	if cmd != nil {
		t.Error("expected nil command for unknown key")
	}
}

func TestUpdate_ListUpdate_CharactersTab(t *testing.T) {
	m := NewModel()
	m.ready = true
	m.activeTab = CharactersTab
	m.charactersState = StateLoaded

	// Send a message that the list might handle
	newModel, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})

	if newModel == nil {
		t.Fatal("expected non-nil model after list update")
	}
}

func TestUpdate_ListUpdate_PlanetsTab(t *testing.T) {
	m := NewModel()
	m.ready = true
	m.activeTab = PlanetsTab
	m.planetsState = StateLoaded

	// Send a message that the list might handle
	newModel, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})

	if newModel == nil {
		t.Fatal("expected non-nil model after list update")
	}
}

func TestUpdate_EmptyCharactersResponse(t *testing.T) {
	m := NewModel()

	msg := CharactersLoadedMsg{
		Response: &api.CharactersResponse{
			Items: []api.Character{},
			Meta:  api.Meta{},
		},
	}

	newModel, _ := m.Update(msg)
	m = newModel.(Model)

	if len(m.characters) != 0 {
		t.Errorf("expected 0 characters, got %d", len(m.characters))
	}

	if m.charactersState != StateLoaded {
		t.Error("expected StateLoaded even with empty response")
	}
}

func TestUpdate_EmptyPlanetsResponse(t *testing.T) {
	m := NewModel()

	msg := PlanetsLoadedMsg{
		Response: &api.PlanetsResponse{
			Items: []api.Planet{},
			Meta:  api.Meta{},
		},
	}

	newModel, _ := m.Update(msg)
	m = newModel.(Model)

	if len(m.planets) != 0 {
		t.Errorf("expected 0 planets, got %d", len(m.planets))
	}

	if m.planetsState != StateLoaded {
		t.Error("expected StateLoaded even with empty response")
	}
}

func TestUpdate_TabSwitchWithoutLoading(t *testing.T) {
	m := NewModel()
	m.ready = true
	m.activeTab = CharactersTab
	m.planetsState = StateLoaded

	// Switch to already loaded planets tab
	newModel, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("tab")})
	m = newModel.(Model)

	if m.activeTab != PlanetsTab {
		t.Error("expected tab to switch to PlanetsTab")
	}

	if cmd != nil {
		t.Error("expected no command when switching to already loaded tab")
	}
}

func TestUpdate_Key2WithLoadedPlanets(t *testing.T) {
	m := NewModel()
	m.ready = true
	m.planetsState = StateLoaded

	newModel, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("2")})
	m = newModel.(Model)

	if m.activeTab != PlanetsTab {
		t.Error("expected PlanetsTab")
	}

	if cmd != nil {
		t.Error("expected no command when planets already loaded")
	}
}

func TestUpdate_SpinnerTick(t *testing.T) {
	m := NewModel()
	m.ready = true

	// Create a spinner tick message
	tickCmd := m.spinner.Tick
	tickMsg := tickCmd()

	newModel, cmd := m.Update(tickMsg)

	if newModel == nil {
		t.Fatal("expected non-nil model after spinner tick")
	}

	if cmd == nil {
		t.Error("expected command from spinner tick")
	}
}

func TestUpdate_MultipleListUpdates(t *testing.T) {
	m := NewModel()
	m.ready = true
	m.activeTab = CharactersTab
	m.charactersState = StateIdle

	// Should not update list when not loaded
	newModel, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})
	m = newModel.(Model)

	if m.charactersState != StateIdle {
		t.Error("state should remain idle")
	}

	// Load characters
	m.charactersState = StateLoaded
	m.characters = []api.Character{{ID: 1, Name: "Goku"}}

	// Now list should update
	newModel, _ = m.Update(tea.KeyMsg{Type: tea.KeyDown})

	if newModel == nil {
		t.Fatal("expected non-nil model")
	}
}
