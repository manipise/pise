func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	maxs := len(height) - 1
	big := 0
	for i < maxs || j < 0 {

		indexdif := j - i
		minval := 0
		if height[i] > height[j] {
			minval = height[j]
			j--
		} else {
			minval = height[i]
			i++
		}
		area := indexdif * minval
		if big < area {
			big = area
		}
	}
	return big
}
