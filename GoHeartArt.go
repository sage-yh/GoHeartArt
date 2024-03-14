package main

import "fmt"

func main() {
	printBoxWithRedHeartAndText("I am looking for you")
}

func printBoxWithRedHeartAndText(text string) {
	// Print top of the box
	fmt.Println("+++++++++++++++++++++++++++++")

	// Print sides of the box
	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Println("|       ❤️                   |")
		} else if i == 3 {
			fmt.Println("|", text, "|")
		} else {
			fmt.Println("|                             |")
		}
	}

	// Print bottom of the box
	fmt.Println("+++++++++++++++++++++++++++++")
}
