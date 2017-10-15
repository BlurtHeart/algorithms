package reality

// Given n non-negative integers representing an elevation map
// where the width of each bar is 1, compute how much water it
// is able to trap after raining.
// Given [0,1,0,2,1,0,1,3,2,1,2,1], return 6.
func Trap(height []int) int {
	left, right := 0, len(height)-1 // left and right index of height
	ret := 0                        // total capacity
	maxLeft, maxRight := 0, 0       // max value of left side and right side
	for left <= right {
		if height[left] <= height[right] {
			if height[left] >= maxLeft {
				maxLeft = height[left]
			} else {
				ret += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				ret += maxRight - height[right]
			}
			right--
		}
	}
	return ret
}
