package sys_test

import (
	"gr8/sys"
	"testing"
)

const TVAL uint32 = 254

func TestCLS(t *testing.T) {
	chip := sys.NewChip8()

	chip.Video[0] = TVAL
	if chip.Video[0] != TVAL {
		t.Errorf("Failed to set VRAM: %v != %v", chip.Video[0], TVAL)
	}

	chip.CLS()
	if chip.Video[0] != 0 {
		t.Errorf("Failed to clear VRAM: %v != %v", chip.Video[0], 0)
	}
}
