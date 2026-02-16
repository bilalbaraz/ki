package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := NewClient()

	if client == nil {
		t.Fatal("expected non-nil client")
	}

	if client.httpClient == nil {
		t.Error("expected non-nil http client")
	}

	if client.baseURL != BaseURL {
		t.Errorf("expected baseURL to be %s, got %s", BaseURL, client.baseURL)
	}

	if client.httpClient.Timeout != 10*time.Second {
		t.Errorf("expected timeout to be 10s, got %v", client.httpClient.Timeout)
	}
}

func TestGetCharacters_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/characters" {
			t.Errorf("expected path /characters, got %s", r.URL.Path)
		}

		page := r.URL.Query().Get("page")
		if page != "1" {
			t.Errorf("expected page=1, got %s", page)
		}

		limit := r.URL.Query().Get("limit")
		if limit != "10" {
			t.Errorf("expected limit=10, got %s", limit)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(CharactersResponse{
			Items: []Character{
				{ID: 1, Name: "Goku", Race: "Saiyan"},
			},
			Meta: Meta{
				TotalItems:   1,
				CurrentPage:  1,
				ItemsPerPage: 10,
			},
		})
	}))
	defer server.Close()

	client := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	resp, err := client.GetCharacters(1, 10)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp == nil {
		t.Fatal("expected non-nil response")
	}

	if len(resp.Items) != 1 {
		t.Errorf("expected 1 item, got %d", len(resp.Items))
	}

	if resp.Items[0].Name != "Goku" {
		t.Errorf("expected name Goku, got %s", resp.Items[0].Name)
	}
}

func TestGetCharacters_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	resp, err := client.GetCharacters(1, 10)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if resp != nil {
		t.Error("expected nil response on error")
	}
}

func TestGetCharacters_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json"))
	}))
	defer server.Close()

	client := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	resp, err := client.GetCharacters(1, 10)
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}

	if resp != nil {
		t.Error("expected nil response on decode error")
	}
}

func TestGetPlanets_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/planets" {
			t.Errorf("expected path /planets, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(PlanetsResponse{
			Items: []Planet{
				{ID: 1, Name: "Earth", IsDestroyed: false},
			},
			Meta: Meta{
				TotalItems:   1,
				CurrentPage:  1,
				ItemsPerPage: 10,
			},
		})
	}))
	defer server.Close()

	client := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	resp, err := client.GetPlanets(1, 10)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp == nil {
		t.Fatal("expected non-nil response")
	}

	if len(resp.Items) != 1 {
		t.Errorf("expected 1 item, got %d", len(resp.Items))
	}

	if resp.Items[0].Name != "Earth" {
		t.Errorf("expected name Earth, got %s", resp.Items[0].Name)
	}
}

func TestGetPlanets_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	resp, err := client.GetPlanets(1, 10)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if resp != nil {
		t.Error("expected nil response on error")
	}
}

func TestGetCharacter_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/characters/1" {
			t.Errorf("expected path /characters/1, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Character{
			ID:   1,
			Name: "Goku",
			Race: "Saiyan",
		})
	}))
	defer server.Close()

	client := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	char, err := client.GetCharacter(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if char == nil {
		t.Fatal("expected non-nil character")
	}

	if char.Name != "Goku" {
		t.Errorf("expected name Goku, got %s", char.Name)
	}
}

func TestGetCharacter_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	client := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	char, err := client.GetCharacter(1)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if char != nil {
		t.Error("expected nil character on error")
	}
}

func TestGetPlanet_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/planets/1" {
			t.Errorf("expected path /planets/1, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Planet{
			ID:          1,
			Name:        "Namek",
			IsDestroyed: false,
		})
	}))
	defer server.Close()

	client := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	planet, err := client.GetPlanet(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if planet == nil {
		t.Fatal("expected non-nil planet")
	}

	if planet.Name != "Namek" {
		t.Errorf("expected name Namek, got %s", planet.Name)
	}
}

func TestGetPlanet_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	}))
	defer server.Close()

	client := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	planet, err := client.GetPlanet(1)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if planet != nil {
		t.Error("expected nil planet on error")
	}
}

func TestGetCharacters_Pagination(t *testing.T) {
	tests := []struct {
		name  string
		page  int
		limit int
	}{
		{"first page", 1, 10},
		{"second page", 2, 20},
		{"large limit", 1, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(CharactersResponse{
					Items: []Character{},
					Meta: Meta{
						CurrentPage:  tt.page,
						ItemsPerPage: tt.limit,
					},
				})
			}))
			defer server.Close()

			client := &Client{
				httpClient: server.Client(),
				baseURL:    server.URL,
			}

			resp, err := client.GetCharacters(tt.page, tt.limit)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if resp.Meta.CurrentPage != tt.page {
				t.Errorf("expected page %d, got %d", tt.page, resp.Meta.CurrentPage)
			}
		})
	}
}

func TestGetPlanets_Pagination(t *testing.T) {
	tests := []struct {
		name  string
		page  int
		limit int
	}{
		{"first page", 1, 10},
		{"third page", 3, 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(PlanetsResponse{
					Items: []Planet{},
					Meta: Meta{
						CurrentPage:  tt.page,
						ItemsPerPage: tt.limit,
					},
				})
			}))
			defer server.Close()

			client := &Client{
				httpClient: server.Client(),
				baseURL:    server.URL,
			}

			resp, err := client.GetPlanets(tt.page, tt.limit)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if resp.Meta.CurrentPage != tt.page {
				t.Errorf("expected page %d, got %d", tt.page, resp.Meta.CurrentPage)
			}
		})
	}
}

func TestGetCharacter_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{invalid}"))
	}))
	defer server.Close()

	client := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	char, err := client.GetCharacter(1)
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}

	if char != nil {
		t.Error("expected nil character on decode error")
	}
}

func TestGetPlanet_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("not json"))
	}))
	defer server.Close()

	client := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	planet, err := client.GetPlanet(1)
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}

	if planet != nil {
		t.Error("expected nil planet on decode error")
	}
}

func TestClient_BaseURL(t *testing.T) {
	client := NewClient()

	if client.baseURL != "https://dragonball-api.com/api" {
		t.Errorf("expected base URL https://dragonball-api.com/api, got %s", client.baseURL)
	}
}

func TestGetCharacters_EmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(CharactersResponse{
			Items: []Character{},
			Meta:  Meta{},
		})
	}))
	defer server.Close()

	client := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	resp, err := client.GetCharacters(1, 10)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(resp.Items) != 0 {
		t.Errorf("expected 0 items, got %d", len(resp.Items))
	}
}

func TestGetPlanets_EmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(PlanetsResponse{
			Items: []Planet{},
			Meta:  Meta{},
		})
	}))
	defer server.Close()

	client := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	resp, err := client.GetPlanets(1, 10)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(resp.Items) != 0 {
		t.Errorf("expected 0 items, got %d", len(resp.Items))
	}
}

func TestGetCharacters_MultipleItems(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(CharactersResponse{
			Items: []Character{
				{ID: 1, Name: "Goku"},
				{ID: 2, Name: "Vegeta"},
				{ID: 3, Name: "Piccolo"},
			},
			Meta: Meta{TotalItems: 3},
		})
	}))
	defer server.Close()

	client := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	resp, err := client.GetCharacters(1, 10)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(resp.Items) != 3 {
		t.Errorf("expected 3 items, got %d", len(resp.Items))
	}
}

func TestGetPlanets_MultipleItems(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(PlanetsResponse{
			Items: []Planet{
				{ID: 1, Name: "Earth"},
				{ID: 2, Name: "Namek"},
			},
			Meta: Meta{TotalItems: 2},
		})
	}))
	defer server.Close()

	client := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	resp, err := client.GetPlanets(1, 10)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(resp.Items) != 2 {
		t.Errorf("expected 2 items, got %d", len(resp.Items))
	}
}

func TestGetCharacters_RequestCreationError(t *testing.T) {
	client := &Client{
		httpClient: &http.Client{},
		baseURL:    ":",
	}

	resp, err := client.GetCharacters(1, 10)
	if err == nil {
		t.Fatal("expected error for invalid URL, got nil")
	}

	if resp != nil {
		t.Error("expected nil response on request creation error")
	}
}

func TestGetPlanets_RequestCreationError(t *testing.T) {
	client := &Client{
		httpClient: &http.Client{},
		baseURL:    ":",
	}

	resp, err := client.GetPlanets(1, 10)
	if err == nil {
		t.Fatal("expected error for invalid URL, got nil")
	}

	if resp != nil {
		t.Error("expected nil response on request creation error")
	}
}

func TestGetCharacter_RequestCreationError(t *testing.T) {
	client := &Client{
		httpClient: &http.Client{},
		baseURL:    ":",
	}

	char, err := client.GetCharacter(1)
	if err == nil {
		t.Fatal("expected error for invalid URL, got nil")
	}

	if char != nil {
		t.Error("expected nil character on request creation error")
	}
}

func TestGetPlanet_RequestCreationError(t *testing.T) {
	client := &Client{
		httpClient: &http.Client{},
		baseURL:    ":",
	}

	planet, err := client.GetPlanet(1)
	if err == nil {
		t.Fatal("expected error for invalid URL, got nil")
	}

	if planet != nil {
		t.Error("expected nil planet on request creation error")
	}
}

func TestGetCharacters_BodyCloseError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", "1")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	_, err := client.GetCharacters(1, 10)
	if err == nil {
		t.Fatal("expected error for incomplete body, got nil")
	}
}

func TestGetPlanets_BodyCloseError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", "1")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	_, err := client.GetPlanets(1, 10)
	if err == nil {
		t.Fatal("expected error for incomplete body, got nil")
	}
}

func TestGetCharacter_BodyCloseError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", "1")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	_, err := client.GetCharacter(1)
	if err == nil {
		t.Fatal("expected error for incomplete body, got nil")
	}
}

func TestGetPlanet_BodyCloseError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", "1")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := &Client{
		httpClient: server.Client(),
		baseURL:    server.URL,
	}

	_, err := client.GetPlanet(1)
	if err == nil {
		t.Fatal("expected error for incomplete body, got nil")
	}
}

func TestGetCharacters_NetworkError(t *testing.T) {
	client := &Client{
		httpClient: &http.Client{
			Timeout: 1 * time.Nanosecond,
		},
		baseURL: "http://unreachable-host-12345.test",
	}

	resp, err := client.GetCharacters(1, 10)
	if err == nil {
		t.Fatal("expected network error, got nil")
	}

	if resp != nil {
		t.Error("expected nil response on network error")
	}
}

func TestGetPlanets_NetworkError(t *testing.T) {
	client := &Client{
		httpClient: &http.Client{
			Timeout: 1 * time.Nanosecond,
		},
		baseURL: "http://unreachable-host-12345.test",
	}

	resp, err := client.GetPlanets(1, 10)
	if err == nil {
		t.Fatal("expected network error, got nil")
	}

	if resp != nil {
		t.Error("expected nil response on network error")
	}
}

func TestGetCharacter_NetworkError(t *testing.T) {
	client := &Client{
		httpClient: &http.Client{
			Timeout: 1 * time.Nanosecond,
		},
		baseURL: "http://unreachable-host-12345.test",
	}

	char, err := client.GetCharacter(1)
	if err == nil {
		t.Fatal("expected network error, got nil")
	}

	if char != nil {
		t.Error("expected nil character on network error")
	}
}

func TestGetPlanet_NetworkError(t *testing.T) {
	client := &Client{
		httpClient: &http.Client{
			Timeout: 1 * time.Nanosecond,
		},
		baseURL: "http://unreachable-host-12345.test",
	}

	planet, err := client.GetPlanet(1)
	if err == nil {
		t.Fatal("expected network error, got nil")
	}

	if planet != nil {
		t.Error("expected nil planet on network error")
	}
}
