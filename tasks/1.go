package main

func main() {

}

func twoSum(nums []int, target int) []int {
	// [rem]ind
	remMap := make(map[int]int)
	for i, v := range nums {
		if val, ok := remMap[v]; ok {
			return []int{val, i}
		}
		rem := target - v
		remMap[rem] = i
	}
	return []int{-1, -1}
}
