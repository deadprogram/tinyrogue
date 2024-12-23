package tinyrogue

import (
	"github.com/firefly-zero/firefly-go/firefly"
)

// GetRandomInt returns an integer from 0 to the number - 1
func GetRandomInt(num int) int {
	return int(firefly.GetRandom()) % (num - 1)
}

// GetDiceRoll returns an integer from 1 to the number
func GetDiceRoll(num int) int {
	if num < 1 {
		num = 1
	}
	x := int(firefly.GetRandom()) % num
	return x + 1
}

// Return a number between two numbers inclusive.
func GetRandomBetween(low int, high int) int {
	return GetDiceRoll(high-low) + high
}
