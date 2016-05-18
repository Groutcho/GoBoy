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
		args := waitCommand()

		cmd := args[0]

		switch cmd {
		case "reg":
			cpu.DumpRegisters()
		case "continue":
			fmt.Print("executing...\n")
			return
		case "quit":
			os.Exit(0)
		case "s":
			cpu.Step()
		case "tdt":
			lcd.PrintTileInformation(args[1:])
		case "cpy":
			lcd.CopyTileMap()
		case "lcd":
			lcd.PrintVideoInformation()
		case "btm":
			lcd.PrintBackgroundTileMap()
		case "brk":
			addr, err := strconv.ParseInt(args[1], 0, 16)
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
