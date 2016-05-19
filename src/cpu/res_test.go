package cpu

import "testing"
import . "memory"
import . "common"

// res  3, C - reset bit 3 of register C
func TestCB_99_res(t *testing.T) {
	SetF(0x00)
	SetC(0xFF)
	xCB_99_res()

	if set := GetBit(GetC(), 3); set != 0 {
		t.Error("TestCB_99_res() failed: expected C(3) unset.")
	}
	testFlags(t, false, false, false, false)
}

// res  0, (HL) - reset bit 0 of [HL]
func TestCB_86_res(t *testing.T) {
	SetF(0x00)
	SetHL(0x0050)
	Write(0x0050, 0xFF)
	xCB_86_res()

	if set := GetBit(Get(0x0050), 0); set != 0 {
		t.Error("TestCB_DD_set() failed: expected [HL](4) unset.")
	}

	testFlags(t, false, false, false, false)
}
