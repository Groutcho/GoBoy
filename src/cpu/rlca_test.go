package cpu

import "testing"

func TestRlca(t* testing.T) {
	SetF(0x00)
	SetA(0x00)
	x07_rlca()
	testRegister(t, REG_A, 0x00)
	testFlags(t, false, false, false, false)

	SetF(0x00)
	SetA(0xFF)
	x07_rlca()
	testRegister(t, REG_A, 0xFE)
	testFlags(t, false, false, false, true)
}