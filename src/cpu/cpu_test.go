package cpu

import "testing"
import . "memory"

func TestFetch(t* testing.T) {
	arr := make([]byte, 4)
	arr[0] = 0x00
	arr[1] = 0x58
	arr[2] = 0xCB
	arr[3] = 0x98

	LoadProgram(arr)
	SetPC(0x0000)

	startingPC := GetPC()
	currentPC := startingPC

	opcode := uint16(0x50)
	opcode = Fetch()
	if opcode != 0x00 {
		t.Errorf("TestFetch(): at first Fetch(), expected 00, got %02x", opcode)
	}

	if currentPC = GetPC(); currentPC != (startingPC + 1) {
		t.Errorf("TestFetch(): at first Fetch(), expected PC @ 0x%04x, got 0x%04x", (startingPC + 1), currentPC)
	}

	opcode = Fetch()
	if opcode != 0x58 {
		t.Errorf("TestFetch(): at second Fetch(), expected 58, got %02x", opcode)
	}

	if currentPC = GetPC(); currentPC != (startingPC + 2) {
		t.Errorf("TestFetch(): at second Fetch(), expected PC @ 0x%04x, got 0x%04x", (startingPC + 2), currentPC)
	}

	opcode = Fetch()
	if opcode != 0x197 {
		t.Errorf("TestFetch(): at last Fetch(), expected extended opcode conversion CB 98 -> 197, got %02x", opcode)
	}

	if currentPC = GetPC(); currentPC != (startingPC + 4) {
		t.Errorf("TestFetch(): at last Fetch(), expected PC @ 0x%04x, got 0x%04x", (startingPC + 4), currentPC)
	}
}

func op0() int {
	SetA(0x27)
	return 1
}

func op1() int {
	SetA(0x44)
	return 1
}

func op2() int {
	SetA(0x38)
	return 1
}

func op3() int {
	SetA(0xFF)
	return 1
}

func TestExecuteNext(t* testing.T) {
	SetPC(0x0000)

	arr := make([]byte, 4)
	arr[0] = 0xF8
	arr[1] = 0xF9
	arr[2] = 0xFA
	arr[3] = 0xFB

	LoadProgram(arr)
	SetPC(0x0000)

	dispatch_table[0xF8] = op0
	dispatch_table[0xF9] = op1
	dispatch_table[0xFA] = op2
	dispatch_table[0xFB] = op3

	ExecuteNext()
	if a := GetA(); a != 0x27 {
		t.Errorf("TestExecuteNext(): expected value 0x27 in register A, got 0x%02x", a)
	}

	ExecuteNext()
	if a := GetA(); a != 0x44 {
		t.Errorf("TestExecuteNext(): expected value 0x44 in register A, got 0x%02x", a)
	}

	ExecuteNext()
	if a := GetA(); a != 0x38 {
		t.Errorf("TestExecuteNext(): expected value 0x38 in register A, got 0x%02x", a)
	}

	ExecuteNext()
	if a := GetA(); a != 0xFF {
		t.Errorf("TestExecuteNext(): expected value 0xFF in register A, got 0x%02x", a)
	}
}

func TestLdImmediateInstructions(t* testing.T) {
	program := []byte{
		0x3E, 0x25, // ld A 0x25
		0x06, 0x65, // ld B 0x65
		0x0E, 0x77, // ld C 0x77
		0x16, 0xFE, // ld D 0xFE
		0x1E, 0xEB, // ld E 0xEB
		0x26, 0x01, // ld H 0x01
		0x2E, 0x9D, // ld L 0x9D
	}

	LoadProgram(program)
	SetPC(0x0000)

	ExecuteNext()
	ExecuteNext()
	ExecuteNext()
	ExecuteNext()
	ExecuteNext()
	ExecuteNext()
	ExecuteNext()

	if a:= GetA(); a != 0x25 {
		t.Errorf("TestLdImmediateInstructions() failed: expected A = 0x25, got 0x%02x", a)
	}
	if b:= GetB(); b != 0x65 {
		t.Errorf("TestLdImmediateInstructions() failed: expected B = 0x65, got 0x%02x", b)
	}
	if c:= GetC(); c != 0x77 {
		t.Errorf("TestLdImmediateInstructions() failed: expected C = 0x77, got 0x%02x", c)
	}
	if d:= GetD(); d != 0xFE {
		t.Errorf("TestLdImmediateInstructions() failed: expected D = 0xFE, got 0x%02x", d)
	}
	if e:= GetE(); e != 0xEB {
		t.Errorf("TestLdImmediateInstructions() failed: expected E = 0xEB, got 0x%02x", e)
	}
	if h:= GetH(); h != 0x01 {
		t.Errorf("TestLdImmediateInstructions() failed: expected H = 0x01, got 0x%02x", h)
	}
	if l:= GetL(); l != 0x9D {
		t.Errorf("TestLdImmediateInstructions() failed: expected L = 0x9D, got 0x%02x", l)
	}
}

func TestLdRegistersInstructions(t* testing.T) {
	program := []byte{
		0x78, // ld A, B
		0x79, // ld A, C
		0x7A, // ld A, D
		0x7B, // ld A, E
		0x7C, // ld A, H
		0x7D, // ld A, L
		0x7F, // ld A, A
	}

	LoadProgram(program)
	SetPC(0x0000)

	SetA(0x58)
	SetB(0x99)
	SetC(0xFE)
	SetD(0xFF)
	SetE(0xCD)
	SetH(0x58)
	SetL(0x00)

	ExecuteNext()
	if a:= GetA(); a != 0x99 {
		t.Errorf("TestLdRegistersInstructions() failed: expected A = 0x99, got 0x%02x", a)
	}
	ExecuteNext()
	if a:= GetA(); a != 0xFE {
		t.Errorf("TestLdRegistersInstructions() failed: expected A = 0xFE, got 0x%02x", a)
	}
	ExecuteNext()
	if a:= GetA(); a != 0xFF {
		t.Errorf("TestLdRegistersInstructions() failed: expected A = 0xFF, got 0x%02x", a)
	}
	ExecuteNext()
	if a:= GetA(); a != 0xCD {
		t.Errorf("TestLdRegistersInstructions() failed: expected A = 0xCD, got 0x%02x", a)
	}
	ExecuteNext()
	if a:= GetA(); a != 0x58 {
		t.Errorf("TestLdRegistersInstructions() failed: expected A = 0x58, got 0x%02x", a)
	}
	ExecuteNext()
	if a:= GetA(); a != 0x00 {
		t.Errorf("TestLdRegistersInstructions() failed: expected A = 0x00, got 0x%02x", a)
	}
}

func TestLdReferenceHLInstructions(t* testing.T) {
	program := []byte{
		0x7E, // ld   A, (HL)
		0x46, // ld   B, (HL)
		0x4E, // ld   C, (HL)
		0x56, // ld   D, (HL)
		0x5E, // ld   E, (HL)
		0x66, // ld   H, (HL)
		0x6E, // ld   L, (HL)
	}

	LoadProgram(program)
	SetPC(0x0000)

	SetHL(0x1000)
	Set(0x1000, 0xFB)

	ExecuteNext()
	if a:= GetA(); a != 0xFB {
		t.Errorf("TestLdRegistersInstructions() failed: expected A = 0xFB, got 0x%02x", a)
	}
	ExecuteNext()
	if b:= GetB(); b != 0xFB {
		t.Errorf("TestLdRegistersInstructions() failed: expected B = 0xFB, got 0x%02x", b)
	}
	ExecuteNext()
	if c:= GetC(); c != 0xFB {
		t.Errorf("TestLdRegistersInstructions() failed: expected C = 0xFB, got 0x%02x", c)
	}
	ExecuteNext()
	if d:= GetD(); d != 0xFB {
		t.Errorf("TestLdRegistersInstructions() failed: expected D = 0xFB, got 0x%02x", d)
	}
	ExecuteNext()
	if e:= GetE(); e != 0xFB {
		t.Errorf("TestLdRegistersInstructions() failed: expected E = 0xFB, got 0x%02x", e)
	}
	ExecuteNext()
	if h:= GetH(); h != 0xFB {
		t.Errorf("TestLdRegistersInstructions() failed: expected H = 0xFB, got 0x%02x", h)
	}
	// tricky part: the register HL is overwritten by the previous instruction
	ExecuteNext()
	if hl:= GetHL(); hl != 0xFB00 {
		t.Errorf("TestLdRegistersInstructions() failed: expected HL = 0xFB00, got 0x%04x", hl)
	}
}

func TestLdToReferenceHLInstructions(t* testing.T) {
	program := []byte{
        0x70, // ld   (HL), B
        0x71, // ld   (HL), C
        0x72, // ld   (HL), D
        0x73, // ld   (HL), E
        0x74, // ld   (HL), H
        0x75, // ld   (HL), L
        0x77, // ld   (HL), A
        0x36, 0x88, // ld   (HL), 0x88
	}

	LoadProgram(program)
	SetPC(0x0000)

	SetHL(0x1000)
	Set(0x1000, 0xFB)

	SetA(0x58)
	SetB(0x99)
	SetC(0xFE)
	SetD(0xFF)
	SetE(0xCD)
	SetH(0x58)
	SetL(0x00)

	ExecuteNext()
	if v:= Get(GetHL()); v != 0x99 {
		t.Errorf("TestLdToReferenceHLInstructions() failed: expected [HL] = 0x99, got 0x%02x", v)
	}
	ExecuteNext()
	if v:= Get(GetHL()); v != 0xFE {
		t.Errorf("TestLdToReferenceHLInstructions() failed: expected [HL] = 0xFE, got 0x%02x", v)
	}
	ExecuteNext()
	if v:= Get(GetHL()); v != 0xFF {
		t.Errorf("TestLdToReferenceHLInstructions() failed: expected [HL] = 0xFF, got 0x%02x", v)
	}
	ExecuteNext()
	if v:= Get(GetHL()); v != 0xCD {
		t.Errorf("TestLdToReferenceHLInstructions() failed: expected [HL] = 0xCD, got 0x%02x", v)
	}
	ExecuteNext()
	if v:= Get(GetHL()); v != 0x58 {
		t.Errorf("TestLdToReferenceHLInstructions() failed: expected [HL] = 0x58, got 0x%02x", v)
	}
	ExecuteNext()
	if v:= Get(GetHL()); v != 0x00 {
		t.Errorf("TestLdToReferenceHLInstructions() failed: expected [HL] = 0x00, got 0x%02x", v)
	}
	ExecuteNext()
	if v:= Get(GetHL()); v != 0x58 {
		t.Errorf("TestLdToReferenceHLInstructions() failed: expected [HL] = 0x58, got 0x%02x", v)
	}
	ExecuteNext()
	if v:= Get(GetHL()); v != 0x88 {
		t.Errorf("TestLdToReferenceHLInstructions() failed: expected [HL] = 0x88, got 0x%02x", v)
	}
}

// measure the duration of a ld instruction
func BenchmarkLdRegistersInstructions(b *testing.B) {
	Set(0x0000, 0x78) // ld A, B

	for i := 0; i < b.N; i++ {
		SetPC(0x0000)
		ExecuteNext()
	}
}
