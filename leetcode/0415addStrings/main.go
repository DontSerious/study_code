package main

import "strconv"

func main() {
	println(addStrings("1", "9"))
}

func addStrings(num1 string, num2 string) string {
	res := ""
	i, j := len(num1)-1, len(num2)-1
	carry := 0

	for i >= 0 || j >= 0 {
		n1, n2 := 0, 0
		if i >= 0 {
			n1 = int(num1[i] - '0')
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
		}

		tmp := n1 + n2 + carry
		carry = tmp / 10
		res = strconv.Itoa(tmp%10) + res
		i--
		j--
	}

	if carry > 0 {
		res = "1" + res
	}

	return res
}
