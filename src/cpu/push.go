package cpu

import . "memory"

// SP=SP-2,  (SP)=BC
func xC5_push() int {
	SP--
	Write(SP, B)
	SP--
	Write(SP, C)

	return 4
}

func xD5_push() int {
	SP--
	Write(SP, D)
	SP--
	Write(SP, E)

	return 4
}

func xE5_push() int {
	SP--
	Write(SP, H)
	SP--
	Write(SP, L)

	return 4
}

func xF5_push() int {
	SP--
	Write(SP, A)
	SP--
	Write(SP, F)

	return 4
}
