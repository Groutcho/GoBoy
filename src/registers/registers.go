/*
This package contains all functions relative to manipulating the registers
of the GoBoy emulator.
*/
package registers

import (

)

/*
The Game boy 6 registers.

All registers are 16bit, but the first four registers
can also be accessed as two 8bit subregisters
This gives us 8 more 8bit registers: A, F, B, C, D, E, H, L.

The lower 8bits of the AF register is the "flags" (F) register.
It contains 4 flags:

 Bit   Name  Set  Clr Expl
 7     zf    Z    NZ  Zero flag
 6     n     -    -   Add/Sub-flag (BCD)
 5     h     -    -   Half Carry flag (BCD)
 4     cy    C    NC  Carry flag
 0-3   -     -    -   Not used (always zero)

The Stack Pointer (SP) and Program Counter (PC) cannot be
accessed outside their 16bit range.
*/
type Registers struct {
	AF uint16
	BC uint16
	DE uint16
	HL uint16
	SP uint16 /* stack pointer */
	PC uint16 /* program counter */
}

var registers Registers

func init() {
	registers = Registers{AF: 0x0000,
		BC: 0x0000,
		DE: 0x0000,
		HL: 0x0000,
		SP: 0x0000,
		PC: 0x0000}
}

// set the 8 high bits of <target> with <value>.
func setHighBits(value uint8, target uint16) uint16 {
	return target | (uint16(value) << 8)
}

// set the 8 low bits of <target> with <value>.
func setLowBits(value uint8, target uint16) uint16 {
	return target | uint16(value)
}

// return the 8 high bits of <value>.
func getHighBits(value uint16) uint8 {
	return uint8((value & 0xFF00) >> 8)
}

// return the 8 low bits of <value>.
func getLowBits(value uint16) uint8 {
	return uint8(value & 0x00FF)
}

// Set the AF register with the given 16bit value.
func SetAF(value uint16) {
	registers.AF = value
}

// Return the value of the AF register.
func GetAF() uint16 {
	return registers.AF
}

// Set the BC register with the given 16bit value.
func SetBC(value uint16) {
	registers.BC = value
}

// Return the value of the BC register.
func GetBC() uint16 {
	return registers.BC
}

// Set the DE register with the given 16bit value.
func SetDE(value uint16) {
	registers.DE = value
}

// Return the value of the DE register.
func GetDE() uint16 {
	return registers.DE
}

// Set the HL register with the given 16bit value.
func SetHL(value uint16) {
	registers.HL = value
}

// Return the value of the HL register.
func GetHL() uint16 {
	return registers.HL
}

// Set the SP register with the given 16bit value.
func SetSP(value uint16) {
	registers.SP = value
}

// Return the value of the SP register.
func GetSP() uint16 {
	return registers.SP
}

// Set the PC register with the given 16bit value.
func SetPC(value uint16) {
	registers.PC = value
}

// Increment the PC.
func IncPC() {
	registers.PC += 1
}

// Return the value of the PC register.
func GetPC() uint16 {
	return registers.PC
}

// Set the value of the upper high bits of the AF register.
func SetA(value uint8) {
	registers.AF = setHighBits(value, registers.AF)
}

// Return the A register value as a 8bit unsigned integer.
func GetA() uint8 {
	return getHighBits(registers.AF)
}

// Set the value of the lower high bits of the AF register.
func SetF(value uint8) {
	registers.AF = setLowBits(value, registers.AF)
}

// Return the F register value as a 8bit unsigned integer.
func GetF() uint8 {
	return getLowBits(registers.AF)
}

// Set the value of the upper high bits of the BC register.
func SetB(value uint8) {
	registers.BC = setHighBits(value, registers.BC)
}

// Return the B register value as a 8bit unsigned integer.
func GetB() uint8 {
	return getHighBits(registers.BC)
}

// Set the value of the lower high bits of the BC register.
func SetC(value uint8) {
	registers.BC = setLowBits(value, registers.BC)
}

// Return the F register value as a 8bit unsigned integer.
func GetC() uint8 {
	return getLowBits(registers.BC)
}

// Set the value of the upper high bits of the DE register.
func SetD(value uint8) {
	registers.DE = setHighBits(value, registers.DE)
}

// Return the B register value as a 8bit unsigned integer.
func GetD() uint8 {
	return getHighBits(registers.DE)
}

// Set the value of the lower high bits of the DE register.
func SetE(value uint8) {
	registers.DE = setLowBits(value, registers.DE)
}

// Return the F register value as a 8bit unsigned integer.
func GetE() uint8 {
	return getLowBits(registers.DE)
}

// Set the value of the upper high bits of the HL register.
func SetH(value uint8) {
	registers.HL = setHighBits(value, registers.HL)
}

// Return the H register value as a 8bit unsigned integer.
func GetH() uint8 {
	return getHighBits(registers.HL)
}

// Set the value of the lower high bits of the HL register.
func SetL(value uint8) {
	registers.HL = setLowBits(value, registers.HL)
}

// Return the L register value as a 8bit unsigned integer.
func GetL() uint8 {
	return getLowBits(registers.HL)
}

// Return the value of the Zero flag.
func GetFlagZf() bool {
	return (registers.AF & 0x0080) != 0
}

// Return the value of the N (Add/Sub) flag.
func GetFlagN() bool {
	return (registers.AF & 0x0040) != 0
}

// Return the value of the H (Half carry) flag.
func GetFlagH() bool {
	return (registers.AF & 0x0020) != 0
}

// Return the value of the CY (carry) flag.
func GetFlagCy() bool {
	return (registers.AF & 0x0010) != 0
}

// Set the value of the Zero flag.
func SetFlagZf(b bool) {
	if b {
		registers.AF |= (1 << 7)
	} else {
		registers.AF &= ^(uint16(1 << 7))
	}
}

// Set the value of the N (Add/Sub) flag.
func SetFlagN(b bool) {
	if b {
		registers.AF |= (1 << 6)
	} else {
		registers.AF &= ^(uint16(1 << 6))
	}
}

// Set the value of the H (Half carry) flag.
func SetFlagH(b bool) {
	if b {
		registers.AF |= (1 << 5)
	} else {
		registers.AF &= ^(uint16(1 << 5))
	}
}

// Set the value of the CY (carry) flag.
func SetFlagCy(b bool) {
	if b {
		registers.AF |= (1 << 4)
	} else {
		registers.AF &= ^(uint16(1 << 4))
	}
}

// Reset all registers to 0x0000.
func Reset() {
	registers.AF = 0x0000
	registers.BC = 0x0000
	registers.BC = 0x0000
	registers.HL = 0x0000
	registers.SP = 0x0000
	registers.PC = 0x0000
}