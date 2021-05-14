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

type Bishop struct {
	Position Position
	White    bool
}

func (b *Bishop) Range() int {
	return 8
}

func (b *Bishop) Symbol() string {
	return "B"
}

func (b *Bishop) Score() float64 {
	return 3.3
}

func (b *Bishop) SetPosition(p Position) {
	b.Position = p
}

func (b *Bishop) Clone() Piece {
	newB := *b
	return &newB
}

func (b *Bishop) GetPosition() Position {
	return b.Position
}

func (b *Bishop) GetDirections() []Direction {
	return []Direction{
		{-1, -1},
		{1, 1},
		{1, -1},
		{-1, 1},
	}
}

func (b *Bishop) CaptureDirections() []Direction {
	return nil
}

func (b *Bishop) CapturesDifferently() bool {
	return false
}

func (b *Bishop) IsWhite() bool {
	return b.White
}
