package main

/*
	定义栈的数据结构，
	请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，
	调用 min、push 及 pop 的时间复杂度都是 O(1)。
*/

type MinStack struct {
	data []int
	minS []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (stack *MinStack) Push(x int)  {
	stack.data = append(stack.data, x)
	minN := len(stack.minS)
	if minN == 0 {
		stack.minS = append(stack.minS, x)
	} else {
		if x < stack.minS[minN - 1] {
			stack.minS = append(stack.minS, x)
		} else {
			stack.minS = append(stack.minS, stack.minS[minN - 1])
		}
	} 
}

func (stack *MinStack) Pop()  {
	stack.data = stack.data[:len(stack.data) - 1]
	stack.minS = stack.minS[:len(stack.minS) - 1]
}

func (stack *MinStack) Top() int {
	return stack.data[len(stack.data) - 1]
}

func (stack *MinStack) Min() int {
	return stack.minS[len(stack.minS) - 1]
}

func main()  {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	stack.Min()
	stack.Pop()
	stack.Top()
	stack.Min()
}