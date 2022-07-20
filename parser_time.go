package main

import (
	"errors"
	"time"
)

var ErrTimeParser = errors.New("is not a time")

type Time struct {
	Str string
}

func (t *Time) Parser() (time.Time, error) {
	for _, layout := range TimeLayouts {
		parse, err := time.Parse(layout, t.Str)
		if err != nil {
			continue
		}
		return parse, nil
	}
	return time.Now(), ErrTimeParser
}
