package main

import "lcd"
import t "time"
import "memory"
import "cpu"
import . "common"

func main() {

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

    square := ""  +
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

    lcd.SetTile(0, lcd.MakeTile(blank), 0)
    lcd.SetTile(1, lcd.MakeTile(block), 0)
    
    lcd.SetTile(0, lcd.MakeTile(blank), 1)
    lcd.SetTile(1, lcd.MakeTile(square), 1)
    lcdc := memory.GetLCDC()
    lcdc = SetBit(lcdc, 4, 0)
    lcdc = SetBit(lcdc, 7, 1)
    memory.SetLCDC(lcdc)
    // lcd.SetTile(2, lcd.MakeTile(t2))

    x := 10
    y := 0

    for {
        lcd.SetBackgroundTile(x, y, 0)
        lcd.SetBackgroundTile(x+1, y, 0)
        lcd.SetBackgroundTile(x+2, y, 0)
        lcd.SetBackgroundTile(x+1, y+1, 0)
        y++
        lcd.SetBackgroundTile(x, y, 1)
        lcd.SetBackgroundTile(x+1, y, 1)
        lcd.SetBackgroundTile(x+2, y, 1)
        lcd.SetBackgroundTile(x+1, y+1, 1)
        lcd.Update()
        t.Sleep(400 * t.Millisecond)
    }

    lcd.Stop()
}