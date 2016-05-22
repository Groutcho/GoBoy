package cpu

import (
	"memory"
	"testing"
)

func TestJp(t *testing.T) {
	ResetSystem()

	PC = 0x200
	memory.Write16(0x200, 0x1234)
	xC3_jp()
	CheckRegister(t, REG_PC, 0x1234)

	PC = 0x200
	SetFlagZf(true)
	memory.Write16(0x200, 0x1234)
	xC2_jp()
	CheckRegister(t, REG_PC, 0x202)

	SetFlagZf(false)
	memory.Write16(0x202, 0x3000)
	xC2_jp()
	CheckRegister(t, REG_PC, 0x3000)

	PC = 0
	SetFlagZf(false)
	memory.Write16(0, 0x3000)
	xCA_jp()
	CheckRegister(t, REG_PC, 0x0002)

	PC = 0
	SetFlagZf(true)
	xCA_jp()
	CheckRegister(t, REG_PC, 0x3000)

	PC = 0
	SetFlagCy(true)
	memory.Write16(0, 0x3000)
	xD2_jp()
	CheckRegister(t, REG_PC, 0x0002)

	PC = 0
	SetFlagCy(false)
	xD2_jp()
	CheckRegister(t, REG_PC, 0x3000)

	PC = 0
	SetFlagCy(false)
	memory.Write16(0, 0x3000)
	xDA_jp()
	CheckRegister(t, REG_PC, 0x0002)

	PC = 0
	SetFlagCy(true)
	xDA_jp()
	CheckRegister(t, REG_PC, 0x3000)

	PC = 0
	SetHL(0x1234)
	xE9_jp()
	CheckRegister(t, REG_PC, 0x1234)
}
