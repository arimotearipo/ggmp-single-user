package teamodels

import (
	"github.com/arimotearipo/ggmp/internal/action"
	"github.com/arimotearipo/ggmp/internal/types"
	tea "github.com/charmbracelet/bubbletea"
)

type PasswordsListModel struct {
	action    *action.Action
	uris      []types.URI
	selected  int
	operation string
	result    string
}

func NewPasswordsListModel(a *action.Action, o string) *PasswordsListModel {
	uris, _ := a.ListURIs()

	return &PasswordsListModel{
		action:    a,
		uris:      append(uris, types.URI{Id: 0, Uri: "BACK"}),
		selected:  0,
		operation: o,
		result:    "",
	}
}

func (m *PasswordsListModel) handleSelection() tea.Model {
	selectedUri := m.uris[m.selected]

	switch m.operation {
	case "Get password":
		u, p, err := m.action.GetPassword(selectedUri)
		if err != nil {
			m.result = err.Error()
		} else {
			m.result = "Username: " + u + "\nPassword: " + p
		}
	case "Delete password":
		return NewPasswordDeleteModel(m.action, selectedUri)
	case "Update password":
		return NewPasswordUpdateModel(m.action, selectedUri)
	}

	return m
}

func (m *PasswordsListModel) Init() tea.Cmd {
	return nil
}

func (m *PasswordsListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "up", "down":
			if msg.String() == "up" {
				m.selected = (m.selected - 1 + len(m.uris)) % len(m.uris)
			} else if msg.String() == "down" {
				m.selected = (m.selected + 1) % len(m.uris)
			}
		case "enter":
			selected := m.uris[m.selected].Uri
			switch selected {
			case "BACK":
				return NewPasswordMenuModel(m.action), nil
			default:
				return m.handleSelection(), nil
			}
		}
	}
	return m, nil
}

func (m *PasswordsListModel) View() string {
	s := "Listing saved login details\n"
	for i, uri := range m.uris {
		if i == m.selected {
			s += "👉 " + uri.Uri + "\n"
		} else {
			s += "   " + uri.Uri + "\n"
		}
	}
	s += m.result
	return s
}
