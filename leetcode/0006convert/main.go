package main

import "fmt"

func main() {
	fmt.Printf("convert(\"PAYPALISHIRING\", 3): %v\n", convert("PAYPALISHIRING", 3))
}

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	arr := make([][]byte, numRows)
	flag := false
	j := 0
	for i := range s {
		if j == numRows-1 || j == 0 {
			flag = !flag
		}

		arr[j] = append(arr[j], s[i])
		if flag {
			j++
		}

		if !flag {
			j--
		}
	}

	// 变为字符串
	res := ""
	for _, v := range arr {
		res += string(v)
	}

	return res
}
