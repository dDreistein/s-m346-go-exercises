package main

import "fmt"

func main() {
	type student struct {
		firstName string
		lastName  string
	}

	type class struct {
		students []student
	}

	type Modules map[int]class
	modules := Modules{
		104: {
			students: []student{
				{firstName: "Max", lastName: "Mustermann"},
				{firstName: "Erika", lastName: "Musterfrau"},
			},
		},
		117: {
			students: []student{
				{firstName: "Hans", lastName: "Schmidt"},
				{firstName: "Anna", lastName: "Schneider"},
			},
		},
		346: {
			students: []student{
				{firstName: "Peter", lastName: "Fischer"},
				{firstName: "Laura", lastName: "Weber"},
			},
		},
	}

	fmt.Println(modules)
}
