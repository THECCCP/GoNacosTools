package lib

import "strings"

func FormatURL(url string) string {
	if strings.HasSuffix(url, "/") {
		return strings.TrimSuffix(url, "/")
	}
	return url
}
