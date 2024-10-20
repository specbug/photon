package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path")
		os.Exit(1)
	}

	fileName := os.Args[1]

	// Read the file
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file", err)
		os.Exit(1)
	}

	fmt.Println("File content:", string(content))
}
