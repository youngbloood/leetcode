package array

import(
	"sort"
)
/*
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/776/

Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

[-4,-1,-1,0,1,2]
A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]


Hint #1
So, we essentially need to find three numbers x, y, and z such that they add up to the given value. If we fix one of the numbers say x, we are left with the two-sum problem at hand!

Hint #2
For the two-sum problem, if we fix one of the numbers, say
x
, we have to scan the entire array to find the next number
y
which is
value - x
where value is the input parameter. Can we change our array somehow so that this search becomes faster?

Hint #3
The second train of thought for two-sum is, without changing the array, can we use additional space somehow? Like maybe a hash map to speed up the search?
*/
func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}
func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	var result [][]int
	for povit := 0; povit <= len(nums)-3; povit++ {
		if povit!=0&&nums[povit]==nums[povit-1]{
			continue
		}
		start, end := povit+1, len(nums)-1
		for start < end {
			if nums[start]+nums[end]+nums[povit] == 0 {
				result = append(result, []int{nums[povit], nums[start], nums[end]})
				start++
				end--
			} else if nums[povit]+nums[start]+nums[end] > 0 {
				end--
			} else {
				start++
			}
		}
	}

LOOP:
	// 去重
	if len(result)>=2{
		for i:=0;i<len(result)-1;i++{
			var equal bool
			for j:=range result[i]{
				if result[i][j]==result[i+1][j]{
					equal=true
				}else{
					equal=false
					break
				}
			}
			if equal{
				result=append(result[:i],result[i+1:]...)
				goto LOOP
			}
		}
	}
	return result
}
