package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// Init initializes the model and returns initial commands
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		LoadCharactersCmd(m.client, m.currentCharactersPage, m.itemsPerPage),
	)
}

// Update handles incoming messages and updates the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true

		// Update list sizes
		listWidth := msg.Width - 4
		listHeight := msg.Height - 10

		m.charactersList.SetSize(listWidth, listHeight)
		m.planetsList.SetSize(listWidth, listHeight)

		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "tab", "shift+tab":
			// Switch tabs
			if m.activeTab == CharactersTab {
				m.activeTab = PlanetsTab
				// Load planets if not loaded yet
				if m.planetsState == StateIdle {
					m.planetsState = StateLoading
					return m, LoadPlanetsCmd(m.client, m.currentPlanetsPage, m.itemsPerPage)
				}
			} else {
				m.activeTab = CharactersTab
			}
			return m, nil

		case "1":
			m.activeTab = CharactersTab
			return m, nil

		case "2":
			m.activeTab = PlanetsTab
			// Load planets if not loaded yet
			if m.planetsState == StateIdle {
				m.planetsState = StateLoading
				return m, LoadPlanetsCmd(m.client, m.currentPlanetsPage, m.itemsPerPage)
			}
			return m, nil

		case "r":
			// Refresh current tab
			if m.activeTab == CharactersTab {
				m.charactersState = StateLoading
				return m, LoadCharactersCmd(m.client, m.currentCharactersPage, m.itemsPerPage)
			} else {
				m.planetsState = StateLoading
				return m, LoadPlanetsCmd(m.client, m.currentPlanetsPage, m.itemsPerPage)
			}

		case "n", "right":
			// Next page
			if m.activeTab == CharactersTab {
				if m.charactersMeta.CurrentPage < m.charactersMeta.TotalPages {
					m.currentCharactersPage++
					m.charactersState = StateLoading
					return m, LoadCharactersCmd(m.client, m.currentCharactersPage, m.itemsPerPage)
				}
			} else {
				if m.planetsMeta.CurrentPage < m.planetsMeta.TotalPages {
					m.currentPlanetsPage++
					m.planetsState = StateLoading
					return m, LoadPlanetsCmd(m.client, m.currentPlanetsPage, m.itemsPerPage)
				}
			}
			return m, nil

		case "p", "left":
			// Previous page
			if m.activeTab == CharactersTab {
				if m.currentCharactersPage > 1 {
					m.currentCharactersPage--
					m.charactersState = StateLoading
					return m, LoadCharactersCmd(m.client, m.currentCharactersPage, m.itemsPerPage)
				}
			} else {
				if m.currentPlanetsPage > 1 {
					m.currentPlanetsPage--
					m.planetsState = StateLoading
					return m, LoadPlanetsCmd(m.client, m.currentPlanetsPage, m.itemsPerPage)
				}
			}
			return m, nil
		}

	case CharactersLoadedMsg:
		m.charactersState = StateLoaded
		m.characters = msg.Response.Items
		m.charactersMeta = msg.Response.Meta

		// Convert characters to list items
		items := make([]list.Item, len(m.characters))
		for i, char := range m.characters {
			items[i] = characterItem{character: char}
		}
		m.charactersList.SetItems(items)

		return m, nil

	case CharactersErrorMsg:
		m.charactersState = StateError
		m.charactersError = msg.Err.Error()
		return m, nil

	case PlanetsLoadedMsg:
		m.planetsState = StateLoaded
		m.planets = msg.Response.Items
		m.planetsMeta = msg.Response.Meta

		// Convert planets to list items
		items := make([]list.Item, len(m.planets))
		for i, planet := range m.planets {
			items[i] = planetItem{planet: planet}
		}
		m.planetsList.SetItems(items)

		return m, nil

	case PlanetsErrorMsg:
		m.planetsState = StateError
		m.planetsError = msg.Err.Error()
		return m, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		cmds = append(cmds, cmd)
	}

	// Update the active list
	if m.activeTab == CharactersTab && m.charactersState == StateLoaded {
		var cmd tea.Cmd
		m.charactersList, cmd = m.charactersList.Update(msg)
		cmds = append(cmds, cmd)
	} else if m.activeTab == PlanetsTab && m.planetsState == StateLoaded {
		var cmd tea.Cmd
		m.planetsList, cmd = m.planetsList.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
