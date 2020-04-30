package other

/*
# Sum of Two Integers
# https://leetcode.com/explore/interview/card/top-interview-questions-medium/114/others/822/

Calculate the sum of two integers a and b, but you are not allowed to use the operator + and -.

Example 1:
Input: a = 1, b = 2
Output: 3
1 => 0001
2 => 0010
1&2 => 0000
1^2 => 0011  // 异或操作，同时为0或1置为0，反之置为1

Example 2:
Input: a = -2, b = 3
Output: 1

2 => 0010
3 => 0011
2&3 => 0010 // 有进位 <<1
2^3 => 0001

*/
func GetSum(a int, b int) int {
	return getSum(a, b)
}

func getSum(a int, b int) int {
	for b != 0 { //循环结束的标志是b==0；这意味着a & b == 0 ，也就是相加不产生进位
		c := a & b //是否有进位
		a = a ^ b  //实现无进位的相加
		b = c << 1 //进位的结果
	}
	return a
}
