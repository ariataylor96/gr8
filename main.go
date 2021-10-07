package main

import (
	"fmt"
	"gr8/sys"
)

func main() {
	chip := sys.NewChip8()

	for i := 0; i < 65535*64; i++ {
		if chip.RandByte() == 254 {
			fmt.Println("Got a 0!")
			return
		}
	}
}
