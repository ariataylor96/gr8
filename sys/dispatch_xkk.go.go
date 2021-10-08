package sys

func XKK(c *Chip8) {
	var (
		x  uint8 = byte((c.Opcode & 0x0F00) >> 8)
		kk uint8 = byte(c.Opcode & 0x00FF)
		f  func(byte, byte)
	)

	switch code := c.Op(); code {
	case 0x3:
		f = c.SE
	case 0x4:
		f = c.SNE
	case 0x6:
		f = c.LD
	case 0x7:
		f = c.ADD
	case 0xC:
		f = c.RNDVX
	}

	f(x, kk)
}
