package sys

import "os"

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

const START_ADDRESS int = 0x200

func (c *Chip8) LoadROMFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	info, _ := file.Stat()
	buf := make([]byte, info.Size())

	file.Read(buf)

	c.LoadROMData(buf)
}

func (c *Chip8) LoadROMData(buf []byte) {
	for idx, val := range buf {
		c.Memory[START_ADDRESS+idx] = val
	}

	c.rom_length = uint16(len(buf))
}

func (c *Chip8) ROMData() []byte {
	return c.Memory[START_ADDRESS : START_ADDRESS+int(c.rom_length)]
}
