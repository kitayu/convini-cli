package main

import (
	"flag"
	"fmt"
)

func main() {

	// コマンド定義
	var v bool
	flag.BoolVar(&v, "v", false, "convini-cli version")

	flag.Parse()

	if v {
		fmt.Println("convini-cli version 1.0.0")
	}
}
