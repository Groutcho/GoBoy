// this package is automatically generated.
package cpu

import . "memory"

// ld   A, B - A=B
func x78_ld() uint8 {
	SetA(GetB())

	return 1
}

// ld   A, C - A=C
func x79_ld() uint8 {
	SetA(GetC())

	return 1
}

// ld   A, D - A=D
func x7A_ld() uint8 {
	SetA(GetD())

	return 1
}

// ld   A, E - A=E
func x7B_ld() uint8 {
	SetA(GetE())

	return 1
}

// ld   A, H - A=H
func x7C_ld() uint8 {
	SetA(GetH())

	return 1
}

// ld   A, L - A=L
func x7D_ld() uint8 {
	SetA(GetL())

	return 1
}

// ld   A, A - A=A
func x7F_ld() uint8 {
	SetA(GetA())

	return 1
}

// ld   A, %1 - A=%1
func x3E_ld() uint8 {
	SetA(FetchOperand8())

	return 2
}

// ld   B, %1 - B=%1
func x06_ld() uint8 {
	SetB(FetchOperand8())

	return 2
}

// ld   C, %1 - C=%1
func x0E_ld() uint8 {
	SetC(FetchOperand8())

	return 2
}

// ld   D, %1 - D=%1
func x16_ld() uint8 {
	SetD(FetchOperand8())

	return 2
}

// ld   E, %1 - E=%1
func x1E_ld() uint8 {
	SetE(FetchOperand8())

	return 2
}

// ld   H, %1 - H=%1
func x26_ld() uint8 {
	SetH(FetchOperand8())

	return 2
}

// ld   L, %1 - L=%1
func x2E_ld() uint8 {
	SetL(FetchOperand8())

	return 2
}

// ld   A, (HL) - A=(HL)
func x7E_ld() uint8 {
	SetA(Get(GetHL()))

	return 2
}

// ld   B, (HL) - B=(HL)
func x46_ld() uint8 {
	SetB(Get(GetHL()))

	return 2
}

// ld   C, (HL) - C=(HL)
func x4E_ld() uint8 {
	SetC(Get(GetHL()))

	return 2
}

// ld   D, (HL) - D=(HL)
func x56_ld() uint8 {
	SetD(Get(GetHL()))

	return 2
}

// ld   E, (HL) - E=(HL)
func x5E_ld() uint8 {
	SetE(Get(GetHL()))

	return 2
}

// ld   H, (HL) - H=(HL)
func x66_ld() uint8 {
	SetH(Get(GetHL()))

	return 2
}

// ld   L, (HL) - L=(HL)
func x6E_ld() uint8 {
	SetL(Get(GetHL()))

	return 2
}

// ld   (HL), B - [HL]=B
func x70_ld() uint8 {
	Set(GetHL(), GetB())

	return 2
}

// ld   (HL), C - [HL]=C
func x71_ld() uint8 {
	Set(GetHL(), GetC())

	return 2
}

// ld   (HL), D - [HL]=D
func x72_ld() uint8 {
	Set(GetHL(), GetD())

	return 2
}

// ld   (HL), E - [HL]=E
func x73_ld() uint8 {
	Set(GetHL(), GetE())

	return 2
}

// ld   (HL), H - [HL]=H
func x74_ld() uint8 {
	Set(GetHL(), GetH())

	return 2
}

// ld   (HL), L - [HL]=L
func x75_ld() uint8 {
	Set(GetHL(), GetL())

	return 2
}

// ld   (HL), A - [HL]=A
func x77_ld() uint8 {
	Set(GetHL(), GetA())

	return 2
}

// ld   (HL), %1 - [HL]=[NN]
func x36_ld() uint8 {
	Set(GetHL(), FetchOperand8())

	return 3
}

// ld   A, (BC) - A=[BC]
func x0A_ld() uint8 {
	SetA(Get(GetBC()))

	return 2
}

// ld   A, (DE) - A=[DE]
func x1A_ld() uint8 {
	SetA(Get(GetDE()))

	return 2
}

// ld   A, (%2) - A=[NN]
func xFA_ld() uint8 {
	SetA(Get(FetchOperand16()))

	return 4
}

// ld   (BC), A - [BC]=A
func x02_ld() uint8 {
	Set(GetBC(), GetA())

	return 2
}

// ld   (DE), A - [DE]=A
func x12_ld() uint8 {
	Set(GetDE(), GetA())

	return 2
}

// ld   (%2), A - [NN]=A
func xEA_ld() uint8 {
	Set(FetchOperand16(), GetA())

	return 4
}

// ld   A, (FF00+%1) - *read from io-port %1 (memory FF00+%1)
func xF0_ld() uint8 {
	// parsing error

	return 3
}

// ld   (FF00+%1), A - *write to io-port %1 (memory FF00+%1)
func xE0_ld() uint8 {
	// parsing error

	return 3
}

// ld   A, (FF00+C) - *read from io-port C (memory FF00+C)
func xF2_ld() uint8 {
	// parsing error

	return 2
}

// ld   (FF00+C), A - *write to io-port C (memory FF00+C)
func xE2_ld() uint8 {
	// parsing error

	return 2
}

// ld   BC, %2 - BC=%2
func x01_ld() uint8 {
	SetBC(FetchOperand16())

	return 3
}

// ld   DE, %2 - DE=%2
func x11_ld() uint8 {
	SetDE(FetchOperand16())

	return 3
}

// ld   HL, %2 - HL=%2
func x21_ld() uint8 {
	SetHL(FetchOperand16())

	return 3
}

// ld   SP, %2 - SP=%2
func x31_ld() uint8 {
	SetSP(FetchOperand16())

	return 3
}

// ld   SP, HL - SP=HL
func xF9_ld() uint8 {
	SetSP(GetHL())

	return 2
}

// ld   HL, SP+%s - *HL = SP +/- %s ;%s is 8bit signed number
func xF8_ld() uint8 {
	SetHL(GetSP())

	return 3
}

// unsupported mnemonic: jr
// unsupported mnemonic: add
// unsupported mnemonic: rrc
// unsupported mnemonic: nop
// unsupported mnemonic: call
// unsupported mnemonic: rl
// unsupported mnemonic: stop
// unsupported mnemonic: inc
// unsupported mnemonic: srl
// unsupported mnemonic: ldi
// unsupported mnemonic: rra
// unsupported mnemonic: push
// unsupported mnemonic: halt
// unsupported mnemonic: cp
// unsupported mnemonic: ret
// unsupported mnemonic: ccf
// unsupported mnemonic: sla
// unsupported mnemonic: sub
// unsupported mnemonic: and
// unsupported mnemonic: swap
// unsupported mnemonic: rr
// unsupported mnemonic: adc
// unsupported mnemonic: rla
// unsupported mnemonic: sbc
// unsupported mnemonic: cpl
// unsupported mnemonic: rrca
// unsupported mnemonic: rst
// unsupported mnemonic: di
// unsupported mnemonic: set
// unsupported mnemonic: res
// unsupported mnemonic: scf
// unsupported mnemonic: daa
// unsupported mnemonic: or
// unsupported mnemonic: ei
// unsupported mnemonic: rlc
// unsupported mnemonic: rlca
// unsupported mnemonic: xor
// unsupported mnemonic: jp
// unsupported mnemonic: ldd
// unsupported mnemonic: sra
// unsupported mnemonic: pop
// unsupported mnemonic: bit
// unsupported mnemonic: reti
// unsupported mnemonic: dec
