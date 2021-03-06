/*
This package contains all functions relative to manipulating the registers
of the GoBoy emulator.
*/
package cpu

import (
	. "common"
)

/*
The Game boy 6

All are 16bit, but the first four registers
can also be accessed as two 8bit subregisters
This gives us 8 more 8bit  A, F, B, C, D, E, H, L.

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
var (
	A  uint8
	F  uint8
	B  uint8
	C  uint8
	D  uint8
	E  uint8
	H  uint8
	L  uint8
	SP uint16 /* stack pointer */
	PC uint16 /* program counter */

	F_SET_0  = 0
	F_SET_1  = 1
	F_SET_IF = 2
	F_IGNORE = 3
	F_8bit   = 4
	F_16bit  = 5
)

// Set the AF register with the given 16bit value.
func SetAF(value uint16) {
	A = uint8(value >> 8)
	F = uint8(value)
}

// Return the value of the AF register.
func GetAF() uint16 {
	return (uint16(A) << 8) | uint16(F)
}

// Increment AF.
func IncAF() {
	SetAF(GetAF() + 1)
}

// Decrement AF.
func DecAF() {
	SetAF(GetAF() - 1)
}

// Set the BC register with the given 16bit value.
func SetBC(value uint16) {
	B = uint8(value >> 8)
	C = uint8(value)
}

// Return the value of the BC register.
func GetBC() uint16 {
	return (uint16(B) << 8) | uint16(C)
}

// Increment BC.
func IncBC() {
	SetBC(GetBC() + 1)
}

// Decrement BC.
func DecBC() {
	SetBC(GetBC() - 1)
}

// Set the DE register with the given 16bit value.
func SetDE(value uint16) {
	D = uint8(value >> 8)
	E = uint8(value)
}

// Return the value of the DE register.
func GetDE() uint16 {
	return (uint16(D) << 8) | uint16(E)
}

// Increment DE.
func IncDE() {
	SetDE(GetDE() + 1)
}

// Decrement DE.
func DecDE() {
	SetDE(GetDE() - 1)
}

// Set the HL register with the given 16bit value.
func SetHL(value uint16) {
	H = uint8(value >> 8)
	L = uint8(value)
}

// Return the value of the HL register.
func GetHL() uint16 {
	return (uint16(H) << 8) | uint16(L)
}

// Increment HL.
func IncHL() {
	SetHL(GetHL() + 1)
}

// Decrement HL.
func DecHL() {
	SetHL(GetHL() - 1)
}

// Set the SP register with the given 16bit value.
func SetSP(value uint16) {
	SP = value
}

// Return the value of the SP register.
func GetSP() uint16 {
	return SP
}

// Set the PC register with the given 16bit value.
func SetPC(value uint16) {
	PC = value
}

// Increment the PC.
func IncPC() {
	PC += 1
}

// Increment the PC.
func IncSP() {
	SP += 1
}

// Decrement the PC.
func DecPC() {
	PC -= 1
}

// Decrement the PC.
func DecSP() {
	SP -= 1
}

// Return the value of the PC register.
func GetPC() uint16 {
	return PC
}

// Return the A register value as a 8bit unsigned integer.
func GetA() uint8 {
	return A
}

// Return the F register value as a 8bit unsigned integer.
func GetF() uint8 {
	return F
}

// Return the B register value as a 8bit unsigned integer.
func GetB() uint8 {
	return B
}

// Return the F register value as a 8bit unsigned integer.
func GetC() uint8 {
	return C
}

// Return the B register value as a 8bit unsigned integer.
func GetD() uint8 {
	return D
}

// Return the F register value as a 8bit unsigned integer.
func GetE() uint8 {
	return E
}

// Return the H register value as a 8bit unsigned integer.
func GetH() uint8 {
	return H
}

// Return the L register value as a 8bit unsigned integer.
func GetL() uint8 {
	return L
}

// Return the value of the Zero flag.
func GetFlagZf() bool {
	return (F & 0x80) != 0
}

// Return the value of the N (Add/Sub) flag.
func GetFlagN() bool {
	return (F & 0x40) != 0
}

// Return the value of the H (Half carry) flag.
func GetFlagH() bool {
	return (F & 0x20) != 0
}

// Return the value of the CY (carry) flag.
func GetFlagCy() bool {
	return (F & 0x10) != 0
}

func GetFlagCyInt() int {
	if GetFlagCy() {
		return 1
	} else {
		return 0
	}
}

func GetFlagZfInt() int {
	if GetFlagZf() {
		return 1
	} else {
		return 0
	}
}

func GetFlagHInt() int {
	if GetFlagH() {
		return 1
	} else {
		return 0
	}
}

func GetFlagNInt() int {
	if GetFlagN() {
		return 1
	} else {
		return 0
	}
}

// return 1 is left + right overflows
func IsAddHalfCarry(left uint8, right uint8) int {
	l := int(GetLowNibble(left))
	r := int(GetLowNibble(right))

	if l+r > 0xF {
		return 1
	} else {
		return 0
	}
}

// return 1 is left - right underflows
func IsSubHalfCarry(left uint8, right uint8) int {
	l := int(GetLowNibble(left))
	r := int(GetLowNibble(right))

	if l-r < 0x0 {
		return 1
	} else {
		return 0
	}
}

// Set the Z, H, N, C flags according to the provided strategies:
//  - F_SET_0: unset the flag
//  - F_SET_1: set the flag
//  - F_SET_IF: set the flag if needed (e.g the value equals to 0,
//    the zero flag is set)
//  - F_IGNORE: leave this flag unchanged
func SetFlags(value int, Z int, N int, H int, C int, size int) {
	if Z == F_SET_0 {
		SetFlagZf(false)
	} else if Z == F_SET_1 {
		SetFlagZf(true)
	} else if Z == F_SET_IF {
		if value == 0 {
			SetFlagZf(true)
		} else {
			SetFlagZf(false)
		}
	}

	if H == F_SET_0 {
		SetFlagH(false)
	} else if H == F_SET_1 {
		SetFlagH(true)
	}

	if N == F_SET_0 {
		SetFlagN(false)
	} else if N == F_SET_1 {
		SetFlagN(true)
	}

	if C == F_SET_0 {
		SetFlagCy(false)
	} else if C == F_SET_1 {
		SetFlagCy(true)
	} else if C == F_SET_IF {
		limit := 0xFF
		if size == F_16bit {
			limit = 0xFFFF
		}

		SetFlagCy(value > limit || value < 0)
	}
}

// Set the value of the Zero flag.
func SetFlagZf(b bool) {
	if b {
		F |= (1 << 7)
	} else {
		F &= ^(uint8(1 << 7))
	}
}

// Set the value of the N (Add/Sub) flag.
func SetFlagN(b bool) {
	if b {
		F |= (1 << 6)
	} else {
		F &= ^(uint8(1 << 6))
	}
}

// Set the value of the H (Half carry) flag.
func SetFlagH(b bool) {
	if b {
		F |= (1 << 5)
	} else {
		F &= ^(uint8(1 << 5))
	}
}

// Set the value of the CY (carry) flag.
func SetFlagCy(b bool) {
	if b {
		F |= (1 << 4)
	} else {
		F &= ^(uint8(1 << 4))
	}
}

// Set the value of the A register.
func SetA(value uint8) {
	A = value
}

// Increment the A register.
func IncA() {
	A += 1
}

// Set the value of the F register.
func SetF(value uint8) {
	F = value
}

// Increment the F register.
func IncF() {
	F += 1
}

// Set the value of the B register.
func SetB(value uint8) {
	B = value
}

// Increment the B register.
func IncB() {
	B += 1
}

// Set the value of the C register.
func SetC(value uint8) {
	C = value
}

// Increment the C register.
func IncC() {
	C += 1
}

// Set the value of the D register.
func SetD(value uint8) {
	D = value
}

// Increment the D register.
func IncD() {
	D += 1
}

// Set the value of the E register.
func SetE(value uint8) {
	E = value
}

// Increment the E register.
func IncE() {
	E += 1
}

// Set the value of the H register.
func SetH(value uint8) {
	H = value
}

// Increment the H register.
func IncH() {
	H += 1
}

// Set the value of the L register.
func SetL(value uint8) {
	L = value
}

// Increment the L register.
func IncL() {
	L += 1
}

// Decrement the A register.
func DecA() {
	A -= 1
}

// Decrement the F register.
func DecF() {
	F -= 1
}

// Decrement the B register.
func DecB() {
	B -= 1
}

// Decrement the C register.
func DecC() {
	C -= 1
}

// Decrement the D register.
func DecD() {
	D -= 1
}

// Decrement the E register.
func DecE() {
	E -= 1
}

// Decrement the H register.
func DecH() {
	H -= 1
}

// Decrement the L register.
func DecL() {
	L -= 1
}

func Swap(value uint8) uint8 {
	return uint8(value<<4 | value>>4)
}

// Reset all to 0x0000.
func Reset() {
	A = 0x00
	F = 0x00
	B = 0x00
	C = 0x00
	D = 0x00
	E = 0x00
	H = 0x00
	L = 0x00
	SP = 0x0000
	PC = 0x0000
}
