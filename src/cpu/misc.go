package cpu

import . "memory"

func init() {
	dispatch_table[0x03F] = x3F_ccf
	dispatch_table[0x037] = x37_scf
	dispatch_table[0x000] = x00_nop
	dispatch_table[0x0F3] = xF3_di
	dispatch_table[0x0FB] = xFB_ei
}

func x3F_ccf() int {
	if GetFlagCy() {
		SetFlagCy(false)
	} else {
		SetFlagCy(true)		
	}

	SetFlagN(false)
	SetFlagH(false)

	return 1
}

func x37_scf() int {
	SetFlagN(false)
	SetFlagH(false)
	SetFlagCy(true)

	return 1
}

func x00_nop() int {
	return 1
}

func xF3_di() int {
	DisableVBlankInterrupt()
	DisableLcdStatInterrupt()
	DisableTimerInterrupt()
	DisableSerialInterrupt()
	DisableJoypadInterrupt()
	return 1
}

func xFB_ei() int {
	EnableVBlankInterrupt()
	EnableLcdStatInterrupt()
	EnableTimerInterrupt()
	EnableSerialInterrupt()
	EnableJoypadInterrupt()
	return 1
}