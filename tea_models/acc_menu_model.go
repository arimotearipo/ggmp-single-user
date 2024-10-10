package teamodels

import (
	"github.com/arimotearipo/ggmp/action"
	tea "github.com/charmbracelet/bubbletea"
)

type AuthMenuModel struct {
	action    *action.Action
	menuIdx   int
	menuItems []string
}

func NewAuthMenuModel(a *action.Action) *AuthMenuModel {
	return &AuthMenuModel{
		action:    a,
		menuIdx:   0,
		menuItems: []string{"Login", "Register", "List accounts", "Delete account", "EXIT"},
	}
}

func (m *AuthMenuModel) Init() tea.Cmd {
	return nil
}

func (m *AuthMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "up":
			m.menuIdx = (m.menuIdx - 1 + len(m.menuItems)) % len(m.menuItems)
		case "down":
			m.menuIdx = (m.menuIdx + 1) % len(m.menuItems)
		case "enter":
			selectedAction := m.menuItems[m.menuIdx]

			switch selectedAction {
			case "Login":
				return NewAccountLoginModel(m.action), nil
			case "Register":
				return NewAccountRegisterModel(m.action), nil
			case "List accounts":
				return NewAccountsListModel(m.action), nil
			case "Delete account":
				return NewAccountDeleteModel(m.action), nil
			case "EXIT":
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m *AuthMenuModel) View() string {
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
