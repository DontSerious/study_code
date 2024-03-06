package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	minNum int
	length int
	stack  []int
}

func Constructor() MinStack {
	return MinStack{
		minNum: math.MaxInt,
		length: 0,
		stack:  []int{},
	}
}

func (this *MinStack) Push(val int) {
	if this.minNum > val {
		this.minNum = val
	}

	this.length++
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	this.length--
	if this.length == 0 {
		this.minNum = math.MaxInt
		this.stack = []int{}
		return
	}

	popNum := this.stack[this.length]
	this.stack = this.stack[:this.length]

	if this.minNum == popNum {
		tmp := this.stack[0]
		for _, x := range this.stack {
			if x < tmp {
				tmp = x
			}
		}
		this.minNum = tmp
	}
}

func (this *MinStack) Top() int {
	return this.stack[this.length-1]
}

func (this *MinStack) GetMin() int {
	return this.minNum
}

func main() {
	obj := Constructor()
	obj.Push(2147483646)
	obj.Push(2147483646)
	obj.Push(2147483647)
	fmt.Printf("obj.Top(): %v\n", obj.Top())
	obj.Pop()
	fmt.Printf("obj.GetMin(): %v\n", obj.GetMin())
	obj.Pop()
	fmt.Printf("obj.GetMin(): %v\n", obj.GetMin())
	obj.Pop()
	obj.Push(2147483647)
	fmt.Printf("obj.Top(): %v\n", obj.Top())
	fmt.Printf("obj.GetMin(): %v\n", obj.GetMin())
	obj.Push(-2147483648)
	fmt.Printf("obj.Top(): %v\n", obj.Top())
	fmt.Printf("obj.GetMin(): %v\n", obj.GetMin())
	obj.Pop()
	fmt.Printf("obj.GetMin(): %v\n", obj.GetMin())
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
