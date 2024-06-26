package util

import (
	"Termbin/model"
	"path/filepath"
	"strings"
)

func GetContentType(url string) (string, error) {
	// 没有后缀的话返回空字符串
	extension := filepath.Ext(url)
	return model.MIMEType[extension], nil
}

func RemoveExt(filename string) (string, error) {
	base := filepath.Base(filename)
	extension := filepath.Ext(base)
	return base[0 : len(base)-len(extension)], nil
}

func IsBrowserUserAgent(userAgent string) bool {
	browsers := []string{"Mozilla", "Chrome", "Edg", "Safari", "Firefox", "Opera"}

	for _, browser := range browsers {
		if strings.Contains(userAgent, browser) {
			return true
		}
	}

	return false
}
