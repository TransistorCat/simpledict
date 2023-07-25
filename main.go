package main

import (
	"fmt"
	"main/dict"
	"os"
	"sync"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
	example: simpleDict hello
			`)
		os.Exit(1)
	}
	word := os.Args[1]

	var wg sync.WaitGroup
	wg.Add(2)

	var caiyunResult, tencentResult string
	go func() {
		caiyunResult = dict.CaiyunQuery(word)
		wg.Done()
	}()

	go func() {
		tencentResult = dict.TencentQuery(word)
		wg.Done()
	}()

	wg.Wait()

	result := caiyunResult + tencentResult
	fmt.Print(result)
}
