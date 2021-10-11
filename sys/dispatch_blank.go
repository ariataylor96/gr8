package sys

func Blank(c *Chip8) {
	var (
		discriminator byte = byte(c.Opcode & 0x00FF)
		f             func()
	)

	switch discriminator {
	case 0xE0:
		f = c.CLS
	case 0xEE:
		f = c.RET
	}

	if f != nil {
		f()
	}
}
