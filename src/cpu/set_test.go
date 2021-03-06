package cpu

import "testing"
import . "memory"
import . "common"

func TestCB_DD_set(t *testing.T) {
	ResetSystem()

	L = 0x00
	xCB_DD_set()

	if set := GetBit(L, 3); set != 1 {
		t.Error("TestCB_DD_set() failed: expected L(3) set.")
	}
}

// set  4, (HL) - set bit 4 of [HL]
func TestCB_E6_set(t *testing.T) {
	ResetSystem()

	SetHL(0x0050)
	Write(0x0050, 0x00)
	xCB_E6_set()
	if set := GetBit(Get(0x0050), 4); set != 1 {
		t.Error("TestCB_DD_set() failed: expected [HL](4) set.")
	}
}
