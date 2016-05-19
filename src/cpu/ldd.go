package cpu

import . "memory"

func x22_ldi() int {
	Write(GetHL(), GetA())
	IncHL()

	return 2
}

func x2A_ldi() int {
	SetA(Get(GetHL()))
	IncHL()

	return 2
}

func x32_ldd() int {
	Write(GetHL(), GetA())
	DecHL()

	return 2
}

func x3A_ldd() int {
	SetA(Get(GetHL()))
	DecHL()

	return 2
}
