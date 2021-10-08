package main

import (
	"fmt"
)

func main() {
	fmt.Println((0x13E5 & 0xF000) >> 12)
	fmt.Println(byte((0x13E5 & 0xF000) >> 12))
}
