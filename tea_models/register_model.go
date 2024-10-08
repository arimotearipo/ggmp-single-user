package teamodels

import (
	"fmt"

	"github.com/arimotearipo/ggmp/action"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type RegisterModel struct {
	action          *action.Action
	menuIdx         int
	menuItems       []string
	username        textinput.Model
	password        textinput.Model
	confirmPassword textinput.Model
	err             string
}

func errorMsg(err error) string {
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	return ""
}

func (m *RegisterModel) blurAllInputs() {
	m.username.Blur()
	m.password.Blur()
	m.confirmPassword.Blur()
}

func (m *RegisterModel) validatePasswords() {
	if m.password.Value() != m.confirmPassword.Value() {
		m.err = "Passwords does not match"
	} else {
		m.err = ""
	}
}

func NewRegisterModel(c *action.Action) *RegisterModel {
	usernameInput := textinput.New()
	usernameInput.Placeholder = "Enter username"
	usernameInput.Focus()

	passwordInput := textinput.New()
	passwordInput.Placeholder = "Enter password"
	passwordInput.EchoMode = textinput.EchoPassword

	confirmPasswordInput := textinput.New()
	confirmPasswordInput.Placeholder = "Confirm password"
	confirmPasswordInput.EchoMode = textinput.EchoPassword

	m := &RegisterModel{
		action:          c,
		menuIdx:         0,
		menuItems:       []string{"Username", "Password", "Confirm Password", "SUBMIT", "BACK"},
		username:        usernameInput,
		password:        passwordInput,
		confirmPassword: confirmPasswordInput,
		err:             "",
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

func (m *RegisterModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *RegisterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				m.username.Focus()
			} else if m.menuIdx == 1 {
				m.password.Focus()
			} else if m.menuIdx == 2 {
				m.confirmPassword.Focus()
			}
		case "enter":
			selected := m.menuItems[m.menuIdx]
			switch selected {
			case "BACK":
				return NewAuthMenuModel(m.action), nil
			case "SUBMIT":
				if m.err != "" {
					return m, nil
				}
				if m.action.Register(m.username.Value(), m.password.Value()) {
					return NewAuthMenuModel(m.action), nil
				}
				return m, tea.Quit
			}
		}
	}

	var cmd tea.Cmd
	selected := m.menuItems[m.menuIdx]
	switch selected {
	case "Username":
		m.username, cmd = m.username.Update(msg)
	case "Password":
		m.password, cmd = m.password.Update(msg)
	case "Confirm Password":
		m.confirmPassword, cmd = m.confirmPassword.Update(msg)
	}
	return m, cmd
}

func (m *RegisterModel) View() string {
	s := ""
	for i, item := range m.menuItems {
		if i == m.menuIdx {
			s += "👉 "
		} else {
			s += "   "
		}

		switch item {
		case "Username":
			s += m.username.View() + "\n"
		case "Password":
			s += m.password.View() + "\n"
		case "Confirm Password":
			s += m.confirmPassword.View() + "\n"
		default:
			s += item + "\n"
		}
	}

	if m.err != "" {
		s += m.err + "\n"
	}
	return s
}
