package dynamic

/*
# Jump Game
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/111/dynamic-programming/807/

Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.

Example 1:
Input: [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:
Input: [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum
             jump length is 0, which makes it impossible to reach the last index.
*/
func CanJump(nums []int) bool {

}
func canJump(nums []int) bool {
	var stepStack stack
	var i int
	for i < len(nums) {
		step := nums[i]
		if step == 0 {
			break
		}
		stepStack.Push(i)
		i += step

	}
}

// stack存放下标索引
type stack []int

func (s stack) Push(val int) {
	s = append(s, val)
}

func (s stack) Pop() int {
	val := s[len(s)-1]
	s = s[:len(s)-1]
	return val
}
