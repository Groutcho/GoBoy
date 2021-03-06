package cpu

import (
	. "common"
	"fmt"
	"io/ioutil"
	"log"
	. "memory"
	"os"
	"time"
)

// The dispatch table is used to redirect a given instruction to its
// implementation by having a direct mapping between the opcode and the array index.
// In the case of extended opcodes (CB XX), the mapping is the actual opcode (XX) + FF.
// example: the instruction CB F8 will sit at index FF+F8 of the dispatch table.
var (
	dispatch_table []instrFunc = make([]instrFunc, 512, 512)

	exitOnStop = false
	CONTINUE   = true

	instructionsCount int
	startTime         time.Time
	cycleCount        uint64
	elapsed           time.Duration

	breakpoints []uint16 = make([]uint16, 16, 16)

	rom []byte
)

func push(value uint16) {
	SP -= 2

	Write(SP, GetHighBits(value))
	Write(SP+1, GetLowBits(value))
}

// Get the opcode at the current PC, increment PC then return the opcode.
// If opcode is an extended opcode, i.e CB XX, return FF + XX after
// incrementing the PC twice.
func Fetch() (uint16, int) {
	opcode := uint16(Get(PC))
	inc := 1
	PC++

	if opcode == 0xCB {
		inc++
		opcode = 256 + uint16(Get(PC))
		PC++
	}

	return opcode, inc
}

// Get the 8bit word at the address pointed by the program counter
// and increment the program counter.
func FetchOperand8() uint8 {
	operand := Get(PC)
	IncPC()
	return operand
}

// Get the 8bit signed word at the address pointed by the program counter
// and increment the program counter.
func FetchOperand8s() int {
	operand := Get(PC)
	IncPC()
	if operand > 128 {
		return int(operand) - 0x100
	}
	return int(operand)
}

// Get the 16bit word at the address pointed by the program counter
// and increment the program counter twice.
func FetchOperand16() uint16 {
	lsb := uint16(Get(PC))
	IncPC()
	msb := uint16(Get(PC))
	IncPC()
	return (msb << 8) | lsb
}

// Perform a call: Decrement the stack pointer of 2 bytes,
// copy the current address in the two allocated bytes, then
// set the program counter to the given address.
func Call(addr uint16) {
	DecSP()
	DecSP()

	Write(GetSP(), GetLowBits(PC))
	Write(GetSP()+1, GetHighBits(PC))

	SetPC(addr)
}

func Step() {
	ExecuteNext()
}

func DumpRegisters() {
	fmt.Printf("A: 0x%02X  F: 0x%02X\n", GetA(), GetF())
	fmt.Printf("B: 0x%02X  C: 0x%02X\n", GetB(), GetC())
	fmt.Printf("D: 0x%02X  E: 0x%02X\n", GetD(), GetE())
	fmt.Printf("H: 0x%02X  L: 0x%02X\n", GetH(), GetL())
	fmt.Printf("SP: 0x%04X\n", GetSP())
	fmt.Printf("PC: 0x%04X\n", PC)
	fmt.Printf("#instructions: %09d\n", instructionsCount)
	fmt.Printf("time: %vs\n", elapsed)
	fmt.Printf("frequency: %fMhz", float64(cycleCount)/elapsed.Seconds()/1000000)
}

func SetBreakpoint(addr uint16) {
	breakpoints = append(breakpoints, addr)
	fmt.Printf("added breakpoint at 0x%04X\n", addr)
}

// execute the next instruction and return the number of cycles taken
// by this instruction, as a multiple of 4, i.e unit cycles and not
// actual CPU cycles. The minimal amount of cycles is 1.
func ExecuteNext() int {
	var opcode uint16
	var inc int

	pc := PC
	for i := 0; i < len(breakpoints); i++ {
		if pc != 0 && breakpoints[i] == pc {
			fmt.Printf("breakpoint reached: 0x%04X\n", pc)
			CONTINUE = false
			return 0
		}
	}

	defer func() {
		if x := recover(); x != nil {
			log.Printf("[0x%04X] %02X (%s)", PC-uint16(inc), opcode, x)
			log.Printf("Dumping memory...")
			dump := GetRange(0x0000, 0xFFFF)
			ioutil.WriteFile("dump.bin", dump, 0644)
			os.Exit(1)
		}
	}()

	opcode, inc = Fetch()
	// fmt.Printf("[0x%04X] %02X\n", PC-uint16(inc), opcode)

	// TODO: activate disassembly on demand
	disasm(opcode, PC, inc)

	if opcode == 0x76 { // halt
		// TODO
		return 1
	} else if opcode == 0x10 { // stop
		if exitOnStop {
			log.Print("Exiting upon STOP (exitOnStop true)")
			DumpRegisters()
			return 0
		}
		return 1
	}

	instructionsCount++
	return dispatch_table[opcode]()
}

func Update() int {
	// execute the next instruction and get its execution time, in microseconds
	wait_microsec := ExecuteNext()
	if wait_microsec > 0 {
		// t.Sleep(t.Duration(wait_microsec) * t.Microsecond)
		return 0
	} else {
		return -1
	}
}

// Starts the execution of the program at any point
func Run() {
	var cycles int
	startTime = time.Now()
	tick := time.NewTicker(time.Millisecond * 2)

	CONTINUE = true
	for {
		if CONTINUE {
			// 1 cycle is around 1 µs
			c := ExecuteNext()
			cycles += c
			cycleCount += uint64(c * 4)
			if cycles >= 2099 {
				cycles = 0
				<-tick.C
			}
		}
	}
}

func Pause() {
	CONTINUE = false
	elapsed = time.Since(startTime)
}

func Continue() {
	CONTINUE = true
}

// Starts the execution of the program in test mode, executing at most
// maxInstructions.
func StartTest(maxInstructions int) {
	ExitOnStop(true)
	ret := 0
	for i := 0; i < maxInstructions; i++ {
		ret = Update()
		if ret == -1 {
			return
		}
	}
}

// When encountering a STOP (0x10) opcode, tell the emulator
// to exit the program and dump the registers. Useful for debug
// purposes when executing custom test roms.
func ExitOnStop(exit bool) {
	exitOnStop = exit
}

// Load given program at address 0x0000
func LoadProgram(program []byte) {
	SetRange(0x0000, program)
}

func LoadRom(rom []byte) {
	rom = rom
	if len(rom) > 0x10000 {
		LoadProgram(rom[:0x10000])
	} else {
		LoadProgram(rom)
	}
}

func Initialize() {
	SetA(0x01)
	SetF(0xB0)
	SetBC(0x0013)
	SetDE(0x00D8)
	SetHL(0x014D)
	SetSP(0xFFFE)
	SetPC(0x100)

	SetTIMA(0x00)
	SetTMA(0x00)
	SetTAC(0x00)
	SetNR10(0x80)
	SetNR11(0xBF)
	SetNR12(0xF3)
	SetNR14(0xBF)
	SetNR21(0x3F)
	SetNR24(0xBF)
	SetNR30(0x7F)
	SetNR31(0xFF)
	// SetNR32(0x9F)
	SetNR33(0xBF)
	SetNR41(0xFF)
	SetNR42(0x00)
	// SetNR43(0x00)
	// TODO: NR30 ??
	SetNR50(0x77)
	SetNR51(0xF3)
	SetNR52(0xF1)
	SetLCDC(0x91)
	SetSCY(0x00)
	SetSCX(0x00)
	SetLYC(0x00)
	SetBGP(0xFC)
	SetOBP0(0xFF)
	SetOBP1(0xFF)
	SetWY(0x00)
	SetWX(0x00)
	SetWX(0x00)
}

type instrFunc func() int
