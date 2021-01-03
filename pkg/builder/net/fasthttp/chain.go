package fasthttp

import (
	"net/http"
	"strings"
	"time"

	"Caesar/pkg/utils"
)

type clientBuilder struct {
	//请求方式
	method string

	//请求内容
	body string

	//HTTP请求中的header信息
	header map[string]string

	//HTTP请求中,携带的cookies
	cookies []*http.Cookie

	//发起请求的client(go 自带的client)
	client *http.Client

	// 连接超时设置
	timeOut time.Duration

	//是否跳过HTTPS证书校验(默认跳过)
	skipVerify bool

	// 用户代理
	UserAgent string
}

// 设置请求方式
func (cb *clientBuilder) SetMethod(method string) *clientBuilder {
	cb.method = method

	supportMethod := []string{http.MethodGet, http.MethodPost, http.MethodConnect, http.MethodDelete, http.MethodHead, http.MethodOptions, http.MethodPatch, http.MethodPut, http.MethodTrace}
	if !utils.StringInSlice(method, supportMethod) {
		cb.method = http.MethodGet
	}

	return cb

}

func (cb *clientBuilder) SetBody(body string) *clientBuilder {
	cb.body = body
	return cb

}

func (cb *clientBuilder) SetSkipVerify(skip bool) *clientBuilder {
	cb.skipVerify = skip
	return cb

}

// 设置超时
func (cb *clientBuilder) SetTimeOut(t time.Duration) *clientBuilder {
	cb.timeOut = t
	return cb
}

// cookie设置
func (cb *clientBuilder) SetCookie(cookies string) *clientBuilder {

	if len(cookies) > 0 {
		if utils.MatchCookie(cookies) {
			nameList := strings.Split(cookies, "; ")

			var cookieSlice []*http.Cookie

			for _, c := range nameList {
				cookie := strings.Split(c, "=")
				newCookie := &http.Cookie{Name: cookie[0], Value: cookie[1]}
				cookieSlice = append(cookieSlice, newCookie)
			}
			cb.cookies = cookieSlice

		} else {
			cb.cookies = nil
		}
	} else {
		cb.cookies = nil
	}

	return cb
}

// user-agent设置
func (cb *clientBuilder) SetUserAgent(ua string) *clientBuilder {
	cb.UserAgent = ua
	return cb
}

// 设置http请求头文件
func (cb *clientBuilder) SetHeader(header map[string]string) *clientBuilder {

	cb.header = header
	return cb
}

func (cb *clientBuilder) FastBuilder() *fastClient {

	c := &fastClient{
		header:    cb.header,
		cookies:   cb.cookies,
		userAgent: cb.UserAgent,
		method:    cb.method,
		body:      cb.body,
	}
	return c
}

//NewClientBuilder 初始化
func NewClientBuilder() *clientBuilder {
	return &clientBuilder{
		skipVerify: true,
	}
}
