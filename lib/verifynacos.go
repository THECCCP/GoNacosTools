package lib

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const faviconURL = "console-ui/public/img/favicon.ico"

func CheckNacos(url string) bool {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("创建请求失败: %v", err)
		return false
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534+ (KHTML, like Gecko) BingPreview/1.0b")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("请求失败: %v", err)
		return false
	}
	//defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("读取响应失败: %v", err)
		return false
	}
	if strings.Contains(string(body), faviconURL) {
		return true
	} else {
		return false
	}
}
