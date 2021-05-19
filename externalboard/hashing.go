package externalboard

import (
	"github.com/notnil/chess"
)

var (
	validMovesMap = make(map[[16]byte][]*chess.Move)
)

// returns moves for that position, as well as if it was gotten from hash
func (p *Position) ValidMovesHash() ([]*chess.Move, bool) {
	hashedMoves := validMovesMap[p.Hash()]
	if hashedMoves == nil {
		moves := p.ValidMoves()
		return moves, false
	}
	return hashedMoves, true
}
