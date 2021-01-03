# 发起请求的核心代码

## v1.0.1 2020/12/11
* 修改base.go,增加serverOpt用来存储引擎的参数:
    1. Paths 要遍历的path路径
    2. Threads 线程数
    
* 移除aftermath函数的结果导出代码，遵守单一职责原则。结果导出代码移到action.go的Export函数中
* force移动到 ../cores/action.go 中


## v1.0.2 2020/12/15
* 修改base.go,增加responseInfo用来存储响应请求的参数:
* 修改base.go,采用工厂模式


## v1.0.2 2020/12/19
* 修改base.go,ServerOpt添加WAF和MVC两个属性，用来传递WEB是否存在WAF或者是MVC架构