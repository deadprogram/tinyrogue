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
	switch num {
	case 0:
		num = 1
	default:
		num -= 1
	}
	x := int(firefly.GetRandom()) % num
	return 1 + x
}

// Return a number between two numbers inclusive.
func GetRandomBetween(low int, high int) int {
	return low + int(firefly.GetRandom())%(high-low)
}
