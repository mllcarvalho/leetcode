package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sortSentence(s string) string {
	var a []string
	a = strings.Split(s, " ")
	ans := make([]string, len(a))
	ans2 := ""
	for _, w := range a {
		i, _ := strconv.Atoi(string(w[len(w)-1]))
		ans[i-1] = w[:len(w)-1]
	}
	for _, word := range ans {
		ans2 += word + " "
	}
	return strings.TrimSpace(ans2)
}

func main() {
	text := "is2 sentence4 This1 a3"
	fmt.Println(sortSentence(text))
}
