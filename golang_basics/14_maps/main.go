package main

import "fmt"

func main() {
	colors := map[string]string{
		"green":  "#7CFC00",
		"red":    "#FF0000",
		"yellow": "#FFFF00",
	}
	colors["white"] = "#FFFFFF"
	delete(colors, "yellow")
	fmt.Println(colors)
	fmt.Println("Iterating through the map:")
	printMap(colors)

}

func printMap(c map[string]string) {
	for k, h := range c {
		fmt.Printf("To Color: %v, Corresponds hex value: %v\n", k, h)
	}
}
