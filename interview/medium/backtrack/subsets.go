package backtrack

import "sort"

/*
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/109/backtracking/796/

Solution
Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/
// TODO: 能完成组合，但是不满足leetcode的要求
func Subsets(nums []int) [][]int {
	return subsets(nums)
}

func subsets(nums []int) [][]int {
	var result [][]int
	result = append(result, []int{})
	for i := 0; i < len(nums); i++ {
		for n := 0; n <= len(nums[:i]); n++ {
			generateSetsWithLeft(nums[:i], n, nums[i], nil, &result)
		}
	}
	return result
}

func generateSetsWithLeft(nums []int, n int, val int, subset []int, result *[][]int) {
	if n == 0 {
		subset = append(subset, val)
		sort.Ints(subset)
		*result = append(*result, subset)
	}
	for i, v := range nums {
		generateSetsWithLeft(nums[i+1:], n-1, val, append(subset, v), result)
	}
}

func generateSets(nums []int, n int, subset []int, result *[][]int) {
	if n == 0 {
		sort.Ints(subset)
		*result = append(*result, subset)
	}
	for i, v := range nums {
		generateSets(nums[i+1:], n-1, append(subset, v), result)
	}
}
