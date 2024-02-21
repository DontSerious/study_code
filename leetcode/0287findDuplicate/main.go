package main

func main() {
	println(findDuplicate([]int{1, 3, 4, 2, 2}))
}

// 在第一个循环结束时，fast和slow都被囚禁在一个循环里，finder和slow一样一次只走一步，会经历一样的路径，当他们相等时就是进入这个循环囚笼的时刻。
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}

	finder := 0
	for {
		finder = nums[finder]
		slow = nums[slow]

		if finder == slow {
			break
		}
	}

	return finder
}
