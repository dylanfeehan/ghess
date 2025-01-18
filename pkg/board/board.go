package board

import (
	"fmt"
	"os"
	"strconv"

	nt "github.com/dylanfeehan/ghess/pkg/notation"

	"github.com/charmbracelet/lipgloss"
)

type Square struct {
	Rank int
	File int
}

type Piece struct {
	Color int
	Type  int
}

type Board [][]*Piece

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

func (b Board) Color(x, y int) int {
	if (x+y)%2 == 0 {
		return WHITE
	}
	return BLACK
}

func (b Board) SquareStyle(x, y int) lipgloss.Style {
	if b.Color(x, y) == WHITE {
		return lipgloss.NewStyle().Background(lipgloss.Color("#ff6eff"))
	} else {
		return lipgloss.NewStyle().Background(lipgloss.Color("#684fff"))
	}
}

func (b Board) Init() Board {
	var board [][]*Piece = make([][]*Piece, 8, 8)

	for i := 0; i < 8; i++ {
		rank := make([]*Piece, 8, 8)
		board[i] = rank
	}

	// zero out entire board
	for rank, row := range board {
		for file, _ := range row {
			board[rank][file] = nil
			board[rank][file] = nil
		}
	}

	//whitePlayer := true

	// initializes top two rows
	pieceColor := BLACK
	for rank := 0; rank < 2; rank++ {
		for file := 0; file < 8; file++ {
			pieceType := P
			if rank == 0 {
				pieceType = whitePlayerInitMap[file]
			}
			board[rank][file] = &Piece{
				Color: pieceColor,
				Type:  pieceType,
			}
		}
	}

	// initializes top two rows
	pieceColor = WHITE
	for rank := 6; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			pieceType := P
			if rank == 7 {
				pieceType = whitePlayerInitMap[file]
			}
			board[rank][file] = &Piece{
				Color: pieceColor,
				Type:  pieceType,
			}
		}
	}
	return board
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

// ExecuteMove  takes input to chess notation for a move
// and returns whether or not the move was executed
func (board Board) ExecuteMove(mv string) bool {
	move, err := nt.Parse(mv)
	if err != nil {
		return false
	}

	piece := PieceStrToInt(move.Piece)
	fileIdx := getFileIndex(move.File, WHITE)
	rankIdx := getRankIndex(move.Rank, WHITE)

	sq1, sq2 := board.GetSquaresForMove(piece, fileIdx, rankIdx)
	if sq1 == nil || sq2 == nil {
		return false
	} else {
		board.executeMove(*sq1, *sq2)
		return true
	}
}

// lots more to do here
func (board Board) executeMove(sq1 Square, sq2 Square) {
	srcPiece := board[sq1.Rank][sq1.File]

	board[sq1.Rank][sq1.File] = nil
	board[sq2.Rank][sq2.File] = srcPiece
}

func getFileIndex(file string, player int) int {
	fileNum := FileStrToInt(file)
	if player == WHITE {
		return fileNum - 1
	} else {
		return 8 - fileNum
	}
}

func getRankIndex(rank string, player int) int {
	rankNum, err := strconv.Atoi(rank)
	if err != nil {
		os.Exit(1)
	}
	if player == BLACK {
		return rankNum - 1
	} else {
		//fmt.Printf("rankNum = %+v\n", rankNum)
		return 8 - rankNum
	}
}

func (board Board) GetSquaresForMove(piece int, fileIdx int, rankIdx int) (*Square, *Square) {
	destSquare := &Square{rankIdx, fileIdx}
	//fmt.Printf("fileIdx = %+v\n", fileIdx)
	//fmt.Printf("rankIdx = %+v\n", rankIdx)
	//
	// missing "If player == pawn" logic
	//player := WHITE
	// these +1 and +2 are dependent on the direction of the pawn which depends on the player
	if board[rankIdx+1][fileIdx] != nil && board[rankIdx+1][fileIdx].Type == PAWN {
		return &Square{rankIdx + 1, fileIdx}, destSquare
	}
	if board[rankIdx+2][fileIdx] != nil && board[rankIdx+2][fileIdx].Type == PAWN {
		return &Square{rankIdx + 2, fileIdx}, destSquare
	}
	//fmt.Println("THe result is unforunately nill")
	return nil, nil
}
