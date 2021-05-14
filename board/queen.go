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

type Queen struct {
	Position Position
	White    bool
}

func (q *Queen) Range() int {
	return 7
}

func (q *Queen) Symbol() string {
	return "Q"
}

func (q *Queen) SetPosition(p Position) {
	q.Position = p
}

func (q *Queen) Score() float64 {
	return 9.2
}

func (q *Queen) Clone() Piece {
	newQ := *q
	return &newQ
}

func (q *Queen) GetPosition() Position {
	return q.Position
}

func (q *Queen) GetDirections() []Direction {
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

func (q *Queen) CaptureDirections() []Direction {
	return nil
}

func (q *Queen) CapturesDifferently() bool {
	return false
}

func (q *Queen) IsWhite() bool {
	return q.White
}
