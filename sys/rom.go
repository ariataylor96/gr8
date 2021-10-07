package sys

import "os"

const ROM_START_ADDRESS int = 0x200

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
		c.Memory[ROM_START_ADDRESS+idx] = val
	}

	c.rom_length = uint16(len(buf))
}

func (c *Chip8) ROMData() []byte {
	return c.Memory[ROM_START_ADDRESS : ROM_START_ADDRESS+int(c.rom_length)]
}
