package teamodels

import (
	"errors"

	"github.com/arimotearipo/ggmp/internal/action"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type EncryptFileModel struct {
	action          *action.Action
	password        textinput.Model
	confirmPassword textinput.Model
	menuIdx         int
	menuItems       []string
	result          string
}

func NewEncryptFileModel(a *action.Action) *EncryptFileModel {
	passwordInput := textinput.New()
	passwordInput.EchoMode = textinput.EchoPassword
	passwordInput.Placeholder = "Password"
	passwordInput.Focus()

	confirmPasswordInput := textinput.New()
	confirmPasswordInput.EchoMode = textinput.EchoPassword
	confirmPasswordInput.Placeholder = "Confirm Password"

	return &EncryptFileModel{
		action:          a,
		password:        passwordInput,
		confirmPassword: confirmPasswordInput,
		menuIdx:         0,
		menuItems:       []string{"Password", "Confirm Password", "ENCRYPT", "BACK"},
		result:          "",
	}
}

func (m *EncryptFileModel) validatePasswords() error {
	if m.password.Value() != m.confirmPassword.Value() {
		return errors.New("passwords do not match")
	}
	return nil
}

func (m *EncryptFileModel) blurAllInputs() {
	m.password.Blur()
	m.confirmPassword.Blur()
}

func (m *EncryptFileModel) currentInput() *textinput.Model {
	switch m.menuIdx {
	case 0:
		return &m.password
	case 1:
		return &m.confirmPassword
	default:
		return nil
	}
}

func (m *EncryptFileModel) Init() tea.Cmd {
	return nil
}

func (m *EncryptFileModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			} else if msg.String() == "down" {
				m.menuIdx = (m.menuIdx + 1) % len(m.menuItems)
			}

			m.blurAllInputs()
			curr := m.currentInput()
			if curr != nil {
				curr.Focus()
			}
		case "enter":
			switch selected {
			case "BACK":
				return NewPasswordMenuModel(m.action), nil
			case "ENCRYPT":
				if err := m.validatePasswords(); err != nil {
					m.result = err.Error()
					return m, nil
				}

				if err := m.action.EncryptDBFile(m.password.Value()); err != nil {
					m.result = err.Error()
				} else {
					m.result = "File successfully encrypted!"
				}
			}
		}
	}

	curr := m.currentInput()
	if curr != nil {
		*curr, cmd = curr.Update(msg)
	}

	return m, cmd
}

func (m *EncryptFileModel) View() string {
	s := "Enter a password to encrypt your file\nNOTE: This password must be remembered and cannot be recovered!\n"

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
