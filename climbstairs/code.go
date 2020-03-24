package climbstairs

import (
	"math/big"
)

/*
You are climbing a stair case. It takes n steps to reach to the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
Note: Given n will be a positive integer.

Example 1:
Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
*/
func ClimbStairs(n int) int {
	if n <= 0 {
		return 0
	}

	var num uint64
	for i := 0; i <= n/2; i++ {
		// n-i:表示1和2的总数量
		// i:表示2的数量
		// n-2*i:表示1的数量
		val := ABig(n-i, i)
		// 计算方式：A(n-i,n-i)/A(i,i)/A(n-2*i,n-2*i)；简化得A(n-i,i)/A(i,i)
		val.Div(val, ABig(i, i))
		num += val.Uint64()
	}
	return int(num)
}

func ABig(n, m int) *big.Int {
	if n == 0 || m == 0 {
		return big.NewInt(1)
	}
	result := big.NewInt(1)
	one := big.NewInt(1)
	for i := 0; i < m; i++ {
		result.Mul(result, one.SetInt64(int64(n)))
		n--
	}
	return result
}

func climbStairs(n int) int {
	switch n {
	case <=0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	}	
	output:=make([]int,n)
	output[0]=1
	output[1]=2
	for i:=2;i<n;i++{
		output[i]=output[i-1]+output[i-2]
	}
	return output[n-1]
}
