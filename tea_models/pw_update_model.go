package teamodels

import (
	"github.com/arimotearipo/ggmp/action"
	tea "github.com/charmbracelet/bubbletea"
)

type UpdatePasswordModel struct {
	action    *action.Action
	menuItems []string
	menuIdx   int
}

func NewUpdatePasswordModel(a *action.Action) *UpdatePasswordModel {
	uris, _ := a.ListURIs()

	return &UpdatePasswordModel{
		action:    a,
		menuItems: append(uris, "BACK"),
		menuIdx:   0,
	}
}

func (m *UpdatePasswordModel) Init() tea.Cmd {
	return nil
}

func (m *UpdatePasswordModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "up", "down":
			if msg.String() == "up" {
				m.menuIdx = (m.menuIdx - 1 + len(m.menuItems)) % len(m.menuItems)
			} else if msg.String() == "down" {
				m.menuIdx = (m.menuIdx + 1) % len(m.menuItems)
			}
		case "enter":
			selected := m.menuItems[m.menuIdx]
			switch selected {
			case "BACK":
				return NewPasswordMenuModel(m.action), nil
			default:
				// TODO: update password
			}
		}
	}
	return m, nil
}

func (m *UpdatePasswordModel) View() string {
	s := ""
	for i, item := range m.menuItems {
		if i == m.menuIdx {
			s += "👉 " + item + "\n"
		} else {
			s += "   " + item + "\n"
		}
	}

	return s
}
