package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	quad_output := bufio.NewReader(os.Stdin)
	quad, _ := quad_output.ReadString('\x00')
	if quad == "" {
		fmt.Println("Not a quad function")
		return
	}
	if quad[len(quad)-1] != '\n'{
		fmt.Println("Not a quad function")
	return
	}
	x, y := 0, 0
	for _, i := range quad {
		if i == '\n' {
			y++
		}
		x++
	}
	x1 := x/y - 1
	x_arg := strconv.Itoa(x1)
	y_arg := strconv.Itoa(y)
	str_exe := []string{"./quadA", "./quadB", "./quadC", "./quadD", "./quadE"}
	matches := []string(nil)
	for _, exe_quad := range str_exe {
		exe_run := exec.Command(exe_quad, x_arg, y_arg)
		exe_output, _ := exe_run.Output()
		if string(exe_output) == quad {
			matches = append(matches, exe_quad[2:])
		}
	}
	if matches == nil {
		fmt.Println("Not a quad function")
		return
	}
	for i, match := range matches {
		fmt.Printf("[%s] [%s] [%s]", match, x_arg, y_arg)
		if i < len(matches)-1 {
			fmt.Printf(" || ")
		}
	}
	fmt.Printf("\n")
}
