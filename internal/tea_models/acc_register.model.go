package teamodels

import (
	"errors"

	"github.com/arimotearipo/ggmp/internal/action"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type AccountRegisterModel struct {
	action          *action.Action
	menuIdx         int
	menuItems       []string
	password        textinput.Model
	confirmPassword textinput.Model
	result          string
}

func (m *AccountRegisterModel) blurAllInputs() {
	m.password.Blur()
	m.confirmPassword.Blur()
}

func (m *AccountRegisterModel) validatePasswords() error {
	if m.password.Value() != m.confirmPassword.Value() {
		return errors.New("passwords do not match")
	}
	return nil
}

func NewAccountRegisterModel(a *action.Action) *AccountRegisterModel {
	passwordInput := textinput.New()
	passwordInput.Placeholder = "Enter master password"
	passwordInput.EchoMode = textinput.EchoPassword
	passwordInput.Focus()

	confirmPasswordInput := textinput.New()
	confirmPasswordInput.Placeholder = "Confirm master password"
	confirmPasswordInput.EchoMode = textinput.EchoPassword

	m := &AccountRegisterModel{
		action:          a,
		menuIdx:         0,
		menuItems:       []string{"Password", "Confirm Password", "SUBMIT", "BACK"},
		password:        passwordInput,
		confirmPassword: confirmPasswordInput,
		result:          "",
	}

	passwordInput.Validate = func(s string) error {
		m.validatePasswords()
		return nil
	}
	confirmPasswordInput.Validate = func(value string) error {
		m.validatePasswords()
		return nil
	}

	return m
}

func (m *AccountRegisterModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *AccountRegisterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				m.password.Focus()
			} else if m.menuIdx == 1 {
				m.confirmPassword.Focus()
			}
		case "enter":
			selected := m.menuItems[m.menuIdx]
			switch selected {
			case "BACK":
				return NewAuthMenuModel(m.action), nil
			case "SUBMIT":
				if err := m.validatePasswords(); err != nil {
					m.result = err.Error()
					return m, nil
				} else {
					m.result = "Master password set"
				}

				if err := m.action.Register(m.password.Value()); err != nil {
					m.result = err.Error()
					return m, nil
				}
				return NewAuthMenuModel(m.action), nil
			}
		}
	}

	var cmd tea.Cmd
	selected := m.menuItems[m.menuIdx]
	switch selected {
	case "Password":
		m.password, cmd = m.password.Update(msg)
	case "Confirm Password":
		m.confirmPassword, cmd = m.confirmPassword.Update(msg)
	}
	return m, cmd
}

func (m *AccountRegisterModel) View() string {
	s := "Register for account\nNOTE: Master password must be remembered and cannot be recovered!\n"
	for i, item := range m.menuItems {
		if i == m.menuIdx {
			s += "👉 "
		} else {
			s += "   "
		}

		switch item {
		case "Password":
			s += m.password.View() + "\n"
		case "Confirm Password":
			s += m.confirmPassword.View() + "\n"
		default:
			s += item + "\n"
		}
	}

	s += m.result

	return s
}
