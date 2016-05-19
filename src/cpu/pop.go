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
	E = Get(SP)
	SP++
	D = Get(SP)
	SP++

	return 3
}

// HL=(SP),  SP=SP+2
func xE1_pop() int {
	L = Get(SP)
	SP++
	H = Get(SP)
	SP++

	return 3
}
