package solve

import (
	"github.com/DemFrogs1/sudoku-solver/pkg/board"
)

func Solve(b board.Board) bool {
	if !checkEmptySpace(b.Board) {
		return true
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b.Board[i][j] == 0 {
				subgrid := b.GetSubGrid(i, j)
				possibleValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

				checkCol(b.Board[i], &possibleValues)
				checkRow(b.Board, j, &possibleValues)
				checkSubGrid(subgrid, &possibleValues)

				if len(possibleValues) > 0 {
					for _, val := range possibleValues {
						b.Board[i][j] = val

						if Solve(b) {
							return true
						} else {
							b.Board[i][j] = 0
						}

					}
				}
				return false
			}
		}
	}

	return false
}

func checkCol(col []int, possibleValues *[]int) {
	for _, n := range col {
		if n != 0 {
			deleteElement(possibleValues, n)
		}
	}
}

func checkRow(b [][]int, rowIndex int, possibleValues *[]int) {
	for i := 0; i < 9; i++ {
		if b[i][rowIndex] != 0 {
			deleteElement(possibleValues, b[i][rowIndex])
		}
	}
}

func checkSubGrid(subgrid [][]int, possibleValues *[]int) {
	for _, row := range subgrid {
		for _, n := range row {
			if n != 0 {
				deleteElement(possibleValues, n)
			}
		}
	}
}

func checkEmptySpace(b [][]int) bool {
	for _, row := range b {
		for _, n := range row {
			if n == 0 {
				return true
			}
		}
	}

	return false
}

func deleteElement(slice *[]int, el int) {
	index := -1

	for i, n := range *slice {
		if n == el {
			index = i
			break
		}
	}

	if index == -1 {
		return
	}

	*slice = append((*slice)[:index], (*slice)[index+1:]...)
}
