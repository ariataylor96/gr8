package sys

import "gr8/font"

func (c *Chip8) LoadFont() {
	for idx, val := range font.FONTSET {
		c.Memory[font.START_ADDRESS+idx] = val
	}
}
