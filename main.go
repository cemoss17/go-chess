package main

import (
	"time"

	"github.com/cemoss17/go-chess/externalboard"
)

func main() {
	it := externalboard.NewIterativeDeepener()
	it.FastGame(true, 25, time.Second*2)
}
