package piscine

func TrimAtoi(s string) int { // this func will take only the numbers from the string
	var runes []rune           // the array that takes the values we entered
	length := 0                // the start of the length
	for i, letter := range s { // the len of the string
		if letter >= '0' && letter <= '9' { // if it found a char between range of  0 > 9 it appends the array
			runes = append(runes, letter) // it appends the array and enter the new value
			length = i + 1                // here to move on and check the next char
		} else if letter == '-' && length == 0 { // if it was a sign in the start of the string it should print it
			runes = append(runes, '-') // prints it
		} else {
			continue // if else go on
		}
	}
	return atoi(runes)
}

func atoi(runes []rune) int {
	LenRune := 0
	result := 0
	for i := range runes {
		i++
		LenRune++
	}
	if LenRune == 0 {
		return 0
	}
	tens := 1
	for k := 0; k < LenRune-1; k++ {
		if runes[k] == '+' || runes[k] == '-' {
			continue
		}
		tens *= 10
	}
	for i := range runes {
		if (runes[0] == '-' || runes[0] == '+') && i == 0 {
			continue
		}
		if runes[i] < '0' || runes[i] > '9' {
			return 0
		}
		numb := 0
		for j := '0'; j < runes[i]; j++ {
			numb++
		}
		result += numb * tens
		tens /= 10
	}
	if runes[0] == '-' {
		result *= -1
	}
	return result
}
