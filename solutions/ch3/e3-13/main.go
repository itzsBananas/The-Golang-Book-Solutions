package main

import (
	"fmt"
)

// Exercise 3.13: Write const declarations for KB, MB, up through YB as compactly as you can.
const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424
	YiB // 1208925819614629174706176
)

// const (
// 	KB = 0x3e8
// 	MB = 0xf4240
// 	GB = 0x3b9aca00
// 	TB = 0xe8d4a51000
// 	PB = 0x38d7ea4c68000
// 	EB = 0xde0b6b3a7640000
// 	ZB = 0x35c9adc5dea00000
// 	YB = 0x35c9adc5dea00000 * 1000
// )

// const (
// KB, MB, GB, TB, PB, EB, ZB, YB = 1e3, 1e6, 1e9, 1e12, 1e15, 1e18, 1e21, 1e24
// )

const (
	KB, MB, GB, TB, PB, EB, ZB, YB = 1000, KB * KB, MB * KB, GB * KB, TB * GB, PB * KB, EB * KB, ZB * KB
)

func main() {
	a := uint64(1)
	for i := 0; i < 8; i++ {
		a = a * 1000
		fmt.Printf("%x\n", a)
	}
}
