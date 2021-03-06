package cpu

import "testing"
import . "memory"

func TestInterrupt(t *testing.T) {
	ResetSystem()

	SetPC(0x5555)
	Write16(0x0060, 0x1234)
	DisableJoypadInterrupt()

	SetIME(false)
	RequestJoypadInterrupt()
	CheckRegister(t, REG_PC, 0x5555)

	SetIME(true)
	RequestJoypadInterrupt()
	CheckRegister(t, REG_PC, 0x5555)

	EnableJoypadInterrupt()
	RequestJoypadInterrupt()
	CheckRegister(t, REG_PC, 0x1234)
}
