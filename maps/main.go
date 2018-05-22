package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	fmt.Println(colors)
	colours := make(map[string]string)
	// add key to the map
	colours["white"] = "#ffffff"
	fmt.Println(colours)
	// delete a key from a map
	delete(colors, "green")
	fmt.Println(colors)
}
