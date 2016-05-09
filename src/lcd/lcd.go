package lcd

import . "common"
import "github.com/veandco/go-sdl2/sdl"
import mem "memory"
import t "time"
import "cpu"

// import "log"

const WHITE uint32 = 0xFFFFFFFF
const DARK_GREY uint32 = 0x44444444
const LIGHT_GREY uint32 = 0xAAAAAAAA
const BLACK uint32 = 0x00000000
const LCD_WIDTH = 256
const LCD_HEIGHT = 256

var window *sdl.Window
var surface *sdl.Surface

var palette = new([4]uint32)

const LCD_ACTIVE = 7
const WDW_ACTIVE = 5
const WDW_MAP = 6
const TDT = 4
const BG_MAP = 3
const BG_WDW_ACTIVE = 0

func init() {
	palette[0] = WHITE
	palette[1] = LIGHT_GREY
	palette[2] = DARK_GREY
	palette[3] = BLACK
}

func fillScreen(color uint32) {
	rect := sdl.Rect{0, 0, int32(LCD_WIDTH), int32(LCD_HEIGHT)}
	surface.FillRect(&rect, color)
}

func SetPixel(x, y int32, color uint32) {
	rect := sdl.Rect{x, y, 1, 1}
	surface.FillRect(&rect, color)
}

func GetTileColor(tile []byte, x, y int) uint32 {
	B := x / 4 + y * 2
	h := 7 - (x % 4) * 2
	l := h - 1

	color := 2 * GetBit(tile[B], uint8(h)) + GetBit(tile[B], uint8(l))
	return palette[color]
}

func DrawTile(tile []byte, x, y int32) {
	for xx := 0; xx < 8; xx++ {
		for yy := 0; yy < 8; yy++ {
			SetPixel(x + int32(xx), y + int32(yy), GetTileColor(tile, xx, yy))
		}
	}
}

func Initialize() {
	var err error

 	sdl.Init(sdl.INIT_EVERYTHING)

	window, err = sdl.CreateWindow("goboy", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
	    LCD_WIDTH, LCD_HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
	    panic(err)
	}

	surface, err = window.GetSurface()
	if err != nil {
	    panic(err)
	}

    fillScreen(WHITE)
}

func GetTile(index uint8, base uint16) []byte {
	addr := uint16(base + uint16(index) * 16)
	return mem.GetRange(addr, 16)
}

func DrawBackground(bgAddress uint16, tileAddr uint16) {
	var x int32 = 0
	var y int32 = 0
	for addr := bgAddress; addr < (bgAddress + 1024); addr++ {
		tile := GetTile(mem.Get(addr), tileAddr)
		DrawTile(tile, x*8, y*8)

		x++
		if x > 31 {
			x = 0
			y++
		}
	}
}

func SetTile(index int, tile []byte, mode int) {
	base := uint16(0x8000)
	if mode == 1 {
		base = uint16(0x8800)
	}
	offset := uint16(len(tile)) * uint16(index)
	mem.SetRange(base + offset, tile)
}

func SetBackgroundTile(x, y, index int) {
	mem.Set(uint16(0x9800 + y * 32 + x), uint8(index))
}

func GetBackgroundTileMap() uint16 {
	addr := uint16(0x9800)
	if IsBitSet(mem.GetLCDC(), BG_MAP) {
		addr = uint16(0x9C00)
	}

	return addr
}

func GetWindowTileMap() uint16 {
	addr := uint16(0x9800)
	if IsBitSet(mem.GetLCDC(), WDW_MAP) {
		addr = uint16(0x9C00)
	}

	return addr
}

func GetTileDataTable() uint16 {
	tdt := uint16(0x8000)
	if IsBitSet(mem.GetLCDC(), TDT) {
		tdt = uint16(0x8800)
	}
	return tdt
}

func DrawWindow() {
	// TODO
}

func Run() {
	for {
		Redraw()
	}
}

// Draw a single frame (144 lines + 10 "lines" of V-Blank (approx 1.1 ms))
func Redraw() {
	Update()

	mem.SetLY(0x00)

	for i := 0; i < 154; i++ {
		// Draw a single scanline
		mem.IncLY()
		if mem.GetLY() == 144 {
			cpu.RequestVBlankInterrupt()
		}
		t.Sleep(108 * t.Microsecond)
	}

	window.UpdateSurface()
}

func Update() {
	lcdc := mem.GetLCDC()

	if IsBitSet(lcdc, LCD_ACTIVE) {
		bgAddr := GetBackgroundTileMap()
		// wdwAddr := GetWindowTileMap()
		tileAddr := GetTileDataTable()

		if IsBitSet(lcdc, BG_WDW_ACTIVE) {
			DrawBackground(bgAddr, tileAddr)
			if IsBitSet(lcdc, WDW_ACTIVE) {
				DrawWindow()
			}
		}
	}
}

func Stop() {
	window.Destroy()
	sdl.Quit()
}