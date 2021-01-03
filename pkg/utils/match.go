package utils

import (
	"net"
	"regexp"
	"strings"
)

/*
	利用正则表达式检查路径是否是文件夹
*/
func MatchDir(name string) bool {
	matchedDir, err := regexp.MatchString(`\..{3,4}$`, name)
	if err != nil {
		return false
	}
	if !matchedDir && !(strings.Contains(name, "=") && strings.Contains(name, "&")) {
		return true

	} else {
		return false
	}
}

/*
	检查path是否是动态文件
*/
func MatchDynamic(name string) bool {
	matchedFIle, err := regexp.MatchString(`\.(php|asp|aspx|jsp|jspx])$`, name)
	if err != nil {
		return false
	}
	if matchedFIle {
		return true

	} else {
		return false
	}
}

/*
	正则校验socks5/http/https代理
*/
func MatchProxy(address string) bool {

	matched, err := regexp.MatchString(`^(https?|socks5)://.*?:[0-9]{1,5}$`, address)
	if err != nil {
		return false
	}
	if matched {
		return true

	} else {
		return false
	}

}

/*
	匹配cookie的规则，比如: username=admin; userid=1; PHPSESSID=9d1q9o4927a42p2thki1ql82p7
*/
func MatchCookie(cookies string) bool {
	matched, err := regexp.MatchString(`^([\w]*?=[\w]*?; )*([\w]*?=[\w]*?)$`, cookies)
	if err != nil {
		return false
	}
	if matched {
		return true

	} else {
		return false
	}
}

/*
正则提取网站标题
*/
func MatchTitle(html string) string {
	matched := regexp.MustCompile(`<title>([\S\s]*?)</title>`)
	results := matched.FindStringSubmatch(html)

	if len(results) > 1 {
		return results[1]
	}

	return ""
}

func IsPrivateIP(ipAddress string) bool {
	/*
		用来判断IP地址是否是私有IP，是的话返回true
	*/

	var ip net.IP

	ip = net.ParseIP(ipAddress)

	var privateIPBlocks []*net.IPNet

	for _, cidr := range []string{
		"127.0.0.0/8",    // IPv4 loopback
		"10.0.0.0/8",     // RFC1918
		"172.16.0.0/12",  // RFC1918
		"192.168.0.0/16", // RFC1918
		"169.254.0.0/16", // RFC3927 link-local
		"::1/128",        // IPv6 loopback
		"fe80::/10",      // IPv6 link-local
		"fc00::/7",       // IPv6 unique local addr
	} {
		_, block, err := net.ParseCIDR(cidr)
		if err != nil {
			return false
		}

		privateIPBlocks = append(privateIPBlocks, block)
	}

	if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
		return true
	}

	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}
