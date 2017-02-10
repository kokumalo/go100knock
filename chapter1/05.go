package main

import (
	"flag"
	"fmt"
	"regexp"
)

var (
	text   string
	word_n int
	char_n int
)

func getWordNgram(text string, n int) [][]string {
	re := regexp.MustCompile("[a-zA-Z]+")
	text_list := re.FindAllString(text, -1)
	ans := make([][]string, 0, len(text_list)-n+1)

	for i := 0; i < cap(ans); i++ {
		sub_ans := make([]string, 0, n)
		for j := 0; j < n; j++ {
			sub_ans = append(sub_ans, text_list[i+j])
		}
		ans = append(ans, sub_ans)
	}
	return ans
}

func getCharNgram(text string, n int) []string {
	sliced_text := []rune(text)
	ans := make([]string, 0, len(sliced_text)-n+1)

	for i := 0; i < cap(ans); i++ {
		ans = append(ans, string(sliced_text[i:i+n]))
	}
	return ans
}

func main() {
	flag.StringVar(&text, "t", "", "解析する文字列")
	flag.IntVar(&word_n, "w", 2, "単語n-gram")
	flag.IntVar(&char_n, "c", 2, "文字n-gram")

	flag.Parse()

	fmt.Println(getWordNgram(text, word_n))
	fmt.Println(getCharNgram(text, char_n))
}
