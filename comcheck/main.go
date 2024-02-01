package main

import (
	"fmt"
	"os"
)

func main() {
	last := os.Args
	for i := range last {
		if last[i] == "galaxy" || last[i] == "01" || last[i] == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
