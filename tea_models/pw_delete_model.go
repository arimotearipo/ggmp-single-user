package teamodels

import (
	"fmt"

	"github.com/arimotearipo/ggmp/action"
	tea "github.com/charmbracelet/bubbletea"
)

type PasswordDeleteModel struct {
	action    *action.Action
	menuItems []string
	menuIdx   int
	uri       string
	result    string
}

func NewPasswordDeleteModel(a *action.Action, uri string) *PasswordDeleteModel {
	return &PasswordDeleteModel{
		action:    a,
		menuItems: []string{"Yes", "No"},
		menuIdx:   1,
		uri:       uri,
		result:    "",
	}
}

func (m *PasswordDeleteModel) Init() tea.Cmd {
	return nil
}

func (m *PasswordDeleteModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			case "No":
				return NewPasswordsListModel(m.action, "Delete password"), nil
			case "Yes":
				err := m.action.DeletePassword(m.uri)
				if err == nil {
					return NewPasswordsListModel(m.action, "Delete password"), nil
				}
				m.result = err.Error()
				return m, nil
			}
		}
	}
	return m, nil
}

func (m *PasswordDeleteModel) View() string {
	s := fmt.Sprintf("Are you sure you want to delete login for %s?\n", m.uri)
	for i, item := range m.menuItems {
		if i == m.menuIdx {
			s += "👉 " + item + "\n"
		} else {
			s += "   " + item + "\n"
		}
	}
	s += m.result + "\n"
	return s
}
