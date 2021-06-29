package main

func main() {

}

func longestOnes(nums []int, k int) int {
	windowStart := 0
	zeroCounter := 0
	windowEnd := 0
	// maxLen := 0
	for windowEnd < len(nums) {
		if nums[windowEnd] == 0 {
			zeroCounter++
		}
		if zeroCounter > k {
			if nums[windowStart] == 0 {
				zeroCounter--
			}
			windowStart++
		}
		windowEnd++
	}
	return windowEnd - windowStart
}
