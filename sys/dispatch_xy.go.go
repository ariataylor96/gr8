package sys

func XY8(c *Chip8, x, y byte) {
	var f func(byte, byte)

	switch discriminator := byte(c.Opcode & 0x000F); discriminator {
	case 0x0:
		f = c.LDR
	case 0x1:
		f = c.ORR
	case 0x2:
		f = c.ANDR
	case 0x3:
		f = c.XORR
	case 0x4:
		f = c.ADDCR
	case 0x5:
		f = c.SUBCR

	case 0x7:
		f = c.SUBNCR

	case 0x6:
		c.SHR(x)
		return
	case 0xE:
		c.SHL(x)
		return
	}

	f(x, y)
}

func XY(c *Chip8) {
	var (
		x byte = byte((c.Opcode & 0x0F00) >> 8)
		y byte = byte((c.Opcode & 0x00F0) >> 4)
		f func(byte, byte)
	)

	switch code := c.Op(); code {
	case 0x5:
		f = c.RSE
	case 0x9:
		f = c.SNER
	case 0x8:
		XY8(c, x, y)
		return
	}

	f(x, y)
}
