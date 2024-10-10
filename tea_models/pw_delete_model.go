package teamodels

import (
	"github.com/arimotearipo/ggmp/action"
	tea "github.com/charmbracelet/bubbletea"
)

type PasswordDeleteModel struct {
	action    *action.Action
	menuItems []string
	menuIdx   int
}

func NewPasswordDeleteModel(a *action.Action) *PasswordDeleteModel {
	uris, _ := a.ListURIs()

	return &PasswordDeleteModel{
		action:    a,
		menuItems: append(uris, "BACK"),
		menuIdx:   0,
	}
}

func (m *PasswordDeleteModel) Init() tea.Cmd {
	return nil
}

func (m *PasswordDeleteModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				err := m.action.DeletePassword(selected)
				if err == nil {
					m.menuItems = append(m.menuItems[:m.menuIdx], m.menuItems[m.menuIdx+1:]...)
				}
			}
		}
	}
	return m, nil
}

func (m *PasswordDeleteModel) View() string {
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
