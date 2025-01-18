package board

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (p *Piece) Text() string {
	if p == nil {
		return " "
	}
	if p.Color == WHITE {
		return whitePieces[p.Type]
	} else {
		return blackPieces[p.Type]
	}
}

func (p *Piece) Style() lipgloss.Style {
	if p == nil {
		return lipgloss.NewStyle() // we'll see if blank style will render properly
	}
	if p.Color == WHITE {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("231"))
	} else {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("232"))
	}
}

func (board Board) renderWhite() string {
	s := ""
	for i, row := range board {
		for j, piece := range row {
			squareStyle := board.SquareStyle(i, j)
			pieceStyle := piece.Style()
			s += fmt.Sprint(squareStyle.Render(pieceStyle.Render(" " + piece.Text() + " ")))
		}
		s += fmt.Sprintln()
	}
	return s

}

func (board Board) renderBlack() string {
	s := ""
	for i := 7; i >= 0; i-- {
		for j := 7; j >= 0; j-- {
			squareStyle := board.SquareStyle(i, j)
			piece := board[i][j]
			pieceStyle := piece.Style()
			s += fmt.Sprint(squareStyle.Render(pieceStyle.Render(" " + piece.Text() + " ")))
		}
		s += fmt.Sprintln()
	}
	return s
}

func (board Board) Render(player int) string {
	// Render board
	if player == WHITE {
		return board.renderWhite()
	} else {
		return board.renderBlack()
	}
}

func (b Board) SquareStyle(x, y int) lipgloss.Style {
	if b.Color(x, y) == WHITE {
		return lipgloss.NewStyle().Background(lipgloss.Color("#ff6eff"))
	} else {
		return lipgloss.NewStyle().Background(lipgloss.Color("#684fff"))
	}
}
