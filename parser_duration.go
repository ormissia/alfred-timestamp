package main

import "time"

type DurationTime struct {
	Str string
}

func (t *DurationTime) Parser() (time.Duration, error) {
	duration, err := time.ParseDuration(t.Str)
	if err != nil {
		du, err := time.ParseDuration(t.Str + "ms")
		if err != nil {
			return 0, err
		}
		duration = du
	}
	return duration, nil
}
