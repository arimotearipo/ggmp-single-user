package teamodels

import (
	"github.com/arimotearipo/ggmp/cmd"
	tea "github.com/charmbracelet/bubbletea"
)

type LoginModel struct {
	cmd       *cmd.Command
	menuIdx   int
	menuItems []string
	username  string
	password  string
}

func NewLoginModel(c *cmd.Command) *LoginModel {
	return &LoginModel{
		cmd:       c,
		menuItems: []string{"Username: ", "Password: ", "SUBMIT", "BACK"},
		menuIdx:   0,
		username:  "",
		password:  "",
	}
}

func (m *LoginModel) Init() tea.Cmd {
	return nil
}

func (m *LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				if m.cmd.Login(m.username, m.password) {
					return NewPasswordMenuModel(m.cmd), nil
				}
				return m, tea.Quit
			}
			if m.menuItems[m.menuIdx] == "BACK" {
				return NewAuthMenuModel(m.cmd), nil
			}
		default:
			if m.menuIdx == 0 {
				m.username += msg.String()
				m.menuItems[m.menuIdx] += msg.String()
			} else if m.menuIdx == 1 {
				m.password += msg.String()
				m.menuItems[m.menuIdx] += "*"
			}
		}
	}
	return m, nil
}

func (m *LoginModel) View() string {
	s := ""
	for i, input := range m.menuItems {
		if i == m.menuIdx {
			s += "👉 "
		} else {
			s += "   "
		}
		s += input + "\n"
	}
	return s
}
