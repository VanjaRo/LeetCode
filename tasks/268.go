package main

func main() {

}

func missingNumber(nums []int) int {
	n := len(nums)
	sum := int(float32(1+n) / float32(2) * float32(n))
	for _, el := range nums {
		sum -= el
	}
	return sum
}
