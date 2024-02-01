package piscine

import (
	"github.com/01-edu/z01"
)

func PrintCombination(arr []rune, data []rune, start int, end int, index int) {
	if index == len(data) {
		for i := 0; i < len(data); i++ {
			z01.PrintRune(data[i])
		}
		var res bool = false
		max := '9'
		index := 0
		r := '0'
		for i := 0; i < len(data); i++ {
			r++
		}
		for k := '0'; k < r; k++ {
			if data[index] != max-((r-1)-k) {
				res = true
			}
			index++
		}
		if res {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}

	for i := start; i <= end && end-i+1 >= len(data)-index; i++ {
		data[index] = arr[i]
		PrintCombination(arr, data, i+1, end, index+1)
	}
}

func PrintCombN(digits int) {
	arr := make([]rune, 0)
	for i := '0'; i <= '9'; i++ {
		arr = append(arr, i)
	}

	data := make([]rune, digits)
	PrintCombination(arr, data, 0, len(arr)-1, 0)
	z01.PrintRune('\n')
}
