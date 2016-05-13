package lcd

import (
	. "common"
	"cpu"
	"github.com/veandco/go-sdl2/sdl"
	// "log"
	mem "memory"
	"os"
	"runtime/pprof"
	"time"
	"unsafe"
)

const (
	WHITE      uint32 = 0xFFFFFFFF
	DARK_GREY  uint32 = 0x44444444
	LIGHT_GREY uint32 = 0xAAAAAAAA
	BLACK      uint32 = 0x00000000
	LCD_WIDTH         = 160
	LCD_HEIGHT        = 144
	SCANLINES         = 153

	SCY_ADDR = 0xFF42
	SCX_ADDR = 0xFF43

	LCD_ACTIVE    = 7
	WDW_ACTIVE    = 5
	WDW_MAP       = 6
	TDT           = 4
	BG_MAP        = 3
	BG_WDW_ACTIVE = 0
	SCALE         = 2
)

var (
	window    *sdl.Window
	renderer  *sdl.Renderer
	screenTex *sdl.Texture
	pixels    = new([LCD_HEIGHT * LCD_WIDTH]uint32)

	palette = new([4]uint32)

	mmap       []byte = nil
	fillBuffer        = 0
	showBuffer        = 1
)

func init() {
	palette[0] = WHITE
	palette[1] = LIGHT_GREY
	palette[2] = DARK_GREY
	palette[3] = BLACK

	mmap = mem.GetMemoryMap()
}

func Initialize() {
	var err error

	sdl.Init(sdl.INIT_EVERYTHING)

	window, err = sdl.CreateWindow("goboy", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		LCD_WIDTH*SCALE, LCD_HEIGHT*SCALE, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	screenTex, err = renderer.CreateTexture(sdl.PIXELFORMAT_ARGB8888, sdl.TEXTUREACCESS_STREAMING, LCD_WIDTH, LCD_HEIGHT)
	if err != nil {
		panic(err)
	}

	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()
}

func RunProfile() {
	f, _ := os.Create("cpuprof.txt")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < 1000; i++ {
		redraw()
	}
}

func Run() {
	for {
		redraw()
	}
}

func Stop() {
	screenTex.Destroy()
	window.Destroy()
	sdl.Quit()
}

func SetTile(index int, tile []byte, mode int) {
	base := uint16(0x8000)
	if mode == 1 {
		base = uint16(0x8800)
	}
	offset := uint16(len(tile)) * uint16(index)
	mem.SetRange(base+offset, tile)
}

func SetBackgroundTile(x, y, index int) {
	addr := uint16(0x9800)
	if IsBitSet(mem.GetLCDC(), BG_MAP) {
		addr = uint16(0x9C00)
	}

	mem.Set(addr+uint16(y*32+x), uint8(index))
}

func SetWindowTile(x, y, index int) {
	addr := uint16(0x9800)
	if IsBitSet(mem.GetLCDC(), WDW_MAP) {
		addr = uint16(0x9C00)
	}

	mem.Set(addr+uint16(y*32+x), uint8(index))
}

func setBufferPixel(x, y, color int, pixels unsafe.Pointer, pitch int) {
	(*[LCD_WIDTH * LCD_HEIGHT]uint32)(pixels)[y*(pitch/4)+x] = palette[color]
}

func getBackgroundTileMap() uint16 {
	addr := uint16(0x9800)
	if IsBitSet(mem.GetLCDC(), BG_MAP) {
		addr = uint16(0x9C00)
	}

	return addr
}

func getWindowTileMap() uint16 {
	addr := uint16(0x9800)
	if IsBitSet(mem.GetLCDC(), WDW_MAP) {
		addr = uint16(0x9C00)
	}

	return addr
}

func getTileDataTable() uint16 {
	tdt := uint16(0x8000)
	if IsBitSet(mem.GetLCDC(), TDT) {
		tdt = uint16(0x8800)
	}
	return tdt
}

func drawWindowLine(y int, mapAddr, tileAddr uint16, pixels unsafe.Pointer, pitch int) {
	x := int(mem.GetWX())
	yy := y + int(mem.GetWY())

	for i := 0; i < LCD_WIDTH; i++ {
		pix := getTilePixel(x, yy, mapAddr, tileAddr)
		setBufferPixel(x, y, pix, pixels, pitch)
		x++
	}
}

// Draw a single line of background
func drawBackgroundLine(y int, mapAddr, tileAddr uint16, pixels unsafe.Pointer, pitch int) {
	x := int(mmap[SCX_ADDR])
	yy := y + int(mmap[SCY_ADDR])

	for i := 0; i < LCD_WIDTH; i++ {
		pix := getTilePixel(x%256, yy%256, mapAddr, tileAddr)
		setBufferPixel(x, y, pix, pixels, pitch)
		x++
	}
}

// Return the color of the background pixel at coordinates x, y
func getTilePixel(x, y int, bgAddr, tileAddr uint16) int {
	// Get the tile corresponding to this coordinate
	tx := x / 8
	ty := y / 8

	// get the tile index in the tile map
	tIndexOffset := ty*32 + tx
	tIndex := mmap[bgAddr+uint16(tIndexOffset)]

	// Get the pixel x,y in the tile itself
	px := x % 8
	py := y % 8

	// get the tile address in the tile data table
	addr := tileAddr + uint16(tIndex*16) + uint16(px>>2+py<<1)

	h := 7 - (px%4)*2
	l := h - 1
	b := mmap[addr]

	color := 2*GetBit(b, uint8(h)) + GetBit(b, uint8(l))
	return int(color)
}

// Draw a single frame (144 lines + 10 "lines" of V-Blank (approx 1.1 ms))
func redraw() {
	lcdc := mem.GetLCDC()

	if !IsBitSet(lcdc, LCD_ACTIVE) {
		return
	}

	var pitch int
	var pixPtr unsafe.Pointer
	err := screenTex.Lock(nil, &pixPtr, &pitch)
	if err != nil {
		panic(err)
	}

	bgAddr := getBackgroundTileMap()
	wdwAddr := getWindowTileMap()
	tileAddr := getTileDataTable()

	drawBgWindow := IsBitSet(lcdc, BG_WDW_ACTIVE)
	drawWindow := IsBitSet(lcdc, WDW_ACTIVE)

	mem.SetLY(0x00)

	// 108 microseconds is the duration of a scanline draw
	tick := time.NewTicker(time.Millisecond * 16)
	defer tick.Stop()

	for y := 0; y < SCANLINES; y++ {
		if y < LCD_HEIGHT {
			if drawBgWindow {
				drawBackgroundLine(y, bgAddr, tileAddr, pixPtr, pitch)
				if drawWindow {
					drawWindowLine(y, wdwAddr, tileAddr, pixPtr, pitch)
				}
			}
		} else if y == LCD_HEIGHT {
			cpu.RequestVBlankInterrupt()
		}
		mem.IncLY()
	}

	<-tick.C
	screenTex.Update(nil, pixPtr, pitch)
	screenTex.Unlock()
	renderer.Clear()
	renderer.Copy(screenTex, nil, nil)
	renderer.Present()
}
