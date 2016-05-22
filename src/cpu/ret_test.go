package cpu

import "testing"
import . "memory"

func TestC9_ret(t *testing.T) {
	ResetSystem()

	SP = 0x0005
	Write16(0x0005, 0x1234)

	xC9_ret()

	CheckRegister(t, REG_PC, 0x1234)
}

func TestD8_ret(t *testing.T) {
	ResetSystem()

	SP = 0x0005
	Write16(0x0005, 0x1234)

	SetFlagCy(false)
	xD8_ret()
	CheckRegister(t, REG_PC, 0x0000)

	SetFlagCy(true)
	xD8_ret()
	CheckRegister(t, REG_PC, 0x1234)
}

func TestD9_reti(t *testing.T) {
	ResetSystem()

	SetIME(false)

	SP = 0x0005
	Write16(0x0005, 0x1234)

	xD9_reti()
	CheckRegister(t, REG_PC, 0x1234)

	if !GetIME() {
		t.Error("TestD9_reti() failed: IME should be set.")
	}
}
