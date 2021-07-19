package main

func main() {

}

func search(nums []int, target int) int {
	if nums[0] <= target {
		for i := 0; i <= len(nums)-1 && nums[i] <= target; i++ {
			if nums[i] == target {
				return i
			}
		}
		return -1
	}

	for i := len(nums) - 1; i >= 0 && nums[i] >= target; i-- {
		if nums[i] == target {
			return i
		}
	}
	return -1
}
