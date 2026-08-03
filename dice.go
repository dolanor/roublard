package main

import (
	"math/rand"
)

// GetRandomBetween returns a number between the two numbers inclusive.
func GetRandomBetween(low, high int) int {
	return GetDiceRoll(high-low) + high
}

// GetDiceRoll returns an integer from 1 to the number
func GetDiceRoll(num int) int {
	return rand.Intn(num) + 1
}
