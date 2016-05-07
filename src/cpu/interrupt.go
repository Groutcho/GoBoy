package cpu

import . "memory"

var IME = false

// set or unset the Interrupt Master Enable (IME)
func SetIME(value bool) {
	IME = value
}

func handleInterrupt(interrupt int) {	
	if IME && GetInterruptEnable(uint8(interrupt)) {
		SetIME(false)
		SetInterruptRequest(uint8(interrupt), false)
		push(GetPC())
		SetPC(GetInterruptHandler(interrupt))
	}
}

func VBlankInterruptEnabled() bool {
	return GetInterruptEnable(VBLANK)
}

func EnableVBlankInterrupt() {
	SetInterruptEnable(VBLANK, true)
}

func DisableVBlankInterrupt() {
	SetInterruptEnable(VBLANK, false)
}

func LcdStatInterruptEnabled() bool {
	return GetInterruptEnable(LCD_STAT)
}

func EnableLcdStatInterrupt() {
	SetInterruptEnable(LCD_STAT, true)
}

func DisableLcdStatInterrupt() {
	SetInterruptEnable(LCD_STAT, false)
}

func TimerInterruptEnabled() bool {
	return GetInterruptEnable(TIMER)
}

func EnableTimerInterrupt() {
	SetInterruptEnable(TIMER, true)
}

func DisableTimerInterrupt() {
	SetInterruptEnable(TIMER, false)
}

func SerialInterruptEnabled() bool {
	return GetInterruptEnable(SERIAL)
}

func EnableSerialInterrupt() {
	SetInterruptEnable(SERIAL, true)
}

func DisableSerialInterrupt() {
	SetInterruptEnable(SERIAL, false)
}

func JoypadInterruptEnabled() bool {
	return GetInterruptEnable(JOYPAD)
}

func EnableJoypadInterrupt() {
	SetInterruptEnable(JOYPAD, true)
}

func DisableJoypadInterrupt() {
	SetInterruptEnable(JOYPAD, false)
}

/*******************************************
 *         Interrupt requests              *
********************************************/

func RequestVBlankInterrupt() {
	SetInterruptRequest(VBLANK, true)
	handleInterrupt(VBLANK)
}

func RemoveVBlankInterrupt() {
	SetInterruptRequest(VBLANK, false)
}

func VBlankInterruptRequested() bool {
	return GetInterruptRequest(VBLANK)
}

func RequestLcdStatInterrupt() {
	SetInterruptRequest(LCD_STAT, true)
	handleInterrupt(LCD_STAT)
}

func RemoveLcdStatInterrupt() {
	SetInterruptRequest(LCD_STAT, false)
}

func LcdStatInterruptRequested() bool {
	return GetInterruptRequest(LCD_STAT)
}

func RequestTimerInterrupt() {
	SetInterruptRequest(TIMER, true)
	handleInterrupt(TIMER)
}

func RemoveTimerInterrupt() {
	SetInterruptRequest(TIMER, false)
}

func TimerInterruptRequested() bool {
	return GetInterruptRequest(TIMER)
}

func RequestSerialInterrupt() {
	SetInterruptRequest(SERIAL, true)
	handleInterrupt(SERIAL)
}

func RemoveSerialInterrupt() {
	SetInterruptRequest(SERIAL, false)
}

func SerialInterruptRequested() bool {
	return GetInterruptRequest(SERIAL)
}

func RequestJoypadInterrupt() {
	SetInterruptRequest(JOYPAD, true)
	handleInterrupt(JOYPAD)
}

func RemoveJoypadInterrupt() {
	SetInterruptRequest(JOYPAD, false)
}

func JoypadInterruptRequested() bool {
	return GetInterruptRequest(JOYPAD)
}