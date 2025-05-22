package main

import (
	"fmt"
)

func gameOfLife(board [][]int) {

	rows := len(board)
	cols := len(board[0])

	extBoard := make([][]int, rows+2)

	for i := range extBoard {
		extBoard[i] = make([]int, cols+2)
	}

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			extBoard[x+1][y+1] = board[x][y]
		}
	}

	for x := 1; x <= rows; x++ {
		for y := 1; y <= cols; y++ {
			aliveNeigh := extBoard[x-1][y-1] + extBoard[x-1][y] + extBoard[x-1][y+1] + extBoard[x][y-1] + extBoard[x][y+1] + extBoard[x+1][y-1] + extBoard[x+1][y] + extBoard[x+1][y+1]

			if aliveNeigh == 3 {
				board[x-1][y-1] = 1
			} else if extBoard[x][y] == 1 && aliveNeigh == 2 {
				board[x-1][y-1] = 1
			} else {
				board[x-1][y-1] = 0
			}

		}
	}

	fmt.Println(board)

}

func main() {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	gameOfLife(board)
}
