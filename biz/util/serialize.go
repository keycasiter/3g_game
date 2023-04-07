package util

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/json"
)

func ToJsonString(ctx context.Context, obj interface{}) string {
	byt, err := json.Marshal(obj)
	if err != nil {
		hlog.CtxErrorf(ctx, "json Marshal err:%v", err)
		return ""
	}
	return string(byt)
}

func ParseJsonObj(ctx context.Context, obj interface{}, str string) {
	err := json.Unmarshal([]byte(str), obj)
	if err != nil {
		hlog.CtxErrorf(ctx, "json Unmarshal err:%v", err)
	}
}
