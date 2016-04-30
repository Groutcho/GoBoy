package cpu

import . "registers"
import . "memory"

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

// execute the next instruction and return the number of cycles taken
// by this instruction, as a multiple of 4, i.e unit cycles and not
// actual CPU cycles. The minimal amount of cycles is 1.
func ExecuteNext() uint8 {
	opcode := Fetch()
	return dispatch_table[opcode]()
}

type instrFunc func() uint8