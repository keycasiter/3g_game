package util

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cast"
)

func ToJsonString(ctx context.Context, obj interface{}) string {
	byt, err := json.Marshal(obj)
	if err != nil {
		hlog.CtxErrorf(ctx, "json Marshal err:%v", err)
		return ""
	}
	return string(byt)
}

func ToIntArrayString(arr []int) string {
	strArr := make([]string, 0)
	for i := 0; i < len(arr); i++ {
		strArr = append(strArr, cast.ToString(arr[i]))
	}
	return ToJsonString(context.Background(), strArr)
}

func ParseJsonObj(ctx context.Context, obj interface{}, str string) {
	err := jsoniter.Unmarshal([]byte(str), obj)
	if err != nil {
		hlog.CtxErrorf(ctx, "json Unmarshal err:%v", err)
	}
}
