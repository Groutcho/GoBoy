package cpu

import "testing"

// D = D-1
func Test15_dec(t* testing.T) {
	SetD(0x15)
	x15_dec()
	CheckRegister(t, REG_D, 0x14)

	SetD(0x00)
	x15_dec()
	CheckRegister(t, REG_D, 0xFF)
	testFlags(t, false, true, true, false)
}