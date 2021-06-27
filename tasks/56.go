package main

func main() {

}

func merge(intervals [][]int) [][]int {
	mpStart := make(map[int]int)
	mpEnd := make(map[int]int)
	max := 0
	for _, interval := range intervals {
		mpStart[interval[0]]++
		mpEnd[interval[1]]++
		max = maxInt(max, interval[1])
	}
	var ret [][]int
	for i := 0; i <= max; i++ {
		if mpStart[i] > 0 {
			sum := 0
			el := make([]int, 2)
			el[0] = i
			for {
				sum += mpStart[i]
				sum -= mpEnd[i]
				if sum == 0 {
					break
				}
				i++
			}
			el[1] = i
			ret = append(ret, el)
		}
	}
	return ret
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
