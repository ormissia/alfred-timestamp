package main

import (
	"errors"
	"strings"
	"time"
)

const (
	nowStr = "now"
)

var ErrNowParser = errors.New("is not a time")

type Now struct {
	Str string
}

func (c *Now) Parser() (time.Time, error) {
	if strings.ToLower(c.Str) == nowStr {
		return time.Now(), nil
	}
	return time.Now(), ErrNowParser
}
