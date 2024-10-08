package teamodels

import (
	"github.com/arimotearipo/ggmp/action"
	tea "github.com/charmbracelet/bubbletea"
)

type RegisterModel struct {
	action          *action.Action
	menuIdx         int
	menuItems       []string
	username        string
	password        string
	confirmPassword string
	err             string
}

func NewRegisterModel(c *action.Action) *RegisterModel {
	return &RegisterModel{
		action:    c,
		menuItems: []string{"Username: ", "Password: ", "Confirm Password: ", "SUBMIT", "BACK"},
	}
}

func (m *RegisterModel) Init() tea.Cmd {
	return nil
}

func (m *RegisterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "up":
			m.menuIdx = (m.menuIdx - 1 + len(m.menuItems)) % len(m.menuItems)
		case "down":
			m.menuIdx = (m.menuIdx + 1) % len(m.menuItems)
		case "backspace":
			m.menuItems[m.menuIdx] = m.menuItems[m.menuIdx][:len(m.menuItems[m.menuIdx])-1]
		case "enter":
			if m.menuItems[m.menuIdx] == "SUBMIT" {
				if m.err != "" {
					return m, nil
				}
				if m.action.Register(m.username, m.password) {
					return NewAuthMenuModel(m.action), nil
				}
				return m, tea.Quit
			}
			if m.menuItems[m.menuIdx] == "BACK" {
				return NewAuthMenuModel(m.action), nil
			}
		default:
			if m.menuIdx == 0 {
				m.username += msg.String()
				m.menuItems[m.menuIdx] += msg.String()
			} else if m.menuIdx == 1 {
				m.password += msg.String()
				m.menuItems[m.menuIdx] += "*"
			} else if m.menuIdx == 2 {
				m.confirmPassword += msg.String()
				m.menuItems[m.menuIdx] += "*"
			}

			if m.password != m.confirmPassword {
				m.err = "Passwords do not match"
			} else {
				m.err = ""
			}
		}
	}
	return m, nil
}

func (m *RegisterModel) View() string {
	s := ""
	for i, input := range m.menuItems {
		if i == m.menuIdx {
			s += "👉 "
		} else {
			s += "   "
		}
		s += input + "\n"
	}

	if m.err != "" {
		s += "Error: " + m.err + "\n"
	}
	return s
}
