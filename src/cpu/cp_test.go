package cpu

import (
	"memory"
	"testing"
)

func TestCp(t *testing.T) {
	ResetSystem()

	A = 0x80
	E = 0x12
	F = 0x00
	xBB_cp()
	testFlags(t, false, true, true, false)

	A = 0x80
	L = 0x80
	F = 0x00
	xBD_cp()
	testFlags(t, true, true, false, false)

	A = 0x55
	L = 0x05
	F = 0x00
	xBD_cp()
	testFlags(t, false, true, false, false)

	A = 0x55
	L = 0xFF
	F = 0x00
	xBD_cp()
	testFlags(t, false, true, true, true)

	A = 0x23
	PC = 0
	memory.Write(0, 0x23)
	xFE_cp()
	testFlags(t, true, true, false, false)

	A = 0x23
	PC = 0
	memory.Write(0, 0x28)
	xFE_cp()
	testFlags(t, true, true, true, false)
}
