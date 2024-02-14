package main

func main() {
	grid := [][]int{{1, 4, 5}, {2, 7, 6}, {6, 8, 7}}
	println(minPathSum(grid))
}

func minPathSum(grid [][]int) int {
	for i := range grid {
		for j := range grid[0] {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}

func min(x int, y int) int {
	if x <= y {
		return x
	}
	return y
}
