package cpu

import "testing"

func TestRlc(t* testing.T) {
	SetF(0x00)
	SetH(0x00)
	xCB_04_rlc()
	CheckRegister(t, REG_H, 0x00)
	testFlags(t, true, false, false, false)

	SetF(0x00)
	SetH(0x80)
	xCB_04_rlc()
	CheckRegister(t, REG_H, 0x01)
	testFlags(t, false, false, false, true)

	SetF(0x00)
	SetH(0x73)
	xCB_04_rlc()
	CheckRegister(t, REG_H, 0xE6)
	testFlags(t, false, false, false, false)
}