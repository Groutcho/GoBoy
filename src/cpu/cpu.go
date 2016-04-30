package cpu

import . "registers"
import . "memory"

var opcode uint16 = 0x0000

// Get the opcode at the current PC, increment PC then return the opcode.
// If opcode is an extended opcode, i.e CB XX, return FF + XX after
// incrementing the PC twice.
func Fetch() uint16 {
	opcode = uint16(Get(GetPC()))
	IncPC()

	if opcode == 0xCB {
		opcode = 0xFF + uint16(Get(GetPC()))
		IncPC()
	}

	return opcode
}