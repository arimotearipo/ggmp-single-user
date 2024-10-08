package teamodels

import (
	"fmt"
	"strings"

	"github.com/arimotearipo/ggmp/action"
	tea "github.com/charmbracelet/bubbletea"
)

// DeletingAccountModel is when you already select the account to delete
type DeletingAccountModel struct {
	cmd       *action.Action
	menuIdx   int
	menuItems []string
	user      string
	password  string
}

func NewDeletingAccountModel(c *action.Action, u string) *DeletingAccountModel {
	return &DeletingAccountModel{
		cmd:       c,
		user:      u,
		menuIdx:   0,
		menuItems: []string{"Enter password: ", "BACK"},
		password:  "",
	}
}

func (m *DeletingAccountModel) Init() tea.Cmd {
	return nil
}

func (m *DeletingAccountModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if len(m.password) == 0 {
				return m, nil
			}
			m.password = m.password[:len(m.password)-1]
		case "enter":
			selected := m.menuItems[m.menuIdx]
			if selected == "BACK" {
				return NewDeleteAccountModel(m.cmd), nil
			}
		default:
			m.password += "*"
		}
	}
	return m, nil

}

func (m *DeletingAccountModel) View() string {
	blurredPassword := strings.Repeat(string("*"), len(m.password))
	s := fmt.Sprintf("Deleting account: %s\n", m.user)
	for idx, item := range m.menuItems {
		if idx == m.menuIdx {
			s += "👉 " + item + " " + blurredPassword + "\n"
		} else {
			s += "   " + item + "\n"
		}
	}
	return s
}

// DeleteAccountModel is the menu for deleting account

type DeleteAccountModel struct {
	cmd       *action.Action
	menuItems []string
	menuIdx   int
}

func NewDeleteAccountModel(c *action.Action) *DeleteAccountModel {
	return &DeleteAccountModel{
		cmd:       c,
		menuItems: append(c.ListAccounts(), "BACK"),
		menuIdx:   0,
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
			selected := m.menuItems[m.menuIdx]
			if selected == "BACK" {
				return NewAuthMenuModel(m.cmd), nil
			}
			// TODO: Prompt for password before deleting
			return NewDeletingAccountModel(m.cmd, selected), nil
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
