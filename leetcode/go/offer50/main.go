package main

import "fmt"

/*
	第一个只出现一次的字符
	在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
*/

func firstUniqChar(s string) byte {
	dic := [26]int{}

	// 创建对应26个字母的哈希表之后遍历字符串，将字符出现的次数保存进去
	for _, v := range s {
		dic[v - 'a']++
	}

	// 再次遍历字符串，从头开始先遍历到哪个在dic中的值只有1的，那他就是答案
	for i, v := range s {
		if dic[v - 'a'] == 1 {
			return s[i]
		}
	}

	return ' '
}

func main() {
	s := "fabaccdez"
	fmt.Printf("firstUniqChar(s): %c\n", firstUniqChar(s))
}