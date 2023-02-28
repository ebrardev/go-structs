package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}
type person struct {
	firstName string
	lastName  string
	number    int
	contact   contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson", number: 1423456789}
	// ebrar := person{firstName: "Ebrar", lastName: "mea", number: 123456789}
	// fmt.Println(alex, ebrar)

	// 	var ebrar person
	// 	ebrar.firstName = "Ebrar"
	// 	ebrar.lastName = "mea"
	// 	ebrar.number = 123456789
	// 	fmt.Println(ebrar)
	// 	fmt.Printf("%+v", ebrar)
	//
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		number:    123456789,
		contact: contactInfo{
			email:   "slsdlsd",
			zipcode: 12345,
		},
	}
	fmt.Printf("%+v", jim)

}
