package lib

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/http"
)

func AddNacosUser(url string) {
	user := fmt.Sprintf("test%s", generateRandomString(5))
	pass := fmt.Sprintf("Test@%s", generateRandomString(5))
	spl := "v1/auth/users?username=" + user + "&password=" + pass
	urls := url + spl
	userAgent := "Nacos-Server"

	// 创建一个自定义的http.Client，忽略SSL错误
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// 创建POST请求
	req, err := http.NewRequest("POST", urls, nil)
	if err != nil {
		fmt.Println("用户添加失败！:", err)
	}

	// 设置请求头
	req.Header.Set("User-Agent", userAgent)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("用户添加失败！", err)

	}
	defer resp.Body.Close()

	// 打印响应结果
	//fmt.Println("响应状态码:", resp.StatusCode)
	re := "用户添加成功！\n用户名：" + user + "\n密码：" + pass
	fmt.Println(re)
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomBytes := make([]byte, length)
	rand.Read(randomBytes)
	for i := 0; i < length; i++ {
		randomBytes[i] = charset[randomBytes[i]%byte(len(charset))]
	}
	return string(randomBytes)
}
