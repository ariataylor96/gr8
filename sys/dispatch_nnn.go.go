package sys

func NNN(c *Chip8) {
	var (
		addr uint16 = c.Opcode & 0x0FFF
		f    func(uint16)
	)

	switch code := c.Op(); code {
	case 0x1:
		f = c.JP
	case 0x2:
		f = c.CALL
	case 0xA:
		f = c.LDI
	case 0xB:
		f = c.JPV0
	}

	f(addr)
}
