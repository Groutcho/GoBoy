package cpu

import "testing"
import . "common"

func TestSetGetAF(t *testing.T) {
	ResetSystem()

	var target uint16 = 0x00FF

	SetAF(target)

	if af := GetAF(); af != target {
		t.Errorf("GetAF() test failed: expected: Ox%04x, got: Ox%04x", target, af)
	}
}

func TestSetGetA(t *testing.T) {
	ResetSystem()

	SetA(0xFF)
	if a := GetA(); a != 0xFF {
		t.Errorf("GetA() test failed: expected: 0xFF, got: Ox%02x", a, F)
	}

	SetA(0x28)
	if a := GetA(); a != 0x28 {
		t.Errorf("GetA() test failed: expected: 0x28, got: Ox%02x", a, F)
	}

	SetA(0x44)
	if a := GetA(); a != 0x44 {
		t.Errorf("GetA() test failed: expected: 0x44, got: Ox%02x", a, F)
	}

	SetA(0x99)
	if a := GetA(); a != 0x99 {
		t.Errorf("GetA() test failed: expected: 0x99, got: Ox%02x", a, F)
	}

	SetA(0xFC)
	if a := GetA(); a != 0xFC {
		t.Errorf("GetA() test failed: expected: 0xFC, got: Ox%02x", a, F)
	}
}

func TestSetGetF(t *testing.T) {
	ResetSystem()

	SetF(0x5C)
	if f := GetF(); f != 0x5C {
		t.Errorf("GetF() test failed: expected: 0x5C, got: Ox%02x", f, F)
	}

	SetF(0x8C)
	if f := GetF(); f != 0x8C {
		t.Errorf("GetF() test failed: expected: 0x8C, got: Ox%02x", f, F)
	}

	SetF(0xFF)
	if f := GetF(); f != 0xFF {
		t.Errorf("GetF() test failed: expected: 0xFF, got: Ox%02x", f, F)
	}

	SetF(0xBB)
	if f := GetF(); f != 0xBB {
		t.Errorf("GetF() test failed: expected: 0xBB, got: Ox%02x", f, F)
	}

	SetF(0x12)
	if f := GetF(); f != 0x12 {
		t.Errorf("GetF() test failed: expected: 0x12, got: Ox%02x", f, F)
	}
}

func TestGetFlagZf(t *testing.T) {
	ResetSystem()
	SetF(0xFF) // 11111111
	if set := GetFlagZf(); !set {
		t.Errorf("GetFlagZf() test failed: Zero flag should be set.", F)
	}

	ResetSystem()
	SetF(0x5c) // 01011100
	if set := GetFlagZf(); set {
		t.Errorf("GetFlagZf() test failed: Zero flag should not be set.", F)
	}

	ResetSystem()
	SetF(0x00)
	if set := GetFlagZf(); set {
		t.Errorf("GetFlagZf() test failed: Zero flag should not be set.", F)
	}
}

func TestGetFlagN(t *testing.T) {
	ResetSystem()
	SetF(0xFF) // 11111111
	if set := GetFlagN(); !set {
		t.Errorf("GetFlagN() test failed: N flag should be set.", F)
	}

	ResetSystem()
	SetF(0x5c) // 01011100
	if set := GetFlagN(); !set {
		t.Errorf("GetFlagN() test failed: N flag should be set.", F)
	}

	ResetSystem()
	SetF(0x35) // 00110101
	if set := GetFlagN(); set {
		t.Errorf("GetFlagN() test failed: N flag should not be set.", F)
	}

	ResetSystem()
	SetF(0x00)
	if set := GetFlagN(); set {
		t.Errorf("GetFlagN() test failed: N flag should not be set.", F)
	}
}

func TestGetFlagH(t *testing.T) {
	ResetSystem()
	SetF(0xFF) // 11111111
	if set := GetFlagH(); !set {
		t.Errorf("GetFlagH() test failed: N flag should be set.", F)
	}

	ResetSystem()
	SetF(0x3c) // 00111100
	if set := GetFlagH(); !set {
		t.Errorf("GetFlagH() test failed: N flag should be set.", F)
	}

	ResetSystem()
	SetF(0x15) // 00010101
	if set := GetFlagH(); set {
		t.Errorf("GetFlagH() test failed: N flag should not be set.", F)
	}

	ResetSystem()
	SetF(0x00)
	if set := GetFlagH(); set {
		t.Errorf("GetFlagH() test failed: N flag should not be set.", F)
	}
}

func TestGetFlagCy(t *testing.T) {
	ResetSystem()
	SetF(0xFF) // 11111111
	if set := GetFlagCy(); !set {
		t.Errorf("GetFlagCy() test failed: CY flag should be set.", F)
	}

	ResetSystem()
	SetF(0x3c) // 00111100
	if set := GetFlagCy(); !set {
		t.Errorf("GetFlagCy() test failed: CY flag should be set.", F)
	}

	ResetSystem()
	SetF(0x05) // 00000101
	if set := GetFlagCy(); set {
		t.Errorf("GetFlagCy() test failed: CY flag should not be set.", F)
	}

	ResetSystem()
	SetF(0x00)
	if set := GetFlagCy(); set {
		t.Errorf("GetFlagCy() test failed: CY flag should not be set.", F)
	}

	ResetSystem()
	SetFlagCy(true)
	if set := GetFlagCy(); !set {
		t.Errorf("GetFlagCy() test failed: CY flag should be set.", F)
	}

	SetFlagCy(false)
	if set := GetFlagCy(); set {
		t.Errorf("GetFlagCy() test failed: CY flag should not be set.", F)
	}
}

func TestIncPC(t *testing.T) {
	PC = 0x0025
	IncPC()

	if pc := GetPC(); pc != 0x0026 {
		t.Errorf("TestIncPC() failed: PC should be 0x0026, got 0x%04x", PC)
	}
}

func TestIncBC(t *testing.T) {
	SetBC(0x0025)
	IncBC()

	if bc := GetBC(); bc != 0x0026 {
		t.Errorf("TestIncBC() failed: BC should be 0x0026, got 0x%04x", bc)
	}
}

func TestIncDE(t *testing.T) {
	SetDE(0x0025)
	IncDE()

	if de := GetDE(); de != 0x0026 {
		t.Errorf("TestIncDE() failed: DE should be 0x0026, got 0x%04x", de)
	}
}

func TestIncHL(t *testing.T) {
	SetHL(0x0025)
	IncHL()

	if hl := GetHL(); hl != 0x0026 {
		t.Errorf("TestIncHL() failed: HL should be 0x0026, got 0x%04x", hl)
	}
}

func TestIncAF(t *testing.T) {
	SetAF(0x0025)
	IncAF()

	if af := GetAF(); af != 0x0026 {
		t.Errorf("TestIncAF() failed: AF should be 0x0026, got 0x%04x", af)
	}
}

func TestDecPC(t *testing.T) {
	PC = 0x0025
	DecPC()

	if pc := GetPC(); pc != 0x0024 {
		t.Errorf("TestDecPC() failed: PC should be 0x0024, got 0x%04x", PC)
	}
}

func TestDecBC(t *testing.T) {
	SetBC(0x0025)
	DecBC()

	if bc := GetBC(); bc != 0x0024 {
		t.Errorf("TestDecBC() failed: BC should be 0x0024, got 0x%04x", bc)
	}
}

func TestDecDE(t *testing.T) {
	SetDE(0x0025)
	DecDE()

	if de := GetDE(); de != 0x0024 {
		t.Errorf("TestDecDE() failed: DE should be 0x0024, got 0x%04x", de)
	}
}

func TestDecHL(t *testing.T) {
	SetHL(0x0025)
	DecHL()

	if hl := GetHL(); hl != 0x0024 {
		t.Errorf("TestDecHL() failed: HL should be 0x0024, got 0x%04x", hl)
	}
}

func TestDecAF(t *testing.T) {
	SetAF(0x0025)
	DecAF()

	if af := GetAF(); af != 0x0024 {
		t.Errorf("TestDecAF() failed: AF should be 0x0024, got 0x%04x", af)
	}
}

func TestSetBit(t *testing.T) {

	values0x00 := [...]uint8{
		0x01,
		0x02,
		0x04,
		0x08,
		0x10,
		0x20,
		0x40,
		0x80,
	}

	for i := 0; i < len(values0x00); i += 2 {
		if actual := SetBit(0x00, uint8(i), 1); actual != values0x00[i] {
			t.Errorf("TestSetBit() failed: expected 0x%02x, got 0x%02x", values0x00[i], actual)
		}
	}

	values0xFF := [...]uint8{
		0xFE,
		0xFD,
		0xFB,
		0xF7,
		0xEF,
		0xDF,
		0xBF,
		0x7F,
	}

	for i := 0; i < len(values0xFF); i += 2 {
		if actual := SetBit(0xFF, uint8(i), 0); actual != values0xFF[i] {
			t.Errorf("TestSetBit() failed: expected 0x%02x, got 0x%02x", values0xFF[i], actual)
		}
	}
}

func TestGetBit(t *testing.T) {
	if set := GetBit(0x0E, 3); /* 00001110 */ set != 1 {
		t.Error("TestGetBit() failed: third bit of 0x0E should be set.")
	}
	if set := GetBit(0x0E, 5); /* 00001110 */ set != 0 {
		t.Error("TestGetBit() failed: third bit of 0x0E should be unset.")
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

	for i := 0; i < len(values); i += 2 {
		if actual := Swap(values[i]); actual != values[i+1] {
			t.Errorf("TestSwap() failed: expected 0x%02x, got 0x%02x", 0xEF, actual)
		}
	}
}
