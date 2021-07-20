package main

import (
	"fmt"
)

func main() {
	str := "/a//b////c/d//././/.."
	fmt.Print(simplifyPath(str))
}

func simplifyPath(path string) string {
	start := 0
	var ret []string
	for start < len(path) {
		for ; start < len(path) && path[start] == '/'; start++ {
		}
		end := start
		for ; end < len(path) && path[end] != '/'; end++ {
		}
		s := path[start:end]
		if s == ".." {
			if len(ret) > 0 {
				ret = ret[:len(ret)-1]
			}
		} else {
			if s != "." && s != "" {
				ret = append(ret, s)
			}
		}
		start = end + 1
	}

	if len(ret) == 0 {
		return "/"
	}
	answ := ""
	for _, el := range ret {
		answ += "/" + el
	}
	return answ
}

// func putInRet(path, str &string, start, end int) {
// 	path +
// }

// /home//.
