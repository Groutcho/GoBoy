package cpu

import . "memory"

// SP=SP-2,  (SP)=BC
func xC5_push() int {
	DecSP()
	DecSP()

	Set(GetSP(), GetC())
	Set(GetSP()+1, GetB())

	return 4
}

func xD5_push() int {
	DecSP()
	DecSP()

	Set(GetSP(), GetE())
	Set(GetSP()+1, GetD())

	return 4
}

func xE5_push() int {
	DecSP()
	DecSP()

	Set(GetSP(), GetL())
	Set(GetSP()+1, GetH())

	return 4
}

func xF5_push() int {
	DecSP()
	DecSP()

	Set(GetSP(), GetF())
	Set(GetSP()+1, GetA())

	return 4
}