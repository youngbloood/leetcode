package pow

/*
Implement pow(x, n), which calculates x raised to the power n (xn).

Example 1:
Input: 2.00000, 10
Output: 1024.00000

Example 2:
Input: 2.10000, 3
Output: 9.26100

Example 3:
Input: 2.00000, -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
Note:

-100.0 < x < 100.0
n is a 32-bit signed integer, within the range [−2^31, 2^31 − 1]
*/
// MyPow

// x^n=x^n-1*x
func MyPow(x float64, n int) float64 {
	if x < (0.00001) && x > (-0.00001) {
		x = 0
	}
	switch x {
	case 0, 1:
		return x
	case -1:
		if n%2 == 0 {
			return 1
		}
		return -1
	}
	if n < 0 {
		n = -1 * n
		return 1 / pow(x, n)
	}
	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	xTemp := float64(1)
	for i := 1; i <= n; i++ {
		xTemp *= x
	}
	return xTemp
}
