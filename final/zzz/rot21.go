package main

func rot21(b rune) rune {
	if b >= 'A' && b < 'E' || b >= 'a' && b < 'e' {
		return b + 21
	}
	if b >= 'E' && b <= 'Z' || b >= 'e' && b <= 'z' {
		return b - 5
	}
	return b
}

func Rot21(str string) string {
	result := ""
	for _, res := range str {
		result += string(rot21(res))
	}
	return result
}

func main() {
	println(('z'))
	println(('n'))
}
