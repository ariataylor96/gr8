package sys

// 00E0: Clear the display
func (c *Chip8) CLS() {
	c.Video = [64 * 32]uint32{}
}

// 00EE: Return from a subroutine
func (c *Chip8) RET() {
	c.SP--
	c.PC = c.Stack[c.SP]
}

// 1nnn: Jump to location addr
func (c *Chip8) JP(addr uint16) {
	c.PC = addr
}

// 2nnn: Call subroutine at addr
func (c *Chip8) CALL(addr uint16) {
	c.Stack[c.SP] = c.PC
	c.SP++
	c.PC = addr
}

// 3xkk: Skip next instruction if register == val
func (c *Chip8) SE(register, val uint8) {
	if c.Registers[register] == val {
		c.Next()
	}
}

// 4xkk: Skip next instruction if register != val
func (c *Chip8) SNE(register, val uint8) {
	if c.Registers[register] != val {
		c.Next()
	}
}

// 5xy0: Skip next instruction if register[x] == register[y]
func (c *Chip8) RSE(x, y uint8) {
	if c.Registers[x] == c.Registers[y] {
		c.Next()
	}
}

// 6xkk: Set register[x] = val
func (c *Chip8) LD(register, val uint8) {
	c.Registers[register] = val
}
