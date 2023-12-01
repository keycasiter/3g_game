package util

import (
	"context"
	"fmt"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
)

func HttpGet(ctx context.Context, url string, headers map[string]interface{}, params map[string]interface{}) (string, error) {
	paramsStr := ""
	for k, v := range params {
		if paramsStr != "" {
			paramsStr += "&"
		}
		paramsStr += fmt.Sprintf("%s=%s", k, cast.ToString(v))
	}
	path := fmt.Sprintf("%s?%s", url, paramsStr)
	resp, err := http.Get(path)
	if err != nil {
		logger.CtxErrorf(ctx, "http get error:%v", err)
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
