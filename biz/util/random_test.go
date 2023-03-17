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
