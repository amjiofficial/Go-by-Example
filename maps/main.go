package main

import (
	"fmt"
	"maps"
)

func main() {

	// 1. Create an empty map
	m := make(map[string]int)

	// 2. Set key-value pairs
	m["k1"] = 7
	m["k2"] = 13

	// 3. Print the map
	fmt.Println("map:", m)

	// 4. Get value by key
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// 5. Get value for non-existing key (zero value)
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// 6. Length of map
	fmt.Println("len:", len(m))

	// 7. Delete a key
	delete(m, "k2")
	fmt.Println("map after delete:", m)

	// 8. Clear all keys
	clear(m)
	fmt.Println("map after clear:", m)

	// 9. Check if key exists
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 10. Declare and initialize map in one line
	n := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	fmt.Println("map n:", n)

	// 11. Compare two maps
	n2 := map[string]int{
		"foo": 1,
		"bar": 2,
	}

	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
