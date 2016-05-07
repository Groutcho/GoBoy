package memory

import (
	"bytes"
	"fmt"
	. "common"
)

/* the game boy address space is 16bit wide */
var RAM = new([65536]byte)

/* Interrupt bit indices */
const VBLANK = 0
const LCD_STAT = 1
const TIMER = 2
const SERIAL = 3
const JOYPAD = 4

const INTRPT_ENABLE_ADDR = 0xFFFF
const INTRPT_REQUEST_ADDR = 0xFF0F

var interruptHandlers = make([]uint16, 5, 5)

func init() {
	interruptHandlers[0] = 0x0040 // VBLANK
	interruptHandlers[1] = 0x0048 // LCD_STAT
	interruptHandlers[2] = 0x0050 // TIMER
	interruptHandlers[3] = 0x0058 // SERIAL
	interruptHandlers[4] = 0x0060 // JOYPAD
}

func Get(addr uint16) byte {
	return RAM[addr]
}

func Get16(addr uint16) uint16 {
	return uint16(uint16(RAM[addr+1]) << 8 | uint16(RAM[addr]))
}

func Set(addr uint16, value byte) {
	RAM[addr] = value
}

func Set16(addr uint16, value uint16) {
	RAM[addr] = GetLowBits(value)
	RAM[addr+1] = GetHighBits(value)
}

func SetRange(from uint16, to uint16, data []byte) {
	k := 0
	for i := from; i <= to; i++ {
		RAM[i] = data[k]
		k++
	}
}

func GetRange(from uint16, to uint16) []byte {
	result := make([]byte, to - from + 1)
	k := 0
	for i := from; i <= to; i++ {
		result[k] = RAM[i]
		k++
	}

	return result
}

func DumpRange(from uint16, to uint16) string {

	var buffer bytes.Buffer

	for i := 0; i < len(RAM[from:to + 1]); i++ {
		if RAM[i] == 0x00 {
			buffer.WriteString("__ ")
		} else {
			buffer.WriteString(fmt.Sprintf("%02X ", RAM[i]))
		}
	}
	return buffer.String()
}

/* reset the RAM to 0 */
func ResetMemory() {
	for i := 0; i < len(RAM); i++ {
		RAM[i] = 0x00
	}
}

/* reset a range of RAM addresses */
func resetRange(start uint16, end uint16) {
	if start == end {
		RAM[start] = 0x00
		return
	}

	for i := start; i <= end; i++ {
		// prevent uint16 overflow,
		// leading to an infinite loop
		if i == 0xFFFF {
			RAM[i] = 0x00
			return
		}
		RAM[i] = 0x00
	}
}

func GetInterruptEnable(bit uint8) bool {
	return (RAM[INTRPT_ENABLE_ADDR] & (0x01 << bit)) != 0
}

func SetInterruptEnable(bit uint8, value bool) {
	if value {
		RAM[INTRPT_ENABLE_ADDR] |= (0x01 << bit)
	} else {
		RAM[INTRPT_ENABLE_ADDR] &= ^(uint8(0x01) << bit)
	}
}

func GetInterruptRequest(bit uint8) bool {
	return (RAM[INTRPT_REQUEST_ADDR] & (0x01 << bit)) != 0
}

func SetInterruptRequest(bit uint8, value bool) {
	if value {
		RAM[INTRPT_REQUEST_ADDR] |= (0x01 << bit)
	} else {
		RAM[INTRPT_REQUEST_ADDR] &= ^(uint8(0x01) << bit)
	}
}

// return the routine address of the given interrupt handler
func GetInterruptHandler(interrupt int) uint16 {
	return Get16(interruptHandlers[interrupt])
}
