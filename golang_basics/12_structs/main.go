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
	alive     bool
	contact   contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		age:       53,
		alive:     true,
		contact: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 33333,
		},
	}
	fmt.Println(alex)
	var john person
	fmt.Println(john)
	fmt.Printf("%+v\n", alex)
	alex.firstName = "Alexis"
	alex.contact.zipCode = 10101
	fmt.Printf("%+v\n", alex)

	alex.print()
	fmt.Println("Failing to change first name:")
	alex.updateName("Fanchon")
	//New Firstname should not be taken into account: pass by value principle of go
	alex.print()

}

func (p person) print() {
	fmt.Printf("ID: %+v\n", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
