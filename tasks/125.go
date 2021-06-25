package main

import (
	"fmt"
)

func main() {
	fmt.Print(isPalindrome("A man, a plan, a canal: Panama"))
}

func isPalindrome(s string) bool {
	for left, right := 0, len(s)-1; left <= right; left, right = left+1, right-1 {
		vLeft := value(s[left])
		vRight := value(s[right])
		if vLeft < 0 {
			right++
			continue
		}
		if vRight < 0 {
			left--
			continue
		}
		if vLeft != vRight {
			return false
		}

	}
	return true
}

func value(m byte) int {
	if m >= '0' && m <= '9' {
		return int((m - '0') + 100)
	}
	if m >= 'A' && m <= 'Z' {
		return int(m - 'A')
	}
	if m >= 'a' && m <= 'z' {
		return int(m - 'a')
	}
	return -1
}

// Same func but with usage of regexp and strings modules

// func isPalindrome(s string) bool {
// 	reg, _ := regexp.Compile("[^a-zA-Z]+")
// 	formattedStr := reg.ReplaceAllString(s, "")
// 	formattedStr = strings.ToLower(strings.Join(strings.Split(formattedStr, " "), ""))
// 	var revesedStr string
// 	for _, c := range formattedStr {
// 		revesedStr = string(c) + revesedStr
// 	}
// 	fmt.Println(formattedStr)
// 	fmt.Println(revesedStr)
// 	return revesedStr == formattedStr
// }
