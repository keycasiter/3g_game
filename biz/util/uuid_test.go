package util

import (
	"fmt"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	fmt.Printf("%s", GenerateUUID())
}
