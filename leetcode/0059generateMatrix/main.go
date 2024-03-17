package main

import "fmt"

func main() {
	fmt.Printf("generateMatrix(3): %v\n", generateMatrix(3))
}

func generateMatrix(n int) [][]int {
	a, b, c, d := 0, 0, n-1, n-1
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	num, end := 1, n*n

	for num <= end {
		for i := a; i <= c; i++ {
			res[b][i] = num
			num++
		}
		b++
		for i := b; i <= d; i++ {
			res[i][c] = num
			num++
		}
		c--
		for i := c; i >= a; i-- {
			res[d][i] = num
			num++
		}
		d--
		for i := d; i >= b; i-- {
			res[i][a] = num
			num++
		}
		a++
	}

	return res
}
