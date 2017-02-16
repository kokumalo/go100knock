package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func shuffle(word string) string {
	if len(word) < 4 {
		return word
	}
	sliced := []rune(word)
	rand.Seed(time.Now().UnixNano())
	for i := len(sliced) - 2; i >= 1; i-- {
		j := rand.Intn(i) + 1
		sliced[i], sliced[j] = sliced[j], sliced[i]
	}
	return string(sliced)

}

func main() {
	text := os.Args[1]

	ans := []string{}
	for _, word := range strings.Fields(text) {
		shuffled := shuffle(word)
		ans = append(ans, shuffled)
	}
	fmt.Println(strings.Join(ans, " "))

}
