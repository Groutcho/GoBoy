package lcd

import . "common"
import "github.com/veandco/go-sdl2/sdl"
import mem "memory"
import t "time"
import "cpu"

const WHITE uint8 = 0xFF
const DARK_GREY uint8 = 0x44
const LIGHT_GREY uint8 = 0xAA
const BLACK uint8 = 0x00
const LCD_WIDTH = 160
const LCD_HEIGHT = 144

var window *sdl.Window
var renderer *sdl.Renderer

var palette = new([4]uint8)

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
	rect := sdl.Rect{0, 0, LCD_WIDTH, LCD_HEIGHT}
	renderer.SetDrawColor(255, 255, 255, 255) // r, g, b, a uint8)
	renderer.FillRect(&rect)
	renderer.Present()
}

func SetPixel(x, y int, color uint8) {
	renderer.SetDrawColor(color, color, color, 255) // r, g, b, a uint8)
	renderer.DrawPoint(x, y)
	// renderer.Present()
}

func GetTileColor(tile []byte, x, y int) uint8 {
	B := x / 4 + y * 2
	h := 7 - (x % 4) * 2
	l := h - 1

	color := 2 * GetBit(tile[B], uint8(h)) + GetBit(tile[B], uint8(l))
	return palette[color]
}

func Initialize() {
	var err error

 	sdl.Init(sdl.INIT_EVERYTHING)

	window, err = sdl.CreateWindow("goboy", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
	    LCD_WIDTH, LCD_HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
	    panic(err)
	}

	renderer, err = sdl.CreateRenderer(window, -1, 0)
	if err != nil {
	    panic(err)
	}

    fillScreen(0xFFFFFFFF)
}

func GetTile(index uint8, base uint16) []byte {
	addr := uint16(base + uint16(index) * 16)
	return mem.GetRange(addr, 16)
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
	addr := uint16(0x9800)
	if IsBitSet(mem.GetLCDC(), BG_MAP) {
		addr = uint16(0x9C00)
	}

	mem.Set(addr + uint16(y * 32 + x), uint8(index))
}

func SetWindowTile(x, y, index int) {
	addr := uint16(0x9800)
	if IsBitSet(mem.GetLCDC(), WDW_MAP) {
		addr = uint16(0x9C00)
	}

	mem.Set(addr + uint16(y * 32 + x), uint8(index))
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

func Run() {
	for {
		Redraw()
	}
}

func DrawWindowLine(y int, wdwAddr, tileAddr uint16) {
	x := int(mem.GetWX())
	yy := y + int(mem.GetWY())

	for i := 0; i < LCD_WIDTH; i++ {
		pix := GetBgPixel(x, yy, wdwAddr, tileAddr)
		SetPixel(x, y, pix)
		x++
	}
}

// Draw a single line of background
func DrawBgLine(y int, bgAddr, tileAddr uint16) {
	x := int(mem.GetSCX())
	yy := y + int(mem.GetSCY())

	for i := 0; i < LCD_WIDTH; i++ {
		pix := GetBgPixel(x % 256, yy % 256, bgAddr, tileAddr)
		SetPixel(x, y, pix)
		x++
	}
}

// Return the color of the background pixel at coordinates x, y
func GetBgPixel(x, y int, bgAddr, tileAddr uint16) uint8 {
	// Get the tile corresponding to this coordinate
	tx := x / 8
	ty := y / 8

	tIndexOffset := ty * 32 + tx
	tIndex := mem.Get(bgAddr + uint16(tIndexOffset))

	tile := GetTile(tIndex, tileAddr)

	// Get the pixel x,y in the tile itself
	px := x % 8
	py := y % 8

	return GetTileColor(tile, int(px), int(py))
}

// Draw a single frame (144 lines + 10 "lines" of V-Blank (approx 1.1 ms))
func Redraw() {
	lcdc := mem.GetLCDC()

	if !IsBitSet(lcdc, LCD_ACTIVE) {
		return
	}

	bgAddr := GetBackgroundTileMap()
	wdwAddr := GetWindowTileMap()
	tileAddr := GetTileDataTable()

	drawBgWindow :=	IsBitSet(lcdc, BG_WDW_ACTIVE)
	drawWindow := IsBitSet(lcdc, WDW_ACTIVE)

	mem.SetLY(0x00)

	for y := 0; y < 154; y++ {
		if y < LCD_HEIGHT {
			if drawBgWindow {
				DrawBgLine(int(y), bgAddr, tileAddr)
				if drawWindow {
					DrawWindowLine(int(y), wdwAddr, tileAddr)
				}
			}
		} else if y == LCD_HEIGHT {
			renderer.Present()
			cpu.RequestVBlankInterrupt()
		}

		t.Sleep(108 * t.Microsecond)
		mem.IncLY()
	}
}

func Stop() {
	window.Destroy()
	sdl.Quit()
}