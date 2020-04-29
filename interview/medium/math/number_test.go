package math_test

import (
	"leetcode/interview/medium/math"
	"testing"
)

func TestTrailingZeroes(t *testing.T) {
	t.Log(math.TrailingZeroes(3))
	t.Log(math.TrailingZeroes(5))
	t.Log(math.TrailingZeroes(10))
	t.Log(math.TrailingZeroes(30))
	t.Log(math.TrailingZeroes(50))
	t.Log(math.TrailingZeroes(2527))
	t.Log(math.TrailingZeroes(1808548329))
}

func TestTitleToNumber(t *testing.T) {
	t.Log(math.TitleToNumber("A"))
	t.Log(math.TitleToNumber("AB"))
	t.Log(math.TitleToNumber("ZY"))
}

func TestDivide(t *testing.T) {
	t.Log(math.Divide(9, 3))
	t.Log(math.Divide(8, 3))
	t.Log(math.Divide(-9, -3))
	t.Log(math.Divide(-8, -3))

	t.Log(math.Divide(-9, 3))
	t.Log(math.Divide(-8, 3))

	t.Log(math.Divide(9, -3))
	t.Log(math.Divide(8, -3))

	t.Log(math.Divide(-2147483648, -1))
}

func TestFractionToDecimal(t *testing.T) {
	t.Log(math.FractionToDecimal(1, 2))
	t.Log(math.FractionToDecimal(2, 1))
	t.Log(math.FractionToDecimal(2, 3))
	t.Log(math.FractionToDecimal(10, 3))

	t.Log(math.FractionToDecimal(4, 9))
	t.Log(math.FractionToDecimal(4, 333))

	t.Log(math.FractionToDecimal(1, 6))

	t.Log(math.FractionToDecimal(1, 333))
}
