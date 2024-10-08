package teamodels

import (
	"github.com/arimotearipo/ggmp/action"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// DeletingAccountModel is when you already select the account to delete
type DeletingAccountModel struct {
	cmd       *action.Action
	menuIdx   int
	menuItems []string
	user      string
	password  textinput.Model
}

func NewDeletingAccountModel(c *action.Action, u string) *DeletingAccountModel {
	passwordInput := textinput.New()
	passwordInput.Placeholder = "Enter password"
	passwordInput.EchoMode = textinput.EchoPassword
	passwordInput.Focus()

	return &DeletingAccountModel{
		cmd:       c,
		user:      u,
		menuIdx:   0,
		menuItems: []string{"Password", "SUBMIT", "BACK"},
		password:  passwordInput,
	}
}

func (m *DeletingAccountModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *DeletingAccountModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		selected := m.menuItems[m.menuIdx]
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "up", "down":
			if msg.String() == "up" {
				m.menuIdx = (m.menuIdx - 1 + len(m.menuItems)) % len(m.menuItems)
			} else {
				m.menuIdx = (m.menuIdx + 1) % len(m.menuItems)
			}
		case "enter":
			switch selected {
			case "BACK":
				return NewDeleteAccountModel(m.cmd), nil
			case "SUBMIT":
				// TODO: handle submit
			}
		}
	}

	if m.menuIdx == 0 {
		m.password, cmd = m.password.Update(msg)
	}
	return m, cmd
}

func (m *DeletingAccountModel) View() string {
	s := ""
	for idx, item := range m.menuItems {
		if idx == m.menuIdx {
			s += "👉 "
		} else {
			s += "   "
		}

		if item == "Password" {
			s += m.password.View() + "\n"
		} else {
			s += item + "\n"
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
