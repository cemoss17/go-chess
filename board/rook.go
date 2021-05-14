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

type Rook struct {
	Position Position
	White    bool
}

func (r *Rook) Range() int {
	return 7
}

func (r *Rook) Symbol() string {
	return "R"
}

func (r *Rook) SetPosition(p Position) {
	r.Position = p
}

func (r *Rook) Score() float64 {
	return 4.9
}

func (r *Rook) Clone() Piece {
	newR := *r
	return &newR
}

func (r *Rook) GetPosition() Position {
	return r.Position
}

func (r *Rook) GetDirections() []Direction {
	return []Direction{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
}

func (r *Rook) CaptureDirections() []Direction {
	return nil
}

func (r *Rook) CapturesDifferently() bool {
	return false
}

func (r *Rook) IsWhite() bool {
	return r.White
}
