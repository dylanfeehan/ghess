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

func getCapturesOf(player int) []*Piece {
	playerCaptures := []*Piece{}
	for _, capture := range captures {
		if capture.Color == player {
			playerCaptures = append(playerCaptures, capture)
		}
	}
	return playerCaptures
}

// player is always on bottom
func (board Board) render() string {
	return ""
}

func (board Board) renderSquare(i, j int) string {
	squareStyle := board.SquareStyle(i, j)
	piece := board[i][j]
	pieceStyle := piece.Style()
	return fmt.Sprint(squareStyle.Render(pieceStyle.Render(" " + piece.Text() + " ")))
}

func renderCaptures(playerCaptures []*Piece) string {
	prisonStyle := lipgloss.NewStyle().Background(lipgloss.Color("#00ADD8"))
	if len(playerCaptures) == 0 {
		return ""
	}
	s := ""
	for _, capture := range playerCaptures {
		pieceStyle := capture.Style()
		s += fmt.Sprint(prisonStyle.Render(pieceStyle.Render(" " + capture.Text() + " ")))
		//s += fmt.Sprint(capture.Text())
	}
	return s
}

func (board Board) Render(player int) string {
	topCaptures := getCapturesOf(player)
	bottomCaptures := getCapturesOf(other(player))
	s := ""
	for i := 0; i < 4; i++ {
		for j := 0; j < 8; j++ {
			s += board.renderSquare(i, j)
		}
		if len(topCaptures) <= 4 {
			s += renderCaptures(topCaptures)
			topCaptures = []*Piece{}
		} else {
			s += renderCaptures(topCaptures[:4])
			topCaptures = topCaptures[4:]
		}
		s += fmt.Sprintln()
	}
	for i := 4; i < 8; i++ {
		for j := 0; j < 8; j++ {
			s += board.renderSquare(i, j)
		}
		if len(bottomCaptures) <= 4 {
			s += renderCaptures(bottomCaptures)
			bottomCaptures = []*Piece{}
		} else {
			s += renderCaptures(bottomCaptures[:4])
			bottomCaptures = bottomCaptures[4:]
		}
		s += fmt.Sprintln()
	}
	return s
}

func (b Board) SquareStyle(x, y int) lipgloss.Style {
	if b.Color(x, y) == WHITE {
		return lipgloss.NewStyle().Background(lipgloss.Color("#ff6eff"))
	} else {
		return lipgloss.NewStyle().Background(lipgloss.Color("#684fff"))
	}
}

func other(player int) int {
	if player == WHITE {
		return BLACK
	} else {
		return WHITE
	}
}
