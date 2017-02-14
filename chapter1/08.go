package main

import (
	"flag"
	"fmt"
)

var text string

func cipher(text string) string {
	char_list := []rune(text)
	ans := []rune{}
	for _, char := range char_list {
		if char >= 'a' && char <= 'z' {
			char = rune(219 - int(char))
		}
		ans = append(ans, char)
	}
	return string(ans)
}

func main() {
	flag.StringVar(&text, "t", "", "input text")
	flag.Parse()

	fmt.Println(cipher(text))
}
