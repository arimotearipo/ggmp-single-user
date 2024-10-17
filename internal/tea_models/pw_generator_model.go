package teamodels

import (
	"strconv"

	"github.com/arimotearipo/ggmp/internal/action"
	"github.com/arimotearipo/ggmp/internal/types"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type PasswordGeneratorModel struct {
	action          *action.Action
	uppercaseLength numberinput
	lowercaseLength numberinput
	specialLength   numberinput
	numericLength   numberinput
	totalLength     numberinput
	menuItems       []string
	menuIdx         int
	result          string
}

func NewPasswordGeneratorModel(a *action.Action) *PasswordGeneratorModel {

	uppercaseLengthInput := NewNumberInput()
	uppercaseLengthInput.SetValue("3")
	uppercaseLengthInput.Focus()

	lowercaseLengthInput := NewNumberInput()
	lowercaseLengthInput.SetValue("3")
	lowercaseLengthInput.Placeholder = "Enter a number"

	specialLengthInput := NewNumberInput()
	specialLengthInput.SetValue("1")
	specialLengthInput.Placeholder = "Enter a number"

	numericLengthInput := NewNumberInput()
	numericLengthInput.SetValue(("1"))
	numericLengthInput.Placeholder = "Enter a number"

	totalLengthInput := NewNumberInput()
	totalLengthInput.SetValue("8")
	totalLengthInput.Placeholder = "Enter a number"

	return &PasswordGeneratorModel{
		action:          a,
		uppercaseLength: uppercaseLengthInput,
		lowercaseLength: lowercaseLengthInput,
		specialLength:   specialLengthInput,
		numericLength:   numericLengthInput,
		totalLength:     totalLengthInput,
		menuItems:       []string{"Uppercase", "Lowercase", "Special characters", "Numeric characters", "Total length", "GENERATE", "BACK"},
		menuIdx:         0,
		result:          "",
	}
}

func (m *PasswordGeneratorModel) currentInput() *numberinput {
	switch m.menuIdx {
	case 0:
		return &m.uppercaseLength
	case 1:
		return &m.lowercaseLength
	case 2:
		return &m.specialLength
	case 3:
		return &m.numericLength
	case 4:
		return &m.totalLength
	default:
		return nil
	}
}

func (m *PasswordGeneratorModel) calibrateTotalLength() {
	sum := m.lowercaseLength.IntValue() + m.uppercaseLength.IntValue() + m.specialLength.IntValue() + m.numericLength.IntValue()

	if m.totalLength.IntValue() < sum {
		sumStr := strconv.Itoa(sum)
		m.totalLength.SetValue(sumStr)
	}
}

func (m *PasswordGeneratorModel) increment() {
	curr := m.currentInput()

	if curr == nil {
		return
	}

	if curr.Value() == "" {
		curr.SetValue("0")
	}

	currentValue, err := strconv.Atoi(curr.Value())
	if err != nil {
		return
	}

	curr.SetValue(strconv.Itoa(currentValue + 1))
}

func (m *PasswordGeneratorModel) decrement() {
	curr := m.currentInput()

	if curr == nil {
		return
	}

	if curr.Value() == "" {
		curr.SetValue("0")
	}

	if curr.Value() == "0" {
		return
	}

	currentValue, err := strconv.Atoi(curr.Value())
	if err != nil {
		return
	}

	curr.SetValue(strconv.Itoa(currentValue - 1))
}

func (m *PasswordGeneratorModel) blurAllInputs() {
	m.uppercaseLength.Blur()
	m.lowercaseLength.Blur()
	m.specialLength.Blur()
	m.numericLength.Blur()
	m.totalLength.Blur()
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
			curr := m.currentInput()
			if curr != nil {
				curr.Focus()
			}
		case "left", "right":
			if msg.String() == "left" {
				m.decrement()
			} else if msg.String() == "right" {
				m.increment()
			}
		case "enter":
			selected := m.menuItems[m.menuIdx]
			if selected == "GENERATE" {
				pw, err := m.action.GeneratePassword(types.PasswordGeneratorConfig{
					UppercaseLength: m.uppercaseLength.IntValue(),
					LowercaseLength: m.lowercaseLength.IntValue(),
					SpecialLength:   m.specialLength.IntValue(),
					NumericLength:   m.numericLength.IntValue(),
					TotalLength:     m.totalLength.IntValue(),
				})

				if err != nil {
					m.result = err.Error()
				} else {
					m.result = pw
				}

			} else if selected == "BACK" {
				return NewAuthMenuModel(m.action), nil
			}
		}
	}

	m.calibrateTotalLength()
	_, cmd = m.currentInput().Update(msg)

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
		case "Total length":
			s += "   " + m.totalLength.View() + "\n"
		default:
			s += "\n"
		}

	}

	s += m.result

	return s
}
