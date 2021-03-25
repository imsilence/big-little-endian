package main

import (
	"fmt"
	"unsafe"
)

func main() {

	type bytes struct {
		low   byte
		hight byte
	}

	var i int16 = 0x0102
	var hight, low byte = 0x01, 0x02

	bs := (*bytes)(unsafe.Pointer(&i))

	if hight == bs.hight && low == bs.low {
		fmt.Println("little-endian")
	} else {
		fmt.Println("big-endian")
	}
}
