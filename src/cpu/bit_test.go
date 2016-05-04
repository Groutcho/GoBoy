package cpu

import "testing"
import . "memory"

// bit  1, (HL) - [HL] & {2^1}
func TestCB_4E_bit(t* testing.T) {
	Set(0x0001, 0x02) // 0000 0010
	SetHL(0x0001)
	xCB_4E_bit()
	if GetFlagZf() {
		t.Error("TestCB_4E_bit() failed: Z flag should be unset.")
	}

	Set(0x0001, 0x01) // 0000 0001
	xCB_4E_bit()
	if !GetFlagZf() {
		t.Error("TestCB_4E_bit() failed: Z flag should be set.")
	}
}

// bit  3, (HL) 
func TestCB_5E_bit(t* testing.T) {
	Set(0x0001, 0x08) // 0000 1000
	SetHL(0x0001)
	xCB_5E_bit()
	if GetFlagZf() {
		t.Error("TestCB_5E_bit() failed: Z flag should be unset.")
	}

	Set(0x0001, 0x01) // 0000 0001
	xCB_5E_bit()
	if !GetFlagZf() {
		t.Error("TestCB_5E_bit() failed: Z flag should be set.")
	}
}