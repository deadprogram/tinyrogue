package tinyrogue

import (
	"crypto/rand"
	"math/big"
)

// GetRandomInt returns an integer from 0 to the number - 1
func GetRandomInt(num int) int {
	x, _ := rand.Int(rand.Reader, big.NewInt(int64(num)))
	return int(x.Int64())
}

// GetDiceRoll returns an integer from 1 to the number
func GetDiceRoll(num int) int {
	if num < 1 {
		num = 1
	}
	x, _ := rand.Int(rand.Reader, big.NewInt(int64(num)))
	return int(x.Int64()) + 1
}

// Return a number between two numbers inclusive.
func GetRandomBetween(low int, high int) int {
	return GetDiceRoll(high-low) + high
}
