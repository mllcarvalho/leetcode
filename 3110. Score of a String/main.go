package main

import (
	"fmt"
	"math"
)

// score calculates the score of the string s
func score(s string) int {
	totalScore := 0
	for i := 1; i < len(s); i++ {
		diff := math.Abs(float64(s[i]) - float64(s[i-1]))
		totalScore += int(diff)
	}
	return totalScore
}

func main() {
	// Test the score function with an example string
	s := "zaz"
	fmt.Printf("The score of the string \"%s\" is: %d\n", s, score(s))
}
