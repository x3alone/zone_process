package mainzz

import (
	"fmt"
)

func RectPerimeter(w, h int) int {
	if w < 0 || h < 0 {
		return -1
	}
	pre := 2 * (w + h)
	return pre
}

func main() {
	fmt.Println(RectPerimeter(10, 2))
	fmt.Println(RectPerimeter(434343, 898989))
	fmt.Println(RectPerimeter(10, -2))
}
