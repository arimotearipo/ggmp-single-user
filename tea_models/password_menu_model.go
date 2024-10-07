package teamodels

import (
	"github.com/arimotearipo/ggmp/cmd"
	tea "github.com/charmbracelet/bubbletea"
)

type PasswordMenuModel struct {
	cmd       *cmd.Command
	menuIdx   int
	menuItems []string
}

func NewPasswordMenuModel(c *cmd.Command) *PasswordMenuModel {
	return &PasswordMenuModel{
		cmd:       c,
		menuItems: []string{"Get password", "Add password", "List URIs", "Delete password", "Update password", "Exit"},
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
			// do something
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
