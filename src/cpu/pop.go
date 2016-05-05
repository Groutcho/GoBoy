package cpu

import . "memory"

// AF=(SP),  SP=SP+2
func xF1_pop() int {
	SetA(Get(GetSP()))
	IncSP()
	SetF(Get(GetSP()))
	IncSP()

	return 3
}

// BC=(SP),  SP=SP+2
func xC1_pop() int {
	SetB(Get(GetSP()))
	IncSP()
	SetC(Get(GetSP()))
	IncSP()

	return 3
}

// DE=(SP),  SP=SP+2
func xD1_pop() int {
	SetD(Get(GetSP()))
	IncSP()
	SetE(Get(GetSP()))
	IncSP()

	return 3
}

// HL=(SP),  SP=SP+2
func xE1_pop() int {
	SetH(Get(GetSP()))
	IncSP()
	SetL(Get(GetSP()))
	IncSP()

	return 3
}