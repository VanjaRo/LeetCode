package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(myAtoi(("   ")))
}

func myAtoi(s string) int {
	// s should begin with number or sign

	// 1. Trim spaces
	sFormatted := strings.TrimLeft(s, " ")

	sign := 1
	startingInd := -1
	endingInd := -1
	signed := false
	// 2. Search and read sign
	for i, c := range sFormatted {
		if startingInd == -1 {
			if (c == rune('+') || c == rune('-')) && signed {
				return 0
			}
			if c == rune('+') {
				signed = true
			} else if c == rune('-') {
				sign = -1
				signed = true
			} else if unicode.IsDigit(c) {
				startingInd = i
			}
		} else {
			if !unicode.IsDigit(c) {
				endingInd = i
				break
			}
		}
	}
	if startingInd == -1 {
		return 0
	}
	if !signed && startingInd > 0 {
		return 0
	}
	if endingInd == -1 && startingInd != -1 {
		endingInd = len(sFormatted)
	}

	retNum, _ := strconv.Atoi(sFormatted[startingInd:endingInd])
	if sign == -1 && retNum > 2<<31 {
		return 2 << 31 * -1
	} else if sign == 1 && retNum > 2<<31-1 {
		return 2<<31 - 1
	} else {
		return retNum * sign
	}
}

// FASTER REALISATION
// func myAtoi(s string) int {
//     for len(s) > 0 && s[0] == ' ' {
//         s = s[1:]
//     }
//     if len(s) == 0 {
//         return 0
//     }
//     neg := false
//     switch s[0] {
//     case '-':
//         neg = true
//         fallthrough
//     case '+':
//         s = s[1:]
//     }
//     var val int64
//     for len(s) > 0 {
//         d := s[0]
//         if d < '0' || d > '9' {
//             break
//         }
//         val *= 10
//         val += int64(d-'0')
//         if val >= 1<<31 {
//             break
//         }
//         s = s[1:]
//     }
//     const (
//         max = 1<<31-1
//         min = -1<<31
//     )
//     if neg {
//         if val = -val; val < min {
//             val = min
//         }
//     } else if val > max {
//         val = max
//     }
//     return int(val)
// }
