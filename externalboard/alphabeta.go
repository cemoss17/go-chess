package externalboard

import (
	"github.com/notnil/chess"
)

func (p *Position) alphabeta(depth, alpha, beta int) (*Position, int) {
	if depth == 0 {
		return p, p.Score()
	}
	var posToReturn *Position
	var alphaBetaPosition *Position
	var abscore int
	score := p.MAX()
	moves := p.ValidMoves()
	if len(moves) == 0 {
		return p, p.Score()
	}
	sortMoves(moves)
	for _, m := range moves {
		nextPos := &Position{p.Update(m)}
		alphaBetaPosition, abscore = nextPos.alphabeta(depth-1, alpha, beta)
		if nextPos.better(abscore, score) {
			posToReturn = alphaBetaPosition
			score = abscore
		}
		if nextPos.Turn() == chess.White {
			if nextPos.better(score, alpha) {
				alpha = score
			}
		} else {
			if nextPos.better(score, beta) {
				beta = score
			}
		}
		if alpha >= beta {
			break
		}
	}
	return posToReturn, score
}

func (p *Position) alphabetaMove(depth, alpha, beta int) (*chess.Move, int) {
	if depth == 0 {
		return nil, p.Score()
	}
	var moveToReturn *chess.Move
	var abscore int
	score := p.MAX()
	moves := p.ValidMoves()
	if len(moves) == 0 {
		return nil, p.Score()
	}
	sortMoves(moves)
	for _, m := range moves {
		nextPos := &Position{p.Update(m)}
		_, abscore = nextPos.alphabetaMove(depth-1, alpha, beta)
		if nextPos.better(abscore, score) {
			moveToReturn = m
			score = abscore
		}
		if nextPos.Turn() == chess.White {
			if nextPos.better(score, alpha) {
				alpha = score
			}
		} else {
			if nextPos.better(score, beta) {
				beta = score
			}
		}
		if alpha >= beta {
			break
		}
	}
	return moveToReturn, score
}

func (p *Position) better(score1, score2 int) bool {
	if p.Turn() == chess.White {
		return score1 > score2
	}
	return score1 < score2
}

func (p *Position) MAX() int {
	if p.Turn() == chess.White {
		return INF
	}
	return -INF
}

/*
func (b *Board) alphabeta(depth int, alpha, beta int) (*Board, int) {
	if depth == 0 {
		return b, b.Evaluate()
	}
	var boardToReturn *Board
	var bab *Board
	var abscore int
	if b.Turn {
		score := -INF
		moves := b.generateAllRawMoves()
		if len(moves) == 0 {
			return b, -INF
		}
		for _, m := range moves {
			nb := b.newBoardFromMove(m)
			bab, abscore = nb.alphabeta(depth-1, alpha, beta)
			if abscore > score {
				boardToReturn = bab
				score = abscore
			}
			if score > alpha {
				alpha = score
			}
			if alpha >= beta {
				break
			}
		}
		return boardToReturn, score
	}
	score := INF
	moves := b.generateAllRawMoves()
	for _, m := range moves {
		nb := b.newBoardFromMove(m)
		if len(moves) == 0 {
			return b, INF
		}
		bab, abscore = nb.alphabeta(depth-1, alpha, beta)
		if abscore < score {
			score = abscore
			boardToReturn = bab
		}
		if score < beta {
			beta = score
		}
		if alpha >= beta {
			break
		}
	}
	return boardToReturn, score
}
*/
