package main

import "fmt"

// Define a struct
type person struct {
	name string
	age  int
}

// Constructor function that returns a pointer to person
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// Create struct using position
	fmt.Println(person{"Bob", 20})

	// Create struct using field names
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields get zero values
	fmt.Println(person{name: "Fred"})

	// Create pointer to struct
	fmt.Println(&person{name: "Ann", age: 40})

	// Use constructor function
	fmt.Println(newPerson("Jon"))

	// Access struct fields
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// Struct pointer auto-dereference
	sp := &s
	fmt.Println(sp.age)

	// Modify struct field
	sp.age = 51
	fmt.Println(sp.age)

	// Anonymous struct
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
