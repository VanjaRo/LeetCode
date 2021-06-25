package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	sortedHash := make(map[string][]string)
	var ret [][]string
	for i := 0; i < len(strs); i++ {
		sortedWord := sortString(strs[i])
		sortedHash[sortedWord] = append(sortedHash[sortedWord], strs[i])
	}
	for el := range sortedHash {
		ret = append(ret, sortedHash[el])
	}
	return ret
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
