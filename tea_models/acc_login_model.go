package teamodels

import (
	"github.com/arimotearipo/ggmp/action"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type LoginModel struct {
	action    *action.Action
	menuIdx   int
	menuItems []string
	username  textinput.Model
	password  textinput.Model
}

func NewLoginModel(a *action.Action) *LoginModel {
	usernameInput := textinput.New()
	usernameInput.Placeholder = "Enter username"
	usernameInput.Focus()

	passwordInput := textinput.New()
	passwordInput.Placeholder = "Enter password"
	passwordInput.EchoMode = textinput.EchoPassword

	return &LoginModel{
		action:    a,
		menuItems: []string{"Username", "Password", "SUBMIT", "BACK"},
		menuIdx:   0,
		username:  usernameInput,
		password:  passwordInput,
	}
}

func (m *LoginModel) blurAllInputs() {
	m.username.Blur()
	m.password.Blur()
}

func (m *LoginModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.menuIdx == 0 {
				m.username.Focus()
			} else if m.menuIdx == 1 {
				m.password.Focus()
			}
		case "enter":
			selected := m.menuItems[m.menuIdx]
			switch selected {
			case "BACK":
				return NewAuthMenuModel(m.action), nil
			case "SUBMIT":
				ok, _ := m.action.Login(m.username.Value(), m.password.Value())
				if ok {
					return NewPasswordMenuModel(m.action), nil
				}
				return m, tea.Quit
			}

		}
	}

	if m.menuIdx == 0 {
		m.username, cmd = m.username.Update(msg)
	} else if m.menuIdx == 1 {
		m.password, cmd = m.password.Update(msg)
	}

	return m, cmd
}

func (m *LoginModel) View() string {
	s := ""
	for i, item := range m.menuItems {
		if i == m.menuIdx {
			s += "👉 "
		} else {
			s += "   "
		}
		if item == "Username" {
			s += m.username.View() + "\n"
		} else if item == "Password" {
			s += m.password.View() + "\n"
		} else {
			s += item + "\n"
		}
	}
	return s
}
