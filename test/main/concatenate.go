package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println()
		return
	}

	result := ""
	for _, arg := range args {
		result += arg
	}

	fmt.Println(result)
}
