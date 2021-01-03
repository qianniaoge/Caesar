# v1.0.1 2020/12/13
* 增加http包，将会用于-r参数解析构造http请求的函数。

# v1.0.2 2020/12/26
* net目录下增加fasthttp文件，用来保存fasthttp的请求代码。
粗略测试的结果，
标准库 GET请求12线程下每秒700并发
fasthttp GET请求12线程下每秒3000并发
基准数据
go test -bench=. -benchmem
fasthttp
```shell script
BenchmarkFastHttpRequest-12            1        16804582317 ns/op       736803176 B/op    536182 allocs/op
```

golang标准库
```shell script
BenchmarkStandRequest-12               1        469205373888 ns/op      1580837544 B/op  5465476 allocs/op
```

*在fasthttp的默认请求头状况下，能实现每秒2000级并发，但是如果额外添加http请求头，比如Connection，则速度变得和golang原生库性能差异不大*