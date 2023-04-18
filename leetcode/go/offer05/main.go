package main

/* 
	替换空格
	请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
*/

func replaceSpace(s string) string {
	b := []byte(s)
    result := make([]byte, 0)
    for i := 0; i < len(b); i++ {
        if b[i] == ' ' {
            result = append(result, []byte("%20")...)
        } else {
            result = append(result, b[i])
        }
    }
    return string(result)
}