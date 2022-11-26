package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var jake person
	fmt.Println(jake)
	fmt.Printf("%+v\n", jake)

	jake.firstName = "Jake"
	jake.lastName = "Jones"
	fmt.Printf("%+v\n", jake)
}
