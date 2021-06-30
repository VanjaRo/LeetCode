package main

func main() {

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	var tmp int
	var j int
	for i := range nums1 {
		if i > m-1 {
			if tmp < nums2[j] {
				nums1[i] = tmp
			}
			nums1[i] = nums2[j]
			j++
		}
		if nums1[i] > nums2[j] {
			tmp
		}

	}
}
