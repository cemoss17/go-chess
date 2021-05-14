package board

var (
	WhiteKingSideCastleMove = []Movement{
		{Position{4, 0}, Position{6, 0}, 0},
		{Position{7, 0}, Position{5, 0}, 0},
	}
	WhiteQueenSideCastleMove = []Movement{
		{Position{4, 0}, Position{2, 0}, 0},
		{Position{0, 0}, Position{3, 0}, 0},
	}
	BlackKingSideCastleMove = []Movement{
		{Position{4, 7}, Position{6, 7}, 0},
		{Position{7, 7}, Position{5, 7}, 0},
	}
	BlackQueenSideCastleMove = []Movement{
		{Position{4, 7}, Position{2, 7}, 0},
		{Position{0, 7}, Position{3, 7}, 0},
	}
)
