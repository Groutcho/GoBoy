package cpu

import "testing"

func TestSrl(t* testing.T) {
	SetC(0x00)
	xCB_39_srl()
	CheckRegister(t, REG_C, 0x00)
	testFlags(t, true, false, false, false)

	SetC(0x80)
	xCB_39_srl()
	CheckRegister(t, REG_C, 0x40)
	testFlags(t, false, false, false, false)

	SetC(0x80)
	xCB_39_srl()
	CheckRegister(t, REG_C, 0x40)
	testFlags(t, false, false, false, false)

	SetC(0xFF)
	xCB_39_srl()
	CheckRegister(t, REG_C, 0x7F)
	testFlags(t, false, false, false, true)
}