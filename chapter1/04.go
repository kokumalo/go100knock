package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	//twos := []int{0, 4, 5, 6, 7, 8, 14, 15, 18}

	re := regexp.MustCompile("[a-zA-Z]+")
	text_list := re.FindAllString(text, -1)

	ans := make([]string, 0, len(text_list))

	for i, t := range text_list {
		switch i {
		case 0, 4, 5, 6, 7, 8, 14, 15, 18:
			ans = append(ans, string([]rune(t)[0]))
		default:
			ans = append(ans, string([]rune(t)[0:2]))
		}
	}
	fmt.Println(ans)
}
