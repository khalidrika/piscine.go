package main

import (
	"fmt"
	"os"
)

func main() {
	// arg := os.Args[1:]
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	} else if len(os.Args) == 1 {
		fmt.Println("File name missing")
		return
	} else {
		os.ReadFile(os.Args[1])
		whrite_on_bite, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Printf("the mistake is %v\n", err.Error())
		}
		fmt.Print(string(whrite_on_bite))
	}
}
