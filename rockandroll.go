package piscine

func RockAndRoll(n int) string {
	if n < 0 {
		return "error: number is negative\n"
	}

	divisibleBy2 := n%2 == 0
	divisibleBy3 := n%3 == 0

	if divisibleBy2 && divisibleBy3 {
		return "rock and roll\n"
	} else if divisibleBy2 {
		return "rock\n"
	} else if divisibleBy3 {
		return "roll\n"
	}

	return "error: non divisible\n"
}
