package externalboard

import (
	"sort"

	"github.com/notnil/chess"
)

func moveOrderingScore(m *chess.Move) int {
	if m.HasTag(chess.Capture) {
		return 2
	}
	if m.HasTag(chess.Check) {
		return 1
	}
	return 0
}

func sortMoves(moves []*chess.Move) {
	sort.Slice(moves, func(i, j int) bool { return moveOrderingScore(moves[i]) > moveOrderingScore(moves[j]) })
}
