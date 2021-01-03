package utils

import (
	"net"
	"net/url"
	"regexp"
	"strings"
)

func UrlToAddressAndPort(host string) (ip, port string) {
	/*
		从请求地址中获取ip和端口
		@param: http://127.0.0.1:8080
		return: 127.0.0.1, 8080
	*/
	u, err := url.Parse(host)
	if err != nil {
		return
	}

	h := strings.Split(u.Host, ":")

	if len(h) == 1 {
		return h[0], "80"
	}

	return h[0], h[1]

}

func GetNewHost(host string) string {
	/*
		从url路径中获取目标地址信息，
		次函数的主要目的是为了避免带路径和带参数的地址的干扰
	*/
	schema := "http://"

	ip, port := UrlToAddressAndPort(host)

	if port == "443" {
		schema = "https://"
	}

	return schema + ip + ":" + port + "/"
}

func DomainToIP(host string) string {
	/*
		将输入的域名转换成ip地址格式
		@param: www.baidu.com
		return: 180.101.49.11
	*/

	addr := strings.Trim(host, " ")

	regStr := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	if match, _ := regexp.MatchString(regStr, addr); match {
		return host
	} else {
		addr, err := net.ResolveIPAddr("ip", host)
		if err != nil {
			return "127.0.0.1"
		}

		return addr.String()
	}

}
