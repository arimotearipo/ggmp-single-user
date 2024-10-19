package teamodels

import (
	"github.com/arimotearipo/ggmp/internal/action"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type PasswordAddModel struct {
	action    *action.Action
	menuItems []string
	menuIdx   int
	uri       textinput.Model
	username  textinput.Model
	password  textinput.Model
	result    string
}

func NewPasswordAddModel(a *action.Action) *PasswordAddModel {
	uriInput := textinput.New()
	uriInput.Placeholder = "Enter URI"
	uriInput.Focus()

	usernameInput := textinput.New()
	usernameInput.Placeholder = "Enter username"

	passwordInput := textinput.New()
	passwordInput.Placeholder = "Enter password"
	passwordInput.EchoMode = textinput.EchoPassword

	return &PasswordAddModel{
		action:    a,
		menuItems: []string{"URI", "Username", "Password", "SUBMIT", "BACK"},
		menuIdx:   0,
		uri:       uriInput,
		username:  usernameInput,
		password:  passwordInput,
		result:    "",
	}
}

func (m *PasswordAddModel) currentInput() *textinput.Model {
	switch m.menuIdx {
	case 0:
		return &m.uri
	case 1:
		return &m.username
	case 2:
		return &m.password
	default:
		return nil
	}
}

func (m *PasswordAddModel) blurAllInputs() {
	m.username.Blur()
	m.password.Blur()
	m.uri.Blur()
}

func (m *PasswordAddModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *PasswordAddModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
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
			selected := m.menuItems[m.menuIdx]
			switch selected {
			case "BACK":
				return NewPasswordMenuModel(m.action), nil
			case "SUBMIT":
				if err := m.action.AddPassword(m.uri.Value(), m.username.Value(), m.password.Value()); err != nil {
					m.result = err.Error()
				} else {
					return NewPasswordMenuModel(m.action), nil
				}
			}
		}
	}
	curr := m.currentInput()
	if curr != nil {
		*curr, cmd = m.currentInput().Update(msg)
	}

	return m, cmd
}

func (m *PasswordAddModel) View() string {
	s := "Add new login details to save\n"
	for i, item := range m.menuItems {
		if i == m.menuIdx {
			s += "👉 "
		} else {
			s += "   "
		}

		switch item {
		case "URI":
			s += m.uri.View() + "\n"
		case "Username":
			s += m.username.View() + "\n"
		case "Password":
			s += m.password.View() + "\n"
		default:
			s += item + "\n"
		}
	}

	s += m.result

	return s
}
