package sys

import "gr8/font"

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

// 5xy0: Skip next instruction if Vx == Vy
func (c *Chip8) RSE(x, y uint8) {
	if c.Registers[x] == c.Registers[y] {
		c.Next()
	}
}

// 6xkk: Set Vx = val
func (c *Chip8) LD(register, val uint8) {
	c.Registers[register] = val
}

// 7xkk: Add Vx += val
func (c *Chip8) ADD(register, val uint8) {
	c.Registers[register] += val
}

// 8xy0: Set Vx = Vy
func (c *Chip8) LDR(x, y uint8) {
	c.Registers[x] = c.Registers[y]
}

// 8xy1: Set Vx |= Vy
func (c *Chip8) ORR(x, y uint8) {
	c.Registers[x] |= c.Registers[y]
}

// 8xy2: Set Vx &= Vy
func (c *Chip8) ANDR(x, y uint8) {
	c.Registers[x] &= c.Registers[y]
}

// 8xy3: Set Vx XOR Vy
func (c *Chip8) XORR(x, y uint8) {
	c.Registers[x] ^= c.Registers[y]
}

// 8xy4: Set Vx += Vy, and set VF = carry
// If the result is > 255 then VF is set to 1 as a flag
func (c *Chip8) ADDCR(x, y uint8) {
	Vx, Vy := c.Registers[x], c.Registers[y]
	sum := Vx + Vy

	// Go does not have a ternary operator
	if sum > 255 {
		c.Registers[0xF] = 1
	} else {
		c.Registers[0xF] = 0
	}

	// Only take the lower 8 bits
	c.Registers[x] = sum & 0xFF
}

// 8xy5: Set Vx -= Vy, and set VF = NOT borrow
// Put plainly, VF = Vx > Vy
func (c *Chip8) SUBCR(x, y uint8) {
	Vx, Vy := c.Registers[x], c.Registers[y]

	if Vx > Vy {
		c.Registers[0xF] = 1
	} else {
		c.Registers[0xF] = 0
	}

	c.Registers[x] -= Vy
}

// 8xy6: SHR Vx
// Shifts Vx to the right (divides it by two) and
// saves the least significant bit in VF
func (c *Chip8) SHR(register uint8) {
	// Save the least significant bit in VF
	c.Registers[0xF] = c.Registers[register] & 0x1
	c.Registers[register] >>= 1
}

// 8xy7: Set Vy -= Vx, and set VF = NOT borrow
func (c *Chip8) SUBNCR(x, y uint8) {
	c.SUBCR(y, x)
}

// 8xyE: SHL Vx
// Shifts Vx to the left (multiplies it by two)
// and saves the most significant bit in VF
func (c *Chip8) SHL(register uint8) {
	// Save the most significant bit in VF
	c.Registers[0xF] = c.Registers[register] & 0x80
	c.Registers[register] <<= 1
}

// 9xy0: Skip next instruction if Vx != Vy
func (c *Chip8) SNER(x, y uint8) {
	Vx, Vy := c.Registers[x], c.Registers[y]

	if Vx != Vy {
		c.Next()
	}
}

// Annn: LD I, addr
// Sets index to addr
func (c *Chip8) LDI(addr uint16) {
	c.Index = addr
}

// Bnnn: JP V0, addr
// Jump to location stored in V0 + addr
func (c *Chip8) JV0(addr uint16) {
	c.PC = uint16(c.Registers[0]) + addr
}

// Cxkk - RND Vx, byte
// Sets Vx = (random byte) & val
func (c *Chip8) RNDVX(register, val uint8) {
	c.Registers[register] = c.RandByte() & val
}

// Dxyn - DRW Vx, Vy, nibble
// Display n-byte sprite starting at memory location I
// at coordinates (Vx, Vy), set VF = collision
func (c *Chip8) DRW(x, y, height uint8) {
	Vx, Vy := c.Registers[x], c.Registers[y]
	xPos, yPos := Vx%VIDEO_WIDTH, Vy%VIDEO_HEIGHT

	c.Registers[0xF] = 0

	for row := 0; row < int(height); row++ {
		spriteByte := c.Memory[c.Index+uint16(row)]

		for col := 0; col < 8; col++ {
			spritePixel := spriteByte & (0x80 >> col)
			screenPixel := &(c.Video[(int(yPos)+row)*int(VIDEO_WIDTH)+(int(xPos)+col)])

			// Sprite pixel is on
			if spritePixel != 0 {
				// Pixel already on the screen is on, meaning there's a collision
				if *screenPixel == 0xFFFFFFFF {
					c.Registers[0xF] = 1
				}

				// XOR with the existence of the sprite pixel instead of the status
				*screenPixel ^= 0xFFFFFFFF
			}
		}
	}
}

// Ex9E: SKP Vx
// Skip next instruction if key with value of Vx is pressed
func (c *Chip8) SKPVX(x uint8) {
	Vx := c.Registers[x]

	if c.Keypad[Vx] != 0 {
		c.Next()
	}
}

// ExA1: SKNP Vx
// Skip next instruction if key with value of Vx is not pressed
func (c *Chip8) SKNPVX(x uint8) {
	Vx := c.Registers[x]

	if c.Keypad[Vx] == 0 {
		c.Next()
	}
}

// Fx07 - LD Vx, DT
// Set Vx = delay timer value
func (c *Chip8) LDVXDT(x uint8) {
	c.Registers[x] = c.DelayTimer
}

// Fx0A - LD Vx, K
// Wait for a key press, store the value of the key in Vx
func (c *Chip8) LDVXK(x uint8) {
	for i := 0; i < 16; i++ {
		if c.Keypad[i] != 0 {
			c.Registers[x] = uint8(i)
			return
		}
	}

	c.Back()
}

// Fx15 - LD DT, Vx
// Set delay timer = Vx
func (c *Chip8) LDDTVX(x uint8) {
	c.DelayTimer = c.Registers[x]
}

// Fx18 - LD ST, Vx
// Set sound timer = Vx
func (c *Chip8) LDSTVX(x uint8) {
	c.SoundTimer = c.Registers[x]
}

// Fx1E - ADD I, Vx
// Set I = I + Vx
func (c *Chip8) ADDIVX(x uint8) {
	c.Index += uint16(c.Registers[x])
}

// Fx29 - LD F, Vx
// Set I = location of sprite for digit Vx
func (c *Chip8) LDFVX(x uint8) {
	digit := c.Registers[x]

	c.Index = uint16(font.START_ADDRESS) + uint16(5*digit)
}

// Fx33 - LD B, Vx
// Store BCD representation of Vx in memory locations I, I+1, and I+2
//
// The interpreter takes the decimal value of Vx, and places the hundreds digit
// in memory at location in I, the tens digit at location I+1, and the ones digit
// at location I+2.
func (c *Chip8) LDBVX(x uint8) {
	val := c.Registers[x]

	// Ones-place
	c.Memory[c.Index+2] = val % 10
	val /= 10

	// Tens-place
	c.Memory[c.Index+1] = val % 10
	val /= 10

	// Hundreds-place
	c.Memory[c.Index] = val % 10
}

// Fx55 - LD [I], Vx
// Store registers V0 through Vx in memory starting at location I
func (c *Chip8) LDIVX(x uint8) {
	for i := 0; i <= int(x); i++ {
		c.Memory[int(c.Index)+i] = c.Registers[i]
	}
}

// Fx65 - LD Vx, [I]
// Read registers V0 through Vx from memory starting at location I
func (c *Chip8) LDVXI(x uint8) {
	for i := 0; i <= int(x); i++ {
		c.Registers[i] = c.Memory[int(c.Index)+i]
	}
}
