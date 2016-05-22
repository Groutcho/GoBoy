package cpu

import . "memory"

func ret() {
	PC = Get16(SP)

	SP += 2
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
	SetIME(true)
	ret()
	return 4
}
