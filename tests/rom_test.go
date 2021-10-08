package sys_test

import (
	"gr8/sys"
	"testing"
)

func TestDataLoad(t *testing.T) {
	data := []byte{1, 2, 3, 4, 5}
	chip := sys.NewChip8()

	chip.LoadROMData(data)

	if len(chip.ROMData()) != len(data) {
		t.Errorf("Data length mismatch: %v != %v", len(chip.ROMData()), len(data))
	}

	for idx, val := range chip.ROMData() {
		if data[idx] != val {
			t.Errorf("Data load mismatch: rom[%v]=%v != data[%v]=%v", idx, val, idx, data[idx])
		}
	}
}
