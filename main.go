package main

import "fmt"
import "os"

func main() {
	if os.Args.size() < 2 || os.Args.size() > 2 {
		fmt.Println("Error, wrong number of arguments,  Usage: babel textsource.txt")
		return
	}

	text_source := os.Args[1]
	file, err := os.Open(text_source)
}
