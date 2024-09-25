package main

import "fmt"

func IsCapitalized(s string) bool {
	if len(s) == 0{
		return false
	}
	in := false 
	for i := 0 ; i<len(s) ; i++{
		ch := s[i]
		if !in&& ch != ' '{
			if ch >= 'a'  && ch <= 'z'{
				return false
			}
			in = true
		}else if ch == ' '{
			in =false
		}
	}
	return true
}

func main() {
	// Test cases
	fmt.Println(IsCapitalized("Hello! How are you?")) // false (lowercase "are")
	fmt.Println(IsCapitalized("Hello How Are You"))   // true
	fmt.Println(IsCapitalized("Whats 4this 100K?"))   // true (starts with '4')
	fmt.Println(IsCapitalized("Whatsthis4"))          // true
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))      // true (non-alphabetic characters)
	fmt.Println(IsCapitalized(""))                    // false (empty string)
}
