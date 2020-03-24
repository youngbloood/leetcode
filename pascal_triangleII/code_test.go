package pascal_test

import (
	"fmt"
	pascal "leetcode/pascal_triangleII"
	"os"
	"strconv"
	"testing"
)

func TestGetRow(t *testing.T) {

	// val := uint64(1)
	// for i := 1; i <= 34; i++ {
	// 	val *= uint64(i)
	// }
	// fmt.Println("val = ", val)

	tmpl := "{%s}"
	var all string
	for i := 0; i <= 33; i++ {
		result := pascal.GetRow(i)
		str := "{"
		for _, v := range result {
			str += strconv.Itoa(v) + ","
		}
		str += "},\n"
		all += str
		fmt.Printf("result[%d] = %v\n", i, result)
	}

	str := fmt.Sprintf(tmpl, all)

	file, _ := os.Create("111.txt")
	file.WriteString(str)

}
