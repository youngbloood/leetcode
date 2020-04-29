package math

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
# Factorial Trailing Zeroes
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/113/math/816/

Solution
Given an integer n, return the number of trailing zeroes in n!.

Example 1:

Input: 3
Output: 0
Explanation: 3! = 6, no trailing zero.

Example 2:
Input: 5
Output: 1
Explanation: 5! = 120, one trailing zero.
Note: Your solution should be in logarithmic time complexity.
*/
func TrailingZeroes(n int) int {
	return trailingZeroes(n)
}
func trailingZeroes(n int) int {
	var count int
	for n > 0 {
		count += n / 5
		n /= 5
	}
	return count
}

/*
# Excel Sheet Column Number
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/113/math/817/

Given a column title as appear in an Excel sheet, return its corresponding column number.

For example:

    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
    ...
Example 1:
Input: "A"
Output: 1

Example 2:
Input: "AB"  = 1*(26^1)+2*(26^0)
Output: 28

Example 3:
Input: "ZY"  = 26*(26^1)+25*(26^0)
Output: 701
*/

func TitleToNumber(s string) int {
	return titleToNumber(s)
}
func titleToNumber(s string) int {
	var count int
	n := len(s)
	for i := 0; i < n; i++ {
		num := s[i] - 64
		count += int(num) * (int(math.Pow(float64(26), float64(n-1-i))))
	}
	return count
}

/*
# Divide Two Integers
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/113/math/820/

Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.

Return the quotient after dividing dividend by divisor.

The integer division should truncate toward zero, which means losing its fractional part. For example, truncate(8.345) = 8 and truncate(-2.7335) = -2.

Example 1:

Input: dividend = 10, divisor = 3
Output: 3
Explanation: 10/3 = truncate(3.33333..) = 3.
Example 2:

Input: dividend = 7, divisor = -3
Output: -2
Explanation: 7/-3 = truncate(-2.33333..) = -2.
Note:

Both dividend and divisor will be 32-bit signed integers.
The divisor will never be 0.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose of this problem, assume that your function returns 231 − 1 when the division result overflows.
*/
// 要求不能使用乘，除，模运算
func Divide(dividend int, divisor int) int {
	return divide(dividend, divisor)
}
func divide(dividend int, divisor int) int {
	var quotient int

	if dividend == 0 {
		return 0
	}
	if divisor == 0 {
		panic("divisor is zero")
	}

	if dividend > 0 && divisor > 0 {
		for dividend >= 0 {
			dividend -= divisor
			quotient++
		}
		quotient--
	} else if dividend < 0 && divisor < 0 {
		for dividend <= 0 {
			dividend -= divisor
			quotient++
		}
		quotient--
	} else if dividend < 0 && divisor > 0 {
		for dividend <= 0 {
			dividend += divisor
			quotient--
		}
		quotient++
	} else {
		for dividend >= 0 {
			dividend += divisor
			quotient--
		}
		quotient++
	}

	top := 1 << 31
	if quotient >= top {
		return top - 1
	} else if quotient < -1*top {
		return -1 * top
	}
	return quotient
}

/*
# Fraction to Recurring Decimal
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/113/math/821/


Given two integers representing the numerator and denominator of a fraction, return the fraction in string format.

If the fractional part is repeating, enclose the repeating part in parentheses.

Example 1:
Input: numerator = 1, denominator = 2
Output: "0.5"

Example 2:
Input: numerator = 2, denominator = 1
Output: "2"

Example 3:
Input: numerator = 2, denominator = 3
Output: "0.(6)"


Hint #1
No scary math, just apply elementary math knowledge. Still remember how to perform a long division?

Hint #2
Try a long division on 4/9, the repeating part is obvious. Now try 4/333. Do you see a pattern?

Hint #3
Notice that once the remainder starts repeating, so does the divided result.

Hint #4
Be wary of edge cases! List out as many test cases as you can think of and test your code thoroughly.
*/
func FractionToDecimal(numerator int, denominator int) string {
	return fractionToDecimal(numerator, denominator)
}
func fractionToDecimal(numerator int, denominator int) string {
	quotient := numerator / denominator
	mod := numerator % denominator
	if mod == 0 {
		return strconv.Itoa(quotient)
	}
	mods := findFractionCycle(mod, denominator)
	return fmt.Sprintf("%d.%s", quotient, mods)
}

// 寻找循环数组
func findFractionCycle(mod, denominator int) (mods string) {

	var list []int
	var cycle []int
	var set = make(map[int]int, 0) //val:index
	var first int = -1
	var firstTemp int
	var period int
	var periodTemp int
	var isCycle bool

	for mod != 0 {
		mod *= 10
		inQuotient := mod / denominator

		list = append(list, inQuotient)
		if index, exist := set[inQuotient]; exist {
			if first == -1 {
				first = index
				firstTemp = first
				period = len(list) - 1 - index
				periodTemp = period
			} else {
				first++
				if list[first] == inQuotient && first < len(list) {
					period--
					if period <= 0 {
						cycle = list[firstTemp : firstTemp+periodTemp]
						isCycle = true
						break
					}
					if inQuotient > 0 {
						mod %= denominator
					}
					continue
				}
			}
		} else {
			period++
			periodTemp++
		}
		set[inQuotient] = len(list) - 1 // 设置值
		mods += strconv.Itoa(inQuotient)
		if inQuotient > 0 {
			mod %= denominator
		}
	}

	if isCycle {
		cycleStr := make([]string, len(cycle)+firstTemp)
		for i := 0; i < firstTemp; i++ {
			cycleStr[i] = strconv.Itoa(list[i])
		}

		if cycle != nil {
			for i, v := range cycle {
				cycleStr[i+firstTemp] = strconv.Itoa(v)
			}
		}
		mods = strings.Join(cycleStr[:firstTemp], "") + "(" + strings.Join(cycleStr[firstTemp:], "") + ")"
	} else {
		cycleStr := make([]string, len(cycle)+period)
		for i, v := range list {
			cycleStr[i] = strconv.Itoa(v)
		}
		mods = strings.Join(cycleStr, "")
	}
	return
}
