package main

import (
	"fmt"
)

func main() {
	fmt.Println(countIslands([][]rune{{'#', '#', '.', '.', '#'},
		{'#', '.', '.', '#', '#'},
		{'.', '.', '.', '.', '.'},
		{'#', '#', '.', '#', '.'}}))

	fmt.Println(countIslands([][]rune{{'#'}}))
	fmt.Println(countIslands([][]rune{}))
	fmt.Println(countIslands([][]rune{{'.', '.'}, {'.', '.'}}))
	fmt.Println(sum([]int{1, 2, 4}))
}

func countIslands(grid [][]rune) int {

	if len(grid) == 0 {
		return 0
	}

	row := len(grid)
	col := len(grid[0])

	var visited = make([][]bool, row)

	for i := range visited {
		visited[i] = make([]bool, col)
	}

	var dfs func(i, j int)

	dfs = func(i, j int) {
		// stop conditions
		if i < 0 || i >= row || j < 0 || j >= col {
			return
		}
		if grid[i][j] != '#' || visited[i][j] {
			return
		}
		//

		visited[i][j] = true

		// explore 4 directions
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	island := 0

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			current_cell := grid[i][j]
			if current_cell == '#' && !visited[i][j] {
				island++
				dfs(i, j)
			}

		}
	}

	return island
}

func sum(nums []int) int {
  keep := 0	
  for i := 0; i < len(nums); i++ {
    if len(nums) <= 2 || len(nums) <= 10000 {
      keep += nums[i]
    }
  }
  return keep
}