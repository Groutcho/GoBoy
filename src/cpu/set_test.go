package cpu

import "testing"
import . "memory"

func TestCB_DD_set(t* testing.T) {
	SetL(0x00)
	xCB_DD_set()

	if set := GetBit(GetL(), 3); set != 1 {
		t.Error("TestCB_DD_set() failed: expected L(3) set.")
	}
}

// set  4, (HL) - set bit 4 of [HL]
func TestCB_E6_set(t* testing.T) {
	SetHL(0x0050)
	Set(0x0050, 0x00)
	xCB_E6_set()
	if set := GetBit(Get(0x0050), 4); set != 1 {
		t.Error("TestCB_DD_set() failed: expected [HL](4) set.")		
	}
}