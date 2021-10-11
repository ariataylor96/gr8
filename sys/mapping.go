package sys

import (
	"gr8/util"
)

var (
	nnnCodes []byte = []byte{0x1, 0x2, 0xA, 0xB}
	xkkCodes []byte = []byte{0x3, 0x4, 0x6, 0x7, 0xC}
	xyCodes  []byte = []byte{0x5, 0x9, 0x8}
	xynCodes []byte = []byte{0xD}
	xCodes   []byte = []byte{0xE, 0xF}
)

type m struct {
	codes    *[]byte
	executor func(*Chip8)
}

func (me *m) Matches(code byte) bool {
	return util.InColl(code, me.codes)
}

var buckets = []m{
	m{&nnnCodes, NNN},
	m{&xkkCodes, XKK},
	m{&xyCodes, XY},
	m{&xynCodes, XYN},
	m{&xCodes, X},
}

func (c *Chip8) Execute() {
	code := c.Op()

	for _, bucket := range buckets {
		if bucket.Matches(code) {
			bucket.executor(c)
			return
		}
	}
}
