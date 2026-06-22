package main

import "fmt"

func main() {
	fmt.Println(rotateGrid([][]rune{{'a'},{'b'},{'c'}}))
}

func rotateGrid(grid [][]rune) [][]rune {

	if len(grid) == 0 {
		return [][]rune{}
	}
	row := len(grid)
	col := len(grid[0])
	// fmt.Println(col)
	new := make([][]rune, col)
	// fmt.Println(new)

	for i := range new {
		new[i] = make([]rune, row)
		// fmt.Println(new[i])
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			new[j][row-i-1] = grid[i][j]
		}
	}
	return new
}
