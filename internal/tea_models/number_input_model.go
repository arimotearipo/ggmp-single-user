package teamodels

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type numberinput struct {
	textinput.Model
}

func NewNumberInput() numberinput {
	return numberinput{textinput.New()}
}

func (n *numberinput) IntValue() int {
	val, _ := strconv.Atoi(n.Value())

	return val
}

func (n *numberinput) Init() tea.Cmd {
	return textinput.Blink
}

func (n *numberinput) Update(msg tea.Msg) (*numberinput, tea.Cmd) {
	var cmd tea.Cmd
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		key := keyMsg.String()
		if !strings.Contains("0123456789", key) && key != "backspace" && key != "delete" {
			return n, nil // Ignore non-numeric keys except backspace and delete
		}
	}

	n.Model, cmd = n.Model.Update(msg)
	return n, cmd
}

func (n *numberinput) View() string {
	return n.Model.View()
}
