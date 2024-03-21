package main

import (
	"fmt"
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("arr[:2]: %v\n", arr[:2])
	fmt.Printf("arr[2:]: %v\n", arr[2:])
}

func buildTree(pre []int, in []int) *TreeNode {
	n := len(pre)
	if n == 0 {
		return nil
	}

	leftSize := slices.Index(in, pre[0])
	left := buildTree(pre[1:leftSize+1], in[:leftSize])
	right := buildTree(pre[leftSize+1:], in[1+leftSize:])
	return &TreeNode{pre[0], left, right}
}
