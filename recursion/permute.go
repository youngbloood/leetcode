package recursion

/*
# https://leetcode.com/explore/learn/card/recursion-ii/507/beyond-recursion/2903/
Given a collection of distinct integers, return all possible permutations.

Example:
Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/
func Permute(nums []int) [][]int {
	return permute(nums)
}

func permute(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	return finishPermute(nil, nil, nums)
}

// 基本逻辑和回溯有点相似，但是在递归中left会被修改值，暂时不是到什么原因
func finishPermute(result [][]int, dst, left []int) [][]int {
	if left == nil || len(left) == 0 {
		dstTemp := make([]int, len(dst))
		copy(dstTemp, dst)
		return append(result, dstTemp)
	}
	leftTemp := make([]int, len(left))
	copy(leftTemp, left)
	for i, v := range leftTemp {
		dst = append(dst, v)
		result = finishPermute(result, dst, append(left[:i], left[i+1:]...))
		dst = dst[:len(dst)-1]
		copy(left, leftTemp)
	}
	return result
}
