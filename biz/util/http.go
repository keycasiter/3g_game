package util

import (
	"context"
	"fmt"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/spf13/cast"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

type ProxyMode int

const (
	ProxyMode_NoProxy   = 1
	ProxyMode_KuaiDaili = 2
	ProxyMode_Fixed     = 3
)

var UseProxyMode = ProxyMode_KuaiDaili
var cnt = 0

var username, password string
var ips []string

func HttpGet(ctx context.Context, requestUrl string, headers map[string]interface{}, params map[string]interface{}) (string, error) {
	path := ""
	paramsStr := ""
	for k, v := range params {
		if paramsStr != "" {
			paramsStr += "&"
		}
		paramsStr += fmt.Sprintf("%s=%s", k, url.QueryEscape(cast.ToString(v)))
	}
	path = fmt.Sprintf("%s?%s", requestUrl, paramsStr)

	var client *http.Client

	switch UseProxyMode {
	case ProxyMode_NoProxy:
		client = &http.Client{}
	case ProxyMode_KuaiDaili:
		username, password, ips = UseDps()
		// 代理服务器
		proxy_raw := ips[rand.Intn(len(ips))]
		hlog.CtxInfof(ctx, "[proxy] %s", proxy_raw)
		proxy_str := fmt.Sprintf("http://%s:%s@%s", username, password, proxy_raw)
		proxy, err := url.Parse(proxy_str)

		if err != nil {
			logger.CtxErrorf(ctx, "url Parse error:%v", err)
			return "", err
		}
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
			},
			Timeout: 5 * time.Second,
		}
	case ProxyMode_Fixed:
		// 代理服务器
		ips := []string{
			"http://223.240.209.149:9999",
		}
		proxy_raw := ips[rand.Intn(len(ips))]
		hlog.CtxInfof(ctx, "[proxy] %s", proxy_raw)
		proxy_str := fmt.Sprintf("http://%s:%s@%s", username, password, proxy_raw)
		proxy, err := url.Parse(proxy_str)
		if err != nil {
			logger.CtxErrorf(ctx, "url Parse error:%v", err)
			return "", err
		}
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
			},
			Timeout: 5 * time.Second,
		}
	}

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
