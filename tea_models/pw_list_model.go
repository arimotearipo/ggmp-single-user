package teamodels

import (
	"github.com/arimotearipo/ggmp/action"
	tea "github.com/charmbracelet/bubbletea"
)

type URI = string

type PasswordsListModel struct {
	action    *action.Action
	uris      []URI
	selection int
}

func NewPasswordsListModel(a *action.Action, prevModel, nextModel tea.Model) *PasswordsListModel {
	uris, _ := a.ListURIs()
	return &PasswordsListModel{
		action:    a,
		uris:      append(uris, "BACK"),
		selection: 0,
	}
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
				m.selection = (m.selection - 1 + len(m.uris)) % len(m.uris)
			} else if msg.String() == "down" {
				m.selection = (m.selection + 1) % len(m.uris)
			}
		case "enter":
			selected := m.uris[m.selection]
			switch selected {
			case "BACK":
				return m, nil
			default:
				return m, nil
			}
		}
	}
	return m, nil
}

func (m *PasswordsListModel) View() string {
	s := ""

	for i, uri := range m.uris {
		if i == m.selection {
			s += "👉 " + uri + "\n"
		} else {
			s += "   " + uri + "\n"
		}
	}
	return s
}
