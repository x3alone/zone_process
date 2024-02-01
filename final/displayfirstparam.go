package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		println(os.Args[1])
	} else {
		fmt.Println()
	}
}
