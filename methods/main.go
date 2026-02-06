package main

import "fmt"

// Define a struct
type rect struct {
	width  int
	height int
}

// Method with pointer receiver
func (r *rect) area() int {
	return r.width * r.height
}

// Method with value receiver
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {

	// Create a rect struct
	r := rect{width: 10, height: 5}

	// Call methods on struct value
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// Call methods using pointer
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
