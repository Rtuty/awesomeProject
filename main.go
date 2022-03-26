package main

import (
	"fmt"
	"sort"
)

/* func Contains(number_array []int, 2 int) bool {
	for _, n := range number_array {
		if 2 == n {
			return true
		}
	}
	return false
} */

func main() {
	var number_array = []string{"1", "2", "3"}
	fmt.Println(sort.SearchStrings(number_array, "3"))
}
