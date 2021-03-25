package main

import (
	"fmt"
	"unsafe"
)

func main() {

	type word struct {
		low   byte
		hight byte
	}

	var i int16 = 0x0102
	var hight, low byte = 0x01, 0x02

	w := (*word)(unsafe.Pointer(&i))

	if hight == w.hight && low == w.low {
		fmt.Println("little-endian")
	} else {
		fmt.Println("big-endian")
	}

}
