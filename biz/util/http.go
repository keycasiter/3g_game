package util

import (
	"context"
	"fmt"
	"github.com/bytedance/gopkg/util/logger"
	"io/ioutil"
	"net/http"
)

func HttpGet(ctx context.Context, url string, headers map[string]string, params map[string]string) (string, error) {
	paramsStr := ""
	for k, v := range params {
		if paramsStr != "" {
			paramsStr += "&"
		}
		paramsStr += fmt.Sprintf("%s=%s", k, v)
	}

	resp, err := http.Get(fmt.Sprintf("%s?%s", url, paramsStr))
	if err != nil {
		logger.CtxErrorf(ctx, "http get error:%v", err)
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
