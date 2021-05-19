package externalboard

import (
	"testing"
	"time"

	mg "github.com/dylhunn/dragontoothmg"
)

var (
	WhiteObviousThreeCaptures = "6k1/q4ppp/8/8/3Q3r/4n3/1b6/4K3 w - - 0 1"
	BlackObviousThreeCaptures = "6k1/q4ppp/8/8/3Q3r/4n3/1b6/4K3 b - - 0 1"
	queenMateFen              = "6k1/8/5KP1/8/8/8/8/8 w - - 0 1"
)

/*
func TestMoveOrder(t *testing.T) {
	it := IterativeDeepenerFromFEN(WhiteObviousThreeCaptures)
	m, score := it.deepeningSearch(4, time.Second*1)
	fmt.Printf("Move: %s\n", m)
	if score != INF {
		t.Errorf("Score is %d, not %d\n", score, INF)
	}
}

func TestMoveOrderBlack(t *testing.T) {
	it := IterativeDeepenerFromFEN(BlackObviousThreeCaptures)
	m, score := it.deepeningSearch(4, time.Second*1)
	fmt.Printf("Move: %s\n", m)
	if score != INF {
		t.Errorf("Score is %d, not %d\n", score, INF)
	}
}
*/

func TestFast(t *testing.T) {
	b := mg.ParseFen(startingFEN)
	it := NewIterativeDeepener()

	it.deepeningFast(b, 6, time.Second*1)
}

func TestQueenMate(t *testing.T) {
	b := mg.ParseFen(queenMateFen)
	it := NewIterativeDeepener()
	it.deepeningFast(b, 20, time.Second*1)
}
