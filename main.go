package main

import (
	"fmt"
	"os"

	tip "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	boardutils "github.com/dylanfeehan/ghess/pkg/board"
)

type model struct {
	board     boardutils.Board
	textInput tip.Model
	err       error
	player    int
	//cursor   int              // which to-do list item our cursor is pointing at
	//selected map[int]struct{} // which to-do items are selected
}

func main() {

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	fmt.Println("ghess donezo")
}

func initialModel() model {
	board := boardutils.Board{}
	textInput := tip.New()
	// in the future, Placeholder = if helpEnabled getSmartMove() else random
	textInput.Placeholder = "c5Nxf7#"
	textInput.Focus()
	textInput.CharLimit = 10
	textInput.Width = 20
	return model{
		board:     board.Init(),
		textInput: textInput,
		player:    boardutils.WHITE,
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		}
		switch msg.Type {

		case tea.KeyEnter:
			move := m.textInput.Value()
			ok := m.board.ExecuteMove(move, m.player)
			if ok {
				m.player = flip(m.player)
			}
			m.textInput.Reset()
		}
	case error:
		m.err = msg
		fmt.Println(msg)
		return m, nil
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, cmd
}

func (m model) View() string {
	// Define chess piece and board styles
	//whitePiece := lipgloss.NewStyle().Foreground(lipgloss.Color("231")).Background(lipgloss.Color("242"))
	//blackPiece := lipgloss.NewStyle().Foreground(lipgloss.Color("232")).Background(lipgloss.Color("136"))
	s := ""
	s += m.board.Render(m.player)
	s += m.textInput.View()
	return s
}

func flip(player int) int {
	if player == boardutils.WHITE {
		return boardutils.BLACK
	} else {
		return boardutils.WHITE
	}
}
