package main

import "fmt"

func main() {
	mySlice := []string{"sa", "naber", "cocuklar", "nasil", "hee"}

	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "ben kacar"
}
