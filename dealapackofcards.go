package piscine

import "fmt"

func arrayLength(arr []int) int {
	length := 0
	for range arr {
		length++
	}

	return length
}

func DealAPackOfCards(deck []int) {
	c := arrayLength(deck)
	cardsPerPlayer := c / 4

	for i := 0; i < 4; i++ {
		startIndex := i * cardsPerPlayer
		fmt.Printf("Player %d: %d, %d, %d\n", i+1, deck[startIndex], deck[startIndex+1], deck[startIndex+2])
	}
}
