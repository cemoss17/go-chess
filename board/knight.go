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

type Knight struct {
	Position Position
	White    bool
}

func (k *Knight) Range() int {
	return 1
}

func (k *Knight) Symbol() string {
	return "N"
}

func (k *Knight) Score() float64 {
	return 3.1
}

func (k *Knight) Clone() Piece {
	newKnight := *k
	return &newKnight
}

func (k *Knight) SetPosition(p Position) {
	k.Position = p
}

func (k *Knight) GetPosition() Position {
	return k.Position
}

func (k *Knight) GetDirections() []Direction {
	return []Direction{
		{1, 2},
		{1, -2},
		{-2, -1},
		{-2, 1},
		{2, -1},
		{2, 1},
		{-1, 1},
		{-1, -2},
		{-1, 2},
	}
}

func (k *Knight) IsWhite() bool {
	return k.White
}

func (k *Knight) CaptureDirections() []Direction {
	return nil
}

func (k *Knight) CapturesDifferently() bool {
	return false
}
