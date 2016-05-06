package cpu 

import "testing"
import . "memory"

func TestC9_ret(t* testing.T) {
	SetSP(0x0005)
	Set(0x0005, 0x34)
	Set(0x0006, 0x12)

	xC9_ret()

	testRegister(t, REG_PC, 0x1234)
}

func TestD8_ret(t* testing.T) {
	SetSP(0x0005)
	SetPC(0x0000)
	Set(0x0005, 0x34)
	Set(0x0006, 0x12)

	SetFlagCy(false)
	xD8_ret()
	testRegister(t, REG_PC, 0x0000)

	SetFlagCy(true)
	xD8_ret()
	testRegister(t, REG_PC, 0x1234)
}

func TestD9_reti(t* testing.T) {
	DisableAllInterrupts()

	SetSP(0x0005)
	SetPC(0x0000)
	Set(0x0005, 0x34)
	Set(0x0006, 0x12)

	xD8_ret()
	testRegister(t, REG_PC, 0x1234)

	if 	VBlankInterruptEnabled() &&
		LcdStatInterruptEnabled() &&
		TimerInterruptEnabled() &&
		SerialInterruptEnabled() &&
		JoypadInterruptEnabled() {
			t.Error("TestD9_reti() failed: All interrupts should be enabled.")
		}
}