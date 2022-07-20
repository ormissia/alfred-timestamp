package main

import (
	"errors"
	"reflect"
	"strings"
	"time"
)

const (
	ParserTypeTime      = "Time"
	ParserTypeTimestamp = "Timestamp"
)

var parserTypes = []string{
	ParserTypeTime,
	ParserTypeTimestamp,
}

var ErrStringParser = errors.New("string can not parser")

func Parser(str string) (string, time.Time, error) {
	for _, parser := range ParserFactory(str) {
		t, err := parser.Parser()
		if err != nil {
			continue
		}
		name := reflect.TypeOf(parser).String()
		split := strings.Split(name, ".")
		for _, parserType := range parserTypes {
			if parserType == split[len(split)-1] {
				return parserType, t, nil
			}
		}
	}
	return "", time.Now(), ErrStringParser
}

type ParserTime interface {
	Parser() (time.Time, error)
}

func ParserFactory(str string) []ParserTime {
	parsers := make([]ParserTime, 0)
	parsers = append(parsers,
		&Time{Str: str},
		&Timestamp{Str: str})
	return parsers
}

type ParserDuration interface {
	Parser() (time.Duration, error)
}
