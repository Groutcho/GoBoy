package main

import (
	. "common"
	"console"
	"cpu"
	"fmt"
	"io/ioutil"
	"lcd"
	"memory"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const NORMAL_SPEED = 1

var (
	sigs chan os.Signal
)

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

	// block := "" +
	// 	"33333333" +
	// 	"31111113" +
	// 	"31....13" +
	// 	"31.11313" +
	// 	"31.11313" +
	// 	"31333313" +
	// 	"31111113" +
	// 	"33333333"

	// square := "" +
	// 	"33333333" +
	// 	"3......3" +
	// 	"3.2222.3" +
	// 	"3.2222.3" +
	// 	"3.2222.3" +
	// 	"3.2222.3" +
	// 	"3......3" +
	// 	"33333333"

	grid := "" +
		"11111111" +
		"1......." +
		"1......." +
		"1......." +
		"1......." +
		"1......." +
		"1......." +
		"1......."

	// sprite := "" +
	// 	"...33..." +
	// 	"..3113.." +
	// 	".313313." +
	// 	".313113." +
	// 	".313113." +
	// 	".313113." +
	// 	"..3113.." +
	// 	"...33..."

	// sprite2 := "" +
	// 	"..3333.." +
	// 	".311113." +
	// 	"31111113" +
	// 	"31111113" +
	// 	"33333333" +
	// 	".312123." +
	// 	".311113." +
	// 	"..3333.."

	arrow := "" +
		"33333..." +
		"322223.." +
		".311113." +
		"..311113" +
		"..311113" +
		".322223." +
		"322223.." +
		"33333..."

	// spritetwo := "" +
	// 	"..3333.." +
	// 	".3....3." +
	// 	"......3." +
	// 	".....33." +
	// 	"...33..." +
	// 	"..3....." +
	// 	".3......" +
	// 	"3333333."

	lcdc := uint8(0)
	lcdc = SetBit(lcdc, 7, 1) // LCD ON
	lcdc = SetBit(lcdc, 6, 1) // Win Tiles = 0x9C00
	lcdc = SetBit(lcdc, 5, 0) // Win OFF
	lcdc = SetBit(lcdc, 4, 0) // BG & win TDT = 0x8800
	lcdc = SetBit(lcdc, 3, 0) // BG tile = 0x9800
	lcdc = SetBit(lcdc, 2, 0) // Sprite size = 8*8
	lcdc = SetBit(lcdc, 1, 1) // Sprite display = ON
	lcdc = SetBit(lcdc, 0, 1) // BG & WIN display = ON
	memory.SetLCDC(lcdc)

	lcd.SetTile(0, lcd.MakeTile(blank), 0x8800)
	lcd.SetTile(1, lcd.MakeTile(grid), 0x8800)

	// lcd.SetTile(1, lcd.MakeTile(sprite), 0x8000)
	lcd.SetTile(0, lcd.MakeTile(arrow), 0x8000)

	spriteFlag := uint8(0x00)
	spriteFlag = SetBit(spriteFlag, 7, 0)
	spriteFlag = SetBit(spriteFlag, 6, 1)
	spriteFlag = SetBit(spriteFlag, 5, 1)
	spriteFlag = SetBit(spriteFlag, 4, 0)
	lcd.SetSprite(0, 200, 30, 0, spriteFlag)
	// lcd.SetSprite(1, 50, 30, 0, 0x80)
	// lcd.SetSprite(2, 53, 33, 0, 0x80)
	// lcd.SetSprite(1, 56, 60, 2, 0x80)

	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			lcd.SetBackgroundTile(i, j, 1)
		}
	}

	// for i := 0; i < 32; i++ {
	// 	for j := 0; j < 1; j++ {
	// 		lcd.SetWindowTile(i, j, 2)
	// 	}
	// 	for j := 17; j < 18; j++ {
	// 		lcd.SetWindowTile(i, j, 2)
	// 	}
	// 	for j := 0; j < 32; j++ {
	// 		lcd.SetWindowTile(0, j, 2)
	// 		lcd.SetWindowTile(19, j, 2)
	// 	}
	// }

	// memory.SetWX(64)
	// memory.SetWY(64)

	// lcd.SetSprite(1, 20, 45, 2, 0x80)
	// lcd.SetSprite(2, 30, 74, 1, 0x80)

	go moveSprites()
	// go fall()
	lcd.Run()
	// lcd.RunProfile()

	lcd.Stop()
}

func moveSprites() {
	tick := time.NewTicker(time.Millisecond * 50)
	for {
		<-tick.C
		x := memory.Get(0xFE03)
		memory.Set(0xFE03, x-1)
	}
}

func fall() {
	tick := time.NewTicker(time.Millisecond * 100)
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
		time.Sleep(1 * time.Microsecond)
	}
}

func signalHandler() {
	for {
		<-sigs
		console.Prompt()
	}
}

func main() {
	// testTetris()
	// return

	sigs = make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go signalHandler()

	memory.ResetMemory()
	lcd.Initialize()
	cpu.Initialize()

	loadRom(os.Args[1])

	cpu.SetPC(0x0000)

	go lcd.Run()
	// go pollSerial()
	console.Prompt()
	cpu.Run()

	lcd.Stop()
}
