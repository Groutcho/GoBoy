package cpu

import "testing"
import . "memory"

// AF=(SP),  SP=SP+2
func TestF1_pop(t* testing.T) {
	ResetMemory()

	SetAF(0x0000)
	SetSP(0x0020)

	Set(0x0020, 0xF8)
	Set(0x0021, 0xAB)

	t.Log(DumpRange(0x0015, 0x0024))
	t.Log("15")
	
	xF1_pop()

	CheckRegister(t, REG_SP, 0x0022)
	CheckRegister(t, REG_AF, 0xABF8)
}