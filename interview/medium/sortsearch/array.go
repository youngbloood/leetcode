package sortsearch

import (
	"sort"
)

/*
# Top K Frequent Elements
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/799/

Given a non-empty array of integers, return the k most frequent elements.

Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
Example 2:

Input: nums = [1], k = 1
Output: [1]
Note:

You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
It's guaranteed that the answer is unique, in other words the set of the top k frequent elements is unique.
You can return the answer in any order.
*/

func TopKFrequent(nums []int, k int) []int {
	return topKFrequent(nums, k)
}

func topKFrequent(nums []int, k int) []int {
	// 统计元素的频次
	freqMap := make(map[int]int, 0)
	for i := range nums {
		if _, ok := freqMap[nums[i]]; ok {
			freqMap[nums[i]]++
		} else {
			freqMap[nums[i]] = 1
		}
	}

	// 桶排序
	bucket := make(kfreqs, 0, len(nums))

	for k, v := range freqMap {
		bucket = append(bucket, kfreq{
			num: v,
			val: k,
		})
	}
	sort.Sort(bucket)

	return bucket.GetKFreq(k)
}

type kfreq struct {
	num int // 数量
	val int
}

type kfreqs []kfreq

func (k kfreqs) Less(i, j int) bool { return k[i].num > k[j].num }
func (k kfreqs) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }
func (k kfreqs) Len() int           { return len(k) }

func (k kfreqs) GetKFreq(i int) []int {
	if !(len(k) >= i) {
		return nil
	}
	var result []int
	for iter := 0; iter < i; iter++ {
		result = append(result, k[iter].val)
	}

	return result
}

/*
# Kth Largest Element in an Array
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/800/

Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:

Input: [3,2,1,5,6,4] and k = 2
Output: 5
Example 2:

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
Note:
You may assume k is always valid, 1 ≤ k ≤ array's length.
*/
func FindKthLargest(nums []int, k int) int {
	return findKthLargest(nums, k)
}

func findKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return -1
	}
	sort.Ints(nums)
	return nums[len(nums)-k]
}

/*
# Find Peak Element
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/801/

A peak element is an element that is greater than its neighbors.

Given an input array nums, where nums[i] ≠ nums[i+1], find a peak element and return its index.

The array may contain multiple peaks, in that case return the index to any one of the peaks is fine.

You may imagine that nums[-1] = nums[n] = -∞.

Example 1:

Input: nums = [1,2,3,1]
Output: 2
Explanation: 3 is a peak element and your function should return the index number 2.
Example 2:

Input: nums = [1,2,1,3,5,6,4]
Output: 1 or 5
Explanation: Your function can return either index number 1 where the peak element is 2,
             or index number 5 where the peak element is 6.
Note:

Your solution should be in logarithmic complexity.
*/
func FindPeakElement(nums []int) int {
	return findPeakElement(nums)
}
func findPeakElement(nums []int) int {
	if nums == nil {
		return -1
	}
	if len(nums) <= 1 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		// 对0的处理
		if i == 0 && nums[i] > nums[i+1] {
			return i
		} else if i == 0 {
			continue
		}
		// 对n-1的处理
		if i == len(nums)-1 && nums[i] > nums[i-1] {
			return i
		} else if i == len(nums)-1 {
			continue
		}

		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}

	}
	return -1
}

/*
# Search for a Range
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/802/

Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
*/
func SearchRange(nums []int, target int) []int {
	return searchRange(nums, target)
}
func searchRange(nums []int, target int) []int {
	start, end := 0, len(nums)-1
	var povit = -1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			povit = mid
			break
		}
	}
	return searchForRange(nums, povit)
}

func searchForRange(nums []int, povit int) []int {
	if povit == -1 {
		return []int{-1, -1}
	}
	origin := povit
	left, right := povit, povit
	for left >= 0 && nums[left] == nums[origin] {
		left--
	}
	for right <= len(nums)-1 && nums[right] == nums[origin] {
		right++
	}
	return []int{left + 1, right - 1}
}

/*
# Merge Intervals
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/110/sorting-and-searching/803/

Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.
*/
func Merge(intervals [][]int) [][]int {
	return merge(intervals)
}
func merge(intervals [][]int) [][]int {
	var result = intervals
	var hasMerged bool = true
	for hasMerged {
		result, hasMerged = mergeIntervals(result)
	}
	return result
}

func mergeIntervals(intervals [][]int) ([][]int, bool) {
	if len(intervals) <= 1 {
		return intervals, false
	}
	var result [][]int

	var hasMerged bool

	for i, val := range intervals {
		if i == 0 {
			result = append(result, val)
			continue
		}

		// 循环遍历result，判断其中值是否有修改过
		var isChanged bool
		for j := range result {
			// {}
			resLeft, resRight := result[j][0], result[j][1]
			// ||
			nowLeft, nowRight := val[0], val[1]

			if nowLeft <= resRight && nowLeft >= resLeft && nowRight >= resRight {
				// {|}|
				result[j][1] = nowRight
				hasMerged = true
				isChanged = true
				break
			} else if nowLeft <= resLeft && nowRight >= resRight {
				// |{}|
				result[j][1] = nowRight
				result[j][0] = nowLeft
				hasMerged = true
				isChanged = true
				break
			} else if nowLeft >= resLeft && nowRight <= resRight {
				// {||}
				isChanged = true
				break
			} else if nowRight >= resLeft && nowRight <= resRight && nowLeft <= resLeft {
				// |{|}
				result[j][0] = nowLeft
				hasMerged = true
				isChanged = true
				break
			}
		}
		// 未修改过，添加至result中
		if !isChanged {
			result = append(result, val)
		}
	}
	return result, hasMerged
}
