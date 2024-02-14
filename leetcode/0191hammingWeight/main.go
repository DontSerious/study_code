package main

func main() {
	println(hammingWeight(uint32(11)))
}

func hammingWeight(num uint32) int {
	var res int

	for num != 0 {
		res += int(num & 1)
		num >>= 1
	}

	return res
}

// func hammingWeight(num uint32) int {
// 	res := 0

// 	for num != 0 {
// 		res++
// 		num &= num - 1
// 	}

// 	return res
// }
