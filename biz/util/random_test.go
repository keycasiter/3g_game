package util

import (
	"fmt"
	"math"
	"testing"
)

func TestRandom(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%.0f\n", math.Round(Random(0, 10000)))
	}
}

// 1000 sample -> true:795 false:205
// 10000 sample -> true:7959 false:2041
func TestGenerateRate(t *testing.T) {
	tr := 0
	f := 0
	for i := 0; i < 10000; i++ {
		res := GenerateRate(1.2)
		if res == true {
			tr++
		} else {
			f++
		}
	}
	fmt.Printf("true:%d\nfalse:%d", tr, f)
}

func TestGenerateHitIdxArr(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", GenerateHitIdxArr(2, 3))
	}
}

func TestGenerateHitOneIdx(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", GenerateHitOneIdx(1))
	}
}

func TestGenerateHitIdxMap(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", GenerateHitIdxMap(2, 3))
	}
}

func TestGenerateHitTwoOrThreeIdxMap(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", GenerateHitTwoOrThreeIdxArr())
	}
}
