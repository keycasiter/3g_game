package util

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"log"

	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/auth"
	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/client"
	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/signtype"
)

// 私密代理使用示例

// 接口鉴权说明：
// 接口鉴权方式为必填项, 目前支持的鉴权方式有"token" 和 "hmacsha1"两种
// 可选值为signtype.TOKEN和signtype.HmacSha1 或直接传"token"或"hmacsha1"

// 返回值说明:
// 所有返回值都包括两个值，第一个为目标值，第二个为error类型, 值为nil说明成功，不为nil说明失败

func UseDps() (username string, password string, ipPool []string) {
	auth := auth.Auth{SecretID: "ozf9w2zl2j0o74a2osyc", SecretKey: "zbm8r8pfemhu5ctvrpzm9e9v3bxx6y5i"}
	client := client.Client{Auth: auth}

	// 获取订单到期时间, 返回时间字符串
	expireTime, err := client.GetOrderExpireTime(signtype.HmacSha1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("expire time: ", expireTime)

	// 获取User-Agent, 返回ua列表
	ua, err := client.GetUA(10, signtype.HmacSha1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("User-Agent: ", ua)

	// 获取ip白名单, 返回ip切片, 类型为[]string
	ipWhitelist, err := client.GetIPWhitelist(signtype.Token)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("ip whitelist: ", ipWhitelist)

	//设置ip白名单，参数类型为[]string
	_, err = client.SetIPWhitelist([]string{"test_ip1", "test_ip2"}, signtype.HmacSha1)
	if err != nil {
		log.Println(err)
	}

	// 提取私密代理, 参数有: 提取数量、鉴权方式及其他参数(放入map[string]interface{}中, 若无则传入nil)
	// (具体有哪些其他参数请参考帮助中心: "https://www.kuaidaili.com/doc/api/getdps/")
	params := map[string]interface{}{"format": "json", "area": "北京,上海"}
	ips, err := client.GetDps(1, signtype.HmacSha1, params)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("ips: ", ips)

	// 检测私密代理有效性， 返回map[string]bool, ip:true/false
	valids, err := client.CheckDpsValid(ips, signtype.HmacSha1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("valids: ", valids)

	// 获取私密代理剩余时间(单位为秒), 返回map[string]string, ip:seconds
	seconds, err := client.GetDpsValidTime(ips, signtype.Token)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("seconds: ", seconds)

	// 获取计数版ip余额（仅私密代理计数版）
	balance, err := client.GetIPBalance(signtype.HmacSha1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("balance: ", balance)

	// 获取订单访问代理IP的鉴权信息
	proxyAuthorization, err := client.GetProxyAuthorization(1, signtype.HmacSha1)
	fmt.Println("proxyAuthorization: ", proxyAuthorization)
	username = proxyAuthorization["username"]
	password = proxyAuthorization["password"]
	for ip, valid := range valids {
		if !valid {
			continue
		}
		ipPool = append(ipPool, ip)
	}
	return username, password, ipPool
}

// 隧道代理使用示例

// 接口鉴权说明：
// 接口鉴权方式为必填项, 目前支持的鉴权方式有"simple" 和 "hmacsha1"两种
// 可选值为signtype.SIMPLE和signtype.HmacSha1 或直接传"simple"或"hmacsha1"

// 返回值说明:
// 所有返回值都包括两个值，第一个为目标值，第二个为error类型, 值为nil说明成功，不为nil说明失败
// https://tps.kdlapi.com/api/gettps/?secret_id=oe2x78yrganqi1xraao1&num=1&signature=2bmzc5uyweff4hxc1n1e4gmouoja2kly&pt=1&format=json&sep=1
func UseTps() (username string, password string, ipPool []string, err error) {
	auth := auth.Auth{SecretID: "oe2x78yrganqi1xraao1", SecretKey: "2bmzc5uyweff4hxc1n1e4gmouoja2kly"}
	client := client.Client{Auth: auth}

	// 获取订单到期时间, 返回时间字符串
	//expireTime, err := client.GetOrderExpireTime(signtype.HmacSha1)
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println("expire time: ", expireTime)

	// 获取ip白名单, 返回ip切片, 类型为[]string
	//ipWhitelist, err := client.GetIPWhitelist(signtype.SIMPLE)
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println("ip whitelist: ", ipWhitelist)

	//设置ip白名单，参数类型为[]string
	//_, err = client.SetIPWhitelist([]string{"test_ip1", "test_ip2"}, signtype.HmacSha1)
	//if err != nil {
	//	log.Println(err)
	//}

	// 改变当前隧道ip
	//newIP, err := client.ChangeTpsIP(signtype.SIMPLE)
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println("new_ip: ", newIP)

	// 获取隧道代理当前的ip
	//ip, err := client.TpsCurrentIP(signtype.SIMPLE)
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println("current_ip: ", ip)

	// 获取订单访问代理IP的鉴权信息
	proxyAuthorization, err := client.GetProxyAuthorization(1, signtype.HmacSha1)
	//fmt.Println("proxyAuthorization: ", proxyAuthorization)

	//提取隧道代理IP, 参数有: 提取数量、鉴权方式及其他参数(放入map[string]interface{}中, 若无则传入nil)
	//(具体有哪些其他参数请参考帮助中心: "https://www.kuaidaili.com/doc/api/gettps/")
	params := map[string]interface{}{"format": "json"}
	ips, err := client.GetTpsIp(1, signtype.HmacSha1, params)
	if err != nil {
		hlog.Errorf("%v", err)
		return "", "", nil, err
	}
	return proxyAuthorization["username"], proxyAuthorization["password"], ips, nil
}
