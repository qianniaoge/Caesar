package relation

//EachPath 对应的是路径字典的json格式.
type EachPath struct {
	Hits int64  `json:"hits"`
	Path string `json:"path"`
}

//TagPath 用来存储hits.
type TagPath struct {
	EachPath
	Tag string
}

// ResultPtah 获取到的http响应信息
type ResultPtah struct {
	Code    int
	Address string
	Title   string
	Length  int
}

// StorePath 为了hits保存的数据
type StorePath struct {
	TagPath
	ResultPtah
}
