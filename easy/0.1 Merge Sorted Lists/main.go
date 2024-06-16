package main

import "fmt"

func main() {

	list1 := []int{1, 5, 7}
	list2 := []int{2, 4, 8}

	fmt.Println(mergeSortedLists(list1, list2))
}

func mergeSortedLists(list1 []int, list2 []int) []int {
	var mergedList []int
	i := 0
	j := 0

	for i < len(list1) && j < len(list2) {
		if list1[i] < list2[j] {
			mergedList = append(mergedList, list1[i])
			i++
		} else {
			mergedList = append(mergedList, list2[j])
			j++
		}
	}

	for i < len(list1) {
		mergedList = append(mergedList, list1[i])
		i++
	}

	for j < len(list2) {
		mergedList = append(mergedList, list2[j])
		j++
	}

	return mergedList
}
