package main

func main() {

}

func productExceptSelf(nums []int) []int {
	prefix := 1
	suffix := 1
	n := len(nums)
	arr := make([]int, n)
	for i := range arr {
		arr[i] = 1
	}
	for i := 1; i <= len(nums); i++ {
		arr[i-1] *= prefix
		arr[n-i] *= suffix
		prefix *= nums[i-1]
		suffix *= nums[n-i]
	}
	return arr
}
