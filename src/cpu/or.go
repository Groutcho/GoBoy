// this file is automatically generated by instruction_generator.py
package cpu

import . "memory"

// or   B - A=A | B
func xB0_or() int {
	result := GetA() | GetB()
	SetA(result)

	SetFlags(int(result), F_SET_IF, F_SET_0, F_SET_0, F_SET_0, F_8bit)

	return 1
}

// or   C - A=A | C
func xB1_or() int {
	result := GetA() | GetC()
	SetA(result)

	SetFlags(int(result), F_SET_IF, F_SET_0, F_SET_0, F_SET_0, F_8bit)

	return 1
}

// or   D - A=A | D
func xB2_or() int {
	result := GetA() | GetD()
	SetA(result)

	SetFlags(int(result), F_SET_IF, F_SET_0, F_SET_0, F_SET_0, F_8bit)

	return 1
}

// or   E - A=A | E
func xB3_or() int {
	result := GetA() | GetE()
	SetA(result)

	SetFlags(int(result), F_SET_IF, F_SET_0, F_SET_0, F_SET_0, F_8bit)

	return 1
}

// or   H - A=A | H
func xB4_or() int {
	result := GetA() | GetH()
	SetA(result)

	SetFlags(int(result), F_SET_IF, F_SET_0, F_SET_0, F_SET_0, F_8bit)

	return 1
}

// or   L - A=A | L
func xB5_or() int {
	result := GetA() | GetL()
	SetA(result)

	SetFlags(int(result), F_SET_IF, F_SET_0, F_SET_0, F_SET_0, F_8bit)

	return 1
}

// or   (HL) - A=A | (HL)
func xB6_or() int {
	result := GetA() | Get(GetHL())
	SetA(result)

	SetFlags(int(result), F_SET_IF, F_SET_0, F_SET_0, F_SET_0, F_8bit)

	return 2
}

// or   A - A=A | A
func xB7_or() int {
	result := GetA() | GetA()
	SetA(result)

	SetFlags(int(result), F_SET_IF, F_SET_0, F_SET_0, F_SET_0, F_8bit)

	return 1
}

// or   %1 - A=A | %1
func xF6_or() int {
	result := GetA() | FetchOperand8()
	SetA(result)

	SetFlags(int(result), F_SET_IF, F_SET_0, F_SET_0, F_SET_0, F_8bit)

	return 2
}

