package main

import (
	"GoNacosTools/lib"
	"flag"
	"fmt"
)

func main() {
	url := flag.String("u", "http://127.0.0.1:8848", "使用 -u 指定URL")
	method := flag.String("m", "", "使用 -m 指定方法")
	flag.Parse()
	res := lib.GetNacosUrl(*url)
	switch *method {
	case "":
		fmt.Println("请使用 -m 参数指定方法\n-m getuser 获取用户名密码\n-m adduser 添加用户")
	case "getuser":
		lib.GetNacosUser(res)
	case "adduser":
		lib.AddNacosUser(res)
	}
}
