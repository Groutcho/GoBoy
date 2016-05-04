package cpu

import "testing"
// import . "memory"

func TestXCD_call(t* testing.T) {
	program := []byte{
		0x45, // garbage
		0x45, // garbage
		0x45, // garbage
		0x45, // garbage
		0x45, // garbage
		0x45, // garbage
		0xCD, 0xFF, 0x15, // call 0xFF15
	}

	SetPC(0x0006)
	ExecuteNext()
	if GetPC() != 0xFF15 {
		t.Errorf("TestxCD_call() failed: expected PC @ 0xFF15, got 0x%04x", GetPC())
	}
}