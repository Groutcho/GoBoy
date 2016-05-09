package cpu

import "testing"

func Test37_scf(t* testing.T) {
	SetFlagCy(false)
	x37_scf()

	if !GetFlagCy() {
		t.Error("Test37_scf() failed: expected CY set.")
	}
}

func Test27_daa(t* testing.T) {
	SetF(0x00)
	SetA(0x2A)
	x27_daa()
	CheckRegister(t, REG_A, 0x30)
	testFlags(t, false, false, true, false)

	SetF(0x00)
	SetFlagH(true)
	SetA(0x20)
	x27_daa()
	CheckRegister(t, REG_A, 0x26)
	testFlags(t, false, false, true, false)

	SetF(0x00)
	SetFlagCy(false)
	SetA(0xA0)
	x27_daa()
	CheckRegister(t, REG_A, 0x00)
	testFlags(t, false, false, false, true)

	SetF(0x00)
	SetFlagCy(true)
	SetA(0x20)
	x27_daa()
	CheckRegister(t, REG_A, 0x80)
	testFlags(t, false, false, false, true)
}