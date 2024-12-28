package tinyrogue

// GetRandomInt returns an integer from 0 to the number - 1
func GetRandomInt(num int) int {
	return int(getRandom()) % (num - 1)
}

// GetDiceRoll returns an integer from 1 to the number
func GetDiceRoll(num int) int {
	switch {
	case num < 2:
		return int(getRandom() % uint32(2))
	}

	return int(getRandom()%uint32(num)) + 1
}

// Return a number between two numbers inclusive.
func GetRandomBetween(low int, high int) int {
	return low + int(getRandom())%(high-low)
}
