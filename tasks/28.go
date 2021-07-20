package main

func main() {

}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for ind := range haystack {
		if len(haystack) < ind+len(needle) {
			break
		}
		if haystack[ind:ind+len(needle)] == needle {
			return ind
		}
	}
	return -1
}
