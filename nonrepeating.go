package main

import "fmt"

/**
 * 寻找最长不含重复字符的子串
 * 对于每个字母x  
 * LastOccurred[x] 不存在 或者小于 start 时 无需操作
 * LastOccurred[x] >= start 更新 start
 * 更新 LastOccurred[x] 更新 maxLength
 */
func lengthOfNonRepeatingString(s string) int {
	LastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {

		if Last1, ok := LastOccurred[ch]; ok && Last1 >= start {
			start = Last1 + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		LastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingString("abcabacb"))
	fmt.Println(lengthOfNonRepeatingString("bbbb"))
	fmt.Println(lengthOfNonRepeatingString("b"))
	fmt.Println(lengthOfNonRepeatingString(""))
	fmt.Println(lengthOfNonRepeatingString("abcderf"))
}
