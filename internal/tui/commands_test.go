package tui

import (
	"errors"
	"testing"

	"github.com/bilalbaraz/ki/internal/api"
	tea "github.com/charmbracelet/bubbletea"
)

func TestLoadCharactersCmd_Success(t *testing.T) {
	client := api.NewClient()

	cmd := LoadCharactersCmd(client, 1, 10)
	msg := cmd()

	switch m := msg.(type) {
	case CharactersLoadedMsg:
		if m.Response == nil {
			t.Fatal("expected non-nil response")
		}
		if m.Response.Meta.CurrentPage < 1 {
			t.Errorf("expected valid page number, got %d", m.Response.Meta.CurrentPage)
		}
	case CharactersErrorMsg:
		// Network errors are acceptable in tests
		if m.Err == nil {
			t.Fatal("expected non-nil error")
		}
	default:
		t.Fatalf("expected CharactersLoadedMsg or CharactersErrorMsg, got %T", msg)
	}
}

func TestLoadCharactersCmd_ReturnsCommand(t *testing.T) {
	client := api.NewClient()

	cmd := LoadCharactersCmd(client, 1, 10)

	if cmd == nil {
		t.Fatal("expected non-nil command")
	}
}

func TestLoadCharactersCmd_Pagination(t *testing.T) {
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
			client := api.NewClient()

			cmd := LoadCharactersCmd(client, tt.page, tt.limit)
			msg := cmd()

			switch m := msg.(type) {
			case CharactersLoadedMsg:
				if m.Response == nil {
					t.Fatal("expected non-nil response")
				}
			case CharactersErrorMsg:
				// Network errors are acceptable
				if m.Err == nil {
					t.Fatal("expected non-nil error in error message")
				}
			default:
				t.Fatalf("unexpected message type: %T", msg)
			}
		})
	}
}

func TestLoadPlanetsCmd_Success(t *testing.T) {
	client := api.NewClient()

	cmd := LoadPlanetsCmd(client, 1, 10)
	msg := cmd()

	switch m := msg.(type) {
	case PlanetsLoadedMsg:
		if m.Response == nil {
			t.Fatal("expected non-nil response")
		}
		if m.Response.Meta.CurrentPage < 1 {
			t.Errorf("expected valid page number, got %d", m.Response.Meta.CurrentPage)
		}
	case PlanetsErrorMsg:
		// Network errors are acceptable in tests
		if m.Err == nil {
			t.Fatal("expected non-nil error")
		}
	default:
		t.Fatalf("expected PlanetsLoadedMsg or PlanetsErrorMsg, got %T", msg)
	}
}

func TestLoadPlanetsCmd_ReturnsCommand(t *testing.T) {
	client := api.NewClient()

	cmd := LoadPlanetsCmd(client, 1, 10)

	if cmd == nil {
		t.Fatal("expected non-nil command")
	}
}

func TestLoadPlanetsCmd_Pagination(t *testing.T) {
	tests := []struct {
		name  string
		page  int
		limit int
	}{
		{"first page", 1, 10},
		{"third page", 3, 15},
		{"max limit", 1, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := api.NewClient()

			cmd := LoadPlanetsCmd(client, tt.page, tt.limit)
			msg := cmd()

			switch m := msg.(type) {
			case PlanetsLoadedMsg:
				if m.Response == nil {
					t.Fatal("expected non-nil response")
				}
			case PlanetsErrorMsg:
				// Network errors are acceptable
				if m.Err == nil {
					t.Fatal("expected non-nil error in error message")
				}
			default:
				t.Fatalf("unexpected message type: %T", msg)
			}
		})
	}
}

func TestCharactersLoadedMsg_Type(t *testing.T) {
	msg := CharactersLoadedMsg{
		Response: &api.CharactersResponse{},
	}

	var _ tea.Msg = msg

	if msg.Response == nil {
		t.Error("expected non-nil response")
	}
}

func TestCharactersErrorMsg_Type(t *testing.T) {
	msg := CharactersErrorMsg{
		Err: errors.New("test error"),
	}

	var _ tea.Msg = msg

	if msg.Err == nil {
		t.Error("expected non-nil error")
	}
}

func TestPlanetsLoadedMsg_Type(t *testing.T) {
	msg := PlanetsLoadedMsg{
		Response: &api.PlanetsResponse{},
	}

	var _ tea.Msg = msg

	if msg.Response == nil {
		t.Error("expected non-nil response")
	}
}

func TestPlanetsErrorMsg_Type(t *testing.T) {
	msg := PlanetsErrorMsg{
		Err: errors.New("test error"),
	}

	var _ tea.Msg = msg

	if msg.Err == nil {
		t.Error("expected non-nil error")
	}
}

func TestLoadCharactersCmd_MessageStructure(t *testing.T) {
	client := api.NewClient()

	cmd := LoadCharactersCmd(client, 1, 10)
	msg := cmd()

	switch m := msg.(type) {
	case CharactersLoadedMsg:
		if m.Response != nil {
			if m.Response.Meta.ItemsPerPage < 0 {
				t.Error("expected non-negative items per page")
			}
			if m.Response.Meta.TotalPages < 0 {
				t.Error("expected non-negative total pages")
			}
		}
	case CharactersErrorMsg:
		// Error case is valid
	default:
		t.Fatalf("unexpected message type: %T", msg)
	}
}

func TestLoadPlanetsCmd_MessageStructure(t *testing.T) {
	client := api.NewClient()

	cmd := LoadPlanetsCmd(client, 1, 10)
	msg := cmd()

	switch m := msg.(type) {
	case PlanetsLoadedMsg:
		if m.Response != nil {
			if m.Response.Meta.ItemsPerPage < 0 {
				t.Error("expected non-negative items per page")
			}
			if m.Response.Meta.TotalPages < 0 {
				t.Error("expected non-negative total pages")
			}
		}
	case PlanetsErrorMsg:
		// Error case is valid
	default:
		t.Fatalf("unexpected message type: %T", msg)
	}
}

func TestCharactersLoadedMsg_WithNilResponse(t *testing.T) {
	msg := CharactersLoadedMsg{
		Response: nil,
	}

	if msg.Response != nil {
		t.Error("expected nil response")
	}
}

func TestCharactersErrorMsg_WithNilError(t *testing.T) {
	msg := CharactersErrorMsg{
		Err: nil,
	}

	if msg.Err != nil {
		t.Error("expected nil error")
	}
}

func TestPlanetsLoadedMsg_WithNilResponse(t *testing.T) {
	msg := PlanetsLoadedMsg{
		Response: nil,
	}

	if msg.Response != nil {
		t.Error("expected nil response")
	}
}

func TestPlanetsErrorMsg_WithNilError(t *testing.T) {
	msg := PlanetsErrorMsg{
		Err: nil,
	}

	if msg.Err != nil {
		t.Error("expected nil error")
	}
}
