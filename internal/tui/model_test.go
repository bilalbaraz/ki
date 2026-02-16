package tui

import (
	"testing"

	"github.com/bilalbaraz/ki/internal/api"
	"github.com/charmbracelet/bubbles/list"
)

func TestNewModel(t *testing.T) {
	m := NewModel()

	if m.client == nil {
		t.Error("expected non-nil client")
	}

	if m.activeTab != CharactersTab {
		t.Errorf("expected activeTab to be CharactersTab, got %v", m.activeTab)
	}

	if m.characters == nil {
		t.Error("expected non-nil characters slice")
	}

	if m.planets == nil {
		t.Error("expected non-nil planets slice")
	}

	if m.charactersState != StateIdle {
		t.Errorf("expected charactersState to be StateIdle, got %v", m.charactersState)
	}

	if m.planetsState != StateIdle {
		t.Errorf("expected planetsState to be StateIdle, got %v", m.planetsState)
	}

	if m.currentCharactersPage != 1 {
		t.Errorf("expected currentCharactersPage to be 1, got %d", m.currentCharactersPage)
	}

	if m.currentPlanetsPage != 1 {
		t.Errorf("expected currentPlanetsPage to be 1, got %d", m.currentPlanetsPage)
	}

	if m.itemsPerPage != 20 {
		t.Errorf("expected itemsPerPage to be 20, got %d", m.itemsPerPage)
	}

	if m.ready {
		t.Error("expected ready to be false initially")
	}
}

func TestNewModel_ListsInitialized(t *testing.T) {
	m := NewModel()

	if len(m.charactersList.Items()) != 0 {
		t.Error("expected empty characters list initially")
	}

	if len(m.planetsList.Items()) != 0 {
		t.Error("expected empty planets list initially")
	}
}

func TestCharacterItem_FilterValue(t *testing.T) {
	char := api.Character{
		ID:   1,
		Name: "Goku",
		Race: "Saiyan",
	}
	item := characterItem{character: char}

	if item.FilterValue() != "Goku" {
		t.Errorf("expected FilterValue to be 'Goku', got '%s'", item.FilterValue())
	}
}

func TestCharacterItem_Title(t *testing.T) {
	char := api.Character{
		ID:   1,
		Name: "Vegeta",
		Race: "Saiyan",
	}
	item := characterItem{character: char}

	if item.Title() != "Vegeta" {
		t.Errorf("expected Title to be 'Vegeta', got '%s'", item.Title())
	}
}

func TestCharacterItem_Description(t *testing.T) {
	tests := []struct {
		name        string
		char        api.Character
		expectedSub string
	}{
		{
			name: "with race and affiliation",
			char: api.Character{
				Race:        "Saiyan",
				Affiliation: "Z Fighter",
			},
			expectedSub: "Saiyan • Z Fighter",
		},
		{
			name: "empty race",
			char: api.Character{
				Race:        "",
				Affiliation: "Z Fighter",
			},
			expectedSub: "Unknown • Z Fighter",
		},
		{
			name: "empty affiliation",
			char: api.Character{
				Race:        "Namekian",
				Affiliation: "",
			},
			expectedSub: "Namekian • Unknown",
		},
		{
			name: "both empty",
			char: api.Character{
				Race:        "",
				Affiliation: "",
			},
			expectedSub: "Unknown • Unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := characterItem{character: tt.char}
			desc := item.Description()

			if desc != tt.expectedSub {
				t.Errorf("expected Description to be '%s', got '%s'", tt.expectedSub, desc)
			}
		})
	}
}

func TestCharacterItem_ListItem(t *testing.T) {
	char := api.Character{
		ID:   1,
		Name: "Piccolo",
		Race: "Namekian",
	}
	item := characterItem{character: char}

	var _ list.Item = item
}

func TestPlanetItem_FilterValue(t *testing.T) {
	planet := api.Planet{
		ID:   1,
		Name: "Earth",
	}
	item := planetItem{planet: planet}

	if item.FilterValue() != "Earth" {
		t.Errorf("expected FilterValue to be 'Earth', got '%s'", item.FilterValue())
	}
}

func TestPlanetItem_Title(t *testing.T) {
	planet := api.Planet{
		ID:   2,
		Name: "Namek",
	}
	item := planetItem{planet: planet}

	if item.Title() != "Namek" {
		t.Errorf("expected Title to be 'Namek', got '%s'", item.Title())
	}
}

func TestPlanetItem_Description(t *testing.T) {
	tests := []struct {
		name        string
		planet      api.Planet
		expectedSub string
	}{
		{
			name: "planet not destroyed",
			planet: api.Planet{
				IsDestroyed: false,
			},
			expectedSub: "Status: Active",
		},
		{
			name: "planet destroyed",
			planet: api.Planet{
				IsDestroyed: true,
			},
			expectedSub: "Status: Destroyed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := planetItem{planet: tt.planet}
			desc := item.Description()

			if desc != tt.expectedSub {
				t.Errorf("expected Description to be '%s', got '%s'", tt.expectedSub, desc)
			}
		})
	}
}

func TestPlanetItem_ListItem(t *testing.T) {
	planet := api.Planet{
		ID:   1,
		Name: "Vegeta",
	}
	item := planetItem{planet: planet}

	var _ list.Item = item
}

func TestTabConstants(t *testing.T) {
	if CharactersTab != 0 {
		t.Errorf("expected CharactersTab to be 0, got %d", CharactersTab)
	}

	if PlanetsTab != 1 {
		t.Errorf("expected PlanetsTab to be 1, got %d", PlanetsTab)
	}
}

func TestLoadingStateConstants(t *testing.T) {
	if StateIdle != 0 {
		t.Errorf("expected StateIdle to be 0, got %d", StateIdle)
	}

	if StateLoading != 1 {
		t.Errorf("expected StateLoading to be 1, got %d", StateLoading)
	}

	if StateLoaded != 2 {
		t.Errorf("expected StateLoaded to be 2, got %d", StateLoaded)
	}

	if StateError != 3 {
		t.Errorf("expected StateError to be 3, got %d", StateError)
	}
}

func TestModel_Fields(t *testing.T) {
	m := NewModel()

	m.width = 100
	m.height = 50
	m.ready = true

	if m.width != 100 {
		t.Errorf("expected width to be 100, got %d", m.width)
	}

	if m.height != 50 {
		t.Errorf("expected height to be 50, got %d", m.height)
	}

	if !m.ready {
		t.Error("expected ready to be true")
	}
}

func TestModel_TabSwitching(t *testing.T) {
	m := NewModel()

	if m.activeTab != CharactersTab {
		t.Errorf("expected initial tab to be CharactersTab, got %d", m.activeTab)
	}

	m.activeTab = PlanetsTab
	if m.activeTab != PlanetsTab {
		t.Errorf("expected tab to be PlanetsTab, got %d", m.activeTab)
	}

	m.activeTab = CharactersTab
	if m.activeTab != CharactersTab {
		t.Errorf("expected tab to be CharactersTab, got %d", m.activeTab)
	}
}

func TestModel_StateTransitions(t *testing.T) {
	m := NewModel()

	m.charactersState = StateLoading
	if m.charactersState != StateLoading {
		t.Errorf("expected charactersState to be StateLoading, got %d", m.charactersState)
	}

	m.charactersState = StateLoaded
	if m.charactersState != StateLoaded {
		t.Errorf("expected charactersState to be StateLoaded, got %d", m.charactersState)
	}

	m.charactersState = StateError
	if m.charactersState != StateError {
		t.Errorf("expected charactersState to be StateError, got %d", m.charactersState)
	}

	m.planetsState = StateLoading
	if m.planetsState != StateLoading {
		t.Errorf("expected planetsState to be StateLoading, got %d", m.planetsState)
	}
}

func TestModel_ErrorMessages(t *testing.T) {
	m := NewModel()

	m.charactersError = "network error"
	if m.charactersError != "network error" {
		t.Errorf("expected charactersError to be 'network error', got '%s'", m.charactersError)
	}

	m.planetsError = "api error"
	if m.planetsError != "api error" {
		t.Errorf("expected planetsError to be 'api error', got '%s'", m.planetsError)
	}
}

func TestModel_Pagination(t *testing.T) {
	m := NewModel()

	m.currentCharactersPage = 5
	if m.currentCharactersPage != 5 {
		t.Errorf("expected currentCharactersPage to be 5, got %d", m.currentCharactersPage)
	}

	m.currentPlanetsPage = 3
	if m.currentPlanetsPage != 3 {
		t.Errorf("expected currentPlanetsPage to be 3, got %d", m.currentPlanetsPage)
	}

	m.itemsPerPage = 50
	if m.itemsPerPage != 50 {
		t.Errorf("expected itemsPerPage to be 50, got %d", m.itemsPerPage)
	}
}

func TestModel_DataStorage(t *testing.T) {
	m := NewModel()

	m.characters = []api.Character{
		{ID: 1, Name: "Goku"},
		{ID: 2, Name: "Vegeta"},
	}

	if len(m.characters) != 2 {
		t.Errorf("expected 2 characters, got %d", len(m.characters))
	}

	m.planets = []api.Planet{
		{ID: 1, Name: "Earth"},
	}

	if len(m.planets) != 1 {
		t.Errorf("expected 1 planet, got %d", len(m.planets))
	}
}

func TestModel_MetaData(t *testing.T) {
	m := NewModel()

	m.charactersMeta = api.Meta{
		TotalItems:   100,
		ItemCount:    20,
		ItemsPerPage: 20,
		TotalPages:   5,
		CurrentPage:  1,
	}

	if m.charactersMeta.TotalItems != 100 {
		t.Errorf("expected TotalItems to be 100, got %d", m.charactersMeta.TotalItems)
	}

	if m.charactersMeta.TotalPages != 5 {
		t.Errorf("expected TotalPages to be 5, got %d", m.charactersMeta.TotalPages)
	}

	m.planetsMeta = api.Meta{
		TotalItems:   50,
		ItemCount:    10,
		ItemsPerPage: 10,
		TotalPages:   5,
		CurrentPage:  1,
	}

	if m.planetsMeta.TotalItems != 50 {
		t.Errorf("expected TotalItems to be 50, got %d", m.planetsMeta.TotalItems)
	}
}

func TestCharacterItem_EmptyName(t *testing.T) {
	char := api.Character{
		ID:   1,
		Name: "",
	}
	item := characterItem{character: char}

	if item.Title() != "" {
		t.Errorf("expected empty Title, got '%s'", item.Title())
	}

	if item.FilterValue() != "" {
		t.Errorf("expected empty FilterValue, got '%s'", item.FilterValue())
	}
}

func TestPlanetItem_EmptyName(t *testing.T) {
	planet := api.Planet{
		ID:   1,
		Name: "",
	}
	item := planetItem{planet: planet}

	if item.Title() != "" {
		t.Errorf("expected empty Title, got '%s'", item.Title())
	}

	if item.FilterValue() != "" {
		t.Errorf("expected empty FilterValue, got '%s'", item.FilterValue())
	}
}

func TestCharacterItem_SpecialCharactersInFields(t *testing.T) {
	char := api.Character{
		ID:          1,
		Name:        "Son Goku",
		Race:        "Saiyan-Human",
		Affiliation: "Z-Fighter",
	}
	item := characterItem{character: char}

	if item.Title() != "Son Goku" {
		t.Errorf("expected Title 'Son Goku', got '%s'", item.Title())
	}

	desc := item.Description()
	if desc != "Saiyan-Human • Z-Fighter" {
		t.Errorf("expected Description 'Saiyan-Human • Z-Fighter', got '%s'", desc)
	}
}

func TestPlanetItem_AllFields(t *testing.T) {
	planet := api.Planet{
		ID:          1,
		Name:        "Planet Vegeta",
		IsDestroyed: true,
		Description: "Homeworld of the Saiyans",
	}
	item := planetItem{planet: planet}

	if item.Title() != "Planet Vegeta" {
		t.Errorf("expected Title 'Planet Vegeta', got '%s'", item.Title())
	}

	if item.FilterValue() != "Planet Vegeta" {
		t.Errorf("expected FilterValue 'Planet Vegeta', got '%s'", item.FilterValue())
	}

	desc := item.Description()
	if desc != "Status: Destroyed" {
		t.Errorf("expected Description 'Status: Destroyed', got '%s'", desc)
	}
}

func TestModel_MultipleStateChanges(t *testing.T) {
	m := NewModel()

	states := []LoadingState{StateIdle, StateLoading, StateLoaded, StateError, StateIdle}

	for _, state := range states {
		m.charactersState = state
		if m.charactersState != state {
			t.Errorf("expected charactersState to be %d, got %d", state, m.charactersState)
		}
	}
}

func TestModel_SpinnerInitialized(t *testing.T) {
	m := NewModel()

	if m.spinner.Spinner.FPS == 0 {
		t.Error("expected spinner to be initialized with non-zero FPS")
	}
}

func TestCharacterItem_LongFields(t *testing.T) {
	longName := "This is a very long character name that exceeds normal length"
	longRace := "This is a very long race name"
	longAffiliation := "This is a very long affiliation name"

	char := api.Character{
		ID:          1,
		Name:        longName,
		Race:        longRace,
		Affiliation: longAffiliation,
	}
	item := characterItem{character: char}

	if item.Title() != longName {
		t.Error("Title should handle long names")
	}

	desc := item.Description()
	expected := longRace + " • " + longAffiliation
	if desc != expected {
		t.Error("Description should handle long fields")
	}
}

func TestModel_ZeroPagination(t *testing.T) {
	m := NewModel()

	m.currentCharactersPage = 0
	m.currentPlanetsPage = 0

	if m.currentCharactersPage != 0 {
		t.Error("should allow zero page number")
	}

	if m.currentPlanetsPage != 0 {
		t.Error("should allow zero page number")
	}
}

func TestModel_NegativePagination(t *testing.T) {
	m := NewModel()

	m.currentCharactersPage = -1
	m.currentPlanetsPage = -1

	if m.currentCharactersPage != -1 {
		t.Error("should store negative page number")
	}

	if m.currentPlanetsPage != -1 {
		t.Error("should store negative page number")
	}
}
