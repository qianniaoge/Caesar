package units

import (
	"Caesar/pkg/builder/net/stand"
	"github.com/valyala/fasthttp"
	"net/http"
	"testing"
)

func doRequest(url string) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	client.Do(req, resp)

	bodyBytes := resp.Body()
	println(string(bodyBytes))
	// User-Agent: fasthttp
	// Body:
}

func TestRequests(T *testing.T) {
	headers := map[string]string{
		"User-Agent": "Mozilla()",
	}

	cookies := "username=admin"
	//proxy := "http://127.0.0.1:8080"

	newRequest := stand.NewClientBuilder().SetMethod(http.MethodGet).SetHeader(headers).SetCookie(cookies).SetTimeOut(3).SetSkipVerify(true).StandBuilder()
	_, _, body, _ := newRequest.FastGet("http://27.211.65.98:8081/")
	println(string(body))
}
