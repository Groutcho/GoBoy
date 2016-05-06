// this file is automatically generated by instruction_generator.py
package cpu

import . "memory"


// and  B - A=A & B
func xA0_and() int {
	result := GetA() & GetB()
    SetA(result)

    SetFlags(int(result), F_SET_IF, F_SET_0, F_SET_1, F_SET_0, F_8bit)

	return 1
}

// and  C - A=A & C
func xA1_and() int {
	result := GetA() & GetC()
    SetA(result)

    SetFlags(int(result), F_SET_IF, F_SET_0, F_SET_1, F_SET_0, F_8bit)

	return 1
}

// and  D - A=A & D
func xA2_and() int {
	result := GetA() & GetD()
    SetA(result)

    SetFlags(int(result), F_SET_IF, F_SET_0, F_SET_1, F_SET_0, F_8bit)

	return 1
}

// and  E - A=A & E
func xA3_and() int {
	result := GetA() & GetE()
    SetA(result)

    SetFlags(int(result), F_SET_IF, F_SET_0, F_SET_1, F_SET_0, F_8bit)

	return 1
}

// and  H - A=A & H
func xA4_and() int {
	result := GetA() & GetH()
    SetA(result)

    SetFlags(int(result), F_SET_IF, F_SET_0, F_SET_1, F_SET_0, F_8bit)

	return 1
}

// and  L - A=A & L
func xA5_and() int {
	result := GetA() & GetL()
    SetA(result)

    SetFlags(int(result), F_SET_IF, F_SET_0, F_SET_1, F_SET_0, F_8bit)

	return 1
}

// and  (HL) - A=A & (HL)
func xA6_and() int {
	result := GetA() & Get(GetHL())
    SetA(result)

    SetFlags(int(result), F_SET_IF, F_SET_0, F_SET_1, F_SET_0, F_8bit)

	return 2
}

// and  A - A=A & A
func xA7_and() int {
	result := GetA() & GetA()
    SetA(result)

    SetFlags(int(result), F_SET_IF, F_SET_0, F_SET_1, F_SET_0, F_8bit)

	return 1
}

// and  %1 - A=A & %1
func xE6_and() int {
	result := GetA() & FetchOperand8()
    SetA(result)

    SetFlags(int(result), F_SET_IF, F_SET_0, F_SET_1, F_SET_0, F_8bit)

	return 2
}
