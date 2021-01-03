package engine

import (
	"net/http"
	"time"

	"Caesar/internal/relation"
)

/*
用来保存扫描后的结构体
*/
type application struct {
	Store   []relation.StorePath
	Results []relation.ResultPtah // result.json保存
}

// 引擎参数
type ServerOpt struct {
	Paths   []relation.TagPath
	Threads int
	WAF     bool // 是否存在WAF
}

// 不存在页面的http回响信息
type ResponseInfo struct {
	Header http.Header
	Body   []byte
}

type RequestInfo struct {
	Address   string
	Method    string
	Header    map[string]string
	Cookies   string
	Body      string
	Proxy     string
	UserAgent []string
	Timeout   time.Duration
}
