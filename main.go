package main

import (
	"fmt"

	"github.com/DemFrogs1/sudoku-solver/pkg/board"
	"github.com/DemFrogs1/sudoku-solver/pkg/solve"
)

func main() {
	initialBoard := board.GenerateBoard()
	initialBoard.SeedBoard()

	fmt.Println("Initial Board:")
	printFormattedBoard(*initialBoard)

	solve.Solve(*initialBoard)

	fmt.Println("\nSolved Board:")
	printFormattedBoard(*initialBoard)
}

func printFormattedBoard(b board.Board) {
	for i, row := range b.Board {
		if i%3 == 0 && i != 0 {
			fmt.Println("------+-------+------")
		}
		for j, val := range row {
			if j%3 == 0 && j != 0 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
	fmt.Println()
}
