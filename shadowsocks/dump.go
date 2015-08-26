package shadowsocks

import (
	"fmt"
)

func DumpMemoryEx(buf []byte, len int, limited bool) {
	const (
		align_bytes = 16
	)

	fmt.Printf("\n")
	fmt.Printf("Address    0  1  2  3  4  5  6  7  8  9  A  B  C  D  E  F   Memory\n")
	fmt.Printf("-----------------------------------------------------------------------------\n")
	var lines int
	lines = (len + align_bytes - 1) / align_bytes

	if !limited {
		for l := 0; l < lines; l++ {
			fmt.Printf("%08X  ", l*align_bytes)
			for i := 0; i < align_bytes; i++ {
				fmt.Printf("%02X ", buf[l*align_bytes+i])
			}
			fmt.Printf("  ")
			for i := 0; i < align_bytes; i++ {
				ch := buf[l*align_bytes+i]
				if ch < 32 {
					fmt.Printf(".")
				} else {
					fmt.Printf("%c", buf[l*align_bytes+i])
				}
			}
			fmt.Printf("\n")
		}
	} else {
		for l := 0; l < lines; l++ {
			offset := l * align_bytes
			fmt.Printf("%08X  ", l*align_bytes)
			for i := 0; i < align_bytes; i++ {
				if (offset + i) < len {
					fmt.Printf("%02X ", buf[l*align_bytes+i])
				} else {
					fmt.Printf("-- ")
				}
			}
			fmt.Printf("  ")
			for i := 0; i < align_bytes; i++ {
				if (offset + i) < len {
					ch := buf[l*align_bytes+i]
					if ch < 32 {
						fmt.Printf(".")
					} else if ch <= 127 {
						fmt.Printf("%c", ch)
					} else {
						fmt.Printf("?")
					}
				} else {
					fmt.Printf("?")
				}
			}
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\n")
	fmt.Printf("len = %d bytes\n", len)
	fmt.Printf("\n")
}

func DumpMemory(buf []byte, len int) {
	DumpMemoryEx(buf, len, false)
}
