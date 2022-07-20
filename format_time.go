package main

import "time"

type FormatToString struct {
	Time time.Time
}

func (t *FormatToString) Format() []string {
	res := make([]string, 0, len(TimeLayouts))
	for _, layout := range TimeLayouts {
		timeStr := t.Time.Format(layout)
		res = append(res, timeStr)
	}
	return res
}
