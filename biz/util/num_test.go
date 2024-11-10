package util

import (
	"fmt"
	"testing"

	"github.com/spf13/cast"
)

func TestFloat64Remain(t *testing.T) {
	fmt.Printf("%v", cast.ToString(Float64Remain(0.29123123123)*100))
}
