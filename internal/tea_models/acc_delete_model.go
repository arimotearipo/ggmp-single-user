package teamodels

import (
	"github.com/arimotearipo/ggmp/internal/action"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ConfirmDeleteAccountModel struct {
	action    *action.Action
	menuIdx   int
	menuItems []string
}

func NewConfirmDeleteAccountModel(a *action.Action) *ConfirmDeleteAccountModel {
	return &ConfirmDeleteAccountModel{
		action:    a,
		menuIdx:   0,
		menuItems: []string{"YES", "NO", "BACK"},
	}
}

func (m *ConfirmDeleteAccountModel) Init() tea.Cmd {
	return nil
}

func (m *ConfirmDeleteAccountModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if selected == "BACK" {
				return NewPasswordMenuModel(m.action), nil
			}

			if selected == "YES" {
				// TODO: encrypt file action
			}
			return NewAccountDeleteModel(m.action), nil
		}
	}

	return m, nil
}

func (m *ConfirmDeleteAccountModel) View() string {
	s := "Would you like to encrypt and backup your data first?\n"
	for i, item := range m.menuItems {
		if m.menuIdx == i {
			s += "👉 " + item + "\n"
		} else {
			s += "   " + item + "\n"
		}
	}
	return s
}

// DeletingAccountModel is when you already select the account to delete
type AccountDeleteModel struct {
	action    *action.Action
	menuIdx   int
	menuItems []string
	password  textinput.Model
	result    string
}

func NewAccountDeleteModel(a *action.Action) *AccountDeleteModel {
	passwordInput := textinput.New()
	passwordInput.Placeholder = "Enter master password"
	passwordInput.EchoMode = textinput.EchoPassword
	passwordInput.Focus()

	return &AccountDeleteModel{
		action:    a,
		menuIdx:   0,
		menuItems: []string{"Password", "SUBMIT", "BACK"},
		password:  passwordInput,
		result:    "",
	}
}

func (m *AccountDeleteModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *AccountDeleteModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

			if m.menuIdx == 0 {
				m.password.Focus()
			} else {
				m.password.Blur()
			}
		case "enter":
			switch selected {
			case "BACK":
				return NewPasswordMenuModel(m.action), nil
			case "SUBMIT":
				err := m.action.DeleteMasterAccount(m.password.Value())
				if err != nil {
					m.result = err.Error()
					return m, nil
				}
				return NewAuthMenuModel(m.action), nil
			}
		}
	}

	if m.menuIdx == 0 {
		m.password, cmd = m.password.Update(msg)
	}
	return m, cmd
}

func (m *AccountDeleteModel) View() string {
	s := "Enter your master password\n"
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
	s += m.result
	return s
}
