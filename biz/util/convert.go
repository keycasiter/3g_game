package util

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/spf13/cast"
	"strconv"
	"strings"
)

func String2Float64(x string) float64 {
	res, err := strconv.ParseFloat(x, 64)
	if err != nil {
		hlog.CtxErrorf(context.Background(), "String2Float64 err:%v", err)
	}
	return res
}

func StringToIntArray(str string) []int64 {
	strArr := strings.Split(str, ",")
	intArr := make([]int64, 0)
	for _, i := range strArr {
		intArr = append(intArr, cast.ToInt64(i))
	}
	return intArr
}

func StructToMap(obj interface{}) map[string]interface{} {
	jsonBytes, _ := json.Marshal(obj)
	data := make(map[string]interface{})
	json.Unmarshal(jsonBytes, &data)
	return data
}
