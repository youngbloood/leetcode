package matrix_test

import (
	"leetcode/matrix"
	"testing"

	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func initMatrix(vals ...[]int) [][]int {
	for i := 0; i < len(vals)-1; i++ {
		if len(vals[i]) != len(vals[i+1]) {
			panic("长度不相等")
		}
	}
	result := make([][]int, len(vals))
	for i := range vals {
		result[i] = vals[i]
	}
	return result
}

func TestSearchMatrix(t *testing.T) {
	mtx := initMatrix(
		[]int{1, 4, 7, 11, 15},
		[]int{2, 5, 8, 12, 19},
		[]int{3, 6, 9, 16, 22},
		[]int{10, 13, 14, 17, 24},
		[]int{18, 21, 23, 26, 30},
	)
	assert.So(matrix.SearchMatrix(mtx, 5), should.Equal, true).Println()
}
