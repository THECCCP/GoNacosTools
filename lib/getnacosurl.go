package lib

func GetNacosUrl(url string) string {
	if CheckNacos(url) {
		return url
	} else {
		urls := url + "/nacos/"
		return urls
	}
}
