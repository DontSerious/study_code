package main

func main() {
	var nums = []int{1, 2, 3, 4}
	twoSum(nums, 4)
}

// 求两数之和
func twoSum(nums []int, target int) []int {
    hashTable := make(map[int]int)
	for i, num := range nums {
		if ti, ok := hashTable[target - num]; ok {
			return []int{ti, i}
		}
		// 保存已遍历过的数字和它的下表，后面如果需要用到就可以很快的查到它的下标
		hashTable[num] = i
	}
	return nil
}