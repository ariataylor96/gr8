package sys_test

import (
	"gr8/font"
	"gr8/sys"
	"testing"
)

func TestFontInitialization(t *testing.T) {
	chip := sys.NewChip8()

	for idx, val := range chip.Memory[font.START_ADDRESS : font.START_ADDRESS+font.SIZE] {
		if val != font.FONTSET[idx] {
			t.Errorf("Fontset load mismatch: rom[%v]=%v != data[%v]=%v", idx, val, idx, font.FONTSET[idx])
		}
	}
}
