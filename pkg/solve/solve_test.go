// pkg/solve/solves_test.go

package solve

import (
	"reflect"
	"testing"

	"github.com/DemFrogs1/sudoku-solver/pkg/board"
)

func TestChecks(t *testing.T) {
	board := board.GenerateBoard()
	board.SeedBoard()

	possibleValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	rowIndex := 2

	expected := []int{1, 2, 4}

	checkRow(board.Board, rowIndex, &possibleValues)
	checkCol(board.Board[0], &possibleValues)

	subgrid := board.GetSubGrid(0, 0)
	checkSubGrid(subgrid, &possibleValues)

	if !reflect.DeepEqual(possibleValues, expected) {
		t.Errorf("Unexpected result. Expected: %v, Got: %v", expected, possibleValues)
	}
}
