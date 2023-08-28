package main

import (
	"GoNacosTools/lib"
	"flag"
	"fmt"
)

func main() {
	url := flag.String("u", "", "使用 -u 指定URL")
	method := flag.String("m", "", "使用 -m 指定方法")
	flag.Parse()
	urls := lib.FormatURL(*url)
	res := lib.GetNacosUrl(urls)
	switch *method {
	case "":
		fmt.Println("请使用 -m 参数指定方法\n-m getuser 获取用户名密码\n-m adduser 添加用户")
	case "getuser":
		lib.GetNacosUser(res)
	case "adduser":
		lib.AddNacosUser(res)
	}
}
