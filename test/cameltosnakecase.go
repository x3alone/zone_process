package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	for i := 0; i < len(s); i++ {
		if !(s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z') || s[i] >= '0' && s[i] <= '9' {
			return s
		}
		if (s[i] >= 'A' && s[i] <= 'Z' && s[i+1] >= 'A' && s[i+1] <= 'Z') || i == len(s)-1 && s[i] >= 'A' && s[i] <= 'Z' {
			return s
		}
	}
	res := ""
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'A' && s[i] <= 'Z') && i != 0 {
			res += "_" + string(s[i])
		} else {
			res += string(s[i])
		}
	}
	return res
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
