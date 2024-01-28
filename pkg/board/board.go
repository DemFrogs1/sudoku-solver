package board

const (
	MIN_RANGE = 17
	MAX_RANGE = 35
)

type Board struct {
	Board [][]int
}

func GenerateBoard() *Board {
	board := make([][]int, 9)
	for i := range board {
		board[i] = make([]int, 9)
	}

	return &Board{board}
}

func (b *Board) SeedBoard() {
	b.Board = [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	//TODO: Randomly seed the sudoku board
	// amountToGenerate := rand.Intn(MAX_RANGE-MIN_RANGE) + MIN_RANGE

	// for i := 0; i < amountToGenerate; i++ {
	// 	// randNum := rand.Intn(9)
	// }
}

func (b *Board) GetSubGrid(row, col int) [][]int {
	subgrid := make([][]int, 3)
	for i := range subgrid {
		subgrid[i] = make([]int, 3)
	}

	startRow, startCol := row-row%3, col-col%3

	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			subgrid[i][j] = b.Board[i+startRow][j+startCol]
		}
	}

	return subgrid
}
