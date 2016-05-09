package tests

import "testing"
import "cpu"
import "memory"
import "io/ioutil"

const romDirectory = "./roms/"

const REG_A = 0
const REG_F = 1
const REG_AF = 2
const REG_B = 3
const REG_C = 4
const REG_BC = 5
const REG_D = 6
const REG_E = 7
const REG_DE = 8
const REG_H = 9
const REG_L = 10
const REG_HL = 11
const REG_SP = 12
const REG_PC = 13

func LoadTestRom(romName string) {
	rom, err := ioutil.ReadFile(romDirectory + romName + ".gb")
	if err != nil {
	    panic(err)
	}

	cpu.LoadProgram(rom)
}

func LaunchRomTest(romName string, maxInstr int) {
	memory.ResetMemory()
	cpu.Initialize()
	
	LoadTestRom(romName)
	cpu.StartTest(maxInstr)
}

//
// ROM 001
//	ld A, 0xF0
//	inc A (7 times)
//
func TestRom001(t* testing.T) {
	LaunchRomTest("t001", 20)

	if a := cpu.GetA(); a != 0xF7 {
		t.Errorf("t001: expected A = 0xF7, got 0x%02X", a)
	}
}

//
// ROM 002
//	ld DE, 0x1234
//	push DE
// 	ld DE, 0x9999
//	pop DE
//
func TestRom002(t* testing.T) {
	LaunchRomTest("t002", 20)
	if de := cpu.GetDE(); de != 0x1234 {
		t.Errorf("t002: expected DE = 0x1234, got 0x%04X", de)
	}
}

//
// ROM 003
//	ld HL, 0x0020
//	ld A, 0x88
//	ld (HL), A 
//
func TestRom003(t* testing.T) {
	LaunchRomTest("t003", 20)
	cpu.CheckRegister(t, REG_HL, 0x0020)
	cpu.CheckRegister(t, REG_A, 0x88)
	if a := memory.Get(0x0020); a != 0x88 {
		t.Errorf("t002: expected [HL] = 0x88, got 0x%02X", a)
	}
}

//
// ROM 004
//	ld A, 0x47
//	jp 0x0110
//	STOP 
//	...
//	NOP
//	...
//	ld A, 0x99 (0x0110)
//	STOP
func TestRom004(t* testing.T) {
	LaunchRomTest("t004", 20)
	cpu.CheckRegister(t, REG_PC, 0x0113)
	cpu.CheckRegister(t, REG_A, 0x99)
}