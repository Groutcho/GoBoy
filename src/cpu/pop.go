package cpu

import . "memory"

// AF=(SP),  SP=SP+2
func xF1_pop() int {
	F = Get(SP)
	SP++
	A = Get(SP)
	SP++

	return 3
}

// BC=(SP),  SP=SP+2
func xC1_pop() int {
	C = Get(SP)
	SP++
	B = Get(SP)
	SP++

	return 3
}

// DE=(SP),  SP=SP+2
func xD1_pop() int {
	SetE(Get(GetSP()))
	IncSP()
	SetD(Get(GetSP()))
	IncSP()

	return 3
}

// HL=(SP),  SP=SP+2
func xE1_pop() int {
	SetL(Get(GetSP()))
	IncSP()
	SetH(Get(GetSP()))
	IncSP()

	return 3
}
