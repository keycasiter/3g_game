package util

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"strconv"
)

func String2Float64(x string) float64 {
	res, err := strconv.ParseFloat(x, 64)
	if err != nil {
		hlog.CtxErrorf(context.Background(), "String2Float64 err:%v", err)
	}
	return res
}
