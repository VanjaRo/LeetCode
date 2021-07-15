package main

import "fmt"

func main() {
	arr := []int{4, 5, 0, -2, -3, 1}
	fmt.Println(subarraysDivByK(arr, 5))
}

func subarraysDivByK(nums []int, k int) int {
	ret, rem, sum := 0, 0, 0
	reminders := map[int]int{0: 1}
	for _, x := range nums {
		sum += x

		rem = sum % k
		if _, ok := reminders[rem]; ok {
			ret += reminders[rem]
		}
		reminders[rem]++
	}
	return ret
}
