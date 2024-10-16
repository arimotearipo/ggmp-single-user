package teamodels

import (
	"github.com/arimotearipo/ggmp/internal/action"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type PasswordGeneratorModel struct {
	action          *action.Action
	uppercaseLength textinput.Model
	lowercaseLength textinput.Model
	specialLength   textinput.Model
	numericLength   textinput.Model
	menuItems       []string
	menuIdx         int
	result          string
}

func NewPasswordGeneratorModel(a *action.Action) *PasswordGeneratorModel {
	uppercaseLengthInput := textinput.New()
	uppercaseLengthInput.Placeholder = "Enter a number"
	uppercaseLengthInput.Focus()

	lowercaseLengthInput := textinput.New()
	lowercaseLengthInput.Placeholder = "Enter a number"

	specialLengthInput := textinput.New()
	specialLengthInput.Placeholder = "Enter a number"

	numericLengthInput := textinput.New()
	numericLengthInput.Placeholder = "Enter a number"

	return &PasswordGeneratorModel{
		action:          a,
		uppercaseLength: uppercaseLengthInput,
		lowercaseLength: lowercaseLengthInput,
		specialLength:   specialLengthInput,
		numericLength:   numericLengthInput,
		menuItems:       []string{"Uppercase", "Lowercase", "Special characters", "Numeric characters", "GENERATE", "BACK"},
		menuIdx:         0,
		result:          "",
	}
}

func (m *PasswordGeneratorModel) blurAllInputs() {
	m.uppercaseLength.Blur()
	m.lowercaseLength.Blur()
	m.specialLength.Blur()
	m.numericLength.Blur()
}

func (m *PasswordGeneratorModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *PasswordGeneratorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

			m.blurAllInputs()
			selected := m.menuItems[m.menuIdx]
			switch selected {
			case "Uppercase":
				m.uppercaseLength.Focus()
			case "Lowercase":
				m.lowercaseLength.Focus()
			case "Special characters":
				m.specialLength.Focus()
			case "Numeric characters":
				m.numericLength.Focus()
			default:
				m.blurAllInputs()
			}
		case "left", "right":
			// TODO: find an efficient way to handle shifting values by left/right keys
		case "enter":
			selected := m.menuItems[m.menuIdx]
			if selected == "GENERATE" {
				// TODO: generate keys logic
			} else if selected == "BACK" {
				return NewAuthMenuModel(m.action), nil
			}
		}
	}

	selected := m.menuItems[m.menuIdx]
	switch selected {
	case "Uppercase":
		m.uppercaseLength, cmd = m.uppercaseLength.Update(msg)
	case "Lowercase":
		m.lowercaseLength, cmd = m.lowercaseLength.Update(msg)
	case "Special characters":
		m.specialLength, cmd = m.specialLength.Update(msg)
	case "Numeric characters":
		m.numericLength, cmd = m.numericLength.Update(msg)
	}

	return m, cmd
}

func (m *PasswordGeneratorModel) View() string {
	s := ""
	for i, item := range m.menuItems {
		if i == m.menuIdx {
			s += "👉 " + item
		} else {
			s += "   " + item
		}

		switch item {
		case "Uppercase":
			s += "   " + m.uppercaseLength.View() + "\n"
		case "Lowercase":
			s += "   " + m.lowercaseLength.View() + "\n"
		case "Special characters":
			s += "   " + m.specialLength.View() + "\n"
		case "Numeric characters":
			s += "   " + m.numericLength.View() + "\n"
		default:
			s += "\n"
		}

	}

	return s
}
