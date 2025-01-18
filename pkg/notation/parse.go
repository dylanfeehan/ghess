package notation

import (
	"errors"
)

type Move struct {
	Piece string
	Rank  string
	File  string
}

// allowed characters:
// +: check
// #: checkmate
// x: capture
// not yet, these are more for rebuilding the game, not necessarily commands of what to play.. we can internally build up a set of these and print it out at the end of the game though

// 1-8: ranks
// a-h: files
// rnqkbp: pieces
var validPieces map[string]struct{} = map[string]struct{}{
	"r": struct{}{},
	"n": struct{}{},
	"q": struct{}{},
	"k": struct{}{},
	"p": struct{}{},
	"b": struct{}{},
}
var validRanks map[string]struct{} = map[string]struct{}{
	"1": struct{}{},
	"2": struct{}{},
	"3": struct{}{},
	"4": struct{}{},
	"5": struct{}{},
	"6": struct{}{},
	"7": struct{}{},
	"8": struct{}{},
}
var validFiles map[string]struct{} = map[string]struct{}{
	"a": struct{}{},
	"b": struct{}{},
	"c": struct{}{},
	"d": struct{}{},
	"e": struct{}{},
	"f": struct{}{},
	"g": struct{}{},
	"h": struct{}{},
}

func validRank(rank string) bool {
	_, ok := validRanks[rank]
	return ok
}

func validFile(file string) bool {
	_, ok := validFiles[file]
	return ok
}

func validPiece(piece string) bool {
	_, ok := validPieces[piece]
	return ok
}

// what, where
// not handling single or double disambiguations yet
func Parse(move string) (*Move, error) {
	if (len(move)) == 0 {
		return nil, errors.New("IDFK")
	}
	if len(move) == 2 {
		file := string(move[0])
		rank := string(move[1])
		if validFile(file) && validRank(rank) {
			return &Move{
				Piece: "p",
				Rank:  rank,
				File:  file,
			}, nil
		} else {
			return nil, errors.New("Invalid file or rank")
		}
	}

	piece := string(move[0])
	file := string(move[1])
	rank := string(move[2])
	if validPiece(piece) && validFile(file) && validRank(rank) {
		return &Move{
			Piece: "p",
			Rank:  rank,
			File:  file,
		}, nil
	} else {
		return nil, errors.New("Invalid piece or file or rank")
	}
}
