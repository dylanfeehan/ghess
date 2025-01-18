package board

const (
	WHITE = iota
	BLACK = iota
)

const (
	ROOK   = iota
	KNIGHT = iota
	BISHOP = iota
	KING   = iota
	QUEEN  = iota
	PAWN   = iota
)

const (
	R = ROOK
	N = KNIGHT
	B = BISHOP
	K = KING
	Q = QUEEN
	P = PAWN
)

var whitePlayerInitMap []int = []int{R, N, B, Q, K, B, N, R}
var blackPlayerInitMap []int = []int{R, N, B, K, Q, B, N, R}

var whitePieces map[int]string = map[int]string{
	R: "♖",
	N: "♘",
	B: "♗",
	K: "♔",
	Q: "♕",
	P: "♙",
}

var blackPieces map[int]string = map[int]string{
	R: "♜",
	N: "♞",
	B: "♝",
	K: "♚",
	Q: "♛",
	P: "♟",
}

var strToInt map[string]int = map[string]int{
	"r": R,
	"n": N,
	"b": B,
	"k": K,
	"q": Q,
	"p": P,
}

var fileToInt map[string]int = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
}

func FileStrToInt(file string) int {
	return fileToInt[file]
}

func PieceStrToInt(piece string) int {
	return strToInt[piece]
}
