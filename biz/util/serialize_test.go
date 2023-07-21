package util

import (
	"context"
	"fmt"
	"testing"
)

func TestToJsonString(t *testing.T) {
	fmt.Printf("%s", ToJsonString(context.Background(), []int{1, 2, 3}))
}

func TestToIntArrayString(t *testing.T) {
	fmt.Printf("%s", ToIntArrayString([]int{1, 2, 3}))
}
