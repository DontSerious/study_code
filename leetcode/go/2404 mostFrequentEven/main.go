package main

func main() {
	print(mostFrequentEven([]int{8154,9139,8194,3346,5450,9190,133,8239,4606,8671,8412,6290}))
}

// 出现最频繁的偶数元素
func mostFrequentEven(nums []int) int {
    hashmap := make(map[int]int)
	for _, num := range nums {
		if num % 2 == 0 {
			hashmap[num]++
		}
	}

	// res保存结果，time保存出现次数
	res, time := -1, 0
	for k, v := range hashmap {
		/* 	但凡出现一次都会修改res和time
			出现次数一旦大于当前保存次数就修改res和time，
			遍历到的v次数和time一致时根据k大小是否小于res进行修改
		*/
		if res == -1 || time < v || time == v && k < res {
			res = k
			time = v
		}
	}

	return res
}