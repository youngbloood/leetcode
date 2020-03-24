package climbstairs_test

import (
	"fmt"
	"leetcode/climbstairs"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	fmt.Println(climbstairs.ClimbStairs(44))
}

func TestA(t *testing.T) {
	fmt.Println(climbstairs.A(27, 17))
}

func TestABig(t *testing.T) {
	fmt.Println(climbstairs.ABig(27, 17))
}
