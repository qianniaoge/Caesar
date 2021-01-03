package units

import (
	"Caesar/internal/relation"
	"Caesar/pkg/builder/net/stand"
	"testing"
)

func TestProxy5u(t *testing.T) {

	// http://www.data5u.com/vipip/tunnel.html 无忧代理，动态转发
	// 5u的动态端口转发效率比较低，20个请求花费106.20s

	//proxyUsername := "【这里替换成你的IP提取码】"
	proxyUsername := ""
	//proxyPwd := "【这里替换成你的动态转发密码】"
	proxyPwd := ""
	proxyIp := "tunnel.data5u.com:56789"
	dynamicProxy := "http://" + proxyUsername + ":" + proxyPwd + "@" + proxyIp

	proxy := []string{dynamicProxy}[0]

	for i := 0; i < 20; i++ {
		build := stand.NewClientBuilder().SetUserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:76.0) Gecko/20100101 Firefox/76.0").SetHeader(relation.Browser.Headers).SetCookie(relation.Browser.Cookie).SetTimeOut(relation.Browser.TimeOut).SetProxy(proxy).SetSkipVerify(true).StandBuilder()
		_, _, body, _ := build.Get("http://myip.ipip.net/")
		println(string(body))
	}

}
