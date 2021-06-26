package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
}

func subarraySum(nums []int, k int) int {
	sumFreq := map[int]int{0: 1}
	sum := 0
	count := 0
	for _, v := range nums {
		sum += v
		count += sumFreq[sum-k]
		sumFreq[sum]++
	}
	return count
}
