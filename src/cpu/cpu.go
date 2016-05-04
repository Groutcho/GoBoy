package cpu

import . "memory"
import t "time"

// The dispatch table is used to redirect a given instruction to its
// implementation by having a direct mapping between the opcode and the array index.
// In the case of extended opcodes (CB XX), the mapping is the actual opcode (XX) + FF.
// example: the instruction CB F8 will sit at index FF+F8 of the dispatch table.
var dispatch_table []instrFunc = make([]instrFunc, 512, 512)

// Get the opcode at the current PC, increment PC then return the opcode.
// If opcode is an extended opcode, i.e CB XX, return FF + XX after
// incrementing the PC twice.
func Fetch() uint16 {
	opcode := uint16(Get(GetPC()))
	IncPC()

	if opcode == 0xCB {
		opcode = 0xFF + uint16(Get(GetPC()))
		IncPC()
	}

	return opcode
}

// Get the 8bit word at the address pointed by the program counter
// and increment the program counter.
func FetchOperand8() uint8 {
	operand := Get(GetPC())
	IncPC()
	return operand
}

// Get the 16bit word at the address pointed by the program counter
// and increment the program counter twice.
func FetchOperand16() uint16 {
	operand0 := uint16(Get(GetPC()))
	IncPC()
	operand1 := uint16(Get(GetPC()))
	IncPC()
	return (operand0 << 8) | operand1
}

// Perform a call: Decrement the stack pointer of 2 bytes,
// copy the current address in the two allocated bytes, then
// set the program counter to the given address.
func Call(addr uint16) {
	DecSP()
	DecSP()

	Set(GetSP(), 	 getLowBits(GetPC()))
	Set(GetSP() + 1, getHighBits(GetPC()))

	SetPC(addr)
}

// execute the next instruction and return the number of cycles taken
// by this instruction, as a multiple of 4, i.e unit cycles and not
// actual CPU cycles. The minimal amount of cycles is 1.
func ExecuteNext() int {
	opcode := Fetch()
	return dispatch_table[opcode]()
}

// Starts the execution of the program at 0x0000
func Start() {
	SetPC(0x0000)
	wait_microsec := 0

	for {
		// execute the next instruction and get its execution time, in microseconds
		wait_microsec = ExecuteNext()
		t.Sleep(t.Duration(wait_microsec) * t.Microsecond)
	}
}

type instrFunc func() int