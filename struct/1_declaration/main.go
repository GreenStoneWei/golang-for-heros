package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// method 1: declare first then assign value
	var p person
	p.name = "deliveryhero"
	p.age = 11
	fmt.Println(p)

	// method 2: assign value based on input order
	p1 := person{"deliveryhero", 11}
	fmt.Println(p1)

	// method 3: assign value by field
	p2 := person{age: 11, name: "deliveryhero"}
	fmt.Println(p2)

	// method 4: use new() to allocate a pointer
	p4 := new(person) // type of p4 is *person
	p4.name = "deliveryhero"
	p4.age = 11

	// method 5: a pointer to the struct
	p5 := &person{ // type of p5 is *person
		name: "deliveryhero",
		age:  11,
	}

	// Omitted fields will be zero-valued.
	omitP := person{name: "deliveryhero"}
	fmt.Println(omitP)

	// comparison: if they are of the same type and contain the same fields values
	fmt.Println(p1 == p2) // true
	fmt.Println(p4 == p5) // false
}
