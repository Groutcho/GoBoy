package cpu

import "testing"
import . "memory"

func Test22_ldi(t *testing.T) {
	ResetSystem()

	SetA(0x89)
	SetHL(0x8789)
	x22_ldi()

	if hl := Get(GetHL() - 1); hl != 0x89 {
		t.Errorf("Test22_ldi() failed: expected [HL-1] = 0x89, got 0x%02X", hl)
	}

	CheckRegister(t, REG_HL, 0x878A)
	CheckRegister(t, REG_A, 0x89)
}
