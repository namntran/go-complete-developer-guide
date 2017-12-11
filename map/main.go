package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4b7000",
		"white": "fffffff",
	}
	printMap(colors)
}
func printMap(c map[string]string) { // c map is short syntax of golang for 'color' map
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
