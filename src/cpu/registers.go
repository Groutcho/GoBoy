/*
This package contains all functions relative to manipulating the registers
of the GoBoy emulator.
*/
package cpu

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
	A uint8
	F uint8
	B uint8
	C uint8
	D uint8
	E uint8
	H uint8
	L uint8
	SP uint16 /* stack pointer */
	PC uint16 /* program counter */
}

var registers Registers

func init() {
	registers = Registers{
		A: 0x00,
		B: 0x00,
		C: 0x00,
		D: 0x00,
		E: 0x00,
		F: 0x00,
		H: 0x00,
		L: 0x00,		
		SP: 0x0000,
		PC: 0x0000}
}

// set the 8 high bits of <target> with <value>.
func setHighBits(value uint8, target uint16) uint16 {
	return (target & 0x00FF) | (uint16(value) << 8)
}

// set the 8 low bits of <target> with <value>.
func setLowBits(value uint8, target uint16) uint16 {
	return (target & 0xFF00) | uint16(value)
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
	registers.A = uint8(value >> 8)
	registers.F = uint8(value)
}

// Return the value of the AF register.
func GetAF() uint16 {
	return (uint16(registers.A) << 8) | uint16(registers.F)
}

// Set the BC register with the given 16bit value.
func SetBC(value uint16) {
	registers.B = uint8(value >> 8)
	registers.C = uint8(value)
}

// Return the value of the BC register.
func GetBC() uint16 {
	return (uint16(registers.B) << 8) | uint16(registers.C)
}

// Set the DE register with the given 16bit value.
func SetDE(value uint16) {
	registers.D = uint8(value >> 8)
	registers.E = uint8(value)
}

// Return the value of the DE register.
func GetDE() uint16 {
	return (uint16(registers.D) << 8) | uint16(registers.E)
}

// Set the HL register with the given 16bit value.
func SetHL(value uint16) {
	registers.H = uint8(value >> 8)
	registers.L = uint8(value)
}

// Return the value of the HL register.
func GetHL() uint16 {
	return (uint16(registers.H) << 8) | uint16(registers.L)
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

// Increment the PC.
func IncSP() {
	registers.SP += 1
}

// Decrement the PC.
func DecPC() {
	registers.PC -= 1
}

// Decrement the PC.
func DecSP() {
	registers.SP -= 1
}

// Return the value of the PC register.
func GetPC() uint16 {
	return registers.PC
}

// Return the A register value as a 8bit unsigned integer.
func GetA() uint8 {
	return registers.A
}

// Return the F register value as a 8bit unsigned integer.
func GetF() uint8 {
	return registers.F
}

// Return the B register value as a 8bit unsigned integer.
func GetB() uint8 {
	return registers.B
}

// Return the F register value as a 8bit unsigned integer.
func GetC() uint8 {
	return registers.C
}

// Return the B register value as a 8bit unsigned integer.
func GetD() uint8 {
	return registers.D
}

// Return the F register value as a 8bit unsigned integer.
func GetE() uint8 {
	return registers.E
}

// Return the H register value as a 8bit unsigned integer.
func GetH() uint8 {
	return registers.H
}

// Return the L register value as a 8bit unsigned integer.
func GetL() uint8 {
	return registers.L
}

// Return the value of the Zero flag.
func GetFlagZf() bool {
	return (registers.F & 0x80) != 0
}

// Return the value of the N (Add/Sub) flag.
func GetFlagN() bool {
	return (registers.F & 0x40) != 0
}

// Return the value of the H (Half carry) flag.
func GetFlagH() bool {
	return (registers.F & 0x20) != 0
}

// Return the value of the CY (carry) flag.
func GetFlagCy() bool {
	return (registers.F & 0x10) != 0
}

// Set the value of the Zero flag.
func SetFlagZf(b bool) {
	if b {
		registers.F |= (1 << 7)
	} else {
		registers.F &= ^(uint8(1 << 7))
	}
}

// Set the value of the N (Add/Sub) flag.
func SetFlagN(b bool) {
	if b {
		registers.F |= (1 << 6)
	} else {
		registers.F &= ^(uint8(1 << 6))
	}
}

// Set the value of the H (Half carry) flag.
func SetFlagH(b bool) {
	if b {
		registers.F |= (1 << 5)
	} else {
		registers.F &= ^(uint8(1 << 5))
	}
}

// Set the value of the CY (carry) flag.
func SetFlagCy(b bool) {
	if b {
		registers.F |= (1 << 4)
	} else {
		registers.F &= ^(uint8(1 << 4))
	}
}


// Set the value of the A register.
func SetA(value uint8) {
	registers.A = value
}

// Increment the A register.
func IncA() {
	registers.A += 1
}

// Set the value of the F register.
func SetF(value uint8) {
	registers.F = value
}

// Increment the F register.
func IncF() {
	registers.F += 1
}

// Set the value of the B register.
func SetB(value uint8) {
	registers.B = value
}

// Increment the B register.
func IncB() {
	registers.B += 1
}

// Set the value of the C register.
func SetC(value uint8) {
	registers.C = value
}

// Increment the C register.
func IncC() {
	registers.C += 1
}

// Set the value of the D register.
func SetD(value uint8) {
	registers.D = value
}

// Increment the D register.
func IncD() {
	registers.D += 1
}

// Set the value of the E register.
func SetE(value uint8) {
	registers.E = value
}

// Increment the E register.
func IncE() {
	registers.E += 1
}

// Set the value of the H register.
func SetH(value uint8) {
	registers.H = value
}

// Increment the H register.
func IncH() {
	registers.H += 1
}

// Set the value of the L register.
func SetL(value uint8) {
	registers.L = value
}

// Increment the L register.
func IncL() {
	registers.L += 1
}

// Decrement the A register.
func DecA() {
	registers.A -= 1
}

// Decrement the F register.
func DecF() {
	registers.F -= 1
}

// Decrement the B register.
func DecB() {
	registers.B -= 1
}

// Decrement the C register.
func DecC() {
	registers.C -= 1
}

// Decrement the D register.
func DecD() {
	registers.D -= 1
}

// Decrement the E register.
func DecE() {
	registers.E -= 1
}

// Decrement the H register.
func DecH() {
	registers.H -= 1
}

// Decrement the L register.
func DecL() {
	registers.L -= 1
}

func Swap(value uint8) uint8 {
	return uint8(value << 4 | value >> 4)
}

// Reset all registers to 0x0000.
func Reset() {
	registers.A = 0x00
	registers.F = 0x00
	registers.B = 0x00
	registers.C = 0x00
	registers.D = 0x00
	registers.E = 0x00
	registers.H = 0x00
	registers.L = 0x00
	registers.SP = 0x0000
	registers.PC = 0x0000
}