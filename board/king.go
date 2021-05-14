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

type King struct {
	Position Position
	White    bool
}

func (k *King) Range() int {
	return 1
}

func (k *King) Symbol() string {
	return "K"
}

func (k *King) IsWhite() bool {
	return k.White
}

func (k *King) Score() float64 {
	return 0.0
}

func (k *King) SetPosition(p Position) {
	k.Position = p
}

func (k *King) Clone() Piece {
	newK := *k
	return &newK
}

func (k *King) GetPosition() Position {
	return k.Position
}

func (k *King) GetDirections() []Direction {
	return []Direction{
		{-1, -1},
		{1, 1},
		{1, -1},
		{-1, 1},
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
}

func (k *King) CaptureDirections() []Direction {
	return nil
}

func (k *King) CapturesDifferently() bool {
	return false
}
