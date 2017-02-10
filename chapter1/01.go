package main

import (
	"fmt"
)

func main() {
	text := "パタトクカシーー"
	text_slice := []rune(text)
	ans := make([]rune, 0, len(text))
	ans = append(ans, text_slice[0])
	ans = append(ans, text_slice[2])
	ans = append(ans, text_slice[4])
	ans = append(ans, text_slice[6])
	fmt.Println(string(ans))
}
