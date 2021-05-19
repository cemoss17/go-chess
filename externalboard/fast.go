package externalboard

import (
	"fmt"
	"math/bits"
	"time"

	mg "github.com/dylhunn/dragontoothmg"
)

const (
	QUEEN_SCORE  = 900
	BISHOP_SCORE = 330
	KNIGHT_SCORE = 320
	PAWN_SCORE   = 100
	ROOK_SCORE   = 500
)

var (
	StartingFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
)

func MAX(wtomove bool) int {
	if wtomove {
		return -INF
	}
	return INF
}

func boardScore(b *mg.Board) int {
	return bitboardsScore(b.White) - bitboardsScore(b.Black)
}

func bitboardsScore(bb mg.Bitboards) int {
	return bits.OnesCount64(bb.Knights)*KNIGHT_SCORE +
		bits.OnesCount64(bb.Pawns)*PAWN_SCORE +
		bits.OnesCount64(bb.Bishops)*PAWN_SCORE +
		bits.OnesCount64(bb.Rooks)*ROOK_SCORE +
		bits.OnesCount64(bb.Queens)*QUEEN_SCORE
}

func endScore(b *mg.Board) int {
	if !b.OurKingInCheck() {
		return 0 // stalemate
	}
	if b.Wtomove {
		return -INF
	}
	return INF
}

func boardBetter(wtomove bool, sc1, sc2 int) bool {
	if wtomove {
		return sc1 > sc2
	}
	return sc1 < sc2
}

func (it *IterativeDeepener) FastGame(playerWhite bool, maxDepth int, timeout time.Duration) {
	b := mg.ParseFen(StartingFEN)
	for {
		if playerWhite == b.Wtomove {
			mvStr := getMove()
			playerMove, err := mg.ParseMove(mvStr)
			if err != nil {
				fmt.Printf("Invalid move %s, try again\n", mvStr)
			}
			moves := b.GenerateLegalMoves()
			if len(moves) == 0 {
				fmt.Println("Game over")
				break
			}
			for _, mv := range moves {
				if playerMove == mv {
					b.Apply(mv)
					continue
				}
			}
		} else {
			move, _ := it.deepeningFast(b, maxDepth, timeout)
			b.Apply(move)
			fmt.Printf("Engine played %s\n", move.String())
		}
	}
}

func (it *IterativeDeepener) deepeningFast(b mg.Board, toDepth int, maxDur time.Duration) (mg.Move, int) {
	var nodesSearched int
	var move mg.Move
	var score int
	startTime := time.Now()
	for i := 0; i <= toDepth; i++ {
		move, score = it.alphabetaFast(b, i, -INF, INF, &nodesSearched)
		fmt.Printf("Depth: %d\tmove: %s\tscore: %d\n", i, move.String(), score)
		if time.Now().Sub(startTime) > maxDur {
			break
		}
	}
	fmt.Printf("Nodes Searched: %d\n", nodesSearched)
	return move, score
}

func (it *IterativeDeepener) alphabetaFast(b mg.Board, depth, alpha, beta int, nodesSearched *int) (mg.Move, int) {
	(*nodesSearched)++
	if depth == 0 {
		return 0, boardScore(&b)
	}
	var abscore int
	moves := b.GenerateLegalMoves()
	if len(moves) == 0 {
		return 0, endScore(&b)
	}
	prevTopMove := it.boardMap[b.Hash()]
	if prevTopMove != 0 {
		for i, move := range moves {
			if move == prevTopMove {
				moves[0], moves[i] = moves[i], moves[0]
			}
		}
	}
	var moveToReturn = moves[0]
	var nextTopMove = moves[0]

	wtomove := b.Wtomove
	score := MAX(wtomove)
	topmovescore := score
	for _, m := range moves {
		unapply := b.Apply(m)
		_, abscore = it.alphabetaFast(b, depth-1, alpha, beta, nodesSearched)
		if boardBetter(wtomove, abscore, topmovescore) {
			topmovescore = abscore
			nextTopMove = m
		}
		unapply()
		if boardBetter(wtomove, abscore, score) {
			moveToReturn = m
			score = abscore
		}
		if b.Wtomove {
			if boardBetter(wtomove, score, alpha) {
				alpha = score
			}
		} else {
			if boardBetter(wtomove, score, beta) {
				beta = score
			}
		}
		if alpha >= beta {
			break
		}
	}
	it.boardMap[b.Hash()] = nextTopMove
	return moveToReturn, score
}
