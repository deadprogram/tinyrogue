//go:build tinygo

package tinyrogue

import (
	"github.com/firefly-zero/firefly-go/firefly"
)

func getRandom() uint32 {
	return firefly.GetRandom()
}

func logError(msg string) {
	firefly.LogError(msg)
}

func logDebug(msg string) {
	firefly.LogDebug(msg)
}
