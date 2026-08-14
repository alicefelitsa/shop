package main

import "fmt"

//交叉编译
//go env -w CGO_ENABLED=0
//go env -w GOOS=linux
//go env -w GOOS=windows
//go env -w GOPROXY=https://goproxy.cn,direct
//Get-Content app.log -Wait win系统，实时读取日志

func main() {
	fmt.Println(123456)
}
