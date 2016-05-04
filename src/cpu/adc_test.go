package cpu

import "testing"
import . "memory"

func Test89_adc(t* testing.T) {
	SetF(0x00)
	SetA(0xF0)
	SetC(0x0F)
	x89_adc()

	if a:= GetA(); a != 0xFF {
		t.Errorf("Test89_adc() failed: expected A = 0xFF, got 0x%02x", a)
	}

	testFlags(t, false, false, false, false)

	SetF(0x00)
	SetA(0xF1)
	SetC(0x0F)
	x89_adc()

	if a:= GetA(); a != 0x00 {
		t.Errorf("Test89_adc() failed: expected A = 0x00, got 0x%02x", a)
	}

	testFlags(t, false, false, true, true)

	SetF(0x00)
	SetA(0x00)
	SetC(0x00)
	x89_adc()

	if a:= GetA(); a != 0x00 {
		t.Errorf("Test89_adc() failed: expected A = 0x00, got 0x%02x", a)
	}

	testFlags(t, true, false, false, false)
}

func Test8E_adc(t* testing.T) {
	SetF(0x00)
	SetA(0x20)
	SetHL(0x15)
	Set(0x15, 0x88)

	x8E_adc()

	if a:= GetA(); a != 0xA8 {
		t.Errorf("Test8E_adc() failed: expected A = 0xA8, got 0x%02x", a)
	}

	testFlags(t, false, false, false, false)

	SetF(0x00)
	SetA(0x21)
	SetHL(0x15)
	Set(0x15, 0xFF)

	x8E_adc()

	if a:= GetA(); a != 0x20 {
		t.Errorf("Test8E_adc() failed: expected A = 0x20, got 0x%02x", a)
	}

	testFlags(t, false, false, true, true)
}

func TestCE_adc(t* testing.T) {
	SetF(0x00)
	SetA(0x20)
	SetPC(0x0000)
	Set(0x0000, 0x88)

	xCE_adc()

	if a:= GetA(); a != 0xA8 {
		t.Errorf("TestCE_adc() failed: expected A = 0xA8, got 0x%02x", a)
	}

	testFlags(t, false, false, false, false)
}