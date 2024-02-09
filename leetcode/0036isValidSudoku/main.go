package main

func main() {
	println(isValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}))
}

func isValidSudoku(board [][]byte) bool {
	var rows, col [9][9]int
	var subboxes [3][3][9]int

	for i, row := range board {
		for j, x := range row {
			if x == '.' {
				continue
			}

			index := x - '1'
			rows[i][index]++
			col[j][index]++
			subboxes[i/3][j/3][index]++
			if rows[i][index] > 1 || col[j][index] > 1 || subboxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}
