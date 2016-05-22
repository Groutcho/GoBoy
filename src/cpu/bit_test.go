package cpu

import "testing"
import . "memory"

// bit  1, (HL) - [HL] & {2^1}
func TestCB_4E_bit(t *testing.T) {
	ResetSystem()

	Write(0x0001, 0x02) // 0000 0010
	SetHL(0x0001)
	xCB_4E_bit()
	if GetFlagZf() {
		t.Error("TestCB_4E_bit() failed: Z flag should be unset.")
	}

	Write(0x0001, 0x01) // 0000 0001
	xCB_4E_bit()
	if !GetFlagZf() {
		t.Error("TestCB_4E_bit() failed: Z flag should be set.")
	}
}

// bit  3, (HL)
func TestCB_5E_bit(t *testing.T) {
	ResetSystem()

	Write(0x0001, 0x08) // 0000 1000
	SetHL(0x0001)
	xCB_5E_bit()
	if GetFlagZf() {
		t.Error("TestCB_5E_bit() failed: Z flag should be unset.")
	}

	Write(0x0001, 0x01) // 0000 0001
	xCB_5E_bit()
	if !GetFlagZf() {
		t.Error("TestCB_5E_bit() failed: Z flag should be set.")
	}
}

// bit  7, H
func TestCB_7C_bit(t *testing.T) {
	ResetSystem()

	SetHL(0x0000)
	xCB_7C_bit()
	if !GetFlagZf() {
		t.Error("TestCB_7C_bit() failed: Z flag should be set.")
	}

	SetHL(0x8000)
	xCB_7C_bit()
	if GetFlagZf() {
		t.Error("TestCB_7C_bit() failed: Z flag should be unset.")
	}
}
