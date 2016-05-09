package cpu

import "testing"
import . "memory"

// inc  H - H = H+1
func Test24_inc(t* testing.T) {
	SetH(0x00)
	x24_inc()

	CheckRegister(t, REG_H, 0x01)
	testFlags(t, false, false, false, false)
}

func Test34_inc(t* testing.T) {
	SetHL(0x0025)
	Set(0x0025, 0x88)
	x34_inc()

	CheckRegister(t, REG_HL, 0x25)

	if Get(GetHL()) != 0x89 {
		t.Errorf("Test34_inc() failed: expected [HL] = 0x89, got 0x%02X", Get(GetHL()))
	}
}