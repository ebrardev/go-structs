package main

import "fmt"

func main() {

	// var colors map[string]string
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// 	"blue":  "#0000ff",
	// delete(colors, "white")
	//

	colors := map[string]string{

		"red":   "#ff0000",
		"green": "#4bf745",
		"blue":  "#0000ff",
	}
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
