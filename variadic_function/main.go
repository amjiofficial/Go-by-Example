package main

import "fmt"

// Variadic function: accepts any number of integers
func sum(nums ...int) {
	fmt.Print(nums, " ")

	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {

	// Call with individual arguments
	sum(1, 2)
	sum(1, 2, 3)

	// Call using a slice
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
