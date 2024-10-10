package teamodels

import (
	"github.com/arimotearipo/ggmp/action"
	tea "github.com/charmbracelet/bubbletea"
)

type PasswordGetModel struct {
	action      *action.Action
	menuItems   []string
	menuIdx     int
	loginDetail [2]string
}

func NewPasswordGetModel(a *action.Action) *PasswordGetModel {
	uris, _ := a.ListURIs()

	return &PasswordGetModel{
		action:      a,
		menuItems:   append(uris, "BACK"),
		menuIdx:     0,
		loginDetail: [2]string{},
	}
}

func (m *PasswordGetModel) Init() tea.Cmd {
	return nil
}

func (m *PasswordGetModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case "enter":
			selected := m.menuItems[m.menuIdx]
			switch selected {
			case "BACK":
				return NewPasswordMenuModel(m.action), nil
			default:
				u, p, err := m.action.GetPassword(selected)
				if err == nil {
					m.loginDetail = [2]string{u, p}
				}
				return m, nil
			}
		}
	}
	return m, cmd
}

func (m *PasswordGetModel) View() string {
	s := ""
	for i, item := range m.menuItems {
		if i == m.menuIdx {
			s += "👉 " + item + "\n"
		} else {
			s += "   " + item + "\n"
		}
	}

	if m.loginDetail[0] != "" {
		s += "Username: " + m.loginDetail[0] + "\n"
		s += "Password: " + m.loginDetail[1] + "\n"
	}
	return s
}
