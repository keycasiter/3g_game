package util

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	jsoniter "github.com/json-iterator/go"
)

func ToJsonString(ctx context.Context, obj interface{}) string {
	byt, err := jsoniter.Marshal(obj)
	if err != nil {
		hlog.CtxErrorf(ctx, "json Marshal err:%v", err)
		return ""
	}
	return string(byt)
}

func ParseJsonObj(ctx context.Context, obj interface{}, str string) {
	err := jsoniter.Unmarshal([]byte(str), obj)
	if err != nil {
		hlog.CtxErrorf(ctx, "json Unmarshal err:%v", err)
	}
}
