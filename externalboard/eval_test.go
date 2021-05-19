package externalboard

import (
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
