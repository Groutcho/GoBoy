package cpu

import . "common"

func init() {
	dispatch_table[0x03F] = x3F_ccf
	dispatch_table[0x037] = x37_scf
	dispatch_table[0x000] = x00_nop
	dispatch_table[0x0F3] = xF3_di
	dispatch_table[0x0FB] = xFB_ei
}

func x3F_ccf() int {
	if GetFlagCy() {
		SetFlagCy(false)
	} else {
		SetFlagCy(true)
	}

	SetFlagN(false)
	SetFlagH(false)

	return 1
}

func x37_scf() int {
	SetFlagN(false)
	SetFlagH(false)
	SetFlagCy(true)

	return 1
}

func x00_nop() int {
	return 1
}

func xF3_di() int {
	SetIME(false)
	return 1
}

func xFB_ei() int {
	SetIME(true)
	return 1
}

func x2F_cpl() int {

	SetA(GetA() ^ 0xFF)

	SetFlagH(true)
	SetFlagN(true)

	return 1
}

// When this instruction is executed, the A register is BCD corrected
// using the contents of the flags. The exact process is the following:
// if the least significant four bits of A contain a non-BCD digit
// (i. e. it is greater than 9) or the H flag is set, then $06 is added
// to the register. Then the four most significant bits are checked.
// If this more significant digit also happens to be greater than 9 or
// the C flag is set, then $60 is added.
//
// If the second addition was needed, the C flag is set after execution,
// otherwise it is reset. The N flag is preserved, P/V is parity and
// the others are altered by definition.
// source: http://z80-heaven.wikidot.com/instructions-set:daa
func x27_daa() int {
	a := GetA()
	if GetLowNibble(a) > 0x9 || GetFlagH() {
		SetFlagH(true)
		a += 0x06
	}

	if GetHighNibble(a) > 0x9 || GetFlagCy() {
		SetFlagCy(true)
		a += 0x60
	}

	SetA(a)

	return 1
}