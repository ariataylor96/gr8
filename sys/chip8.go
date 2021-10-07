package sys

type Chip8 struct {
	Registers  [16]uint8
	Memory     [4096]uint8
	Index      uint16
	PC         uint16
	Stack      [16]uint16
	SP         uint8
	DelayTimer uint8
	SoundTimer uint8
	Keypad     [16]uint8
	Video      [64 * 32]uint32
	Opcode     uint16

	rom_length uint16
}

func NewChip8() Chip8 {
	res := Chip8{}

	res.LoadFont()

	return res
}

func (c *Chip8) Next() {
	c.PC += 2
}
