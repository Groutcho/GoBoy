package main

import "lcd"
import t "time"
import "memory"
import "cpu"
import "os"
import "io/ioutil"

// import "log"
import "fmt"
import . "common"

const NORMAL_SPEED = 1

func loadRom(filename string) {
	rom, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	cpu.LoadProgram(rom)
}

func testTetris() {

	lcd.Initialize()
	cpu.Initialize()
	memory.ResetMemory()

	blank := "" +
		"........" +
		"........" +
		"........" +
		"........" +
		"........" +
		"........" +
		"........" +
		"........"

	block := "" +
		"33333333" +
		"31111113" +
		"31....13" +
		"31.11313" +
		"31.11313" +
		"31333313" +
		"31111113" +
		"33333333"

	square := "" +
		"33333333" +
		"3......3" +
		"3.2222.3" +
		"3.2222.3" +
		"3.2222.3" +
		"3.2222.3" +
		"3......3" +
		"33333333"

		// mire := ""  +
		//  "....3333" +
		//  "....3333" +
		//  "....3333" +
		//  "....3333" +
		//  "2222...." +
		//  "2222...." +
		//  "2222...." +
		//  "2222...."

	lcd.SetTile(0, lcd.MakeTile(blank), 1)
	lcd.SetTile(1, lcd.MakeTile(block), 1)
	lcd.SetTile(2, lcd.MakeTile(square), 1)

	// for i := 0; i < 32; i++ {
	// 	lcd.SetWindowTile(i, 0, 2)
	// 	lcd.SetWindowTile(i, 1, 2)
	// 	lcd.SetWindowTile(i, 2, 2)
	// }

	for i := 0; i < 32; i++ {
		for j := 0; j < 10; j++ {
			lcd.SetBackgroundTile(i, j, 1)
		}
	}

	for i := 0; i < 32; i++ {
		for j := 10; j < 20; j++ {
			lcd.SetWindowTile(i, j, 2)
		}
	}

	memory.SetLCDC(0x91)
	lcdc := memory.GetLCDC()
	memory.SetLCDC(SetBit(lcdc, 7, 1))

	go fall()
	lcd.Run()
	// lcd.RunProfile()

	lcd.Stop()
}

func fall() {
	tick := t.NewTicker(t.Millisecond * 100)
	for {
		<-tick.C
		memory.IncSCY()
	}
}

func pollSerial() {
	for {
		if memory.Get(0xFF02) == 0x81 {
			fmt.Printf("%c", memory.Get(0xFF01))
		}
		t.Sleep(1 * t.Microsecond)
	}
}

func main() {
	testTetris()
	return

	memory.ResetMemory()
	lcd.Initialize()
	cpu.Initialize()

	loadRom(os.Args[1])

	go lcd.Run()
	// go pollSerial()
	cpu.Run()

	lcd.Stop()
}
