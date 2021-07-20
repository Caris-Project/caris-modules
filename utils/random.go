package utils

import (
	"math/rand"
	"time"
)

// RandomIntBetweenRange ...
func RandomIntBetweenRange(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
