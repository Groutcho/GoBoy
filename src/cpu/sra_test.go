package cpu

import "testing"

func TestSra(t *testing.T) {
	ResetSystem()

	SetB(0x00)
	xCB_28_sra()
	CheckRegister(t, REG_B, 0x00)
	testFlags(t, true, false, false, false)

	SetB(0x80) // 1000 0000
	xCB_28_sra()
	CheckRegister(t, REG_B, 0xC0)
	testFlags(t, false, false, false, false)

	SetB(0x81) // 1000 0001
	xCB_28_sra()
	CheckRegister(t, REG_B, 0xC0)
	testFlags(t, false, false, false, true)

	SetB(0xFE) // 1111 1110
	xCB_28_sra()
	CheckRegister(t, REG_B, 0xFF)
	testFlags(t, false, false, false, false)
}
