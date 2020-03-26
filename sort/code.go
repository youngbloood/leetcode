package sort

/*
Given an array of integers nums, sort the array in ascending order.


Example 1:
Input: nums = [5,2,3,1]
Output: [1,2,3,5]
Example 2:

Input: nums = [5,1,1,2,0,0]
Output: [0,0,1,1,2,5]


Constraints:(约束条件)

1 <= nums.length <= 50000
-50000 <= nums[i] <= 50000
*/
func SortArray(nums []int) []int {
	return topDownSort(nums)
}

func TopDownSort(nums []int) []int {
	return topDownSort(nums)
}

/*
1.In the first step, we divide the list into two sublists.  (Divide)

2.Then in the next step, we recursively sort the sublists in the previous step.  (Conquer)

3.Finally we merge the sorted sublists in the above step repeatedly to obtain the final list of sorted elements.  (Combine)
*/
func topDownSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	pivot := len(nums) / 2
	left := topDownSort(nums[:pivot])
	right := topDownSort(nums[pivot:])
	return merge(left, right)
}

func bottomUpSort(nums []int) []int {

	// var ret []int
	// numsLen := len(nums)
	// var left, right []int
	// for len(ret) != numsLen {

	// }
	return nil
}

func merge(left, right []int) []int {
	deal := make([]int, len(left)+len(right))
	var index, leftPtr, rightPtr int

	for leftPtr < len(left) && rightPtr < len(right) {
		if left[leftPtr] < right[rightPtr] {
			deal[index] = left[leftPtr]
			leftPtr++
		} else {
			deal[index] = right[rightPtr]
			rightPtr++
		}
		index++
	}
	for leftPtr < len(left) {
		deal[index] = left[leftPtr]
		index++
		leftPtr++
	}
	for rightPtr < len(right) {
		deal[index] = right[rightPtr]
		index++
		rightPtr++
	}

	return deal
}

// QuickSort. 快速排序
func QuickSort(in []int) {
	qSort(in, 0, len(in)-1)
}

func qSort(in []int, start, end int) {
	if start < end {
		p := partition(in, start, end)
		qSort(in, start, p-1)
		qSort(in, p+1, end)
	}
}
func partition(in []int, start, end int) int {
	pivot := in[end]
	i := start
	for j := start; j < end; j++ {
		if in[j] < pivot {
			in[i], in[j] = in[j], in[i]
			i++
		}
	}
	in[i], in[end] = in[end], in[i]
	return i
}
