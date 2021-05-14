package board

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/notnil/chess"
)

const (
	SCORE_MAX = 100000
	INF       = 1000000
)

func (b *Board) draw() {
	fen, err := chess.FEN(b.FEN())
	if err != nil {
		log.Printf("FEN is: %s\n", b.FEN())
		log.Printf("Error validating FEN: %v\n", err)
	}
	game := chess.NewGame(fen)
	fmt.Println(game.Position().Board().Draw())
}

func (b *Board) Draw() string {
	drawStr := " A B C D E F G H\n"
	for j := 7; j >= 0; j-- {
		drawStr += fmt.Sprintf("%d", j+1)
		for i := 0; i < 8; i++ {
			piece := b.bb[i][j]
			if piece == 0 {
				drawStr += "- "
				continue
			}
			drawStr = drawStr + pieceToFEN(piece) + " "
		}
		drawStr += "\n"
	}
	return drawStr
}

var (
	StartingPosition = &Board{
		bb: [][]int8{
			{4, 1, 0, 0, 0, 0, -1, -4},
			{2, 1, 0, 0, 0, 0, -1, -2},
			{3, 1, 0, 0, 0, 0, -1, -3},
			{5, 1, 0, 0, 0, 0, -1, -5},
			{6, 1, 0, 0, 0, 0, -1, -6},
			{3, 1, 0, 0, 0, 0, -1, -3},
			{2, 1, 0, 0, 0, 0, -1, -2},
			{4, 1, 0, 0, 0, 0, -1, -4},
		},
		CanWhiteKingSideCastle:  true,
		CanBlackKingSideCastle:  true,
		CanWhiteQueenSideCastle: true,
		CanBlackQueenSideCastle: true,
		Turn:                    true,
		children:                nil,
	}
)

func (b *Board) FEN() string {
	fenStr := ""

	var zeroCount int

	for j := 7; j >= 0; j-- {
		for i := 0; i < 8; i++ {
			piece := b.bb[i][j]
			if piece != 0 {
				if zeroCount != 0 {
					fenStr = fenStr + fmt.Sprintf("%d", zeroCount)
					zeroCount = 0
				}
				fenStr = fenStr + pieceToFEN(piece)
			} else {
				zeroCount++
			}
		}
		if zeroCount != 0 {
			fenStr = fenStr + fmt.Sprintf("%d", zeroCount)
			zeroCount = 0
		}
		if j != 0 {
			fenStr = fenStr + "/"
		}
	}

	if b.Turn {
		fenStr += " w "
	} else {
		fenStr += " b "
	}

	castleCount := 0

	if b.CanWhiteKingSideCastle {
		castleCount++
		fenStr += "K"
	}

	if b.CanWhiteQueenSideCastle {
		castleCount++
		fenStr += "Q"
	}

	if b.CanBlackKingSideCastle {
		castleCount++
		fenStr += "k"
	}

	if b.CanBlackQueenSideCastle {
		castleCount++
		fenStr += "q"
	}

	if castleCount == 0 {
		fenStr += "-"
	}

	fenStr += " "

	//
	fenStr += "- 2 2"

	return fenStr
}

func (b *Board) canWhiteKingSideCastle() bool {
	return b.CanWhiteKingSideCastle && b.bb[4][0] == 6 && b.bb[5][0] == 0 && b.bb[6][0] == 0 && b.bb[7][0] == 4 && !b.areAnyPositionsCapturable(
		Position{4, 0},
		Position{5, 0},
		Position{6, 0},
	)
}

func (b *Board) canWhiteQueenSideCastle() bool {
	return b.CanWhiteQueenSideCastle && b.bb[4][0] == 6 && b.bb[3][0] == 0 && b.bb[2][0] == 0 && b.bb[1][0] == 0 && b.bb[0][0] == 4 && !b.areAnyPositionsCapturable(
		Position{4, 0},
		Position{3, 0},
		Position{2, 0},
	)
}

func (b *Board) canBlackKingSideCastle() bool {
	return b.CanBlackKingSideCastle && b.bb[4][7] == -6 && b.bb[5][7] == 0 && b.bb[6][7] == 0 && b.bb[7][7] == -4 && !b.areAnyPositionsCapturable(
		Position{4, 7},
		Position{5, 7},
		Position{6, 7},
	)
}

func (b *Board) canBlackQueenSideCastle() bool {
	return b.CanBlackQueenSideCastle && b.bb[4][7] == -6 && b.bb[3][7] == 0 && b.bb[2][7] == 0 && b.bb[1][7] == 0 && b.bb[0][7] == -4 && !b.areAnyPositionsCapturable(
		Position{4, 7},
		Position{3, 7},
		Position{2, 7},
	)
}

type Board struct {
	bb                      BitBoard
	PrevMove                Move
	CanWhiteKingSideCastle  bool
	CanWhiteQueenSideCastle bool
	CanBlackKingSideCastle  bool
	CanBlackQueenSideCastle bool
	Turn                    bool // white's turn if true, black's turn if false
	parent                  *Board
	children                []*Board
}

func (b *Board) isPositionCapturable(pos Position) bool {
	return b.canSideCapturePos(!b.Turn, pos)
}

func (b *Board) areAnyPositionsCapturable(positions ...Position) bool {
	for _, pos := range positions {
		if ok := b.isPositionCapturable(pos); ok {
			return true
		}
	}
	return false
}

func (b *Board) printValidMoves() {
	for _, nb := range b.children {
		fmt.Println(nb.PrevMove.String())
	}
}

func (b *Board) copyBoard() Board {
	newbb := make([][]int8, 8, 8)
	for i, _ := range b.bb {
		newbb[i] = make([]int8, len(b.bb[i]))
		copy(newbb[i], b.bb[i])
	}
	return Board{
		bb:       newbb,
		PrevMove: b.PrevMove,
		Turn:     b.Turn,
	}
}

func (b *Board) printBoard() {
	fmt.Printf("Turn: %v\n", b.Turn)
	for _, row := range b.bb {
		fmt.Printf("%+v\n", row)
	}
	fmt.Println("-----")
}

func (b *Board) trickleUp(par *Board) *Board {
	var currBoard *Board
	currBoard = b
	for currBoard.parent != par {
		currBoard = currBoard.parent
	}
	return currBoard
}

type GameResult struct {
	Done   bool
	Result int
}

func (b *Board) printChildren() {
	for _, ch := range b.children {
		ch.printBoard()
	}
}

func (gr *GameResult) String() string {
	if gr.Result == 1 {
		return "white win"
	} else if gr.Result == -1 {
		return "black win"
	}
	return "draw"
}

func turnToGameResult(turn bool) *GameResult {
	if turn {
		return &GameResult{true, 1}
	}
	return &GameResult{true, -1}
}

func (b *Board) minimax() (*Board, int) {
	if b.children == nil {
		return b, b.Evaluate()
	}
	var boardToReturn *Board
	if b.Turn {
		score := -INF
		for _, ch := range b.children {
			_, currScore := ch.minimax()
			if currScore > score {
				boardToReturn = ch
				score = currScore
			}
		}
		return boardToReturn, score
	}
	score := INF
	for _, ch := range b.children {
		_, currScore := ch.minimax()
		if currScore < score {
			boardToReturn = ch
			score = currScore
		}
	}
	return boardToReturn, score
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func turnToScore(b int) float64 {
	if b == 0 {
		return 0.0
	}
	if b == 1 {
		return 10000.0
	}
	return -10000.0
}

func (b *Board) alphabeta(depth int, alpha, beta int) (*Board, int) {
	if depth == 0 {
		return b, b.Evaluate()
	}
	var boardToReturn *Board
	var bab *Board
	var abscore int
	if b.Turn {
		score := -INF
		moves := b.generateAllRawMoves()
		if len(moves) == 0 {
			return b, -INF
		}
		for _, m := range moves {
			nb := b.newBoardFromMove(m)
			bab, abscore = nb.alphabeta(depth-1, alpha, beta)
			if abscore > score {
				boardToReturn = bab
				score = abscore
			}
			if score > alpha {
				alpha = score
			}
			if alpha >= beta {
				break
			}
		}
		return boardToReturn, score
	}
	score := INF
	moves := b.generateAllRawMoves()
	for _, m := range moves {
		nb := b.newBoardFromMove(m)
		if len(moves) == 0 {
			return b, INF
		}
		bab, abscore = nb.alphabeta(depth-1, alpha, beta)
		if abscore < score {
			score = abscore
			boardToReturn = bab
		}
		if score < beta {
			beta = score
		}
		if alpha >= beta {
			break
		}
	}
	return boardToReturn, score
}

type BitBoard [][]int8

func (b *Board) newBoardFromMove(mv Move) *Board {
	newb := b.copyBoard()
	newb.Turn = !newb.Turn
	for _, movement := range mv {
		newb.bb[movement.FinalPosition.X][movement.FinalPosition.Y] = newb.bb[movement.InitialPosition.X][movement.InitialPosition.Y]
		//newb.bb[movement.FinalPosition.X][movement.FinalPosition.Y] = movement.FinalValue
		newb.bb[movement.InitialPosition.X][movement.InitialPosition.Y] = 0
	}
	newb.PrevMove = mv
	newb.parent = b
	return &newb
}

func (b *Board) printAllPossibleMoves() {
	for _, ch := range b.children {
		fmt.Println(ch.PrevMove.String())
	}
}

func (b *Board) AdvanceBoard(m Move) (*Board, bool) {
	var isValid bool
	var childrenIndex int
	for i, board := range b.children {
		if compareMoves(board.PrevMove, m) {
			isValid = true
			childrenIndex = i
			break
		}
	}
	if !isValid {
		return nil, false
	}
	b.Turn = !b.Turn
	return b.children[childrenIndex], true
}

func loadStartingPosition() *Board {
	return StartingPosition
}

func (b *Board) PickRandomMove() Move {
	return b.children[0].PrevMove
}

func getMove(turn bool) (Move, error) {
	fmt.Println("Enter Move: ")
	var moveStr string
	fmt.Scanln(&moveStr)

	// castling move, handle as an exception
	if moveStr == "0-0" || moveStr == "0-0-0" {
		if turn {
			if moveStr == "0-0" {
				return WhiteKingSideCastleMove, nil
			}
			return WhiteQueenSideCastleMove, nil
		}
		if moveStr == "0-0" {
			return BlackKingSideCastleMove, nil
		}
		return BlackQueenSideCastleMove, nil
	}

	tokens := strings.Split(moveStr, "-")
	if len(tokens) != 2 {
		return Move{}, fmt.Errorf("Move has to involve 2 squares")
	}
	initPosStr := tokens[0]
	finalPosStr := tokens[1]
	initPos, err := parsePosition(initPosStr)
	if err != nil {
		return Move{}, err
	}
	finalPos, err := parsePosition(finalPosStr)
	if err != nil {
		return Move{}, err
	}
	return []Movement{Movement{initPos, finalPos, 0}}, nil
}

func parsePosition(posStr string) (Position, error) {
	if len(posStr) != 2 {
		return Position{}, fmt.Errorf("Square has to involve 2 parts, [a-h] and [1-8]")
	}
	x := LetterToX[posStr[0]]
	vertStr := string(posStr[1])
	vert, err := strconv.Atoi(vertStr)
	if err != nil {
		return Position{}, fmt.Errorf("Cannot parse the vertical number")
	}
	pos := Position{X: x, Y: int8(vert - 1)}
	fmt.Printf("posStr: %s, Position: %+v\n", posStr, pos)
	return pos, nil
}

func (b *Board) goToDepth(depth int, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth == 0 {
		return
	}
	if b.children == nil {
		moves := b.generateAllRawMoves()
		if len(moves) == 0 {
			return
		}
		for _, m := range moves {
			b.children = append(b.children, b.newBoardFromMove(m))
		}
	}
	wg.Add(len(b.children))
	for _, ch := range b.children {
		go ch.goToDepth(depth-1, wg)
	}
}

func (b *Board) Depth(depth int) {
	var wg sync.WaitGroup
	if b.children == nil {
		moves := b.generateAllRawMoves()
		if len(moves) == 0 {
			return
		}
		for _, m := range moves {
			b.children = append(b.children, b.newBoardFromMove(m))
		}
	}
	wg.Add(len(b.children))
	for _, ch := range b.children {

		go ch.goToDepth(depth, &wg)
	}
	wg.Wait()
}

func compareMoves(m1, m2 Move) bool {
	return m1[0].InitialPosition == m2[0].InitialPosition && m1[0].FinalPosition == m2[0].FinalPosition
}

func GameAlphaBeta(playerPosition bool, depth int) {

	b := loadStartingPosition()
	for {
		if b.children == nil {
			moves := b.generateAllRawMoves()
			log.Printf("%+v\n", moves)
			if len(moves) == 0 {
				log.Println("Game over")
				return
			}
			for _, m := range moves {
				b.children = append(b.children, b.newBoardFromMove(m))
			}
		}
		if b.Turn == playerPosition {
			for {
				mv, err := getMove(playerPosition)
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				newBoard, ok := b.AdvanceBoard(mv)
				if !ok {
					fmt.Println("Could not advace board, invalid move")
					fmt.Println(b.Draw())
					continue
				}
				b = newBoard
				break
			}
		} else {
			fb, score := b.alphabeta(depth, -INF, INF)
			nb := fb.trickleUp(b)
			fmt.Println(nb.Draw())
			fmt.Printf("Move: %s\n", nb.PrevMove.String())
			fmt.Printf("Score: %d\n", score)
			newBoard, ok := b.AdvanceBoard(nb.PrevMove)
			if !ok {
				log.Printf("Could not advance board\n")
				return
			}
			b = newBoard
		}
	}
}

func GameMinimax(playerPosition bool, depth int) {

	b := loadStartingPosition()
	for {
		if b.children == nil {
			moves := b.generateAllRawMoves()
			log.Printf("%+v\n", moves)
			if len(moves) == 0 {
				log.Println("Game over")
				return
			}
			for _, m := range moves {
				b.children = append(b.children, b.newBoardFromMove(m))
			}
		}
		if b.Turn == playerPosition {
			log.Printf("Number of children: %d\n", len(b.children))
			for {
				mv, err := getMove(playerPosition)
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				newBoard, ok := b.AdvanceBoard(mv)
				if !ok {
					fmt.Println("Could not advace board, invalid move")
					b.printBoard()
					continue
				}
				b = newBoard
				break
			}
		} else {
			b.Depth(depth)
			nb, score := b.minimax()
			fmt.Printf("Move: %s\n", nb.PrevMove.String())
			fmt.Printf("Score is %d\n", score)
			newBoard, ok := b.AdvanceBoard(nb.PrevMove)
			if !ok {
				log.Printf("Could not advance board\n")
				return
			}
			b = newBoard
		}
	}
}

/*
func Game(playerPosition bool, depth int) {
	b := loadStartingPosition()
	for {
		b.printBoard()
		if b.children == nil {
			children, gr := b.GenerateChildren()
			if gr != nil {
				log.Printf("Game ended in a %s\n", gr.String())
				break
			}
			b.children = children
		}
		if b.Turn == playerPosition {
			for {
				mv, err := getMove()
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				newBoard, ok := b.AdvanceBoard(mv)
				if !ok {
					fmt.Println("Could not advace board, invalid move")
					b.printBoard()
					continue
				}
				b = newBoard
				break
			}
		} else {
			b.Depth(depth)
			nextBoard, score := b.alphabeta(depth, -10000000000.0, 10000000000.0)
			nb := nextBoard.trickleUp(b)
			fmt.Printf("Move: %s\n", nb.PrevMove.String())
			fmt.Println("Score is %f\n", score)
			newBoard, ok := b.AdvanceBoard(nb.PrevMove)
			if !ok {
				log.Printf("Could not advance board\n")
				return
			}
			b = newBoard
		}
	}
}
*/

func (b *Board) canSideCapturePos(turn bool, pos Position) bool {
	var x int8
	var y int8
	for x = 0; x < 8; x++ {
		for y = 0; y < 8; y++ {
			piece := b.bb[x][y]
			if piece == 0 {
				continue
			}
			if (piece > 0 && turn) || (piece < 0 && (!turn)) {
				if ok := b.canPieceCapture(piece, x, y, pos); ok {
					return true
				}
			}
		}
	}
	return false
}

func (b *Board) canPieceCapture(piece, x, y int8, pos Position) bool {
	dirs := pieceDirs(piece)
	initX := x
	initY := y
	for _, dir := range dirs {
		if isPawn(piece) {
			break
		}
		x = initX
		y = initY
		pr := pieceRange(piece, x, y)
		for i := 0; i < pr; i++ {
			x += dir.X
			y += dir.Y
			if !inBounds(x, y) {
				break
			}
			pieceVal := b.bb[x][y]
			if pieceVal == 0 {
				if y == 0 || y == 8 {
					// check if promotion
					if x == pos.X && y == pos.Y {
						return true
					}
					continue
				}
				if x == pos.X && y == pos.Y {
					return true
				}
				continue
			} else if pieceVal > 0 {
				if piece > 0 {
					break
				}
				if pieceCapDiff(piece) {
					break
				}
				if pos.X == x && pos.Y == y {
					return true
				}
				break
			} else {
				if piece < 0 {
					break
				}
				if pieceCapDiff(piece) {
					break
				}
				if pos.X == x && pos.Y == y {
					return true
				}
				break
			}
		}
	}
	if pieceCapDiff(piece) {
		capDirs := pieceCapDirs(piece)
		for _, dir := range capDirs {
			x = initX
			y = initY
			x += dir.X
			y += dir.Y
			if !inBounds(x, y) {
				continue
			}
			if pos.X == x && pos.Y == y {
				return true
			}
		}
	}
	return false
}

func (b *Board) generateAllRawMoves() []Move {
	moves := []Move{}
	var x int8
	var y int8
	for x = 0; x < 8; x++ {
		for y = 0; y < 8; y++ {
			piece := b.bb[x][y]
			if piece == 0 {
				continue
			}
			if (piece > 0 && b.Turn) || (piece < 0 && (!b.Turn)) {
				firstMoves := b.generateRawMoves(piece, x, y, turnQueen(b.Turn))
				for _, fm := range firstMoves {
					ok := b.generateAllRawMovesAfterMove(fm)
					if !ok {
						moves = append(moves, fm)
					}
				}
			}
		}
	}
	return moves
}

func turnQueen(turn bool) int8 {
	if turn {
		return 5
	}
	return -5
}

func (b *Board) generateAllRawMovesAfterMove(mv Move) bool {
	var i int8
	var j int8
	for i = 0; i < 8; i++ {
		for j = 0; j < 8; j++ {
			sp := b.bb[i][j]
			if sp == 0 {
				continue
			}
			if (sp > 0 && !b.Turn) || (sp < 0 && b.Turn) {
				ok := b.generateRawMovesAfterMove(sp, i, j, mv)
				if ok {
					return true
				}
			}
		}
	}
	return false
}

func copyPosition(p Position) Position {
	return Position{
		X: p.X,
		Y: p.Y,
	}
}

func (b *Board) generateRawMovesAfterMove(piece, x, y int8, mv Move) bool {
	if mv[0].FinalPosition.X == x && mv[0].FinalPosition.Y == y {
		return false
	}
	if mv[0].InitialPosition.X == x && mv[0].InitialPosition.Y == y {
		x = mv[0].FinalPosition.X
		y = mv[0].FinalPosition.Y
	}
	dirs := pieceDirs(piece)
	initX := x
	initY := y
	for _, dir := range dirs {
		x = initX
		y = initY
		pr := pieceRange(piece, x, y)
		for i := 0; i < pr; i++ {
			x += dir.X
			y += dir.Y
			if !inBounds(x, y) {
				break
			}

			pieceVal := b.getPieceValAfterMove(x, y, mv)
			if pieceVal == 0 {
				continue
			} else if pieceVal > 0 {
				if piece > 0 {
					break
				}
				if pieceVal == 6 {
					return true
				}
				break
			} else {
				if piece < 0 {
					break
				}
				if pieceVal == -6 {
					return true
				}
				break
			}
		}
	}
	if pieceCapDiff(piece) {
		capDirs := pieceCapDirs(piece)
		for _, dir := range capDirs {
			x = initX
			y = initY
			x += dir.X
			y += dir.X
			if !inBounds(x, y) {
				continue
			}
			pieceVal := b.getPieceValAfterMove(x, y, mv)
			if pieceVal == 0 {
				continue
			}
			if (piece > 0) != (pieceVal > 0) {
				if pieceVal == 6 || pieceVal == -6 {
					return true
				}
			}
		}
	}
	return false
}
func (b *Board) getPieceValAfterMove(x, y int8, mv Move) int8 {
	if mv[0].InitialPosition.X == x && mv[0].InitialPosition.Y == y {
		return 0
	} else if mv[0].FinalPosition.X == x && mv[0].FinalPosition.Y == y {
		return b.bb[mv[0].InitialPosition.X][mv[0].InitialPosition.Y]
	}
	return b.bb[x][y]
}

func (b *Board) generateRawMoves(piece, x, y, promotion int8) []Move {
	moves := []Move{}
	dirs := pieceDirs(piece)
	initX := x
	initY := y
	for _, dir := range dirs {
		x = initX
		y = initY
		pr := pieceRange(piece, x, y)
		for i := 0; i < pr; i++ {
			x += dir.X
			y += dir.Y
			if !inBounds(x, y) {
				break
			}
			pieceVal := b.bb[x][y]
			if pieceVal == 0 {
				if y == 0 || y == 8 {
					// check if promotion
					moves = append(moves, []Movement{{Position{initX, initY}, Position{x, y}, promotion}})
					continue
				}
				moves = append(moves, []Movement{{Position{initX, initY}, Position{x, y}, 0}})
				continue
			} else if pieceVal > 0 {
				if piece > 0 {
					break
				}
				if pieceCapDiff(piece) {
					break
				}
				moves = append(moves, []Movement{{Position{initX, initY}, Position{x, y}, 0}})
				break
			} else {
				if piece < 0 {
					break
				}
				if pieceCapDiff(piece) {
					break
				}
				moves = append(moves, []Movement{{Position{initX, initY}, Position{x, y}, 0}})
				break
			}
		}
	}
	if pieceCapDiff(piece) {
		capDirs := pieceCapDirs(piece)
		for _, dir := range capDirs {
			x = initX
			y = initY
			x += dir.X
			y += dir.Y
			if !inBounds(x, y) {
				continue
			}
			pieceVal := b.bb[x][y]
			if pieceVal == 0 {
				continue
			}
			if (piece > 0) != (pieceVal > 0) {
				moves = append(moves, []Movement{{Position{initX, initY}, Position{x, y}, 0}})
			}
		}
	}
	if isKing(piece) {
		castling_moves := b.castlingMoves(piece)
		for _, mv := range castling_moves {
			moves = append(moves, mv)
		}
	}
	return moves
}

func (b *Board) castlingMoves(piece int8) []Move {
	moves := []Move{}
	if piece > 0 {
		if b.canWhiteKingSideCastle() {
			moves = append(moves, WhiteKingSideCastleMove)
		}
		if b.canWhiteQueenSideCastle() {
			moves = append(moves, WhiteQueenSideCastleMove)
		}
		return moves
	}
	if b.canBlackKingSideCastle() {
		moves = append(moves, BlackKingSideCastleMove)
	}
	if b.canBlackQueenSideCastle() {
		moves = append(moves, BlackQueenSideCastleMove)
	}
	return moves
}

func (b *Board) Evaluate() int {
	score := 0
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			score += pieceScore(b.bb[x][y])
		}
	}
	return score
}

func (b *Board) EvaluateWithMove(mv Move) int {
	score := 0
	var x int8
	var y int8
	for x = 0; x < 8; x++ {
		for y = 0; y < 8; y++ {
			if x == mv[0].FinalPosition.X && y == mv[0].FinalPosition.Y {
				score += pieceScore(b.bb[x][y])
				continue
			}
			if x == mv[0].InitialPosition.X && y == mv[0].InitialPosition.Y {
				continue
			}
			score += pieceScore(b.bb[x][y])
		}
	}
	return score
}
