package main

import "strings"

/*
	翻转单词顺序

	输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。
	为简单起见，标点符号和普通字母一样处理。
	例如输入字符串"I am a student. "，则输出"student. a am I"。
*/

func reverseWords(s string) string {
	s = strings.TrimSpace(s)

	arr := strings.Fields(s)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return strings.Join(arr, " ")
}