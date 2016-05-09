package main

import "lcd"
import t "time"
import "memory"
import "cpu"
import "os"
import "io/ioutil"
// import "log"
import "fmt"

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

    block := ""  +
     "33333333" +
     "31111113" +
     "31....13" +
     "31.11313" +
     "31.11313" +
     "31333313" +
     "31111113" +
     "33333333"

    // square := ""  +
    //  "33333333" +
    //  "3......3" +
    //  "3.2222.3" +
    //  "3.2222.3" +
    //  "3.2222.3" +
    //  "3.2222.3" +
    //  "3......3" +
    //  "33333333"

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

    // lcd.SetTile(0, lcd.MakeTile(blank), 1)
    // lcd.SetTile(1, lcd.MakeTile(square), 1)
    memory.SetLCDC(0x91)
    // lcdc := memory.GetLCDC()
    // lcdc = SetBit(lcdc, 4, 1)
    // lcdc = SetBit(lcdc, 7, 1)
    // memory.SetLCDC(lcdc)
    // lcd.SetTile(2, lcd.MakeTile(t2))

    // x := 10
    y := 0

    go lcd.Run()
    for {
        for x := 0; x < 32; x++ {
            lcd.SetBackgroundTile(x, y, 1)
        }
        y++
        // lcd.SetBackgroundTile(x+1, y, 0)
        // lcd.SetBackgroundTile(x+2, y, 0)
        // lcd.SetBackgroundTile(x+1, y+1, 0)
        // y++
        // lcd.SetBackgroundTile(x, y, 1)
        // lcd.SetBackgroundTile(x+1, y, 1)
        // lcd.SetBackgroundTile(x+2, y, 1)
        // lcd.SetBackgroundTile(x+1, y+1, 1)
        t.Sleep(100 * t.Millisecond)
    }

    lcd.Stop()
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
    // testTetris()

    memory.ResetMemory()
    lcd.Initialize()
    cpu.Initialize()

    loadRom(os.Args[1])

    go lcd.Run()
    // go pollSerial()
    cpu.Run()

    lcd.Stop()
}