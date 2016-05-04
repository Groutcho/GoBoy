package cpu

import "testing"

func Test80_add(t* testing.T) {
	SetF(0x00)
	SetA(0x05)
	SetB(0x06)
	x80_add()

	if a:= GetA(); a != 0x05 + 0x06 {
		t.Errorf("Test80_add() failed: expected A = 0x0A, got 0x%02x", a)
	}

	testFlags(t, false, false, false, false)

	SetF(0x00)
	SetA(0xFF)
	SetB(0x05)
	x80_add()

	if a:= GetA(); a != 0x04 {
		t.Errorf("Test80_add() failed: expected A = 0x04, got 0x%02x", a)
	}

	testFlags(t, false, false, true, true)

	SetF(0x00)
	SetA(0x00)
	SetB(0x00)
	x80_add()

	if a:= GetA(); a != 0x00 {
		t.Errorf("Test80_add() failed: expected A = 0x00, got 0x%02x", a)
	}

	testFlags(t, true, false, false, false)
}

// add  HL, DE - HL = HL+DE
func Test19_add(t* testing.T) {
	SetF(0x00)
	SetHL(0x0FDE)
	SetDE(0x21FF)
	x19_add()

	if hl:= GetHL(); hl != 0x31DD {
		t.Errorf("Test19_add() failed: expected HL = 0x31DD, got 0x%04x", hl)
	}

	testFlags(t, false, false, true, false)

	SetF(0x00)
	SetHL(0xFFDE)
	SetDE(0x21FF)
	x19_add()

	if hl:= GetHL(); hl != 0x21DD {
		t.Errorf("Test19_add() failed: expected HL = 0x21DD, got 0x%04x", hl)
	}

	testFlags(t, false, false, true, true)
}

	