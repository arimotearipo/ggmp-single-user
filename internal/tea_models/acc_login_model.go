package teamodels

import (
	"github.com/arimotearipo/ggmp/internal/action"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type AccountLoginModel struct {
	action    *action.Action
	menuIdx   int
	menuItems []string
	password  textinput.Model
	result    string
}

func NewAccountLoginModel(a *action.Action) *AccountLoginModel {
	passwordInput := textinput.New()
	passwordInput.Placeholder = "Enter master password"
	passwordInput.EchoMode = textinput.EchoPassword
	passwordInput.Focus()

	return &AccountLoginModel{
		action:    a,
		menuItems: []string{"Master Password", "SUBMIT", "BACK"},
		menuIdx:   0,
		password:  passwordInput,
		result:    "",
	}
}

func (m *AccountLoginModel) blurAllInputs() {
	m.password.Blur()
}

func (m *AccountLoginModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *AccountLoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				m.password.Focus()
			}
		case "enter":
			selected := m.menuItems[m.menuIdx]
			switch selected {
			case "BACK":
				return NewAuthMenuModel(m.action), nil
			case "SUBMIT":
				if err := m.action.Login(m.password.Value()); err != nil {
					m.result = err.Error()
					return m, nil
				}
				return NewPasswordMenuModel(m.action), nil
			}

		}
	}

	if m.menuIdx == 0 {
		m.password, cmd = m.password.Update(msg)
	}

	return m, cmd
}

func (m *AccountLoginModel) View() string {
	s := "Enter your credentials\n"
	for i, item := range m.menuItems {
		if i == m.menuIdx {
			s += "👉 "
		} else {
			s += "   "
		}
		if item == "Master Password" {
			s += m.password.View() + "\n"
		} else {
			s += item + "\n"
		}
	}
	s += m.result
	return s
}
