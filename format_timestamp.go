package main

import (
	"strconv"
	"time"
)

type FormatToTimestamp struct {
	Time time.Time
}

func (c *FormatToTimestamp) Format() []string {
	res := make([]string, 0, 4)
	res = append(res, strconv.Itoa(int(c.Time.Unix())))
	res = append(res, strconv.Itoa(int(c.Time.UnixMilli())))
	res = append(res, strconv.Itoa(int(c.Time.UnixMicro())))
	res = append(res, strconv.Itoa(int(c.Time.UnixNano())))
	return res
}
