package main

import "os"

// TODO 增加错误日志
// TODO 增加函数

func main() {
	if len(os.Args) > 1 {
		Run(os.Args[1])
	}
	// Run("2022-07-19 16:27:19 + 1h")
}
