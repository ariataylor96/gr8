package sys

import "gr8/interfaces"

const (
	VIDEO_WIDTH  byte = interfaces.VIDEO_WIDTH
	VIDEO_HEIGHT byte = interfaces.VIDEO_HEIGHT
)

type Chip8 struct {
	interfaces.Chip8

	romLength uint16
}

func NewChip8() Chip8 {
	res := Chip8{}

	res.LoadFont()

	return res
}

func (c *Chip8) Next() {
	c.PC += 2
}

func (c *Chip8) Back() {
	c.PC -= 2
}

func (c *Chip8) Op() byte {
	return byte((c.Opcode & 0xF000) >> 12)
}

func (c *Chip8) Cycle() {
	// This glob of bit hacks just takes two bytes and forms a
	// 16-bit word - i.e. 0xAF and 0xE2 becomes 0xAFE2
	c.Opcode = uint16(c.Memory[c.PC])<<8 | uint16(c.Memory[c.PC+1])

	// Since an instruction is 16 bits, move forward 2 bytes
	c.PC += 2

	// This is handwaving a lot here - "do work"
	c.Execute()

	// Lastly, decrement our timers
	if c.DelayTimer > 0 {
		c.DelayTimer -= 1
	}

	if c.SoundTimer > 0 {
		c.SoundTimer -= 1
	}
}
