package director

import (
	"net/http"
	"time"

	"Caesar/internal/relation"
	"Caesar/pkg/builder/net/fasthttp"
	"Caesar/pkg/builder/net/stand"
	"Caesar/pkg/utils"
)

func GenerateGet(target string, random bool) (int, http.Header, []byte, error) {
	if random {
		build := stand.NewClientBuilder().SetUserAgent(utils.GetRandomElement(relation.Browser.UserAgent)).SetHeader(relation.Browser.Headers).SetCookie(relation.Browser.Cookie).SetProxy(relation.Browser.Proxy).SetTimeOut(relation.Browser.TimeOut).SetSkipVerify(true).StandBuilder()
		return build.Get(target)
	} else {
		build := stand.NewClientBuilder().SetUserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:76.0) Gecko/20100101 Firefox/76.0").SetHeader(relation.Browser.Headers).SetCookie(relation.Browser.Cookie).SetProxy(relation.Browser.Proxy).SetTimeOut(relation.Browser.TimeOut).SetSkipVerify(true).StandBuilder()
		return build.Get(target)
	}

}

func GenerateNormalGet(target string) (int, http.Header, []byte, error) {
	build := stand.NewClientBuilder().SetUserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:76.0) Gecko/20100101 Firefox/76.0").SetHeader(relation.Browser.Headers).SetCookie(relation.Browser.Cookie).SetTimeOut(relation.Browser.TimeOut).SetSkipVerify(true).StandBuilder()
	return build.Get(target)
}

func GenerateHttp(target string, httpMethod string, body map[string]string, headers map[string]string, cookies string, proxy string) (int, http.Header, []byte, error) {
	build := stand.NewClientBuilder().SetHeader(headers).SetCookie(cookies).SetProxy(proxy).SetTimeOut(relation.Browser.TimeOut).SetSkipVerify(true).StandBuilder()
	return build.Get(target)
}

func UnitTest(target string, httpMethod string, UserAgent string, header map[string]string, cookies string, proxy string, timeout time.Duration, body string) (int, http.Header, []byte, error) {
	build := stand.NewClientBuilder().SetMethod(httpMethod).SetUserAgent(UserAgent).SetHeader(header).SetCookie(cookies).SetProxy(proxy).SetTimeOut(timeout).SetBody(body).SetSkipVerify(true).StandBuilder()
	return build.RawHttp(target)
}

func FastHttpRequest(target string, httpMethod string, UserAgent string) (int, http.Header, []byte, error) {
	build := fasthttp.NewClientBuilder().SetMethod(httpMethod).SetUserAgent(UserAgent).SetSkipVerify(true).FastBuilder()
	return build.RawRequest(target)
}
