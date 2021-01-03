package main

import (
	"Caesar/cmd"
	"Caesar/internal/library/boot"
)

func init() {
	// 配置全局变量路径
	boot.SetPaths()
	// 配置基础参数
	boot.SetConf()

}

func main() {
	cmd.Execute()

}
