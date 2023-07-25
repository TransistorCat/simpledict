package main

import (
	"fmt"
	"main/dict"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
	example: simpleDict hello
			`)
		os.Exit(1)
	}
	word := os.Args[1]
	// word := "hello"
	result := dict.CaiyunQuery(word) + dict.TencentQuery(word)
	fmt.Print(result)

}
