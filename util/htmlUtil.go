package util

import "regexp"

//去除SCRIPT

var re, _ = regexp.Compile("<script.*?>(.*?)</script>")

func ReplaceHtml(html string) string {

	res := re.ReplaceAll([]byte(html), []byte(""))
	return string(res)
}
