package main

import (
	"fmt"
)

func main() {
	text1 := "パトカー"
	text2 := "タクシー"

	sliced_text1 := []rune(text1)
	sliced_text2 := []rune(text2)
	ans := make([]rune, 0, len(sliced_text1)+len(sliced_text2))

	for i := 0; i < len(sliced_text1); i++ {
		ans = append(ans, sliced_text1[i])
		ans = append(ans, sliced_text2[i])
	}

	fmt.Println(string(ans))
}
