package memory

/* the game boy address space is 16bit wide */
var RAM = new([65536]byte)

/* Interrupt bit indices */
const VBLANK = 0
const LCD_STAT = 1
const TIMER = 2
const SERIAL = 3
const JOYPAD = 4

const INTRPT_ENABLE_ADDR = 0xFFFF
const INTRPT_REQUEST_ADDR = 0xFF0F

/* reset the RAM to 0 */
func reset() {
	for i := 0; i < len(RAM); i++ {
		RAM[i] = 0x00
	}
}

/* reset a range of RAM addresses */
func resetRange(start uint16, end uint16) {
	if start == end {
		RAM[start] = 0x00
		return
	}

	for i := start; i <= end; i++ {
		// prevent uint16 overflow,
		// leading to an infinite loop
		if i == 0xFFFF {
			RAM[i] = 0x00
			return
		}
		RAM[i] = 0x00
	}
}

func getInterruptEnable(bit uint8) bool {
	return (RAM[INTRPT_ENABLE_ADDR] & (0x01 << bit)) != 0
}

func setInterruptEnable(bit uint8, value bool) {
	if value {
		RAM[INTRPT_ENABLE_ADDR] |= (0x01 << bit)
	} else {
		RAM[INTRPT_ENABLE_ADDR] &= ^(uint8(0x01) << bit)
	}
}

func getInterruptRequest(bit uint8) bool {
	return (RAM[INTRPT_REQUEST_ADDR] & (0x01 << bit)) != 0
}

func setInterruptRequest(bit uint8, value bool) {
	if value {
		RAM[INTRPT_REQUEST_ADDR] |= (0x01 << bit)
	} else {
		RAM[INTRPT_REQUEST_ADDR] &= ^(uint8(0x01) << bit)
	}
}

/*******************************************
 *         Interrupt enables               *
********************************************/

func VBlankInterruptEnabled() bool {
	return getInterruptEnable(VBLANK)
}

func EnableVBlankInterrupt() {
	setInterruptEnable(VBLANK, true)
}

func DisableVBlankInterrupt() {
	setInterruptEnable(VBLANK, false)
}

func LcdStatInterruptEnabled() bool {
	return getInterruptEnable(LCD_STAT)
}

func EnableLcdStatInterrupt() {
	setInterruptEnable(LCD_STAT, true)
}

func DisableLcdStatInterrupt() {
	setInterruptEnable(LCD_STAT, false)
}

func TimerInterruptEnabled() bool {
	return getInterruptEnable(TIMER)
}

func EnableTimerInterrupt() {
	setInterruptEnable(TIMER, true)
}

func DisableTimerInterrupt() {
	setInterruptEnable(TIMER, false)
}

func SerialInterruptEnabled() bool {
	return getInterruptEnable(SERIAL)
}

func EnableSerialInterrupt() {
	setInterruptEnable(SERIAL, true)
}

func DisableSerialInterrupt() {
	setInterruptEnable(SERIAL, false)
}

func JoypadInterruptEnabled() bool {
	return getInterruptEnable(JOYPAD)
}

func EnableJoypadInterrupt() {
	setInterruptEnable(JOYPAD, true)
}

func DisableJoypadInterrupt() {
	setInterruptEnable(JOYPAD, false)
}

/*******************************************
 *         Interrupt requests              *
********************************************/

func RequestVBlankInterrupt() {
	setInterruptRequest(VBLANK, true)
}

func RemoveVBlankInterrupt() {
	setInterruptRequest(VBLANK, false)
}

func VBlankInterruptRequested() bool {
	return getInterruptRequest(VBLANK)
}

func RequestLcdStatInterrupt() {
	setInterruptRequest(LCD_STAT, true)
}

func RemoveLcdStatInterrupt() {
	setInterruptRequest(LCD_STAT, false)
}

func LcdStatInterruptRequested() bool {
	return getInterruptRequest(LCD_STAT)
}

func RequestTimerInterrupt() {
	setInterruptRequest(TIMER, true)
}

func RemoveTimerInterrupt() {
	setInterruptRequest(TIMER, false)
}

func TimerInterruptRequested() bool {
	return getInterruptRequest(TIMER)
}

func RequestSerialInterrupt() {
	setInterruptRequest(SERIAL, true)
}

func RemoveSerialInterrupt() {
	setInterruptRequest(SERIAL, false)
}

func SerialInterruptRequested() bool {
	return getInterruptRequest(SERIAL)
}

func RequestJoypadInterrupt() {
	setInterruptRequest(JOYPAD, true)
}

func RemoveJoypadInterrupt() {
	setInterruptRequest(JOYPAD, false)
}

func JoypadInterruptRequested() bool {
	return getInterruptRequest(JOYPAD)
}