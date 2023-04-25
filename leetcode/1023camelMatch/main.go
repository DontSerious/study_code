package main

func main() {
	camelMatch(nil, "")
}

func camelMatch(queries []string, pattern string) (res []bool) {
	check := func(a, b string) bool {
		m, n := len(a), len(b)
		i, j := 0, 0
		/* 
			i, j双指针
			要求忽视query所有多余的小写字母
			保证query驼峰等于Pattern
		*/
		for ; j < n ; i, j = i + 1, j + 1 {
			/* 忽视遇到下一个大写字母之前多余的小写字母 */
			for i < m && a[i] != b[j] && (a[i] >= 'a' && a[i] <= 'z') {
				i++;
			}
			// 进行比较，退出条件
			if i == m || a[i] != b[j] {
				return false
			}
		}
		/* 
			到达pattern末尾时，判断query剩余的字母是否都为小写
		 	防止出现后续驼峰情况
		*/
		for i < m && a[i] >= 'a' && a[i] <= 'z' {
			i++
		}
		// 遍历query到最后，没遇到新的驼峰则返回true
		return i == m
	}
	
	// 调用check方法往res数组添加数据
	for _, query := range queries {
		res = append(res, check(query, pattern))
	}

	return
}