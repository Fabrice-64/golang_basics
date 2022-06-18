package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Jones",
		age:       33,
		contact: contactInfo{
			email:   "alex@mail.com",
			zipCode: 67000,
		},
	}
	fmt.Println("Original ID:")
	alex.print()
	alex.updateName("Alexandre")
	fmt.Println("Change the firstname:")
	alex.print()

}

func (p person) print() {
	fmt.Printf("Person data: %v\n", p)
}

func (pointedPers *person) updateName(newFirstName string) {
	(*pointedPers).firstName = newFirstName
}
