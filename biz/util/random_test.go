package util

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%0.2f\n", Random(-0.1, 0.1))
	}
}

// 1000 sample -> true:795 false:205
// 10000 sample -> true:7959 false:2041
func TestGenerateRate(t *testing.T) {
	tr := 0
	f := 0
	for i := 0; i < 10000; i++ {
		res := GenerateRate(0.8)
		if res == true {
			tr++
		} else {
			f++
		}
	}
	fmt.Printf("true:%d\nfalse:%d", tr, f)
}
