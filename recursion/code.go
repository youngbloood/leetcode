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

/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example:
Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
*/
func LetterCombinations(digits string) []string {
	return letterCombinations(digits)
}

var phoneDigits = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	return letterCombination(nil, "", digits)
}

func letterCombination(result []string, src, digits string) []string {
	if len(digits) == 0 {
		result = append(result, src)
		return result
	}
	waitStr := phoneDigits[string(digits[0])]
	for _, letter := range waitStr {
		src += string(letter)
		result = letterCombination(result, src, digits[1:])
		src = src[:len(src)-1]
	}
	return result
}
