package cpu

import "testing"

func TestSla(t* testing.T) {
	SetE(0x82) // 1000 0010
	xCB_23_sla()
	testRegister(t, REG_E, 0x4) // 0000 0100
	testFlags(t, false, false, false, true)

	SetE(0xFF) // 1111 1111
	xCB_23_sla()
	testRegister(t, REG_E, 0xFE) // 1111 1110
	testFlags(t, false, false, false, true)

	SetE(0x7F) // 0111 1111
	xCB_23_sla()
	testRegister(t, REG_E, 0xFE) // 1111 1110
	testFlags(t, false, false, false, false)

	SetE(0x00)
	xCB_23_sla()
	testRegister(t, REG_E, 0x00)
	testFlags(t, true, false, false, false)

	SetE(0x80) // 1000 0000
	xCB_23_sla()
	testRegister(t, REG_E, 0x00)
	testFlags(t, true, false, false, true)
}