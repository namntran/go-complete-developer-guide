package main

import "fmt"

type contactInfo struct {
	email    string
	postCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Smith",
		contactInfo: contactInfo{
			email:    "jim@gmail.com",
			postCode: 4000,
		},
	}
	jimPointer := &jim
	jimPointer.updateName("Jimmy") // pointer
	jim.print()
}
func (pointerToPerson *person) updateName(newFirstName string) { //*person - this is a type description-it means we're working with a pointer to a person.
	(*pointerToPerson).firstName = newFirstName //*pointerToPerson - this is an actual operator - it means we want to manipulate the value the pointer is referencing. It takes this pointer and turns it into an actual value.
}
func (p person) print() {
	fmt.Printf("%+v", p)
}

// Turn address into value with *address
// Turn value into address with &value
