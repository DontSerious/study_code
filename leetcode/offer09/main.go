package main

/* 
	用两个栈实现一个队列。
	队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead 
	分别完成在队列尾部插入整数和在队列头部删除整数的功能。
	(若队列中没有元素，deleteHead 操作返回 -1 )
*/

type CQueue struct {
	stackHead []int
	stackTail []int
}

func Constructor() CQueue {
	var queue CQueue
	return queue
}

// 将队尾栈的数据放到队头栈
func (queue *CQueue) Trans() int {
	count := len(queue.stackTail)
	for i := count; i > 0; i-- {
		queue.stackHead = append(queue.stackHead, queue.stackTail[i - 1])
		count--
		queue.stackTail = queue.stackTail[:count]
	}
	return len(queue.stackHead)
}

func (queue *CQueue) AppendTail(value int)  {
	queue.stackTail = append(queue.stackTail, value)
}

func (queue *CQueue) DeleteHead() (value int) {
	count := len(queue.stackHead)
	if count == 0 {
		if len(queue.stackTail) == 0 {
			return -1
		}
		count = queue.Trans()
	}
	count--
	value = queue.stackHead[count]
	queue.stackHead = queue.stackHead[:count]
	return
}

func main() {
	q := Constructor()
	q.DeleteHead()
	q.AppendTail(12)
	q.DeleteHead()
	q.AppendTail(10)
	q.AppendTail(9)
	q.DeleteHead()
	q.DeleteHead()
	q.DeleteHead()
	q.AppendTail(20)
	q.DeleteHead()
	q.AppendTail(1)
	q.AppendTail(8)
	q.AppendTail(20)
	q.AppendTail(1)
	q.AppendTail(11)
	q.AppendTail(2)
	q.DeleteHead()
	q.DeleteHead()
	q.DeleteHead()
	q.DeleteHead()
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */