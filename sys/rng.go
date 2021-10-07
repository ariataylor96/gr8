package sys

import (
	"math/rand"
	"time"
)

func seedRng() {
	rand.Seed(time.Now().UnixNano())
}

func (c *Chip8) RandByte() uint8 {
	return uint8(rand.Intn(255))
}
