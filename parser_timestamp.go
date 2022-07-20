package main

import (
	"errors"
	"strconv"
	"time"
)

// 1658160037
// 1658160037585
// 1658160037585314
// 1658160037585314000
const (
	TimeStampLengthUnix      = 10
	TimeStampLengthUnixMilli = 13
	TimeStampLengthUnixMicro = 16
	TimeStampLengthUnixNano  = 19
)

var ErrTimestampParser = errors.New("is not a time")

type Timestamp struct {
	Str string
}

func (t *Timestamp) Parser() (time.Time, error) {
	timestampInt, err := strconv.Atoi(t.Str)
	if err != nil {
		return time.Now(), err
	}
	timestamp := int64(timestampInt)
	switch len(t.Str) {
	case TimeStampLengthUnix:
		return time.Unix(timestamp, 0), nil
	case TimeStampLengthUnixMilli:
		return time.UnixMilli(timestamp), nil
	case TimeStampLengthUnixMicro:
		return time.UnixMicro(timestamp), nil
	case TimeStampLengthUnixNano:
		sec := timestampInt / 1e9
		nsec := timestampInt % 1000000000
		return time.Unix(int64(sec), int64(nsec)), nil
	default:
		return time.Now(), ErrTimestampParser
	}
}
