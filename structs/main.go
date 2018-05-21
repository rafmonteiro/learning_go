package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%+v", p)

}

func (p person) updateMe(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@party.com",
			zipCode: 6400,
		},
	}
	jim.updateMe("Jimmy")
	jim.print()
}
