package sys

func X(c *Chip8) {
	var (
		x             byte = byte((c.Opcode & 0x0F00) >> 8)
		discriminator byte = byte(c.Opcode & 0x00FF)
		f             func(byte)
	)

	switch discriminator {
	case 0x9E:
		f = c.SKPVX
	case 0xA1:
		f = c.SKNPVX
	case 0x07:
		f = c.LDVXDT
	case 0x0A:
		f = c.LDVXK
	case 0x15:
		f = c.LDDTVX
	case 0x18:
		f = c.LDSTVX
	case 0x1E:
		f = c.ADDIVX
	case 0x29:
		f = c.LDFVX
	case 0x33:
		f = c.LDBVX
	case 0x55:
		f = c.LDIVX
	case 0x65:
		f = c.LDVXI
	}

	if f != nil {
		f(x)
	}
}
