package cpu

import "testing"
import . "memory"

// AF=(SP),  SP=SP+2
func TestF1_pop(t *testing.T) {
	ResetMemory()

	SetAF(0x0000)
	SetSP(0x0020)

	Write(0x0020, 0xF8)
	Write(0x0021, 0xAB)

	xF1_pop()

	CheckRegister(t, REG_SP, 0x0022)
	CheckRegister(t, REG_AF, 0xABF8)
}
