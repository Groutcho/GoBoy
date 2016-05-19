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

const (
	white = "" +
		"........" +
		"........" +
		"........" +
		"........" +
		"........" +
		"........" +
		"........" +
		"........"

	lgrey = "" +
		"11111111" +
		"11111111" +
		"11111111" +
		"11111111" +
		"11111111" +
		"11111111" +
		"11111111" +
		"11111111"

	dgrey = "" +
		"22222222" +
		"22222222" +
		"22222222" +
		"22222222" +
		"22222222" +
		"22222222" +
		"22222222" +
		"22222222"

	black = "" +
		"33333333" +
		"33333333" +
		"33333333" +
		"33333333" +
		"33333333" +
		"33333333" +
		"33333333" +
		"33333333"

	pattern = "" +
		"...33..." +
		"..3..3.." +
		"33....33" +
		"33......" +
		"33333..." +
		"....3333" +
		"......33" +
		"...3333."

	cross = "" +
		"33111113" +
		".331113." +
		"..3313.." +
		"...33..." +
		"..3233.." +
		".322233." +
		"32222233" +
		"22222222"

	block = "" +
		"33333333" +
		"31111113" +
		"31....13" +
		"31.11313" +
		"31.11313" +
		"31333313" +
		"31111113" +
		"33333333"
)

var (
	sigs chan os.Signal
)

func loadRom(filename string) {
	rom, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	cpu.LoadRom(rom)
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
	sigs = make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	for {
		<-sigs
		console.Prompt()
	}
}

func showCheckerboard() {
	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			if j%2 == 0 {
				if i%2 == 0 {
					lcd.SetBackgroundTile(i, j, 0)
				} else {
					lcd.SetBackgroundTile(i, j, 3)
				}
			} else {
				if i%2 == 0 {
					lcd.SetBackgroundTile(i, j, 3)
				} else {
					lcd.SetBackgroundTile(i, j, 0)
				}

			}
		}
	}
	time.Sleep(time.Second * 1)
}

func showColors() {
	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			lcd.SetBackgroundTile(i, j, 0)
		}
	}

	time.Sleep(time.Second * 1)

	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			lcd.SetBackgroundTile(i, j, 1)
		}
	}

	time.Sleep(time.Second * 1)

	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			lcd.SetBackgroundTile(i, j, 2)
		}
	}

	time.Sleep(time.Second * 1)

	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			lcd.SetBackgroundTile(i, j, 3)
		}
	}

	time.Sleep(time.Second * 1)
}

func showCrosses() {
	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			lcd.SetBackgroundTile(i, j, 0)
		}
	}

	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			lcd.SetBackgroundTile(i, j, 4)
		}
	}
	time.Sleep(time.Second * 2)
}

func showScatter() {
	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			if i%2 == 1 {
				lcd.SetBackgroundTile(i, j, 2)
			} else if i%3 == 1 {
				lcd.SetBackgroundTile(i, j, 4)
			} else if i%4 == 1 {
				lcd.SetBackgroundTile(i, j, 5)
			}
		}
	}
}

func runTests() {
	lcd.SetTile(0, lcd.MakeTile(white), 0x8000)
	lcd.SetTile(1, lcd.MakeTile(lgrey), 0x8000)
	lcd.SetTile(2, lcd.MakeTile(dgrey), 0x8000)
	lcd.SetTile(3, lcd.MakeTile(black), 0x8000)
	lcd.SetTile(4, lcd.MakeTile(cross), 0x8000)
	lcd.SetTile(5, lcd.MakeTile(block), 0x8000)

	go lcd.Run()
	time.Sleep(time.Second * 1)

	showColors()
	showCheckerboard()
	showCrosses()
	showScatter()

	time.Sleep(time.Second * 5)
}

func run() {
	loadRom(os.Args[1])

	go signalHandler()
	go lcd.Run()

	cpu.Run()

	lcd.Stop()
}

func main() {
	memory.ResetMemory()
	cpu.Initialize()
	lcd.Initialize()

	run()
}
