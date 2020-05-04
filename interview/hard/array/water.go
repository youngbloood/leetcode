package array

/*
# Container With Most Water
# https://leetcode.com/explore/interview/card/top-interview-questions-hard/116/array-and-strings/830/

Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.





The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.



Example:

Input: [1,8,6,2,5,4,8,3,7]
Output: 49



Hint #2
Start with the maximum width container and go to a shorter width container if there is a vertical line longer than the current containers shorter line. This way we are compromising on the width but we are looking forward to a longer length container.
*/

func MaxArea(height []int) int {
	return maxArea(height)
}
func maxArea(height []int) int {
	var area int
	allLenght := len(height) - 1
	for lenght := allLenght; lenght >= 1; lenght-- {
		for i := 0; i <= allLenght-lenght; i++ {
			high := height[i]
			if height[i+lenght] < high {
				high = height[i+lenght]
			}
			if val := high * lenght; val > area {
				area = val
			}
		}
	}
	return area
}

// Leetcode最优解
func maxArea2(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0
	for left < right {
		leftHeight := height[left]
		rightHeight := height[right]
		width := right - left
		if leftHeight < rightHeight {
			if leftHeight*width > maxArea {
				maxArea = leftHeight * width
			}

			left++
		} else {
			if rightHeight*width > maxArea {
				maxArea = rightHeight * width
			}

			right--
		}
	}
	return maxArea
}
