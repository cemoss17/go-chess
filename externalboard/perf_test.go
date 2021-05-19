package externalboard

import (
	"testing"

	mg "github.com/dylhunn/dragontoothmg"
)

var (
	startingFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
)

func BenchmarkPerft(b *testing.B) {
	for i := 0; i < b.N; i++ {
		startingPos.ValidMoves()
	}
}

func BenchmarkPerftFast(b *testing.B) {
	board := mg.ParseFen(startingFEN)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.GenerateLegalMoves()
	}
}
