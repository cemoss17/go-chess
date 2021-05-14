package board

import (
	"fmt"
)

var (
	xToLetter = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	LetterToX = map[byte]int8{
		'a': 0,
		'b': 1,
		'c': 2,
		'd': 3,
		'e': 4,
		'f': 5,
		'g': 6,
		'h': 7,
	}
)

type Piece interface {
	Range() int
	Symbol() string
	GetPosition() Position
	GetDirections() []Direction
	CaptureDirections() []Direction
	CapturesDifferently() bool
	SetPosition(Position)
	Score() float64
	Clone() Piece
	IsWhite() bool
}

func pieceCapDiff(p int8) bool {
	return p == 1 || p == -1
}

func isPawn(p int8) bool {
	return p == 1 || p == -1
}

func isKing(p int8) bool {
	return p == 6 || p == -6
}

func pieceToFEN(p int8) string {
	return pieceFEN[p]
}

var (
	pieceFEN = map[int8]string{
		1:  "P",
		2:  "N",
		3:  "B",
		4:  "R",
		5:  "Q",
		6:  "K",
		-1: "p",
		-2: "n",
		-3: "b",
		-4: "r",
		-5: "q",
		-6: "k",
	}
)

var (
	KnightDirs = []Direction{
		{1, 2},
		{1, -2},
		{-2, -1},
		{-2, 1},
		{2, -1},
		{2, 1},
		{-1, -2},
		{-1, 2},
	}
	BishopDirs = []Direction{
		{-1, -1},
		{1, 1},
		{1, -1},
		{-1, 1},
	}
	RookDirs = []Direction{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	QueenDirs = []Direction{
		{-1, -1},
		{1, 1},
		{1, -1},
		{-1, 1},
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
)

func pieceCapDirs(p int8) []Direction {
	switch p {
	case 1:
		return []Direction{{1, 1}, {-1, 1}}
	case -1:
		return []Direction{{1, -1}, {-1, -1}}
	}
	return []Direction{}
}

func pieceDirs(p int8) []Direction {
	switch p {
	case 1:
		return []Direction{{0, 1}}
	case 2:
		return KnightDirs
	case 3:
		return BishopDirs
	case 4:
		return RookDirs
	case 5:
		return QueenDirs
	case 6:
		// king and queen directions are the same
		return QueenDirs
	case -1:
		return []Direction{{0, -1}}
	case -2:
		return KnightDirs
	case -3:
		return BishopDirs
	case -4:
		return RookDirs
	case -5:
		return QueenDirs
	case -6:
		return QueenDirs
	}
	return []Direction{}
}

func pieceScore(p int8) int {

	switch p {
	case 1:
		return 10
	case 2:
		return 31
	case 3:
		return 32
	case 4:
		return 50
	case 5:
		return 90
	case 6:
		return 0
	case -1:
		return -10
	case -2:
		return -31
	case -3:
		return -32
	case -4:
		return -50
	case -5:
		return -90
	case -6:
		return 0
	}
	return 0
}

func pieceRange(p, x, y int8) int {
	switch p {
	case 1:
		if y == 1 {
			return 2
		}
		return 1
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 7
	case 5:
		return 7
	case 6:
		return 1
	case -1:
		if y == 6 {
			return 2
		}
		return 1
	case -2:
		return 1
	case -3:
		return 7
	case -4:
		return 7
	case -5:
		return 7
	case -6:
		return 1
	}
	return 0
}

func inBounds(x, y int8) bool {
	return x >= 0 && x < 8 && y >= 0 && y < 8
}

type Position struct {
	X int8
	Y int8
}

func (p *Position) Notation() string {
	return xToLetter[p.X] + fmt.Sprintf("%d", p.Y+1)
}

func (p *Position) InBounds() bool {
	return p.X >= 0 && p.X < 8 && p.Y >= 0 && p.Y < 8
}

func (p *Position) AddDirection(d Direction) Position {
	return Position{p.X + d.X, p.Y + d.Y}
}

func (p *Position) ApplyDirection(d Direction) {
	p.X = p.X + d.X
	p.Y = p.Y + d.Y
}

type Direction struct {
	X int8
	Y int8
}

type Movement struct {
	InitialPosition Position
	FinalPosition   Position
	FinalValue      int8
}

type Move []Movement

func (m *Move) String() string {
	if len(*m) == 0 {
		return "empty move"
	}
	if len(*m) == 2 {
		// castling move
	}
	return fmt.Sprintf("%s-%s", (*m)[0].InitialPosition.Notation(), (*m)[0].FinalPosition.Notation())
}
