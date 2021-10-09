package sys

import (
	"math/rand"
	"time"
)

func seedRng() {
	rand.Seed(time.Now().UnixNano())
}

func (c *Chip8) RandByte() byte {
	return byte(rand.Intn(255))
}
