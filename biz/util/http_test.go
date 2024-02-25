package util

import (
	"context"
	"fmt"
	"testing"
)

func TestHttpGet(t *testing.T) {
	params := make(map[string]interface{})

	resp, err := HttpGet(context.Background(), "https://m.jiaoyimao.com/api2/sgzItemList2022/getSgzGameZoneItemList", nil, params)
	if err != nil {
		t.Errorf("http get error:%v", err)
		t.Failed()
	}
	fmt.Printf("reps:%v\n", resp)
}
