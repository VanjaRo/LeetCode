package main

import "fmt"

func main() {
	nums := []int{0, 0, 1}
	moveZeroes(nums)
	fmt.Print(nums)
}

func moveZeroes(nums []int) {
	count := 0
	i := 0
	for i < len(nums) {

		if nums[i] == 0 && i < len(nums)-1 {
			fmt.Println(nums)
			fmt.Println(i)
			nums = append(nums[:i], nums[i+1:]...)
			count++
			continue
		}
		i++

	}
	zeros := make([]int, count)
	nums = append(nums, zeros...)
}
