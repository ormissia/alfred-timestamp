package main

import (
	"strconv"
	"time"
)

type FormatToTimestamp struct {
	Time time.Time
}

func (t *FormatToTimestamp) Format() []string {
	res := make([]string, 0, 4)
	res = append(res, strconv.Itoa(int(t.Time.UnixMilli())))
	res = append(res, strconv.Itoa(int(t.Time.Unix())))
	res = append(res, strconv.Itoa(int(t.Time.UnixMicro())))
	res = append(res, strconv.Itoa(int(t.Time.UnixNano())))
	return res
}
