package main

func main() {

}

func maxArea(height []int) int {
	var i, j, mx int
	j = len(height) - 1
	for i <= j {
		minDimY := min(height[i], height[j])
		space := (j - i) * minDimY
		if space > mx {
			mx = space
		}
		if height[j] > height[i] {
			i++
		} else {
			j--
		}
	}
	return mx
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
