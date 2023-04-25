package main

// 动态规划 + 剪枝
// func maxValue(grid [][]int) int {
// 	// 剪枝 记录已经计算过的数据 防止重复计算 
// 	m, n := len(grid), len(grid[0])
// 	memo := make([][]int, m)
// 	for i := range memo {
// 		memo[i] = make([]int, n)
// 	}
	
// 	var dfs func(i, j int) int
// 	dfs = func(i, j int) int {
// 		if i < 0 || j < 0 {
// 			return 0
// 		}
// 		p := &memo[i][j]
// 		if *p != 0 {
// 			return *p
// 		}
// 		*p = max(dfs(i - 1, j), dfs(i, j - 1)) + grid[i][j]
// 		return *p
// 	}
// 	return dfs(len(grid) - 1, len(grid[0]) - 1)
// }

// 递推 将递归走到底层的步骤去掉 变成循环
func maxValue(grid [][]int) int {
	// 解决这种定义方式没有状态能表示递归边界，即出界的情况。
	// 修改后的 f[i+1][j+1]表示从左上角到(i, j)的最大价值和 此时f[i][0]和f[0][j]就可以表示出界的状态
	m, n := len(grid), len(grid[0])
	f := make([][]int, m + 1)
	for i := range f {
		f[i] = make([]int, n + 1)
	}

	for i, row := range grid {
		for j, x := range row {
			f[i + 1][j + 1] = max(f[i + 1][j], f[i][j + 1])	+ x
		}
	}
	return f[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}