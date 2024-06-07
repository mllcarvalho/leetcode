package main

import "fmt"

func main() {

	chars := []string{"A", "B", "C", "A"}
	fmt.Println(VerifyFirstDuplicated(chars))
}

func VerifyFirstDuplicated(chars []string) string {

	m := make(map[string]bool)

	for _, j := range chars {

		fmt.Println(m[j])

		if m[j] == true {
			return j
		}

		m[j] = true
	}

	return ""
}
