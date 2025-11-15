package arrays

/*
You are given an integer array height of length n.
There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container that holds the most water.
Return the maximum amount of water it can store.

You may not slant the container.

Example:
Input:  height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The vertical lines form a container with max area 49.
*/

// ContainerWithMostWater returns the maximum area of water
// that can be contained between two vertical lines.
// Time Complexity: O(n)
// Space Complexity: O(1)
func ContainerWithMostWater(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		lh, rh := height[left], height[right]
		width := right - left
		area := min(lh, rh) * width

		if lh < rh {
			left++
		} else {
			right--
		}

		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
