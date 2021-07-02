package util

import (
	"strconv"
	"strings"
)

func IsNotBlank(v string) bool {
	return v != "" && strings.TrimSpace(v) != ""
}

func IsBlank(v *string) bool {
	if v == nil {
		return true
	}
	return *v == "" || strings.TrimSpace(*v) == ""
}

func ToInt(v string) (int, error) {
	res, x := strconv.Atoi(v)
	return res, x
}
