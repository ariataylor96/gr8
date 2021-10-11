package interfaces

const (
	VIDEO_WIDTH  byte = 64
	VIDEO_HEIGHT byte = 32
)

type Chip8 struct {
	Registers  [16]byte
	Memory     [4096]byte
	Index      uint16
	PC         uint16
	Stack      [16]uint16
	SP         byte
	DelayTimer byte
	SoundTimer byte
	Keypad     [16]byte
	Video      [int(VIDEO_WIDTH) * int(VIDEO_HEIGHT)]byte
	Opcode     uint16
	opcodes
}
