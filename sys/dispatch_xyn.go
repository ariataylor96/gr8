package sys

func XYN(c *Chip8) {
	var (
		x      byte = byte((c.Opcode & 0x0F00) >> 8)
		y      byte = byte((c.Opcode & 0x00F0) >> 4)
		height byte = byte((c.Opcode & 0x000F))
	)

	c.DRW(x, y, height)
}
