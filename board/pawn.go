package board

/*

type Piece interface {
	Range() int
	Symbol() string
	GetPosition() Position
	GetDirections() []Direction
	CaptureDirections() []Direction
	CapturesDifferently() bool
	IsWhite() bool
}

*/

type Pawn struct {
	Position Position
	White    bool
}

func (pawn *Pawn) GetPosition() Position {
	return pawn.Position
}

func (pawn *Pawn) SetPosition(p Position) {
	pawn.Position = p
}

func (pawn *Pawn) Score() float64 {
	return 1.0
}

func (pawn *Pawn) Clone() Piece {
	newPawn := *pawn
	return &newPawn
}

func (pawn *Pawn) IsWhite() bool {
	return pawn.White
}

func (pawn *Pawn) CapturesDifferently() bool {
	return true
}

func (pawn *Pawn) Range() int {
	return 1
}

func (pawn *Pawn) Symbol() string {
	return ""
}

func (pawn *Pawn) CaptureDirections() []Direction {
	if pawn.White {
		return []Direction{{1, 1}, {-1, 1}}
	}
	return []Direction{{1, -1}, {-1, -1}}
}

func (pawn *Pawn) GetDirections() []Direction {
	if pawn.White {
		if pawn.Position.Y == 1 {
			return []Direction{{0, 2}, {0, 1}}
		}
		return []Direction{{0, 1}}
	}
	if pawn.Position.Y == 6 {
		return []Direction{{0, -2}, {0, -1}}
	}
	return []Direction{{0, -1}}
}
