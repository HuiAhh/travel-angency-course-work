package exception

import "errors"

//空指针异常
func Npe(msg *string) error {
	if msg == nil {
		return errors.New("空指针异常")
	}
	return errors.New(*msg)
}


//输入时间格式有误
func TimeIsError(msg *string) error {
	if (msg == nil) {
		return errors.New("输入时间格式有误")
	}
	return errors.New(*msg)
}