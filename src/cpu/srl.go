// this file is automatically generated by instruction_generator.py
package cpu

import . "memory"

// srl  B - shift B right logical (b7=0)
func xCB_38_srl() int {
	value := GetB()

	SetFlagCy(value & 0x01 == 1)

	value = value >> 1

	SetFlagZf(value == 0)
	SetFlagN(false)
	SetFlagH(false)

	SetB(value)

	return 2
}

// srl  C - shift C right logical (b7=0)
func xCB_39_srl() int {
	value := GetC()

	SetFlagCy(value & 0x01 == 1)

	value = value >> 1

	SetFlagZf(value == 0)
	SetFlagN(false)
	SetFlagH(false)

	SetC(value)

	return 2
}

// srl  D - shift D right logical (b7=0)
func xCB_3A_srl() int {
	value := GetD()

	SetFlagCy(value & 0x01 == 1)

	value = value >> 1

	SetFlagZf(value == 0)
	SetFlagN(false)
	SetFlagH(false)

	SetD(value)

	return 2
}

// srl  E - shift E right logical (b7=0)
func xCB_3B_srl() int {
	value := GetE()

	SetFlagCy(value & 0x01 == 1)

	value = value >> 1

	SetFlagZf(value == 0)
	SetFlagN(false)
	SetFlagH(false)

	SetE(value)

	return 2
}

// srl  H - shift H right logical (b7=0)
func xCB_3C_srl() int {
	value := GetH()

	SetFlagCy(value & 0x01 == 1)

	value = value >> 1

	SetFlagZf(value == 0)
	SetFlagN(false)
	SetFlagH(false)

	SetH(value)

	return 2
}

// srl  L - shift L right logical (b7=0)
func xCB_3D_srl() int {
	value := GetL()

	SetFlagCy(value & 0x01 == 1)

	value = value >> 1

	SetFlagZf(value == 0)
	SetFlagN(false)
	SetFlagH(false)

	SetL(value)

	return 2
}

// srl  (HL) - shift right logical (b7=0)
func xCB_3E_srl() int {
	value := Get(GetHL())

	SetFlagCy(value & 0x01 == 1)

	value = value >> 1

	SetFlagZf(value == 0)
	SetFlagN(false)
	SetFlagH(false)

	Write(GetHL(), value)

	return 4
}

// srl  A - shift A right logical (b7=0
func xCB_3F_srl() int {
	value := GetA()

	SetFlagCy(value & 0x01 == 1)

	value = value >> 1

	SetFlagZf(value == 0)
	SetFlagN(false)
	SetFlagH(false)

	SetA(value)

	return 2
}

