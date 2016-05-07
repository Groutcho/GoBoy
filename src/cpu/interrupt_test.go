package cpu

import "testing"
import . "memory"

func TestInterrupt(t* testing.T) {
	ResetMemory()
	SetPC(0x5555)
	Set16(0x0060, 0x1234)
	DisableJoypadInterrupt()

	SetIME(false)
	RequestJoypadInterrupt()
	testRegister(t, REG_PC, 0x5555)
	
	SetIME(true)
	RequestJoypadInterrupt()
	testRegister(t, REG_PC, 0x5555)

	EnableJoypadInterrupt()
	RequestJoypadInterrupt()
	testRegister(t, REG_PC, 0x1234)
}