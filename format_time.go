package main

import "time"

type FormatToString struct {
	Time time.Time
}

func (c *FormatToString) Format() []string {
	res := make([]string, 0, len(TimeLayouts))
	for _, layout := range TimeLayouts {
		timeStr := c.Time.Format(layout)
		res = append(res, timeStr)
	}
	return res
}
