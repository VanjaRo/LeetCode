package main

import (
	"fmt"
)

func main() {
	str1 := "adc"
	str2 := "sadadasdasdsadasdasdasdadacdadasadadadasd"
	fmt.Print(checkInclusion(str1, str2))
}

func checkInclusion(s1 string, s2 string) bool {
	var mp1, mp2 [26]byte
	s2Len := len(s2)
	s1Len := len(s1)
	if s2Len < s1Len {
		return false
	}

	lPtr := 0
	rPtr := 0
	for i := range s1 {
		mp1[s1[i]-'a']++
		mp2[s2[i]-'a']++
		rPtr++
	}
	for rPtr != s2Len {
		if equal(mp1, mp2) {
			return true
		}
		mp2[s2[lPtr]-'a']--
		lPtr++
		mp2[s2[rPtr]-'a']++
		rPtr++
	}
	return equal(mp1, mp2)
}

func equal(mp1, mp2 [26]byte) bool {
	for i := range mp1 {
		if mp1[i] != mp2[i] {
			return false
		}
	}
	return true
}

// func checkInclusion(s1 string, s2 string) bool {
// 	s2Len := len(s2)
// 	s1Len := len(s1)
// 	if s2Len < s1Len {
// 		return false
// 	}
// 	s1Sorted := sortString(s1)
// 	for i := 0; i+s1Len-1 < s2Len; i++ {
// 		s2SortedSlice := sortString(s2[i : i+s1Len])
// 		if s1Sorted == s2SortedSlice {
// 			return true
// 		}
// 	}
// 	return false
// }

// type sortRunes []rune

// func (sr sortRunes) Less(i, j int) bool {
// 	return sr[i] < sr[j]
// }

// func (sr sortRunes) Swap(i, j int) {
// 	sr[i], sr[j] = sr[j], sr[i]
// }

// func (sr sortRunes) Len() int {
// 	return len(sr)
// }

// func sortString(s string) string {
// 	rs := []rune(s)
// 	sort.Sort(sortRunes(rs))
// 	return string(rs)
// }
