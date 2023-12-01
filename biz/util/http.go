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
	path := ""
	paramsStr := ""
	for k, v := range params {
		if paramsStr != "" {
			paramsStr += "&"
		}
		paramsStr += fmt.Sprintf("%s=%s", k, cast.ToString(v))
	}
	if len(paramsStr) > 0 {
		path = fmt.Sprintf("%s?%s", url, paramsStr)
	} else {
		path = fmt.Sprintf("%s?%s", url, paramsStr)
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		logger.CtxErrorf(ctx, "http NewRequest error:%v", err)
		return "", err
	}
	// 添加请求头
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("authority", "m.jiaoyimao.com")
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		logger.CtxErrorf(ctx, "http request error:%v", err)
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
