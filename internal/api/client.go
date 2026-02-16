package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	// BaseURL is the base URL for the Dragon Ball API
	BaseURL = "https://dragonball-api.com/api"
)

// Client represents the Dragon Ball API client
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// NewClient creates a new Dragon Ball API client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: BaseURL,
	}
}

// GetCharacters fetches all characters with optional pagination
func (c *Client) GetCharacters(page, limit int) (*CharactersResponse, error) {
	url := fmt.Sprintf("%s/characters?page=%d&limit=%d", c.baseURL, page, limit)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch characters: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var charactersResp CharactersResponse
	if err := json.NewDecoder(resp.Body).Decode(&charactersResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &charactersResp, nil
}

// GetPlanets fetches all planets with optional pagination
func (c *Client) GetPlanets(page, limit int) (*PlanetsResponse, error) {
	url := fmt.Sprintf("%s/planets?page=%d&limit=%d", c.baseURL, page, limit)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch planets: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var planetsResp PlanetsResponse
	if err := json.NewDecoder(resp.Body).Decode(&planetsResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &planetsResp, nil
}

// GetCharacter fetches a single character by ID
func (c *Client) GetCharacter(id int) (*Character, error) {
	url := fmt.Sprintf("%s/characters/%d", c.baseURL, id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch character: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var character Character
	if err := json.NewDecoder(resp.Body).Decode(&character); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &character, nil
}

// GetPlanet fetches a single planet by ID
func (c *Client) GetPlanet(id int) (*Planet, error) {
	url := fmt.Sprintf("%s/planets/%d", c.baseURL, id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch planet: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var planet Planet
	if err := json.NewDecoder(resp.Body).Decode(&planet); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &planet, nil
}
