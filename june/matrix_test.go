package june_test

import (
	"leetcode/june"
	"testing"
)

func TestTwoCitySchedCost(t *testing.T) {
	costs := [][]int{
		[]int{10, 20},  // -10
		[]int{30, 200}, // -170
		[]int{400, 50}, // 350
		[]int{30, 20},  // 10
	}
	t.Log(june.TwoCitySchedCost(costs))

	// 259,448,926,184,840,577   770,54,667,139,118,469
	// 排序
	//  2   4   5   1   0   3     1  4   3   5   2   0
	// 926,840,577,448,259,184   54,118,139,469,667,770
	// 最优解  5+0+3，1+4+2 = 1859
	costs = [][]int{
		[]int{259, 770}, // -511
		[]int{448, 54},  // 434
		[]int{926, 667}, // 259
		[]int{184, 139}, // 45
		[]int{840, 118}, // 722
		[]int{577, 469}, // 108
	}
	t.Log(june.TwoCitySchedCost(costs))
}

func TestReconstructQueue(t *testing.T) {
	people := [][]int{
		[]int{7, 0},
		[]int{4, 4},
		[]int{7, 1},
		[]int{5, 0},
		[]int{6, 1},
		[]int{5, 2},
	}
	t.Log(june.ReconstructQueue(people))
}
