package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	target := 4
	fmt.Println(twoSum(nums, target))

}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		j, ok := m[target-v]
		if ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}
