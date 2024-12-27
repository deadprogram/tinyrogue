//go:build !tinygo

package tinyrogue

import "math/rand"

func getRandom() uint32 {
	return rand.Uint32()
}

func logError(msg string) {
	println(msg)
}

func logDebug(msg string) {
	println(msg)
}
