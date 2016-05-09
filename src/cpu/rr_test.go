package cpu

import "testing"

func TestRr(t* testing.T) {
	SetF(0x00)
	SetD(0x00)
	xCB_1A_rr()
	CheckRegister(t, REG_D, 0x00)
	testFlags(t, true, false, false, false)

	SetF(0x00)
	SetFlagCy(true)
	SetD(0x00)
	xCB_1A_rr()
	CheckRegister(t, REG_D, 0x80)
	testFlags(t, false, false, false, false)

	SetF(0x00)
	SetFlagCy(true)
	SetD(0x01)
	xCB_1A_rr()
	CheckRegister(t, REG_D, 0x80)
	testFlags(t, false, false, false, true)

	SetF(0x00)
	SetFlagCy(false)
	SetD(0xFF)
	xCB_1A_rr()
	CheckRegister(t, REG_D, 0x7F)
	testFlags(t, false, false, false, true)

	SetF(0x00)
	SetFlagCy(true)
	SetD(0xFF)
	xCB_1A_rr()
	CheckRegister(t, REG_D, 0xFF)
	testFlags(t, false, false, false, true)
}