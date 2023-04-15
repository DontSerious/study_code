func gardenNoAdj(n int, paths [][]int) []int {
	// 画图
	g := make([][]int, n)
	for _, p := range paths {
		x, y := p[0]-1, p[1]-1
		// 双向连通
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	color := make([]int, n)
	for i, nodes := range g {
		// mask记录相邻节点使用的颜色
		mask := uint8(1)
		for _, node := range nodes {
			mask |= 1 << color[node]
		}
		// 找出最小未使用颜色
		// bits.TrailingZeros8() 八位二进制尾部有几个0
		color[i] = bits.TrailingZeros8(^mask)
	}

	return color
}