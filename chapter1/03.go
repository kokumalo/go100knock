package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."

	re := regexp.MustCompile("[a-zA-z]+")
	text_list := re.FindAllString(text, -1)
	ans := make([]int, 0, len(text_list))

	for _, t := range text_list {
		ans = append(ans, len(t))
	}

	fmt.Println(ans)
}
