package main

import (
	"fmt"
	"strconv"
)

//Person struct define
type Person struct {
	FirstName, LastName string
	Age                 int
	Gender              bool //true : false :: male : female
}

//method (value reciever)
func (p Person) greet() string {
	return "Hello, I am " + p.FirstName + " " + p.LastName + ", I am " + strconv.Itoa(p.Age) + "years old. How you doin?"
}

//method pointer reciever
func (p *Person) hasBirthday() {
	p.Age++
}

func main() {
	person1 := Person{FirstName: "Ankush", LastName: "Malik", Age: 21, Gender: true}
	fmt.Println(person1)
	//Alternate method
	person2 := Person{"Foo", "Bar", 32, false}
	fmt.Println(person2)
	person2.hasBirthday()
	fmt.Println(person2.greet())
}
