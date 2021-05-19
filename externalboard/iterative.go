package externalboard

import (
	"fmt"
	"log"
	"sort"
	"time"

	mg "github.com/dylhunn/dragontoothmg"
	"github.com/notnil/chess"
)

type IterativeDeepener struct {
	moveMap  map[[16]byte]*HashResult
	boardMap map[uint64]mg.Move
	Game     *Game
}

func NewIterativeDeepener() *IterativeDeepener {
	return &IterativeDeepener{
		moveMap:  make(map[[16]byte]*HashResult),
		boardMap: make(map[uint64]mg.Move),
		Game:     NewGame(),
	}
}

func IterativeDeepenerFromFEN(fen string) *IterativeDeepener {
	gf, err := chess.FEN(fen)
	if err != nil {
		log.Fatalf("Invalid FEN: %v\n", err)
	}
	return &IterativeDeepener{
		moveMap: make(map[[16]byte]*HashResult),
		Game:    &Game{chess.NewGame(gf)},
	}
}

// have to keep the entire position to avoid collisions
type HashResult struct {
	moves []*chess.Move
	FEN   string
}

func (it *IterativeDeepener) DeepeningGame(playerColor chess.Color, depth int, timeout time.Duration) {
	game := it.Game
	for {
		pos := game.GetPosition()
		fmt.Println(pos.GetBoard().Draw())
		if pos.Turn() == playerColor {
			for {
				mv := getMove()
				err := game.MoveStr(mv)
				if err != nil {
					log.Printf("Error advancing with move %s: %v\n", mv, err)
					continue
				}
				break
			}
		} else {
			engineMove, score := it.deepeningSearch(depth, timeout)
			fmt.Printf("Score is %d, move selected %s\n", score, engineMove)
			err := game.Move(engineMove)
			if err != nil {
				log.Printf("Engine move cannot be applied: %v\n", err)
			}
		}
	}
}

func (it *IterativeDeepener) deepeningSearch(maxDepth int, timeout time.Duration) (*chess.Move, int) {
	pos := it.Game.GetPosition()
	var move *chess.Move
	var score int
	nodesSearched := 0
	startTime := time.Now()
	for i := 1; i <= maxDepth; i++ {
		move, score = it.alphabetaMove(pos, i, -INF, INF, &nodesSearched)
		dur := time.Now().Sub(startTime)
		if dur > timeout {
			break
		}
		fmt.Printf("Depth: %d, Score: %d, Move: %s\n", i, score, move)
	}
	fmt.Printf("Nodes searched: %d\n", nodesSearched)
	return move, score
}

func (it *IterativeDeepener) alphabetaMove(p *Position, depth, alpha, beta int, nodesSearched *int) (*chess.Move, int) {
	(*nodesSearched)++
	if depth == 0 {
		return nil, p.Score()
	}
	var moveToReturn *chess.Move
	var abscore int
	score := p.MAX()
	moves, _ := it.ValidMovesSorted(p)
	if len(moves) == 0 {
		return nil, p.Score()
	}
	scoreMap := make(map[string]*int)
	for _, m := range moves {
		nextPos := &Position{p.Update(m)}
		_, abscore = it.alphabetaMove(nextPos, depth-1, alpha, beta, nodesSearched)
		scoreToSet := abscore
		scoreMap[m.String()] = &scoreToSet
		if p.better(abscore, score) {
			moveToReturn = m
			score = abscore
		}
		if p.Turn() == chess.White {
			if p.better(score, alpha) {
				alpha = score
			}
		} else {
			if p.better(score, beta) {
				beta = score
			}
		}
		if alpha >= beta {
			break
		}
	}
	if depth > 0 { // this is a tunable parameter
		it.sortMoves(p, moves, scoreMap)
		it.setTopMoves(p, moves)
	}
	return moveToReturn, score
}

func (it *IterativeDeepener) sortMoves(p *Position, moves []*chess.Move, scoreMap map[string]*int) {
	turn := p.Turn()
	sort.Slice(moves, func(i, j int) bool {
		return it.compareMoveOrder(turn, scoreMap[moves[i].String()], scoreMap[moves[j].String()])
	})
}

func (it *IterativeDeepener) compareMoveOrder(t chess.Color, mv1, mv2 *int) bool {
	if mv1 == nil {
		if mv2 == nil {
			return true
		}
		return false
	}
	if mv2 == nil {
		return true
	}
	if t == chess.White {
		return *mv1 > *mv2
	}
	return *mv2 > *mv1
}

func (it *IterativeDeepener) setTopMoves(p *Position, moves []*chess.Move) {
	it.moveMap[p.Hash()] = &HashResult{
		FEN:   p.String(),
		moves: moves,
	}
}

func (it *IterativeDeepener) ValidMovesSorted(p *Position) ([]*chess.Move, bool) {
	hr := it.moveMap[p.Hash()]
	if hr == nil || hr.FEN != p.String() {
		return p.ValidMoves(), false
	}
	return hr.moves, true
}

func printMap(moveMap map[string]*int) {
	for mv, val := range moveMap {
		fmt.Printf("%s: %d, ", mv, *val)
	}
	fmt.Println()
}
