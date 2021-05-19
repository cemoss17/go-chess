package externalboard

import (
	ch "github.com/notnil/chess"
)

var (
	whitePawnMap = map[ch.Square]int{
		ch.H8: 0, ch.H7: 50, ch.H6: 10, ch.H5: 05, ch.H4: 00, ch.H3: 005, ch.H2: 5, ch.H1: 0,
		ch.G8: 0, ch.G7: 50, ch.G6: 10, ch.G5: 05, ch.G4: 00, ch.G3: -05, ch.G2: 10, ch.G1: 0,
		ch.F8: 0, ch.F7: 50, ch.F6: 20, ch.F5: 10, ch.F4: 00, ch.F3: -10, ch.F2: 20, ch.F1: 0,
		ch.E8: 0, ch.E7: 50, ch.E6: 30, ch.E5: 25, ch.E4: 20, ch.E3: 000, ch.E2: -20, ch.E1: 0,
		ch.D8: 0, ch.D7: 50, ch.D6: 30, ch.D5: 25, ch.D4: 20, ch.D3: 000, ch.D2: -20, ch.D1: 0,
		ch.C8: 0, ch.C7: 50, ch.C6: 20, ch.C5: 10, ch.C4: 00, ch.C3: -10, ch.C2: 10, ch.C1: 0,
		ch.B8: 0, ch.B7: 50, ch.B6: 10, ch.B5: 05, ch.B4: 00, ch.B3: -05, ch.B2: 10, ch.B1: 0,
		ch.A8: 0, ch.A7: 50, ch.A6: 10, ch.A5: 05, ch.A4: 00, ch.A3: 005, ch.A2: 5, ch.A1: 0,
	}
	blackPawnMap = map[ch.Square]int{
		ch.H8: 0, ch.H7: -5, ch.H6: -5, ch.H5: 000, ch.H4: -05, ch.H3: -10, ch.H2: -50, ch.H1: 0,
		ch.G8: 0, ch.G7: -10, ch.G6: 05, ch.G5: 000, ch.G4: -05, ch.G3: -10, ch.G2: -50, ch.G1: 0,
		ch.F8: 0, ch.F7: -20, ch.F6: 10, ch.F5: 000, ch.F4: -10, ch.F3: -20, ch.F2: -50, ch.F1: 0,
		ch.E8: 0, ch.E7: 20, ch.E6: 00, ch.E5: -20, ch.E4: -25, ch.E3: -30, ch.E2: -50, ch.E1: 0,
		ch.D8: 0, ch.D7: 20, ch.D6: 00, ch.D5: -20, ch.D4: -25, ch.D3: -30, ch.D2: -50, ch.D1: 0,
		ch.C8: 0, ch.C7: -20, ch.C6: 10, ch.C5: 000, ch.C4: -10, ch.C3: -20, ch.C2: -50, ch.C1: 0,
		ch.B8: 0, ch.B7: -10, ch.B6: 05, ch.B5: 000, ch.B4: -05, ch.B3: -10, ch.B2: -50, ch.B1: 0,
		ch.A8: 0, ch.A7: -5, ch.A6: -5, ch.A5: 000, ch.A4: -05, ch.A3: -10, ch.A2: -50, ch.A1: 0,
	}
	whiteKnightMap = map[ch.Square]int{
		ch.H8: -50, ch.H7: -40, ch.H6: -30, ch.H5: -30, ch.H4: -30, ch.H3: -30, ch.H2: -40, ch.H1: -50,
		ch.G8: -40, ch.G7: -20, ch.G6: 000, ch.G5: 5, ch.G4: 000, ch.G3: 5, ch.G2: -20, ch.G1: -40,
		ch.F8: -30, ch.F7: 000, ch.F6: 10, ch.F5: 15, ch.F4: 15, ch.F3: 10, ch.F2: 0, ch.F1: -30,
		ch.E8: -30, ch.E7: 000, ch.E6: 15, ch.E5: 20, ch.E4: 20, ch.E3: 15, ch.E2: 5, ch.E1: -30,
		ch.D8: -30, ch.D7: 000, ch.D6: 15, ch.D5: 20, ch.D4: 20, ch.D3: 15, ch.D2: 5, ch.D1: -30,
		ch.C8: -30, ch.C7: 000, ch.C6: 10, ch.C5: 15, ch.C4: 15, ch.C3: 10, ch.C2: 0, ch.C1: -30,
		ch.B8: -40, ch.B7: -20, ch.B6: 000, ch.B5: 5, ch.B4: 000, ch.B3: 5, ch.B2: -20, ch.B1: -40,
		ch.A8: -50, ch.A7: -40, ch.A6: -30, ch.A5: -30, ch.A4: -30, ch.A3: -30, ch.A2: -40, ch.A1: -50,
	}
	blackKnightMap = map[ch.Square]int{
		ch.H8: 50, ch.H7: 40, ch.H6: 30, ch.H5: 30, ch.H4: 30, ch.H3: 30, ch.H2: 40, ch.H1: 50,
		ch.G8: 40, ch.G7: 20, ch.G6: 0, ch.G5: -5, ch.G4: 000, ch.G3: -5, ch.G2: 20, ch.G1: 40,
		ch.F8: 30, ch.F7: 00, ch.F6: -10, ch.F5: -15, ch.F4: -15, ch.F3: -10, ch.F2: 00, ch.F1: 30,
		ch.E8: 30, ch.E7: 00, ch.E6: -15, ch.E5: -20, ch.E4: -20, ch.E3: -15, ch.E2: -5, ch.E1: 30,
		ch.D8: 30, ch.D7: 00, ch.D6: -15, ch.D5: -20, ch.D4: -20, ch.D3: -15, ch.D2: -5, ch.D1: 30,
		ch.C8: 30, ch.C7: 00, ch.C6: -10, ch.C5: -15, ch.C4: -15, ch.C3: -10, ch.C2: 00, ch.C1: 30,
		ch.B8: 40, ch.B7: 20, ch.B6: 0, ch.B5: -5, ch.B4: 00, ch.B3: -5, ch.B2: 20, ch.B1: 40,
		ch.A8: 50, ch.A7: 40, ch.A6: 30, ch.A5: 30, ch.A4: 30, ch.A3: 30, ch.A2: 40, ch.A1: 50,
	}
	whiteKingMapStart = map[ch.Square]int{
		ch.H8: -30, ch.H7: -30, ch.H6: -30, ch.H5: -30, ch.H4: -20, ch.H3: -10, ch.H2: 20, ch.H1: 20,
		ch.G8: -40, ch.G7: -40, ch.G6: -40, ch.G5: -40, ch.G4: -30, ch.G3: -20, ch.G2: 20, ch.G1: 35,
		ch.F8: -40, ch.F7: -40, ch.F6: -40, ch.F5: -40, ch.F4: -30, ch.F3: -20, ch.F2: 00, ch.F1: 10,
		ch.E8: -50, ch.E7: -50, ch.E6: -50, ch.E5: -50, ch.E4: -40, ch.E3: -20, ch.E2: 00, ch.E1: 00,
		ch.D8: -50, ch.D7: -50, ch.D6: -50, ch.D5: -50, ch.D4: -40, ch.D3: -20, ch.D2: 00, ch.D1: 00,
		ch.C8: -40, ch.C7: -40, ch.C6: -40, ch.C5: -40, ch.C4: -30, ch.C3: -20, ch.C2: 00, ch.C1: 10,
		ch.B8: -40, ch.B7: -40, ch.B6: -40, ch.B5: -40, ch.B4: -30, ch.B3: -20, ch.B2: 20, ch.B1: 30,
		ch.A8: -30, ch.A7: -30, ch.A6: -30, ch.A5: -30, ch.A4: -20, ch.A3: -10, ch.A2: 20, ch.A1: 20,
	}
	blackKingMapStart = map[ch.Square]int{
		ch.H8: -20, ch.H7: -20, ch.H6: 10, ch.H5: 20, ch.H4: 30, ch.H3: 30, ch.H2: 30, ch.H1: 30,
		ch.G8: -35, ch.G7: -20, ch.G6: 20, ch.G5: 30, ch.G4: 40, ch.G3: 40, ch.G2: 40, ch.G1: 40,
		ch.F8: -10, ch.F7: 000, ch.F6: 20, ch.F5: 30, ch.F4: 40, ch.F3: 40, ch.F2: 40, ch.F1: 40,
		ch.E8: 000, ch.E7: 000, ch.E6: 20, ch.E5: 40, ch.E4: 50, ch.E3: 50, ch.E2: 50, ch.E1: 50,
		ch.D8: 000, ch.D7: 000, ch.D6: 20, ch.D5: 40, ch.D4: 50, ch.D3: 50, ch.D2: 50, ch.D1: 50,
		ch.C8: -10, ch.C7: 000, ch.C6: 20, ch.C5: 30, ch.C4: 40, ch.C3: 40, ch.C2: 40, ch.C1: 40,
		ch.B8: -30, ch.B7: -20, ch.B6: 20, ch.B5: 30, ch.B4: 40, ch.B3: 40, ch.B2: 40, ch.B1: 40,
		ch.A8: -20, ch.A7: -20, ch.A6: 10, ch.A5: 20, ch.A4: 30, ch.A3: 30, ch.A2: 30, ch.A1: 30,
	}
	whiteRookMap = map[ch.Square]int{
		ch.H8: 0, ch.H7: 05, ch.H6: -5, ch.H5: -5, ch.H4: -5, ch.H3: -5, ch.H2: -5, ch.H1: 5,
		ch.G8: 0, ch.G7: 10, ch.G6: 00, ch.G5: 00, ch.G4: 00, ch.G3: 00, ch.G2: 00, ch.G1: 0,
		ch.F8: 0, ch.F7: 10, ch.F6: 00, ch.F5: 00, ch.F4: 00, ch.F3: 00, ch.F2: 00, ch.F1: 0,
		ch.E8: 0, ch.E7: 10, ch.E6: 00, ch.E5: 00, ch.E4: 00, ch.E3: 00, ch.E2: 00, ch.E1: 5,
		ch.D8: 0, ch.D7: 10, ch.D6: 00, ch.D5: 00, ch.D4: 00, ch.D3: 00, ch.D2: 00, ch.D1: 5,
		ch.C8: 0, ch.C7: 10, ch.C6: 00, ch.C5: 00, ch.C4: 00, ch.C3: 00, ch.C2: 00, ch.C1: 0,
		ch.B8: 0, ch.B7: 10, ch.B6: 00, ch.B5: 00, ch.B4: 00, ch.B3: 00, ch.B2: 00, ch.B1: 0,
		ch.A8: 0, ch.A7: 05, ch.A6: -5, ch.A5: -5, ch.A4: -5, ch.A3: -5, ch.A2: -5, ch.A1: 0,
	}
	blackRookMap = map[ch.Square]int{
		ch.H8: 00, ch.H7: 5, ch.H6: 5, ch.H5: 5, ch.H4: 5, ch.H3: 5, ch.H2: -5, ch.H1: 0,
		ch.G8: 00, ch.G7: 0, ch.G6: 0, ch.G5: 0, ch.G4: 0, ch.G3: 0, ch.G2: -10, ch.G1: 0,
		ch.F8: 00, ch.F7: 0, ch.F6: 0, ch.F5: 0, ch.F4: 0, ch.F3: 0, ch.F2: -10, ch.F1: 0,
		ch.E8: -5, ch.E7: 0, ch.E6: 0, ch.E5: 0, ch.E4: 0, ch.E3: 0, ch.E2: -10, ch.E1: 0,
		ch.D8: -5, ch.D7: 0, ch.D6: 0, ch.D5: 0, ch.D4: 0, ch.D3: 0, ch.D2: -10, ch.D1: 0,
		ch.C8: 00, ch.C7: 0, ch.C6: 0, ch.C5: 0, ch.C4: 0, ch.C3: 0, ch.C2: -10, ch.C1: 0,
		ch.B8: 00, ch.B7: 0, ch.B6: 0, ch.B5: 0, ch.B4: 0, ch.B3: 0, ch.B2: -10, ch.B1: 0,
		ch.A8: 00, ch.A7: 5, ch.A6: 5, ch.A5: 5, ch.A4: 5, ch.A3: 5, ch.A2: -5, ch.A1: 0,
	}
	zeroBonus  = map[ch.Square]int{}
	pieceBonus = map[ch.Piece]map[ch.Square]int{
		ch.WhiteKing:   whiteKingMapStart,
		ch.BlackKing:   blackKingMapStart,
		ch.WhiteKnight: whiteKnightMap,
		ch.BlackKnight: blackKnightMap,
		ch.WhiteQueen:  zeroBonus, // for now, no bonus is given to queen positions
		ch.BlackQueen:  zeroBonus, // for now, no bonus is given to queen positions
		ch.WhitePawn:   whitePawnMap,
		ch.BlackPawn:   blackPawnMap,
		ch.WhiteBishop: zeroBonus, // for now, no bonus is given to bishop positions
		ch.BlackBishop: zeroBonus, // for now, no bonus is given to bishop positions
		ch.WhiteRook:   whiteRookMap,
		ch.BlackRook:   blackRookMap,
	}
	pieceValue = map[ch.Piece]int{
		ch.WhiteKnight: 320,
		ch.BlackKnight: -320,
		ch.WhiteQueen:  900,
		ch.BlackQueen:  -900,
		ch.WhitePawn:   100,
		ch.BlackPawn:   -100,
		ch.WhiteBishop: 330,
		ch.BlackBishop: -330,
		ch.WhiteRook:   500,
		ch.BlackRook:   -500,
	}
)

func pieceScore(p ch.Piece) int {
	var score int
	switch p.Type() {
	case ch.Rook:
		score = 500
	case ch.Bishop:
		score = 330
	case ch.Knight:
		score = 320
	case ch.Queen:
		score = 900
	case ch.Pawn:
		score = 100
	}

	if p.Color() == ch.Black {
		score = -score
	}

	return score
}

// calculate the score of the piece by its type and the square it occupies
func pieceScoreBonus(p ch.Piece, sq ch.Square) int {
	return pieceValue[p] + pieceBonus[p][sq]
}
