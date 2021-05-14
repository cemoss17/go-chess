package board

import (
	"fmt"
	"log"
	"testing"
)

func TestPawnMove(t *testing.T) {
	b := StartingPosition
	b.Turn = false
	children, _ := b.GenerateChildren()
	for _, ch := range children {
		fmt.Println(ch.PrevMove.String())
	}
	log.Printf("%#v\n", b)
	b = children[1]
	log.Printf("%#v\n", b)
	children, _ = b.GenerateChildren()
	for _, ch := range children {
		fmt.Println(ch.PrevMove.String())
	}
}

func TestPawnOnly(t *testing.T) {
	b := &Board{
		WhitePieces: map[string]Piece{
			"e2": &Pawn{Position: Position{X: 4, Y: 1}, White: true},
		},
		BlackPieces: map[string]Piece{
			"e7": &Pawn{Position: Position{X: 4, Y: 6}, White: false},
		},
		Turn: true,
	}
	children, _ := b.GenerateChildren()
	for _, ch := range children {
		fmt.Println(ch.PrevMove.String())
	}
	log.Printf("%#v\n", b)
	b = children[1]
	log.Printf("%#v\n", b)
	children, _ = b.GenerateChildren()
	for _, ch := range children {
		fmt.Println(ch.PrevMove.String())
	}
}
