package teamodels

import (
	"github.com/arimotearipo/ggmp/action"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type PasswordUpdateModel struct {
	action    *action.Action
	menuItems []string
	menuIdx   int
	uri       textinput.Model
	username  textinput.Model
	password  textinput.Model
	result    string
}

func NewPasswordUpdateModel(a *action.Action) *PasswordUpdateModel {
	uriInput := textinput.New()
	uriInput.Placeholder = "Enter URI"
	uriInput.Focus()

	usernameInput := textinput.New()
	usernameInput.Placeholder = "Enter username"

	passwordInput := textinput.New()
	passwordInput.Placeholder = "Enter password"
	passwordInput.EchoMode = textinput.EchoPassword

	return &PasswordUpdateModel{
		action:    a,
		menuItems: []string{"URI", "Username", "Password", "UPDATE", "BACK"},
		menuIdx:   0,
		uri:       uriInput,
		username:  usernameInput,
		password:  passwordInput,
		result:    "",
	}
}

func (m *PasswordUpdateModel) Init() tea.Cmd {
	return nil
}

func (m *PasswordUpdateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case "enter":
			selected := m.menuItems[m.menuIdx]
			switch selected {
			case "BACK":
				return NewPasswordsListModel(m.action, "Update password"), nil
			case "UPDATE":
				if err := m.action.UpdatePassword(m.uri.Value(), m.username.Value(), m.password.Value()); err != nil {
					m.result = err.Error()
				} else {
					return NewPasswordsListModel(m.action, "Update password"), nil
				}
			}
		}
	}

	selected := m.menuItems[m.menuIdx]
	switch selected {
	case "URI":
		m.uri, cmd = m.uri.Update(msg)
	case "Username":
		m.username, cmd = m.username.Update(msg)
	case "Password":
		m.password, cmd = m.password.Update(msg)
	}

	return m, cmd
}

func (m *PasswordUpdateModel) View() string {
	s := ""
	for i, item := range m.menuItems {
		if i == m.menuIdx {
			s += "👉 " + item + "\n"
		} else {
			s += "   " + item + "\n"
		}
	}
	s += m.result
	return s
}
