package teamodels

import (
	"github.com/arimotearipo/ggmp/action"
	tea "github.com/charmbracelet/bubbletea"
)

type PasswordMenuModel struct {
	action    *action.Action
	menuIdx   int
	menuItems []string
	selected  string
}

func NewPasswordMenuModel(c *action.Action) *PasswordMenuModel {
	return &PasswordMenuModel{
		action:    c,
		menuIdx:   0,
		menuItems: []string{"Get password", "Add password", "Update password", "Delete password", "LOGOUT"},
	}
}

func (m *PasswordMenuModel) handleSelection() tea.Model {
	selected := m.menuItems[m.menuIdx]
	switch selected {
	case "Get password", "Update password", "Delete password":
		return NewPasswordsListModel(m.action, selected)
	case "Add password":
		return NewPasswordAddModel(m.action)
	}
	return m
}

func (m *PasswordMenuModel) Init() tea.Cmd {
	return nil
}

func (m *PasswordMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			m.selected = m.menuItems[m.menuIdx]
			switch m.selected {
			case "LOGOUT":
				m.action.Logout()
				return NewAuthMenuModel(m.action), nil
			default:
				return m.handleSelection(), nil
			}
		}
	}
	return m, nil
}

func (m *PasswordMenuModel) View() string {
	s := "Select an option\n"
	for i, item := range m.menuItems {
		if i == m.menuIdx {
			s += "👉 " + item + "\n"
		} else {
			s += "   " + item + "\n"
		}
	}
	return s
}
