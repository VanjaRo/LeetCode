package main

func main() {

}

func isReflected(points [][]int) bool {
	if len(points) < 2 {
		return true
	}
	var min, max int
	var mid float32
	var m map[[2]int]bool
	m = make(map[[2]int]bool)
	min = points[0][0]
	max = min
	for i := 0; i < len(points); i++ {
		if points[i][0] < min {
			min = points[i][0]
		} else if points[i][0] > max {
			max = points[i][0]
		}
		m[[2]int{points[i][0], points[i][1]}] = true
	}
	mid = float32(min+max) / 2.0
	for el := 0; el < len(points); el++ {
		refX := int(2*mid) - points[el][0]
		if _, ok := m[[2]int{refX, points[el][1]}]; !ok {
			return false
		}
	}
	return true
}
