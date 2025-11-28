package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	Day   byte
	Month byte
	Year  int
}

type Profile struct {
	FullName
	BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		FullName: FullName{
			FirstName: "Bela",
			LastName:  "Müller",
		},
		BirthDate: BirthDate{
			Day:   2,
			Month: 6,
			Year:  2008,
		},
		NumberOfSiblings: 1,
		ZodiacSign:       '♊',
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
