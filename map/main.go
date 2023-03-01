package main

import "fmt"

func main() {

	// var colors map[string]string
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// 	"blue":  "#0000ff",
	// }

	colors := make(map[string]string)
	colors["white"] = "#ffffff"
	colors["black"] = "#000000"
	delete(colors, "white")

	fmt.Println(colors)
}
