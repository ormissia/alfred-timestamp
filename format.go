package main

import (
	"time"
)

func Format(parserType string, t time.Time) []string {
	res := make([]string, 0)
	formatToString := (&FormatToString{Time: t}).Format()
	formatToTimestamp := (&FormatToTimestamp{Time: t}).Format()
	switch parserType {
	case ParserTypeNow:
		res = append(formatToTimestamp, formatToTimestamp...)
		res = append(formatToTimestamp, formatToString...)
	case ParserTypeTime:
		res = append(formatToTimestamp, formatToString...)
	case ParserTypeTimestamp:
		res = append(formatToString, formatToTimestamp...)
	default:
		res = append(formatToString, formatToTimestamp...)
	}
	return res
}

type FormatTime interface {
	Format() []string
}
