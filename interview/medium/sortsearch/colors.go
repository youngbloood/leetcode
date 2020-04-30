package sortsearch

/*
# Sort Colors
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/798/

Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note: You are not suppose to use the library's sort function for this problem.

Example:

Input: [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
Follow up:

A rather straight forward solution is a two-pass algorithm using counting sort.
First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.
Could you come up with a one-pass algorithm using only constant space?
*/
func SortColors(nums []int) {
	sortColors(nums)
}

func sortColors(nums []int) {
	red, blue := 0, len(nums)-1
	for i := 0; i <= blue; i++ {
		if nums[i] == 0 {
			nums[i], nums[red] = nums[red], nums[i]
			red++
		} else if nums[i] == 2 {
			nums[i], nums[blue] = nums[blue], nums[i]
			// i前移，进入下次循环重新检测i位置的值是否是1
			i--
			blue--
		}
	}
}

func sortColors2(nums []int) {

	var red, white, blue int
	for _, v := range nums {
		switch v {
		case 0:
			red++
		case 1:
			white++
		case 2:
			blue++
		}
	}
	for i := range nums {
		if red > 0 {
			nums[i] = 0
			red--
			continue
		}
		if white > 0 {
			nums[i] = 1
			white--
			continue
		}
		nums[i] = 2
	}
}
