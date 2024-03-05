package util

import (
	"fmt"
	"testing"
)

func Test_useDps(t *testing.T) {
	UseDps()
}

func TestUseTps(t *testing.T) {
	u, p, ips := UseTps()
	fmt.Printf("%v ,%v ,%v", u, p, ips)
}
