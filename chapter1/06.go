package main

import (
	"fmt"
)

func getCharNgram(text string, n int) []string {
	sliced_text := []rune(text)
	ans := make([]string, 0, len(sliced_text)-n+1)

	for i := 0; i < cap(ans); i++ {
		ans = append(ans, string(sliced_text[i:i+n]))
	}
	return ans
}

func set2slice(set map[string]struct{}) []string {
	res := make([]string, 0, len(set))
	for id := range set {
		res = append(res, id)
	}
	return res
}

func slice2set(str []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, t := range str {
		set[t] = struct{}{}
	}
	return set
}

func getSet(text_list []string) []string {
	set := slice2set(text_list)
	return set2slice(set)
}

func join(set_X []string, set_Y []string) []string {
	set := slice2set(append(set_X, set_Y...))
	return set2slice(set)
}

func product(set_X []string, set_Y []string) []string {
	ans := make([]string, 0, len(set_X))
	for _, x := range set_X {
		for _, y := range set_Y {
			if x == y {
				ans = append(ans, x)
				break
			}
		}
	}
	return ans
}

func difference(set_X []string, set_Y []string) []string {
	ans := make([]string, 0, len(set_X))
	for _, x := range set_X {
		flag := false
		for _, y := range set_Y {
			if x == y {
				flag = true
				break
			}
		}
		if !flag {
			ans = append(ans, x)
		}
	}
	return ans
}

func searchSet(set []string, key string) bool {
	for _, t := range set {
		if t == key {
			return true
		}
	}
	return false
}

func main() {
	text1 := "paraparaparadise"
	text2 := "paragraph"

	set_X := getSet(getCharNgram(text1, 2))
	set_Y := getSet(getCharNgram(text2, 2))

	fmt.Println("set_X:  ", set_X)
	fmt.Println("set_Y:  ", set_Y)

	fmt.Println("X∪Y:    ", join(set_X, set_Y))
	fmt.Println("X∩Y:    ", product(set_X, set_Y))

	fmt.Println("X-Y:    ", difference(set_X, set_Y))
	fmt.Println("Y-X:    ", difference(set_Y, set_X))

	fmt.Println("se in X:", searchSet(set_X, "se"))
	fmt.Println("se in Y:", searchSet(set_Y, "se"))
}
