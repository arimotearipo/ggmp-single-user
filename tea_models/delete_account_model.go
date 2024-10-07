package teamodels

import (
	"github.com/arimotearipo/ggmp/cmd"
	tea "github.com/charmbracelet/bubbletea"
)

type DeleteAccountModel struct {
	cmd         *cmd.Command
	menuItems   []string
	menuIdx     int
	idxToDelete int
}

func NewDeleteAccountModel(c *cmd.Command) *DeleteAccountModel {
	return &DeleteAccountModel{
		cmd:         c,
		menuItems:   append(c.ListAccounts(), "BACK"),
		menuIdx:     0,
		idxToDelete: -1,
	}
}

func (m *DeleteAccountModel) Init() tea.Cmd {
	return nil
}

func (m *DeleteAccountModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.menuItems[m.menuIdx] == "BACK" {
				return NewAuthMenuModel(m.cmd), nil
			}
			// TODO: Prompt for password before deleting
			m.cmd.Delete(m.menuItems[m.menuIdx], "")
			return m, nil
		}
	}
	return m, nil
}

func (m *DeleteAccountModel) View() string {
	s := "Delete account:\n"
	for idx, item := range m.menuItems {
		if idx == m.menuIdx {
			s += "❌ " + item + "\n"
		} else {
			s += "   " + item + "\n"
		}
	}
	return s
}
