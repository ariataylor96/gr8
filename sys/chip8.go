package sys

const (
	VIDEO_WIDTH  uint8 = 64
	VIDEO_HEIGHT uint8 = 32
)

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
	Video      [int(VIDEO_WIDTH) * int(VIDEO_HEIGHT)]uint32
	Opcode     uint16

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
