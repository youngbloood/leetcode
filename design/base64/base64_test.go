package base64_test

import (
	"fmt"
	"leetcode/design/base64"
	"testing"
)

func TestBase64Encode(t *testing.T) {
	srcs := []string{
		"6666",
		"asdfakjdfmadkjnjnlnhijkhuij",
		"jinoiau2439-==<KLKL2",
		"12340",
	}
	for i, src := range srcs {
		result := base64.Encode([]byte(src))
		rtemp := make([]byte, len(result))
		copy(rtemp, result)
		rsrc := base64.Decode(result)
		t.Log(fmt.Sprintf("srcs[%d] = %s; encode = %s; decode = %s", i, src, string(rtemp), string(rsrc)))
		// assert.Equal(t, rtemp, rsrc)
	}

}

func BenchmarkBase64Encode(b *testing.B) {
	srcs := []string{
		"6666",
		"asdfakjdfmadkjnjnlnhijkhuij",
		"jinoiau2439-==<KLKL2",
	}

	for ii := 0; ii < b.N; ii++ {
		for _, src := range srcs {
			base64.Encode([]byte(src))
			// t.Log(fmt.Sprintf("srcs[%d] = %s,encode = %s", i, src, string(result)))
		}
	}
}
