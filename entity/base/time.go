package base

import (
	"database/sql/driver"
	"errors"
	"time"
)

const TimeFormat = "2006-01-02 15:04"


type Time time.Time

// 自定义 json 序列化

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) AsInt() int64 {

	 var w time.Time = time.Time(t)
	 return w.Unix()
}

func (t *Time) String() string {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b  )
	b = time.Time(*t).AppendFormat(b, TimeFormat)
	b = append(b  )
	return string(b)
}
func (t *Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(*t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t Time) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format( TimeFormat), nil
}

func (t *Time) Scan(v interface{}) error {
	//fmt.Println(v)
	switch vt := v.(type) {

	case string:
		//fmt.Println(vt)
		// 字符串转成 time.Time 类型
		tTime, _ := time.Parse("2006-01-02 15:04:05", vt)
		*t =  Time(tTime)
	case time.Time:
		*t = Time(vt)
	default:
		//fmt.Println("fuck = ",vt)
		return errors.New("类型处理错误")
	}
	return nil
}
