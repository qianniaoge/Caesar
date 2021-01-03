package generated

import (
	"regexp"
	"strings"

	"Caesar/pkg/buoys"
)

// 从文本中获取http请求信息
func ParseRequestFromFile(requests string) (target string, method string, userAgent string, cookies string, headers map[string]string, data string) {

	var (
		uri        = ""
		newLine    string
		getPostReq = false
		port       = ""
		params     = false
		schema     = "http://"
		host       = ""
	)

	// 保存http头文件
	headers = make(map[string]string)

	if !strings.Contains(requests, "HTTP/") {
		return
	}

	lines := strings.Split(requests, "\n")

	for index := 0; index < len(lines); index++ {
		line := lines[index]

		if strings.HasSuffix(line, "\r") {
			newLine = "\r\n"
		} else {
			newLine = "\n"
		}

		// http请求头第一行文件
		matched := regexp.MustCompile(`^([A-Z]+) (.+) HTTP/[\d.]+$`)
		results := matched.FindStringSubmatch(line)

		if len(results) >= 3 {
			// 处理http请求的第一行 METHOD URI HTTPVersion
			method = results[1]
			uri = results[2]
			getPostReq = true

		} else if matched, _ := regexp.MatchString(`^\S+:`, line); matched {
			// 处理头文件
			n := strings.SplitN(line, ":", 2)
			key := n[0]
			value := strings.TrimLeft(n[1], " ")

			if strings.ToUpper(key) == strings.ToUpper(buoys.COOKIE) {
				// 处理cookie headers
				cookies = value
			} else if strings.ToUpper(key) == strings.ToUpper(buoys.HOST) {
				// 处理host headers
				host = value
				temp := strings.Split(host, ":")
				//address = temp[0]
				port = temp[1]

			} else if strings.ToUpper(key) == strings.ToUpper(buoys.UserAgent) {
				// 处理user-agent headers
				userAgent = value

			} else if strings.ToUpper(key) == strings.ToUpper(buoys.ContentLength) {
				// Avoid to add a static content length header to
				// headers and consider the following lines as
				// POSTed data

				params = true
			} else if strings.ToUpper(key) != strings.ToUpper(buoys.ProxyConnection) && strings.ToUpper(key) != strings.ToUpper(buoys.ContentLength) && strings.ToUpper(key) != strings.ToUpper(buoys.IfModifiedSince) && strings.ToUpper(key) != strings.ToUpper(buoys.IfNoneMatch) {
				// Avoid proxy and connection type related headers
				headers[key] = value
			}

		} else {
			line = strings.Trim(line, "\r")
			line = strings.Trim(line, "\n")
			if len(line) == 0 {
				continue
			}

			if getPostReq && params {
				data = data + line + newLine
			}
		}
	}

	if port == "443" {
		schema = "https://"
	}

	target = schema + host + uri

	return

}
