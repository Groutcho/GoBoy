package cpu

import "testing"
import . "memory"

func TestC5_push(t* testing.T) {
	ResetMemory()
	SetSP(0x0A)
	SetBC(0x1234)

	xC5_push()

	testRegister(t, REG_SP, 0x08)
	testAddress(t, GetSP(), 0x34)
	testAddress(t, GetSP() + 1, 0x12)
}


func TestF5_push(t* testing.T) {
	ResetMemory()

	SetF(0x00)
	SetFlags(0, F_SET_1, F_SET_1, F_SET_1, F_SET_1, 0)
	SetA(0x78)

	SetSP(0x0A)

	xF5_push()
	SetA(0x00)
	SetFlags(0, F_SET_0, F_SET_0, F_SET_0, F_SET_0, 0)
	testFlags(t, false, false, false, false)
	xF1_pop()

	testRegister(t, REG_A, 0x78)
	testFlags(t, true, true, true, true)
}