package api

// Character represents a Dragon Ball character from the API
type Character struct {
	ID              int              `json:"id"`
	Name            string           `json:"name"`
	Ki              string           `json:"ki"`
	MaxKi           string           `json:"maxKi"`
	Race            string           `json:"race"`
	Gender          string           `json:"gender"`
	Description     string           `json:"description"`
	Image           string           `json:"image"`
	Affiliation     string           `json:"affiliation"`
	DeletedAt       interface{}      `json:"deletedAt"`
	OriginPlanet    *OriginPlanet    `json:"originPlanet,omitempty"`
	Transformations []Transformation `json:"transformations,omitempty"`
}

// Planet represents a Dragon Ball planet from the API
type Planet struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	IsDestroyed bool        `json:"isDestroyed"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	DeletedAt   interface{} `json:"deletedAt"`
}

// OriginPlanet represents the origin planet of a character
type OriginPlanet struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	IsDestroyed bool   `json:"isDestroyed"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

// Transformation represents a character transformation
type Transformation struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Image     string      `json:"image"`
	Ki        string      `json:"ki"`
	DeletedAt interface{} `json:"deletedAt"`
}

// Meta represents pagination metadata
type Meta struct {
	TotalItems   int `json:"totalItems"`
	ItemCount    int `json:"itemCount"`
	ItemsPerPage int `json:"itemsPerPage"`
	TotalPages   int `json:"totalPages"`
	CurrentPage  int `json:"currentPage"`
}

// Links represents pagination links
type Links struct {
	First    string `json:"first"`
	Previous string `json:"previous"`
	Next     string `json:"next"`
	Last     string `json:"last"`
}

// CharactersResponse represents the response from the characters endpoint
type CharactersResponse struct {
	Items []Character `json:"items"`
	Meta  Meta        `json:"meta"`
	Links Links       `json:"links"`
}

// PlanetsResponse represents the response from the planets endpoint
type PlanetsResponse struct {
	Items []Planet `json:"items"`
	Meta  Meta     `json:"meta"`
	Links Links    `json:"links"`
}
