package teamodels

import (
	"github.com/arimotearipo/ggmp/cmd"
	tea "github.com/charmbracelet/bubbletea"
)

type AuthMenuModel struct {
	cmd       *cmd.Command
	menuIdx   int
	menuItems []string
}

func NewAuthMenuModel(c *cmd.Command) *AuthMenuModel {
	return &AuthMenuModel{
		cmd:       c,
		menuIdx:   0,
		menuItems: []string{"Login", "Register", "List accounts", "Delete account", "Exit"},
	}
}

func (m *AuthMenuModel) Init() tea.Cmd {
	return nil
}

func (m *AuthMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			m.menuIdx = (m.menuIdx - 1 + len(m.menuItems)) % len(m.menuItems)
		case "down":
			m.menuIdx = (m.menuIdx + 1) % len(m.menuItems)
		case "enter":
			selectedAction := m.menuItems[m.menuIdx]

			switch selectedAction {
			case "Login":
				return NewLoginModel(m.cmd), nil
			case "Register":
				return NewRegisterModel(m.cmd), nil
			case "List accounts":
				return NewListAccountsModel(m.cmd), nil
			case "Delete account":
				return NewDeleteAccountModel(m.cmd), nil
			case "Exit":
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
