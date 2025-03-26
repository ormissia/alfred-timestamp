package main

import "time"

type DurationTime struct {
	Str string
}

func (c *DurationTime) Parser() (time.Duration, error) {
	duration, err := time.ParseDuration(c.Str)
	if err != nil {
		du, err := time.ParseDuration(c.Str + "ms")
		if err != nil {
			return 0, err
		}
		duration = du
	}
	return duration, nil
}
