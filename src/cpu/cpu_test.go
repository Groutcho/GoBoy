package cpu

import "testing"
import . "memory"
import . "registers"

func TestFetch(t* testing.T) {
	arr := make([]byte, 4)
	arr[0] = 0x00
	arr[1] = 0x58
	arr[2] = 0xCB
	arr[3] = 0x98

	SetRange(0x0000, 0x0003, arr)

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