package fasthttp

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"net/http"

	"Caesar/pkg/buoys"
	"Caesar/pkg/utils"

	"github.com/valyala/fasthttp"
	"golang.org/x/text/transform"
)

type fastClient struct {

	//设置http请求方式
	method string

	// 请求内容
	body string

	//HTTP请求中的header信息
	header map[string]string

	//HTTP请求中,携带的cookies
	cookies []*http.Cookie

	//设置UserAgent
	userAgent string
}

//初始化一个 http.Request, 并填充属性
func (c *fastClient) RawRequest(address string) (int, http.Header, []byte, error) {
	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)

	request.SetRequestURI(address)
	request.Header.SetMethod(c.method)

	request.SetBody([]byte(c.body))
	// fasthttp does not automatically request a gzipped response.
	// We must explicitly ask for it.
	//request.Header.Set("Accept-Encoding", "gzip")

	//for k, v := range c.header {
	//	request.Header.Set(k, v)
	//}

	//for _, v := range c.cookies {
	//	request.Header.Cookie(c)
	//}

	if len(c.userAgent) > 0 {
		request.Header.SetUserAgent(c.userAgent)
	}

	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	// Perform the request
	if err := fasthttp.Do(request, response); err != nil {
		return buoys.StatusError, nil, nil, err

	}

	//  开始探测网页编码
	bodyReader := bufio.NewReader(bytes.NewReader(response.Body()))
	e := utils.DetermineEncoding(bodyReader)

	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	body, err := ioutil.ReadAll(utf8Reader)
	return response.StatusCode(), nil, body, err
}
