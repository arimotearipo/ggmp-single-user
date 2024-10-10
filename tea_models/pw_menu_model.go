package teamodels

import (
	"github.com/arimotearipo/ggmp/action"
	tea "github.com/charmbracelet/bubbletea"
)

type PasswordMenuModel struct {
	action    *action.Action
	menuIdx   int
	menuItems []string
}

func NewPasswordMenuModel(c *action.Action) *PasswordMenuModel {
	return &PasswordMenuModel{
		action:    c,
		menuItems: []string{"Get password", "Add password", "Delete password", "Update password", "EXIT"},
	}
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
		case "up":
			m.menuIdx = (m.menuIdx - 1 + len(m.menuItems)) % len(m.menuItems)
		case "down":
			m.menuIdx = (m.menuIdx + 1) % len(m.menuItems)
		case "backspace":
			m.menuItems[m.menuIdx] = m.menuItems[m.menuIdx][:len(m.menuItems[m.menuIdx])-1]
		case "enter":
			selected := m.menuItems[m.menuIdx]
			switch selected {
			case "Add password":
				return NewPasswordAddModel(m.action), nil
			case "Get password":
				return NewPasswordGetModel(m.action), nil
			case "Delete password":
				return NewPasswordDeleteModel(m.action), nil
			case "Update password":
				return NewPasswordUpdateModel(m.action), nil
			case "EXIT":
				return NewAuthMenuModel(m.action), nil
			}
		}
	}
	return m, nil
}

func (m *PasswordMenuModel) View() string {
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
