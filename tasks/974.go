package main

import "fmt"

func main() {
	arr := []int{-1, 2, 9}
	fmt.Println(subarraysDivByK(arr, 2))
}

func subarraysDivByK(nums []int, k int) int {
	ret, rem, sum := 0, 0, 0
	reminders := make([]uint16, k)
	reminders[0] = 1
	for _, x := range nums {
		sum += x

		rem = sum % k
		if rem < 0 {
			rem += k
		}
		ret += int(reminders[rem])
		reminders[rem]++
	}
	return ret
}
