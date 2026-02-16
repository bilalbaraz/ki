package tui

import (
	"github.com/bilalbaraz/ki/internal/api"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
)

// Tab represents the active tab in the UI
type Tab int

const (
	CharactersTab Tab = iota
	PlanetsTab
)

// LoadingState represents the current loading state
type LoadingState int

const (
	StateIdle LoadingState = iota
	StateLoading
	StateLoaded
	StateError
)

// Model represents the main application state
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

// NewModel creates a new TUI model
func NewModel() Model {
	// Initialize spinner
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = loadingStyle

	// Initialize lists
	charactersList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	charactersList.Title = "Dragon Ball Characters"
	charactersList.SetShowStatusBar(false)
	charactersList.SetFilteringEnabled(false)
	charactersList.SetShowHelp(false)

	planetsList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	planetsList.Title = "Dragon Ball Planets"
	planetsList.SetShowStatusBar(false)
	planetsList.SetFilteringEnabled(false)
	planetsList.SetShowHelp(false)

	return Model{
		client:                api.NewClient(),
		activeTab:             CharactersTab,
		characters:            []api.Character{},
		planets:               []api.Planet{},
		charactersList:        charactersList,
		planetsList:           planetsList,
		charactersState:       StateIdle,
		planetsState:          StateIdle,
		spinner:               s,
		currentCharactersPage: 1,
		currentPlanetsPage:    1,
		itemsPerPage:          20,
	}
}

// characterItem implements list.Item for Character
type characterItem struct {
	character api.Character
}

func (i characterItem) FilterValue() string {
	return i.character.Name
}

func (i characterItem) Title() string {
	return i.character.Name
}

func (i characterItem) Description() string {
	race := i.character.Race
	if race == "" {
		race = "Unknown"
	}
	affiliation := i.character.Affiliation
	if affiliation == "" {
		affiliation = "Unknown"
	}
	return race + " • " + affiliation
}

// planetItem implements list.Item for Planet
type planetItem struct {
	planet api.Planet
}

func (i planetItem) FilterValue() string {
	return i.planet.Name
}

func (i planetItem) Title() string {
	return i.planet.Name
}

func (i planetItem) Description() string {
	status := "Active"
	if i.planet.IsDestroyed {
		status = "Destroyed"
	}
	return "Status: " + status
}
