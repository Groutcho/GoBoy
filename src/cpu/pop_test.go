package cpu

import "testing"
import . "memory"

// AF=(SP),  SP=SP+2
func TestF1_pop(t* testing.T) {
	ResetMemory()

	SetAF(0x0000)
	SetSP(0x0020)

	Set(0x0020, 0xF8)
	Set(0x0021, 0xAB)

	t.Log(DumpRange(0x0015, 0x0024))
	t.Log("15")
	
	xF1_pop()

	if sp := GetSP(); sp != 0x0022 {
		t.Errorf("TestF1_pop() failed: expected SP @ 0x0022, got 0x%04X", sp)
	}
	if af := GetAF(); af != 0xF8AB {
		t.Errorf("TestF1_pop() failed: expected AF @ 0xF8AB, got 0x%04X", af)		
	}
}