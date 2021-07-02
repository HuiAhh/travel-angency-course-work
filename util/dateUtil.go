package util

import "time"


func GetDayDiff(start *time.Time, end *time.Time) int {
	if start == nil || end == nil {
		return -1
	}
	return int((end.Unix() - start.Unix()) / 86400)
}
