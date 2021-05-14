package main

import (
	"github.com/cemoss17/go-chess/externalboard"
	"github.com/notnil/chess"
)

func main() {
	externalboard.AlphaBetaGame(chess.White, 5)
}
