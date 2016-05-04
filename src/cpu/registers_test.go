package cpu

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

	SetA(0xFF)
	if a := GetA(); a != 0xFF {
		t.Errorf("GetA() test failed: expected: 0xFF, got: Ox%02x", a, registers.F)
	}

	SetA(0x28)
	if a := GetA(); a != 0x28 {
		t.Errorf("GetA() test failed: expected: 0x28, got: Ox%02x", a, registers.F)
	}

	SetA(0x44)
	if a := GetA(); a != 0x44 {
		t.Errorf("GetA() test failed: expected: 0x44, got: Ox%02x", a, registers.F)
	}

	SetA(0x99)
	if a := GetA(); a != 0x99 {
		t.Errorf("GetA() test failed: expected: 0x99, got: Ox%02x", a, registers.F)
	}

	SetA(0xFC)
	if a := GetA(); a != 0xFC {
		t.Errorf("GetA() test failed: expected: 0xFC, got: Ox%02x", a, registers.F)
	}
}

func TestSetGetF(t *testing.T) {
	Reset()

	SetF(0x5C)
	if f := GetF(); f != 0x5C {
		t.Errorf("GetF() test failed: expected: 0x5C, got: Ox%02x", f, registers.F)
	}

	SetF(0x8C)
	if f := GetF(); f != 0x8C {
		t.Errorf("GetF() test failed: expected: 0x8C, got: Ox%02x", f, registers.F)
	}

	SetF(0xFF)
	if f := GetF(); f != 0xFF {
		t.Errorf("GetF() test failed: expected: 0xFF, got: Ox%02x", f, registers.F)
	}

	SetF(0xBB)
	if f := GetF(); f != 0xBB {
		t.Errorf("GetF() test failed: expected: 0xBB, got: Ox%02x", f, registers.F)
	}

	SetF(0x12)
	if f := GetF(); f != 0x12 {
		t.Errorf("GetF() test failed: expected: 0x12, got: Ox%02x", f, registers.F)
	}
}

func TestGetFlagZf(t *testing.T) {
	Reset()
	SetF(0xFF) // 11111111
	if set := GetFlagZf(); !set {
		t.Errorf("GetFlagZf() test failed: Zero flag should be set.", registers.F)
	}

	Reset()
	SetF(0x5c) // 01011100
	if set:= GetFlagZf(); set {
		t.Errorf("GetFlagZf() test failed: Zero flag should not be set.", registers.F)
	}

	Reset()
	SetF(0x00)
	if set:= GetFlagZf(); set {
		t.Errorf("GetFlagZf() test failed: Zero flag should not be set.", registers.F)
	}
}

func TestGetFlagN(t *testing.T) {
	Reset()
	SetF(0xFF) // 11111111
	if set := GetFlagN(); !set {
		t.Errorf("GetFlagN() test failed: N flag should be set.", registers.F)
	}

	Reset()
	SetF(0x5c) // 01011100
	if set:= GetFlagN(); !set {
		t.Errorf("GetFlagN() test failed: N flag should be set.", registers.F)
	}

	Reset()
	SetF(0x35) // 00110101
	if set:= GetFlagN(); set {
		t.Errorf("GetFlagN() test failed: N flag should not be set.", registers.F)
	}

	Reset()
	SetF(0x00)
	if set:= GetFlagN(); set {
		t.Errorf("GetFlagN() test failed: N flag should not be set.", registers.F)
	}
}

func TestGetFlagH(t *testing.T) {
	Reset()
	SetF(0xFF) // 11111111
	if set := GetFlagH(); !set {
		t.Errorf("GetFlagH() test failed: N flag should be set.", registers.F)
	}

	Reset()
	SetF(0x3c) // 00111100
	if set:= GetFlagH(); !set {
		t.Errorf("GetFlagH() test failed: N flag should be set.", registers.F)
	}

	Reset()
	SetF(0x15) // 00010101
	if set:= GetFlagH(); set {
		t.Errorf("GetFlagH() test failed: N flag should not be set.", registers.F)
	}

	Reset()
	SetF(0x00)
	if set:= GetFlagH(); set {
		t.Errorf("GetFlagH() test failed: N flag should not be set.", registers.F)
	}
}

func TestGetFlagCy(t *testing.T) {
	Reset()
	SetF(0xFF) // 11111111
	if set := GetFlagCy(); !set {
		t.Errorf("GetFlagCy() test failed: CY flag should be set.", registers.F)
	}

	Reset()
	SetF(0x3c) // 00111100
	if set:= GetFlagCy(); !set {
		t.Errorf("GetFlagCy() test failed: CY flag should be set.", registers.F)
	}

	Reset()
	SetF(0x05) // 00000101
	if set:= GetFlagCy(); set {
		t.Errorf("GetFlagCy() test failed: CY flag should not be set.", registers.F)
	}

	Reset()
	SetF(0x00)
	if set:= GetFlagCy(); set {
		t.Errorf("GetFlagCy() test failed: CY flag should not be set.", registers.F)
	}

	Reset()
	SetFlagCy(true)
	if set := GetFlagCy(); !set {
		t.Errorf("GetFlagCy() test failed: CY flag should be set.", registers.F)
	}

	SetFlagCy(false)
	if set := GetFlagCy(); set {
		t.Errorf("GetFlagCy() test failed: CY flag should not be set.", registers.F)
	}
}

func TestIncPC(t *testing.T) {
	registers.PC = 0x0025
	IncPC()

	if pc := GetPC(); pc != 0x0026 {
		t.Errorf("TestIncPC() failed: PC should be 0x0026, got 0x%04x", registers.PC)
	}
}

func TestSwap(t *testing.T) {

	values := [...]uint8{
		0xFE, 0xEF,
		0x00, 0x00,
		0xFF, 0xFF,
		0x58, 0x85,
		0x1F, 0xF1,
	}

	for i := 0; i < len(values); i+=2 {
		if actual := Swap(values[i]); actual != values[i+1] {
			t.Errorf("TestSwap() failed: expected 0x%02x, got 0x%02x", 0xEF, actual)
		}
	}
}