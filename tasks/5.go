package main

func main() {

}

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}
	start, end := 0, 0
	for i := range s {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		lenMax := 0
		if len1 > len2 {
			lenMax = len1
		} else {
			lenMax = len2
		}

		if lenMax > end-start {
			start = i - (lenMax-1)/2
			end = i + (lenMax)/2
		}

	}
	return s[start : end+1]
}

func expandAroundCenter(s string, L, R int) int {
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}
