package main

import (
	"fmt"
)

// score calculates the score of the string s
func sumOfMultiples(number int) int {
	var sum int
	for i := 1; i <= number; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	number := 9
	fmt.Println(sumOfMultiples(number))
}
