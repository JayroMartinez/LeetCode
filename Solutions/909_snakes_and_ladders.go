package main

import (
	"fmt"
)

func snakesAndLadders(board [][]int) int {
	n := len(board)
	// Create a 1-indexed 1D array to represent the board.
	cells := make([]int, n*n+1)
	idx := 1
	leftToRight := true
	// Flatten the board: start from the bottom row.
	for row := n - 1; row >= 0; row-- {
		if leftToRight {
			for col := 0; col < n; col++ {
				cells[idx] = board[row][col]
				idx++
			}
		} else {
			for col := n - 1; col >= 0; col-- {
				cells[idx] = board[row][col]
				idx++
			}
		}
		leftToRight = !leftToRight
	}

	// Define a type for BFS queue elements.
	type state struct {
		square int // current square (1-indexed)
		moves  int // number of dice rolls so far
	}

	// BFS starting from square 1.
	queue := []state{{square: 1, moves: 0}}
	visited := make([]bool, n*n+1)
	visited[1] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		// Check if we've reached the last square.
		if curr.square == n*n {
			return curr.moves
		}
		// Try all possible dice rolls (1 through 6).
		for dice := 1; dice <= 6; dice++ {
			nextSquare := curr.square + dice
			if nextSquare > n*n {
				break
			}
			// If there's a snake or ladder, jump to the destination.
			if cells[nextSquare] != -1 {
				nextSquare = cells[nextSquare]
			}
			// If this square hasn't been visited yet, add it to the queue.
			if !visited[nextSquare] {
				visited[nextSquare] = true
				queue = append(queue, state{square: nextSquare, moves: curr.moves + 1})
			}
		}
	}

	return -1
}

func main() {
	board := [][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1}}
	result := snakesAndLadders(board)
	fmt.Println(result)
}
