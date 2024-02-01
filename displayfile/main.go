package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
		os.Exit(0)
	}
	if len(os.Args) != 2 {
		fmt.Println("Too many arguments")
		os.Exit(0)
	}
	fileName := os.Args[1]
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file %s: %s\n", fileName, err)
		os.Exit(1)
	}
	fmt.Printf("%s", content)
}
