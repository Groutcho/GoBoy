package cpu

import "testing"

func TestCp(t* testing.T) {
	SetA(0x80)
	SetE(0x12)
	SetF(0x00)
	xBB_cp()
	testFlags(t, false, true, true, false)

	SetA(0x80)
	SetL(0x80)
	SetF(0x00)
	xBD_cp()
	testFlags(t, true, true, false, false)

	SetA(0x55)
	SetL(0x05)
	SetF(0x00)
	xBD_cp()
	testFlags(t, false, true, false, false)

	SetA(0x55)
	SetL(0xFF)
	SetF(0x00)
	xBD_cp()
	testFlags(t, false, true, true, true)
}