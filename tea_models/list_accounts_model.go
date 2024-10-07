package teamodels

import (
	"github.com/arimotearipo/ggmp/cmd"
	tea "github.com/charmbracelet/bubbletea"
)

type ListAccountsModel struct {
	cmd       *cmd.Command
	menuIdx   int
	menuItems []string
}

func NewListAccountsModel(c *cmd.Command) *ListAccountsModel {
	return &ListAccountsModel{
		cmd:       c,
		menuItems: append(c.ListAccounts(), "BACK"),
	}
}

func (m *ListAccountsModel) Init() tea.Cmd {
	return nil
}

func (m *ListAccountsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			return m, nil
		}
	}
	return m, nil
}

func (m *ListAccountsModel) View() string {
	s := ""
	for idx, item := range m.menuItems {
		if idx == m.menuIdx {
			s += "👉 " + item + "\n"
		} else {
			s += "   " + item + "\n"
		}
	}
	return s
}
