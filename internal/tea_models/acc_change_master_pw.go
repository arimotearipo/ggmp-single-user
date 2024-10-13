package teamodels

import (
	"errors"

	"github.com/arimotearipo/ggmp/internal/action"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type AccountChangeMasterPasswordModel struct {
	a                     *action.Action
	masterPassword        textinput.Model
	confirmMasterPassword textinput.Model
	menuItems             []string
	menuIdx               int
	result                string
}

func NewAccountChangeMasterPasswordModel(a *action.Action) *AccountChangeMasterPasswordModel {
	masterPasswordInput := textinput.New()
	masterPasswordInput.EchoMode = textinput.EchoPassword
	masterPasswordInput.Placeholder = "Enter new master password"
	masterPasswordInput.Focus()

	confirmMasterPasswordInput := textinput.New()
	confirmMasterPasswordInput.EchoMode = textinput.EchoPassword
	confirmMasterPasswordInput.Placeholder = "Confirm new master password"

	return &AccountChangeMasterPasswordModel{
		a:                     a,
		masterPassword:        masterPasswordInput,
		confirmMasterPassword: confirmMasterPasswordInput,
		menuItems:             []string{"Master Password", "Confirm Master Password", "SUBMIT", "BACK"},
		menuIdx:               0,
		result:                "",
	}
}

func (m *AccountChangeMasterPasswordModel) validatePasswords() error {
	if m.masterPassword.Value() != m.confirmMasterPassword.Value() {
		return errors.New("passwords do not match")
	}

	return nil
}

func (m *AccountChangeMasterPasswordModel) blurAllInputs() {
	m.masterPassword.Blur()
	m.confirmMasterPassword.Blur()
}

func (m *AccountChangeMasterPasswordModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *AccountChangeMasterPasswordModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "up", "down":
			if msg.String() == "up" {
				m.menuIdx = (m.menuIdx - 1 + len(m.menuItems)) % len(m.menuItems)
			} else {
				m.menuIdx = (m.menuIdx + 1) % len(m.menuItems)
			}

			m.blurAllInputs()
			if m.menuIdx == 0 {
				m.masterPassword.Focus()
			} else if m.menuIdx == 1 {
				m.confirmMasterPassword.Focus()
			}
		case "enter":
			selected := m.menuItems[m.menuIdx]
			switch selected {
			case "SUBMIT":
				if err := m.validatePasswords(); err != nil {
					m.result = err.Error()
					return m, nil
				} else {
					m.result = ""
				}

				if err := m.a.UpdateMasterPassword(m.masterPassword.Value()); err != nil {
					m.result = err.Error()
				} else {
					m.result = "Master password successfully changed!"
				}

				return m, nil
			case "BACK":
				return NewPasswordMenuModel(m.a), nil
			}
		}
	}

	selected := m.menuItems[m.menuIdx]
	if selected == "Master Password" {
		m.masterPassword, cmd = m.masterPassword.Update(msg)

	} else if selected == "Confirm Master Password" {
		m.confirmMasterPassword, cmd = m.confirmMasterPassword.Update(msg)
	}

	m.validatePasswords()
	return m, cmd
}

func (m *AccountChangeMasterPasswordModel) View() string {
	s := "Insert new master password\n"

	for i, item := range m.menuItems {
		if i == m.menuIdx {
			s += "👉 "
		} else {
			s += "   "
		}
		switch item {
		case "Master Password":
			s += m.masterPassword.View() + "\n"
		case "Confirm Master Password":
			s += m.confirmMasterPassword.View() + "\n"
		default:
			s += item + "\n"
		}
	}

	s += m.result

	return s
}
