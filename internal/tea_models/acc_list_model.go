package teamodels

import (
	"github.com/arimotearipo/ggmp/internal/action"
	tea "github.com/charmbracelet/bubbletea"
)

type AccountsListModel struct {
	cmd       *action.Action
	menuIdx   int
	menuItems []string
}

func NewAccountsListModel(a *action.Action) *AccountsListModel {
	accounts, _ := a.ListMasterAccounts()

	return &AccountsListModel{
		cmd:       a,
		menuItems: append(accounts, "BACK"),
	}
}

func (m *AccountsListModel) Init() tea.Cmd {
	return nil
}

func (m *AccountsListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m *AccountsListModel) View() string {
	s := "Listing registered accounts\n"
	for idx, item := range m.menuItems {
		if idx == m.menuIdx {
			s += "👉 " + item + "\n"
		} else {
			s += "   " + item + "\n"
		}
	}
	return s
}
