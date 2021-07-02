package main

import (
	"fmt"
	"testing"
	"travel-agency/util"
)

//去掉 script标签

func TestReplaceScriptTag(t *testing.T) {
	res := util.ReplaceHtml("script  <script></script> <Script></Script>")
	fmt.Println(res)
}
