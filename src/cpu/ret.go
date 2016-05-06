package cpu

import . "memory"

func ret() {
	h := Get(GetSP() + 1)
	l := Get(GetSP())

	IncSP()
	IncSP()

	SetPC(uint16(h) << 8 | uint16(l))
}

func xC9_ret() int {
	ret()
	return 4
}

func xC8_ret() int {
	if GetFlagZf() {
		ret()
		return 6
	}
	return 3
}

func xC0_ret() int {
	if !GetFlagZf() {
		ret()
		return 6
	}
	return 3
}

func xD0_ret() int {
	if !GetFlagCy() {
		ret()
		return 6
	}
	return 3
}

func xD8_ret() int {
	if GetFlagCy() {
		ret()
		return 6
	}
	return 3
}

func xD9_reti() int {
	EnableAllInterrupts()
	ret()
	return 4
}
