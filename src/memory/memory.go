package memory

import (
	"bytes"
	. "common"
	"fmt"
	"time"
)

/* Interrupt bit indices */
const (
	VBLANK   = 0
	LCD_STAT = 1
	TIMER    = 2
	SERIAL   = 3
	JOYPAD   = 4
)

const (
	P1_ADDR             = 0xFF00
	SB_ADDR             = 0xFF01
	SC_ADDR             = 0xFF02
	DIV_ADDR            = 0xFF04
	TIMA_ADDR           = 0xFF05
	TMA_ADDR            = 0xFF06
	TAC_ADDR            = 0xFF07
	NR10_ADDR           = 0xFF10
	NR11_ADDR           = 0xFF11
	NR12_ADDR           = 0xFF12
	NR13_ADDR           = 0xFF13
	NR14_ADDR           = 0xFF14
	NR21_ADDR           = 0xFF16
	NR22_ADDR           = 0xFF17
	NR23_ADDR           = 0xFF18
	NR24_ADDR           = 0xFF19
	NR30_ADDR           = 0xFF1A
	NR31_ADDR           = 0xFF1B
	NR33_ADDR           = 0xFF1D
	NR34_ADDR           = 0xFF1E
	NR41_ADDR           = 0xFF20
	NR42_ADDR           = 0xFF21
	NR50_ADDR           = 0xFF24
	NR51_ADDR           = 0xFF25
	NR52_ADDR           = 0xFF26
	WAV_ADDR            = 0xFF30
	LCDC_ADDR           = 0xFF40
	STAT_ADDR           = 0xFF41
	SCY_ADDR            = 0xFF42
	SCX_ADDR            = 0xFF43
	LY_ADDR             = 0xFF44
	LYC_ADDR            = 0xFF45
	DMA_ADDR            = 0xFF46
	BGP_ADDR            = 0xFF47
	OBP0_ADDR           = 0xFF48
	OBP1_ADDR           = 0xFF49
	WY_ADDR             = 0xFF4A
	WX_ADDR             = 0xFF4B
	INTRPT_ENABLE_ADDR  = 0xFFFF
	INTRPT_REQUEST_ADDR = 0xFF0F
)

/* the game boy address space is 16bit wide */
var RAM = new([65536]byte)
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
	return uint16(uint16(RAM[addr+1])<<8 | uint16(RAM[addr]))
}

func Set(addr uint16, value byte) {
	if addr == DMA_ADDR {
		startDMATransfer(uint16(value) * 0x100)
	}
	RAM[addr] = value
}

func Set16(addr uint16, value uint16) {
	RAM[addr] = GetLowBits(value)
	RAM[addr+1] = GetHighBits(value)
}

func SetRange(from uint16, data []byte) {
	k := 0
	for i := 0; i < len(data); i++ {
		RAM[from+uint16(i)] = data[k]
		k++
	}
}

// return a slice of the complete RAM
func GetMemoryMap() []byte {
	return RAM[:]
}

func GetRange(from uint16, size uint16) []byte {
	return RAM[from : from+size]
}

func DumpRange(from uint16, to uint16) string {
	var buffer bytes.Buffer

	for i := 0; i < len(RAM[from:to+1]); i++ {
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

func GetP1() byte {
	return RAM[P1_ADDR]
}

func SetP1(value byte) {
	RAM[P1_ADDR] = value
}

func GetSB() byte {
	return RAM[SB_ADDR]
}

func SetSB(value byte) {
	RAM[SB_ADDR] = value
}

func GetSC() byte {
	return RAM[SC_ADDR]
}

func SetSC(value byte) {
	RAM[SC_ADDR] = value
}

func GetDIV() byte {
	return RAM[DIV_ADDR]
}

func SetDIV(value byte) {
	RAM[DIV_ADDR] = value
}

func GetTIMA() byte {
	return RAM[TIMA_ADDR]
}

func SetTIMA(value byte) {
	RAM[TIMA_ADDR] = value
}

func GetTMA() byte {
	return RAM[TMA_ADDR]
}

func SetTMA(value byte) {
	RAM[TMA_ADDR] = value
}

func GetTAC() byte {
	return RAM[TAC_ADDR]
}

func SetTAC(value byte) {
	RAM[TAC_ADDR] = value
}

func GetNR10() byte {
	return RAM[NR10_ADDR]
}

func SetNR10(value byte) {
	RAM[NR10_ADDR] = value
}

func GetNR11() byte {
	return RAM[NR11_ADDR]
}

func SetNR11(value byte) {
	RAM[NR11_ADDR] = value
}

func GetNR12() byte {
	return RAM[NR12_ADDR]
}

func SetNR12(value byte) {
	RAM[NR12_ADDR] = value
}

func GetNR13() byte {
	return RAM[NR13_ADDR]
}

func SetNR13(value byte) {
	RAM[NR13_ADDR] = value
}

func GetNR14() byte {
	return RAM[NR14_ADDR]
}

func SetNR14(value byte) {
	RAM[NR14_ADDR] = value
}

func GetNR21() byte {
	return RAM[NR21_ADDR]
}

func SetNR21(value byte) {
	RAM[NR21_ADDR] = value
}

func GetNR22() byte {
	return RAM[NR22_ADDR]
}

func SetNR22(value byte) {
	RAM[NR22_ADDR] = value
}

func GetNR23() byte {
	return RAM[NR23_ADDR]
}

func SetNR23(value byte) {
	RAM[NR23_ADDR] = value
}

func GetNR24() byte {
	return RAM[NR24_ADDR]
}

func SetNR24(value byte) {
	RAM[NR24_ADDR] = value
}

func GetNR30() byte {
	return RAM[NR30_ADDR]
}

func SetNR30(value byte) {
	RAM[NR30_ADDR] = value
}

func GetNR31() byte {
	return RAM[NR31_ADDR]
}

func SetNR31(value byte) {
	RAM[NR31_ADDR] = value
}

func GetNR33() byte {
	return RAM[NR33_ADDR]
}

func SetNR33(value byte) {
	RAM[NR33_ADDR] = value
}

func GetNR34() byte {
	return RAM[NR34_ADDR]
}

func SetNR34(value byte) {
	RAM[NR34_ADDR] = value
}

func GetNR41() byte {
	return RAM[NR41_ADDR]
}

func SetNR41(value byte) {
	RAM[NR41_ADDR] = value
}

func GetNR42() byte {
	return RAM[NR42_ADDR]
}

func SetNR42(value byte) {
	RAM[NR42_ADDR] = value
}

func GetNR50() byte {
	return RAM[NR50_ADDR]
}

func SetNR50(value byte) {
	RAM[NR50_ADDR] = value
}

func GetNR51() byte {
	return RAM[NR51_ADDR]
}

func SetNR51(value byte) {
	RAM[NR51_ADDR] = value
}

func GetNR52() byte {
	return RAM[NR52_ADDR]
}

func SetNR52(value byte) {
	RAM[NR52_ADDR] = value
}

func GetWAV() byte {
	return RAM[WAV_ADDR]
}

func SetWAV(value byte) {
	RAM[WAV_ADDR] = value
}

func GetLCDC() byte {
	return RAM[LCDC_ADDR]
}

func SetLCDC(value byte) {
	RAM[LCDC_ADDR] = value
}

func GetSTAT() byte {
	return RAM[STAT_ADDR]
}

func SetSTAT(value byte) {
	RAM[STAT_ADDR] = value
}

func GetSCY() byte {
	return RAM[SCY_ADDR]
}

func IncSCY() {
	RAM[SCY_ADDR]++
}

func SetSCY(value byte) {
	RAM[SCY_ADDR] = value
}

func GetSCX() byte {
	return RAM[SCX_ADDR]
}

func SetSCX(value byte) {
	RAM[SCX_ADDR] = value
}

func IncSCX() {
	RAM[SCX_ADDR]++
}

func GetLY() byte {
	return RAM[LY_ADDR]
}

func SetLY(value byte) {
	RAM[LY_ADDR] = value
}

func IncLY() {
	RAM[LY_ADDR]++
}

func GetLYC() byte {
	return RAM[LYC_ADDR]
}

func SetLYC(value byte) {
	RAM[LYC_ADDR] = value
}

func GetDMA() byte {
	return RAM[DMA_ADDR]
}

func SetDMA(value byte) {
	RAM[DMA_ADDR] = value
}

func GetBGP() byte {
	return RAM[BGP_ADDR]
}

func SetBGP(value byte) {
	RAM[BGP_ADDR] = value
}

func GetOBP0() byte {
	return RAM[OBP0_ADDR]
}

func SetOBP0(value byte) {
	RAM[OBP0_ADDR] = value
}

func GetOBP1() byte {
	return RAM[OBP1_ADDR]
}

func SetOBP1(value byte) {
	RAM[OBP1_ADDR] = value
}

func GetWY() byte {
	return RAM[WY_ADDR]
}

func SetWY(value byte) {
	RAM[WY_ADDR] = value
}

func GetWX() byte {
	return RAM[WX_ADDR]
}

func SetWX(value byte) {
	RAM[WX_ADDR] = value
}

// initialize a DMA transfer of 40 * 28 bit from address <addr> to OAM (0xFE00)
func startDMATransfer(addr uint16) {
	tick := time.NewTicker(time.Microsecond * 160)
	var j uint16

	for i := 0; i < 40; i++ {
		j = uint16(i)
		RAM[0xFE00+j] = RAM[addr+j]
	}

	<-tick.C
}
