package recursion

/*
# https://leetcode.com/explore/learn/card/recursion-ii/507/beyond-recursion/2901/
Given n non-negative integers representing the histogram's bar height where the width of each bar is 1, find the area of largest rectangle in the histogram.

Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].

The largest rectangle is shown in the shaded area, which has area = 10 unit.

Example:
Input: [2,1,5,6,2,3]
Output: 10
*/
func LargestRectangleArea(heights []int) int {
	return largestRectangleArea(heights)
}

func largestRectangleArea(heights []int) int {
	var max int
	for i := range heights {
		rectangle := heights[i] * getWidth(heights, i)
		if rectangle > max {
			max = rectangle
		}
	}
	return max
}

func getWidth(heights []int, i int) (width int) {
	left := i
	right := i

	for ; left >= 0; left-- {
		if heights[left] < heights[i] {
			break
		}
	}

	for ; right < len(heights); right++ {
		if heights[right] < heights[i] {
			break
		}
	}
	return right - left - 1
}