package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	//different ways of using structs
	// relying on property order
	alex := person{"Alex", "Anderson"}
	fmt.Println(alex)
	// identify the property
	james := person{firstName: "James", lastName: "Jamesson"}
	fmt.Println(james)
	// updating structs values
	var john person
	john.firstName = "John"
	john.lastName = "Johnson"
	fmt.Println(john)
	fmt.Printf("%+v", john)
}
