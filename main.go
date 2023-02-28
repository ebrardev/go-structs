package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	number    int
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson", number: 1423456789}
	ebrar := person{firstName: "Ebrar", lastName: "mea", number: 123456789}
	fmt.Println(alex, ebrar)

}
