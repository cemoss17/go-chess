package externalboard

import "github.com/notnil/chess"

func pieceScore(p chess.Piece) int {
	var score int
	switch p.Type() {
	case chess.Rook:
		score = 50
	case chess.Bishop:
		score = 32
	case chess.Knight:
		score = 31
	case chess.Queen:
		score = 89
	}

	if p.Color() == chess.Black {
		score = -score
	}

	return score
}
