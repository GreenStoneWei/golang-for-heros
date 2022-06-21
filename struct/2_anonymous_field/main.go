package main

import "fmt"

type person struct {
	name   string
	age    int
	weight int
}

type student struct {
	person     // anonymous field, or embedded field, student contains all the field of person
	name       string
	speciality string
}

func main() {
	mark := student{person{"Mark", 25, 120}, "act", "Computer Science"}

	// could access the field of anonymous field
	fmt.Println("Student age:", mark.age)
	fmt.Println("Student weight:", mark.weight)
	fmt.Println("Student speciality:", mark.speciality)
	fmt.Println("===")
	fmt.Println("Student name", mark.name)
	fmt.Println("Student person name", mark.person.name)
	fmt.Println("===")

	// could directly modify the anonymous field
	mark.speciality = "AI"
	mark.age = 46
	fmt.Printf("mark %+v", mark)

	// custom type and built-in type can be used in anonymous field
	type superpowers []string
	type hero struct {
		name string
		superpowers
	}
	batman := hero{name: "batman"}
	batman.superpowers = append(batman.superpowers, "rich")
	fmt.Printf("batman %+v", batman)
}
