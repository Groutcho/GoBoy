package console

import (
	"bufio"
	"cpu"
	"fmt"
	"lcd"
	"os"
	"strconv"
	"strings"
)

func Prompt() {
	cpu.Pause()
	lcd.Pause()

	defer lcd.Continue()
	defer cpu.Continue()

	for {
		fmt.Print("\n(gb) ")
		cmd := waitCommand()

		if cmd[0] == "reg" {
			cpu.DumpRegisters()
		} else if cmd[0] == "continue" {
			fmt.Print("executing...\n")
			return
		} else if cmd[0] == "quit" {
			os.Exit(0)
		} else if cmd[0] == "s" {
			cpu.Step()
		} else if cmd[0] == "brk" {
			addr, err := strconv.ParseInt(cmd[1], 0, 16)
			if err != nil {
				panic(err)
			}
			cpu.SetBreakpoint(uint16(addr))
		}
	}
}

func waitCommand() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	args := strings.Split(scanner.Text(), " ")
	return args
}
