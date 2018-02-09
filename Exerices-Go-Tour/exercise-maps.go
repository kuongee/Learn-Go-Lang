package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	field := strings.Fields(s)
	m := make(map[string]int)
	for _, str := range field {
		if v, ok := m[str]; ok == false {
			m[str] = 1
		} else { // else grammar check
			m[str] = v + 1
		}
	
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

