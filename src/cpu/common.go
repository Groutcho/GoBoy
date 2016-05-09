package cpu

import "testing"
import . "memory"

const REG_A = 0
const REG_F = 1
const REG_AF = 2
const REG_B = 3
const REG_C = 4
const REG_BC = 5
const REG_D = 6
const REG_E = 7
const REG_DE = 8
const REG_H = 9
const REG_L = 10
const REG_HL = 11
const REG_SP = 12
const REG_PC = 13

func testFlags(t* testing.T, Z bool, N bool, H bool, C bool) {
	if GetFlagZf() != Z {
		if Z {
			t.Error("TestFlags() failed: expected Z set")
		} else {
			t.Error("TestFlags() failed: expected Z unset")
		}
	}

	if GetFlagN() != N {
		if N {
			t.Error("TestFlags() failed: expected N set")
		} else {
			t.Error("TestFlags() failed: expected N unset")
		}
	}

	if GetFlagH() != H {
		if H {
			t.Error("TestFlags() failed: expected H set")
		} else {
			t.Error("TestFlags() failed: expected H unset")
		}
	}

	if GetFlagCy() != C {
		if C {
			t.Error("TestFlags() failed: expected C set")
		} else {
			t.Error("TestFlags() failed: expected C unset")
		}
	}
}

func CheckRegister(t* testing.T, registerCode int, expected int) {

	name := ""
	actual := 0

	switch registerCode {
		case REG_A: name = "A"; actual = int(GetA())
		case REG_B: name = "B"; actual = int(GetB())
		case REG_C: name = "C"; actual = int(GetC())
		case REG_D: name = "D"; actual = int(GetD())
		case REG_E: name = "E"; actual = int(GetE())
		case REG_F: name = "F"; actual = int(GetF())
		case REG_H: name = "H"; actual = int(GetH())
		case REG_L: name = "L"; actual = int(GetL())
		case REG_AF: name = "AF"; actual = int(GetAF())
		case REG_BC: name = "BC"; actual = int(GetBC())
		case REG_DE: name = "DE"; actual = int(GetDE())
		case REG_HL: name = "HL"; actual = int(GetHL())
		case REG_SP: name = "SP"; actual = int(GetSP())
		case REG_PC: name = "PC"; actual = int(GetPC())
	}

	if actual != expected {
		t.Errorf("Register test failed: expected %s @ 0x%04X, got 0x%04X", name, expected, actual)
	}
}

func testAddress(t* testing.T, addr uint16, expected byte) {
	if actual := Get(addr); actual != expected {
		t.Errorf("Address test failed: at 0x%04X, expected 0x%02X, got 0x%02X", addr, expected, actual)
	}
}
