package main

import "fmt"

func main() {
	// 40 '(' 41 ')'
	// 91 '[' 93 ']'
	// 123 '{' 125 '}'
	fmt.Print(isValid("([)]"))
}

func isValid(s string) bool {
	// openedRound := 0
	// openedSquare := 0
	// openedFigure := 0
	var opened []byte
	if s[0] != ')' && s[0] != ']' && s[0] != '}' {
		opened = append(opened, s[0])
	} else {
		return false
	}

	for i := 1; i < len(s); i++ {
		switch s[i] {
		case ')':
			if len(opened) > 0 && abs(')'-opened[len(opened)-1]) < 3 {
				opened = opened[:len(opened)-1]
			} else {
				return false
			}

		case ']':
			if len(opened) > 0 && abs(']'-opened[len(opened)-1]) < 3 {
				opened = opened[:len(opened)-1]
			} else {
				return false
			}

		case '}':
			if len(opened) > 0 && abs('}'-opened[len(opened)-1]) < 3 {
				opened = opened[:len(opened)-1]
			} else {
				return false
			}
		default:
			opened = append(opened, s[i])
		}

	}
	if len(opened) > 0 {
		return false
	}
	return true

}

func abs(a byte) byte {
	if a > 0 {
		return a
	}
	return -a
}
