package externalboard

import (
	"fmt"
	"testing"

	"github.com/notnil/chess"
)

var (
	startingPos = &Position{chess.NewGame().Position()}
)

func TestStarting(t *testing.T) {
	if score := startingPos.Score(); score != 0 {
		t.Errorf("The starting position score should be 0, got %d", score)
	}
}

func TestAlphaBetaStarting(t *testing.T) {
	finalPos, score := startingPos.alphabeta(7, -INF, INF)
	fmt.Printf("Score: %d\n", score)
	fmt.Println(finalPos)
}

func TestAlphaBetaMoveStarting(t *testing.T) {
	move, score := startingPos.alphabetaMove(7, -INF, INF)
	fmt.Printf("Score: %d\n", score)
	fmt.Println(move)
}
