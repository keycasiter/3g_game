package util

import (
	"fmt"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	fmt.Printf("%s", GenerateUUID())
}

func TestGenerateUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", GenerateUID())
	}
}
