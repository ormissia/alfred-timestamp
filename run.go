package main

import (
	"strings"
)

const (
	add = " + "
	sub = " - "
)

func Run(input string) {
	if strings.Contains(input, add) &&
		strings.Contains(input, sub) {
		return
	}
	if !strings.Contains(input, add) &&
		!strings.Contains(input, sub) {
		Transform(strings.TrimSpace(input))
	}

	args := make([]string, 0)
	calType := ""
	if strings.Contains(input, add) {
		args = strings.Split(input, add)
		calType = add
	} else if strings.Contains(input, sub) {
		args = strings.Split(input, sub)
		calType = sub
	}
	if len(args) != 2 {
		return
	}
	for i, _ := range args {
		args[i] = strings.TrimSpace(args[i])
	}
	if args[1] == "" {
		Transform(args[0])
	}
	Calculate(args[0], calType, args[1])

	// inputTime := time.Now()
	// inputType := ParserTypeTime
	//
	// 	it, t, err := Parser(os.Args[1])
	// 	if err == nil {
	// 		inputTime = t
	// 		inputType = it
	// 	}
	//
	// bytes, _ := json.Marshal(output)
	// fmt.Println(string(bytes))

	// item := Item{
	// 	Uid:          "uid" + os.Args[1] + "1",
	// 	Type:         "type" + os.Args[1] + "1",
	// 	Title:        "title" + os.Args[1] + "1",
	// 	Subtitle:     "subtitle" + os.Args[1] + "1",
	// 	Arg:          "arg" + os.Args[1] + "1",
	// 	Autocomplete: "autocomplete" + os.Args[1] + "1",
	// }
	// output.Items = append(output.Items, item)
}

func Transform(inputTime string) {
	parserType, parserTime, err := Parser(inputTime)
	if err != nil {
		return
	}
	OutPut(Format(parserType, parserTime))
}

func Calculate(inputTime, addSub, duration string) {
	parserType, parserTime, err := Parser(inputTime)
	if err != nil {
		return
	}
	du := DurationTime{Str: strings.TrimSpace(addSub) + duration}
	durationTime, err := du.Parser()
	parserTime = parserTime.Add(durationTime)
	OutPut(Format(parserType, parserTime))
}
