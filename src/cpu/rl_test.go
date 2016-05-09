package cpu

import "testing"

func TestRl(t* testing.T) {
	SetF(0x00)
	SetL(0x00)
	xCB_15_rl()
	CheckRegister(t, REG_L, 0x00)
	testFlags(t, true, false, false, false)

	SetF(0x00)
	SetL(0x00)
	SetFlagCy(true)
	xCB_15_rl()
	CheckRegister(t, REG_L, 0x01)
	testFlags(t, false, false, false, false)

	SetF(0x00)
	SetL(0x80)
	SetFlagCy(true)
	xCB_15_rl()
	CheckRegister(t, REG_L, 0x01)
	testFlags(t, false, false, false, true)

	SetF(0x00)
	SetL(0x80)
	SetFlagCy(false)
	xCB_15_rl()
	CheckRegister(t, REG_L, 0x00)
	testFlags(t, true, false, false, true)
}