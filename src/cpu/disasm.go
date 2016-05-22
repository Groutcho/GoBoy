package cpu

import (
	"fmt"
	"memory"
	"os"
)

const (
	NO_ARG = 0
	S8     = 1
	U8     = 2
	U16    = 3
)

type opcodePair struct {
	text     string
	argument int
}

var (
	disasm_table = new([512]opcodePair)
	asm_cache    = make([]string, 64000, 64000)
)

func init() {
	for i := 0; i < 512; i++ {
		disasm_table[i].text = "ld A, 0x%02X"
		disasm_table[i].argument = U8
	}
}

func cache(s string) {
	asm_cache = append(asm_cache, s)
}

// translate the opcode into a mnemonic
func disasm(opcode, pc uint16, inc int) {
	pair := disasm_table[opcode]
	txt := pair.text
	addr := fmt.Sprintf("[0x%04X]", pc-uint16(inc))

	switch pair.argument {
	case U8:
		arg := memory.Get(pc)
		txt = fmt.Sprintf(txt, arg)
	case S8:
		arg := memory.Get(pc)
		txt = fmt.Sprintf(txt, arg)
	case U16:
		arg := memory.Get16(pc)
		txt = fmt.Sprintf(txt, arg)
	}
	result := fmt.Sprintf("%s %s\n", addr, txt)
	cache(result)
	fmt.Print(result)
}

func DumpASM() {
	f, err := os.Create("dump.asm")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i := 0; i < len(asm_cache); i++ {
		f.WriteString(asm_cache[i])
	}
}
