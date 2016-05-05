package cpu

import . "memory"

func x22_ldi() int {
	Set(GetHL(), GetA())
	IncHL()

	return 2
}

func x2A_ldi() int {
	SetA(Get(GetHL()))
	IncHL()

	return 2
}

func x32_ldd() int {
	Set(GetHL(), GetA())
	DecHL()

	return 2
}

func x3A_ldd() int {
	SetA(Get(GetHL()))
	DecHL()

	return 2
}
