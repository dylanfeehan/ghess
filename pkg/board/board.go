package board

import (
	"os"
	"strconv"

	nt "github.com/dylanfeehan/ghess/pkg/notation"
)

type Board [][]*Piece

type Piece struct {
	Color int
	Type  int
}

type Square struct {
	Rank int
	File int
}

func (b Board) Color(x, y int) int {
	if (x+y)%2 == 0 {
		return WHITE
	}
	return BLACK
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

// ExecuteMove  takes input to chess notation for a move
// and returns whether or not the move was executed

// the board representation of the board always has white on the bottom. the flipping is just a rendering trick
func (board Board) ExecuteMove(mv string, player int) bool {
	move, err := nt.Parse(mv)
	if err != nil {
		return false
	}

	piece := PieceStrToInt(move.Piece)
	fileIdx := getFileIndex(move.File, WHITE)
	rankIdx := getRankIndex(move.Rank, WHITE)

	// find the player which is being moved to the location specified in the notation
	moveSource := board.GetMoveSource(piece, fileIdx, rankIdx, player)

	if moveSource == nil {
		return false
	} else {
		board.executeMove(*moveSource, Square{rankIdx, fileIdx})
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

func (board Board) GetMoveSource(piece, fileIdx, rankIdx, player int) *Square {
	switch piece {
	case ROOK:
		return nil
	case KNIGHT:
		return nil
	case BISHOP:
		return nil
	case QUEEN:
		return nil
	case KING:
		return nil
	case PAWN:
		return board.getPawnSquares(piece, fileIdx, rankIdx, player)
	}
	return nil
}

func (board Board) getPawnSquares(piece, file, rank, player int) *Square {
	// The internal representation of the board has black at the "top", e.g. board[0], so -- is looking backwards for BLACK
	if player == WHITE {
		if board[rank+1][file] != nil && board[rank+1][file].Type == PAWN {
			return &Square{rank + 1, file}
		}
		if board[rank+2][file] != nil && board[rank+2][file].Type == PAWN {
			return &Square{rank + 2, file}
		}
	} else {
		if board[rank-1][file] != nil && board[rank-1][file].Type == PAWN {
			return &Square{rank - 1, file}
		}
		if board[rank-2][file] != nil && board[rank-2][file].Type == PAWN {
			return &Square{rank - 2, file}
		}
	}
	return nil
}
