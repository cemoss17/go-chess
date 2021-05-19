package externalboard

import "github.com/notnil/chess"

const (
	INF = 100000
)

type Game struct {
	*chess.Game
}

func NewGame() *Game {
	return &Game{
		chess.NewGame(),
	}
}

type Position struct {
	*chess.Position
}

func (g *Game) GetPosition() *Position {
	return &Position{
		g.Position(),
	}
}

type Board struct {
	*chess.Board
}

func (p *Position) GetBoard() *Board {
	return &Board{
		p.Board(),
	}
}

func (p *Position) Score() int {
	if method := p.Status(); method != chess.NoMethod {
		if method != chess.Checkmate {
			return 0
		}
		if p.Turn() == chess.White {
			return -INF
		}
		return INF
	}

	var score int

	for sq, p := range p.Board().SquareMap() {
		score += pieceScoreBonus(p, sq)
	}

	return score
}
