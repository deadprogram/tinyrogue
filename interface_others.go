//go:build !tinygo

package tinyrogue

func getRandom() uint32 {
	return 0
}

func logError(msg string) {
	println(msg)
}

func logDebug(msg string) {
	println(msg)
}
