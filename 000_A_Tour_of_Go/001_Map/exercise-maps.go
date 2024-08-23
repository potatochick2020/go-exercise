package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	returnValue := make(map[string]int)
	for _, v := range strings.Fields(s) {
		returnValue[v]++
	}
	return returnValue
}

func main() {
	wc.Test(WordCount)
}
