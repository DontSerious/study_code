package main

import "math/bits"

func main() {
	paths := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 3}, {2, 4}}
	gardenNoAdj(4, paths)
}

func gardenNoAdj(n int, paths [][]int) []int {
	// 画图
	// 根据paths生成无向图，第一个下标表示第几个花园，第二个下标表示花园的第几个邻居
	g := make([][]int, n)
	for _, p := range paths {
		x, y := p[0]-1, p[1]-1
		// 双向连通
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	// 位运算解题
	// color数组下标代表第几个花园，值为用了几号颜色
	color := make([]int, n)
	for i, nodes := range g {
		// mask记录相邻节点使用的颜色
		mask := uint8(1)
		// 遍历color查看邻居已使用颜色来决定自己使用的最小编号颜色
		for _, node := range nodes {
			// 已被使用的颜色越多，mask取或之后后面的1就越多
			/* 
				假如邻居的编号为 2和3 且color[1] = 2, color[2] = 3
				可以得出 mask = 1101 后面跟着一个1，恰好一号颜色未被使用
			*/
			mask |= 1 << color[node]
		}
		// 找出最小未使用颜色
		// bits.TrailingZeros8() 八位二进制尾部有几个0
		color[i] = bits.TrailingZeros8(^mask)
	}

	return color
}