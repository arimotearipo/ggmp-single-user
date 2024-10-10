package teamodels

import (
	"github.com/arimotearipo/ggmp/action"
	tea "github.com/charmbracelet/bubbletea"
)

type ListPasswordsModel struct {
	action    *action.Action
	menuItems []string
	menuIdx   int
	prevModel tea.Model
	nextModel tea.Model
}

func NewListPasswordsModel(a *action.Action, prevModel, nextModel tea.Model) *ListPasswordsModel {
	uris, _ := a.ListURIs()
	return &ListPasswordsModel{
		action:    a,
		menuItems: append(uris, "BACK"),
		menuIdx:   0,
		prevModel: prevModel,
		nextModel: nextModel,
	}
}

func (m *ListPasswordsModel) Init() tea.Cmd {
	return nil
}

func (m *ListPasswordsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				return m.prevModel, nil
			default:
				return m.nextModel, nil
			}
		}
	}
	return m, nil
}

func (m *ListPasswordsModel) View() string {
	s := ""

	for i, item := range m.menuItems {
		if i == m.menuIdx {
			s += "👉 " + item + "\n"
		} else {
			s += "   " + item + "\n"
		}
	}
	return s
}
