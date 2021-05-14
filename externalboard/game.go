package externalboard

import (
	"fmt"
	"log"

	"github.com/notnil/chess"
)

func AlphaBetaGame(playerColor chess.Color, depth int) {
	game := NewGame()
	for {
		pos := game.GetPosition()
		fmt.Println(pos.GetBoard().Draw())
		if pos.Turn() == playerColor {
			playerPossibleMoves := pos.ValidMoves()
			fmt.Printf("Possible moves: %q\n", playerPossibleMoves)
			for {
				mv := getMove()
				err := game.MoveStr(mv)
				if err != nil {
					log.Printf("Error advancing with move %s: %v\n", mv, err)
					continue
				}
				break
			}
		} else {
			engineMove, score := pos.alphabetaMove(depth, -INF, INF)
			fmt.Printf("Score is %d, move selected %s\n", score, engineMove)
			err := game.Move(engineMove)
			if err != nil {
				log.Printf("Engine move cannot be applied: %v\n", err)
			}
		}
	}
}

func getMove() string {
	fmt.Println("Enter Move: ")
	var moveStr string
	fmt.Scanln(&moveStr)
	return moveStr
}
