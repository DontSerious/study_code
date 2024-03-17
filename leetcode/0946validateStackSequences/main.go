package main

func main() {
	validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1})
}

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	index := 0

	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && stack[len(stack)-1] == popped[index] {
			stack = stack[:len(stack)-1]
			index++
		}
	}

	return len(stack) == 0
}
