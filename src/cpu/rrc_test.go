package cpu

import "testing"

func TestRrc(t* testing.T) {
	SetF(0x00)
	SetB(0x00)
	xCB_08_rrc()
	testRegister(t, REG_B, 0x00)
	testFlags(t, true, false, false, false)

	SetF(0x00)
	SetB(0x80) // 1000 0000
	xCB_08_rrc()
	testRegister(t, REG_B, 0x40)
	testFlags(t, false, false, false, false)

	SetF(0x00)
	SetB(0xFF) // 1000 0000
	xCB_08_rrc()
	testRegister(t, REG_B, 0xFF)
	testFlags(t, false, false, false, true)

	SetF(0x00)
	SetB(0x01)
	xCB_08_rrc()
	testRegister(t, REG_B, 0x80)
	testFlags(t, false, false, false, true)
}