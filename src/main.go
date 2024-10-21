package main

import (
	"fmt"
	"os"
	"photon/src/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path")
		os.Exit(1)
	}

	fileName := os.Args[1]
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	baseParser := parser.BaseParser{FileContent: &fileContent}
	err = baseParser.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
