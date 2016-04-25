package registers

import "testing"

func TestSetGetAF(t *testing.T) {
	Reset()

	var target uint16 = 0x00FF

	SetAF(target)

	if af := GetAF(); af != target {
		t.Errorf("GetAF() test failed: expected: Ox%04x, got: Ox%04x", target, af)
	}
}

func TestSetGetA(t *testing.T) {
	Reset()

	var target uint8 = 0xFF

	SetA(target)

	if a := GetA(); a != target {
		t.Errorf("GetA() test failed: expected: Ox%02x, got: Ox%02x (AF: 0x%04x)", target, a, registers.AF)
	}
}

func TestSetGetF(t *testing.T) {
	Reset()

	var target uint8 = 0x5C

	SetF(target)

	if f := GetF(); f != target {
		t.Errorf("GetF() test failed: expected: Ox%02x, got: Ox%02x (AF: 0x%04x)", target, f, registers.AF)
	}
}

func TestGetFlagZf(t *testing.T) {
	Reset()
	SetF(0xFF) // 11111111
	if set := GetFlagZf(); !set {
		t.Errorf("GetFlagZf() test failed: Zero flag should be set. (AF: 0x%04x)", registers.AF)
	}

	Reset()
	SetF(0x5c) // 01011100
	if set:= GetFlagZf(); set {
		t.Errorf("GetFlagZf() test failed: Zero flag should not be set. (AF: 0x%04x)", registers.AF)
	}
}