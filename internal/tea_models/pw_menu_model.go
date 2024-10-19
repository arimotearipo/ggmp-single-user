package teamodels

import (
	"github.com/arimotearipo/ggmp/internal/action"
	tea "github.com/charmbracelet/bubbletea"
)

type PasswordMenuModel struct {
	action    *action.Action
	menuIdx   int
	menuItems []string
	selected  string
	result    string
}

func NewPasswordMenuModel(a *action.Action) *PasswordMenuModel {
	return &PasswordMenuModel{
		action:    a,
		menuIdx:   0,
		menuItems: []string{"Get password", "Add password", "Update password", "Delete password", "Change master password", "Delete master account", "Encrypt password file", "LOGOUT"},
		result:    "",
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
			case "Change master password":
				return NewAccountChangeMasterPasswordModel(m.action), nil
			case "Delete master account":
				return NewConfirmDeleteAccountModel(m.action), nil
			case "Encrypt password file":
				return NewEncryptFileModel(m.action), nil
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

	s += m.result

	return s
}
