package main

func main() {
	println(isValid("(){}[]"))
}

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []byte{}

	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			// 如果栈长度为0以及取栈顶不等则返回false
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			// 出栈
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
