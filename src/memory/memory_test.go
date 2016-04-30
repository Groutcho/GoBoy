package memory

import "testing"

func TestInterruptsEnables(t *testing.T) {
	var vblank 	bool
	var timer 	bool
	var lcdStat bool
	var joypad 	bool
	var serial	bool

	resetRange(0xFFFF, 0xFFFF)

	EnableVBlankInterrupt()

	vblank 	= VBlankInterruptEnabled()
	timer 	= TimerInterruptEnabled()
	lcdStat = LcdStatInterruptEnabled()
	joypad 	= JoypadInterruptEnabled()
	serial 	= SerialInterruptEnabled()

	if !vblank {
		t.Error("TestInterrupts() test failed: the V-Blank interrupt should be enabled.")
	}

	if timer || lcdStat || joypad || serial {
		t.Error("TestInterrupts() test failed: enabling V-Blank should not toggle other interrupts.")
	}

	DisableVBlankInterrupt()

	vblank 	= VBlankInterruptEnabled()
	timer 	= TimerInterruptEnabled()
	lcdStat = LcdStatInterruptEnabled()
	joypad 	= JoypadInterruptEnabled()
	serial 	= SerialInterruptEnabled()

	if vblank {
		t.Error("TestInterrupts() test failed: the V-Blank interrupt should be disabled.")
	}

	if timer || lcdStat || joypad || serial {
		t.Error("TestInterrupts() test failed: disabling V-Blank should not toggle other interrupts.")
	}

	EnableTimerInterrupt()

	vblank 	= VBlankInterruptEnabled()
	timer 	= TimerInterruptEnabled()
	lcdStat = LcdStatInterruptEnabled()
	joypad 	= JoypadInterruptEnabled()
	serial 	= SerialInterruptEnabled()

	if !timer {
		t.Error("TestInterrupts() test failed: the Timer interrupt should be enabled.")
	}

	if vblank || lcdStat || joypad || serial {
		t.Error("TestInterrupts() test failed: disabling Timer should not toggle other interrupts.")
	}

	DisableTimerInterrupt()

	vblank 	= VBlankInterruptEnabled()
	timer 	= TimerInterruptEnabled()
	lcdStat = LcdStatInterruptEnabled()
	joypad 	= JoypadInterruptEnabled()
	serial 	= SerialInterruptEnabled()

	if timer {
		t.Error("TestInterrupts() test failed: the Timer interrupt should be disabled.")
	}

	if vblank || lcdStat || joypad || serial {
		t.Error("TestInterrupts() test failed: disabling Timer should not toggle other interrupts.")
	}

	EnableLcdStatInterrupt()

	vblank 	= VBlankInterruptEnabled()
	timer 	= TimerInterruptEnabled()
	lcdStat = LcdStatInterruptEnabled()
	joypad 	= JoypadInterruptEnabled()
	serial 	= SerialInterruptEnabled()

	if !lcdStat {
		t.Error("TestInterrupts() test failed: the LCD STAT interrupt should be enabled.")
	}

	if vblank || timer || joypad || serial {
		t.Error("TestInterrupts() test failed: disabling LCD STAT should not toggle other interrupts.")
	}

	DisableLcdStatInterrupt()

	vblank 	= VBlankInterruptEnabled()
	timer 	= TimerInterruptEnabled()
	lcdStat = LcdStatInterruptEnabled()
	joypad 	= JoypadInterruptEnabled()
	serial 	= SerialInterruptEnabled()

	if lcdStat {
		t.Error("TestInterrupts() test failed: the LCD STAT interrupt should be disabled.")
	}

	if vblank || timer || joypad || serial {
		t.Error("TestInterrupts() test failed: disabling LCD STAT should not toggle other interrupts.")
	}

	EnableSerialInterrupt()

	vblank 	= VBlankInterruptEnabled()
	timer 	= TimerInterruptEnabled()
	lcdStat = LcdStatInterruptEnabled()
	joypad 	= JoypadInterruptEnabled()
	serial 	= SerialInterruptEnabled()

	if !serial {
		t.Error("TestInterrupts() test failed: the Serial interrupt should be enabled.")
	}

	if vblank || timer || joypad || lcdStat {
		t.Error("TestInterrupts() test failed: disabling Serial should not toggle other interrupts.")
	}

	DisableSerialInterrupt()

	vblank 	= VBlankInterruptEnabled()
	timer 	= TimerInterruptEnabled()
	lcdStat = LcdStatInterruptEnabled()
	joypad 	= JoypadInterruptEnabled()
	serial 	= SerialInterruptEnabled()

	if serial {
		t.Error("TestInterrupts() test failed: the Serial interrupt should be disabled.")
	}

	if vblank || timer || joypad || lcdStat {
		t.Error("TestInterrupts() test failed: disabling Serial should not toggle other interrupts.")
	}

	EnableJoypadInterrupt()

	vblank 	= VBlankInterruptEnabled()
	timer 	= TimerInterruptEnabled()
	lcdStat = LcdStatInterruptEnabled()
	joypad 	= JoypadInterruptEnabled()
	serial 	= SerialInterruptEnabled()

	if !joypad {
		t.Error("TestInterrupts() test failed: the Joypad interrupt should be enabled.")
	}

	if vblank || timer || serial || lcdStat {
		t.Error("TestInterrupts() test failed: disabling Joypad should not toggle other interrupts.")
	}

	DisableJoypadInterrupt()

	vblank 	= VBlankInterruptEnabled()
	timer 	= TimerInterruptEnabled()
	lcdStat = LcdStatInterruptEnabled()
	joypad 	= JoypadInterruptEnabled()
	serial 	= SerialInterruptEnabled()

	if joypad {
		t.Error("TestInterrupts() test failed: the Joypad interrupt should be disabled.")
	}

	if vblank || timer || serial || lcdStat {
		t.Error("TestInterrupts() test failed: disabling Joypad should not toggle other interrupts.")
	}
}

func TestInterruptsRequests(t *testing.T) {
	var vblank 	bool
	var timer 	bool
	var lcdStat bool
	var joypad 	bool
	var serial	bool

	resetRange(0xFFFF, 0xFFFF)

	RequestVBlankInterrupt()

	vblank 	= VBlankInterruptRequested()
	timer 	= TimerInterruptRequested()
	lcdStat = LcdStatInterruptRequested()
	joypad 	= JoypadInterruptRequested()
	serial 	= SerialInterruptRequested()

	if !vblank {
		t.Error("TestInterrupts() test failed: the V-Blank interrupt should be Requested.")
	}

	if timer || lcdStat || joypad || serial {
		t.Error("TestInterrupts() test failed: enabling V-Blank should not toggle other interrupts.")
	}

	RemoveVBlankInterrupt()

	vblank 	= VBlankInterruptRequested()
	timer 	= TimerInterruptRequested()
	lcdStat = LcdStatInterruptRequested()
	joypad 	= JoypadInterruptRequested()
	serial 	= SerialInterruptRequested()

	if vblank {
		t.Error("TestInterrupts() test failed: the V-Blank interrupt should be disabled.")
	}

	if timer || lcdStat || joypad || serial {
		t.Error("TestInterrupts() test failed: disabling V-Blank should not toggle other interrupts.")
	}

	RequestTimerInterrupt()

	vblank 	= VBlankInterruptRequested()
	timer 	= TimerInterruptRequested()
	lcdStat = LcdStatInterruptRequested()
	joypad 	= JoypadInterruptRequested()
	serial 	= SerialInterruptRequested()

	if !timer {
		t.Error("TestInterrupts() test failed: the Timer interrupt should be Requested.")
	}

	if vblank || lcdStat || joypad || serial {
		t.Error("TestInterrupts() test failed: disabling Timer should not toggle other interrupts.")
	}

	RemoveTimerInterrupt()

	vblank 	= VBlankInterruptRequested()
	timer 	= TimerInterruptRequested()
	lcdStat = LcdStatInterruptRequested()
	joypad 	= JoypadInterruptRequested()
	serial 	= SerialInterruptRequested()

	if timer {
		t.Error("TestInterrupts() test failed: the Timer interrupt should be Removed.")
	}

	if vblank || lcdStat || joypad || serial {
		t.Error("TestInterrupts() test failed: disabling Timer should not toggle other interrupts.")
	}

	RequestLcdStatInterrupt()

	vblank 	= VBlankInterruptRequested()
	timer 	= TimerInterruptRequested()
	lcdStat = LcdStatInterruptRequested()
	joypad 	= JoypadInterruptRequested()
	serial 	= SerialInterruptRequested()

	if !lcdStat {
		t.Error("TestInterrupts() test failed: the LCD STAT interrupt should be Requested.")
	}

	if vblank || timer || joypad || serial {
		t.Error("TestInterrupts() test failed: disabling LCD STAT should not toggle other interrupts.")
	}

	RemoveLcdStatInterrupt()

	vblank 	= VBlankInterruptRequested()
	timer 	= TimerInterruptRequested()
	lcdStat = LcdStatInterruptRequested()
	joypad 	= JoypadInterruptRequested()
	serial 	= SerialInterruptRequested()

	if lcdStat {
		t.Error("TestInterrupts() test failed: the LCD STAT interrupt should be Removed.")
	}

	if vblank || timer || joypad || serial {
		t.Error("TestInterrupts() test failed: disabling LCD STAT should not toggle other interrupts.")
	}

	RequestSerialInterrupt()

	vblank 	= VBlankInterruptRequested()
	timer 	= TimerInterruptRequested()
	lcdStat = LcdStatInterruptRequested()
	joypad 	= JoypadInterruptRequested()
	serial 	= SerialInterruptRequested()

	if !serial {
		t.Error("TestInterrupts() test failed: the Serial interrupt should be Requested.")
	}

	if vblank || timer || joypad || lcdStat {
		t.Error("TestInterrupts() test failed: disabling Serial should not toggle other interrupts.")
	}

	RemoveSerialInterrupt()

	vblank 	= VBlankInterruptRequested()
	timer 	= TimerInterruptRequested()
	lcdStat = LcdStatInterruptRequested()
	joypad 	= JoypadInterruptRequested()
	serial 	= SerialInterruptRequested()

	if serial {
		t.Error("TestInterrupts() test failed: the Serial interrupt should be Removed.")
	}

	if vblank || timer || joypad || lcdStat {
		t.Error("TestInterrupts() test failed: disabling Serial should not toggle other interrupts.")
	}

	RequestJoypadInterrupt()

	vblank 	= VBlankInterruptRequested()
	timer 	= TimerInterruptRequested()
	lcdStat = LcdStatInterruptRequested()
	joypad 	= JoypadInterruptRequested()
	serial 	= SerialInterruptRequested()

	if !joypad {
		t.Error("TestInterrupts() test failed: the Joypad interrupt should be Requested.")
	}

	if vblank || timer || serial || lcdStat {
		t.Error("TestInterrupts() test failed: disabling Joypad should not toggle other interrupts.")
	}

	RemoveJoypadInterrupt()

	vblank 	= VBlankInterruptRequested()
	timer 	= TimerInterruptRequested()
	lcdStat = LcdStatInterruptRequested()
	joypad 	= JoypadInterruptRequested()
	serial 	= SerialInterruptRequested()

	if joypad {
		t.Error("TestInterrupts() test failed: the Joypad interrupt should be disabled.")
	}

	if vblank || timer || serial || lcdStat {
		t.Error("TestInterrupts() test failed: disabling Joypad should not toggle other interrupts.")
	}
}

func TestGet(t *testing.T) {
	RAM[0xFF21] = 0x28

	if b := Get(0xFF21); b != 0x28 {
		t.Errorf("TestGet() failed: expected 0x28, got 0x%04x", b)
	}
}

func TestSet(t *testing.T) {
	Set(0x5247, 0x99)

	if b := Get(0x5247); b != 0x99 {
		t.Errorf("TestGet() failed: expected 39, got 0x%04x", b)
	}
}

func TestSetRange(t *testing.T) {
	var from uint16 = 0x405F
	var to uint16 = 0x505F
	size := to - from + 1

	data := make([]byte, int(size))
	data[0x00] = 0x66
	// data[0x55] = 0x77
	// data[0xF0] = 0x79

	SetRange(from, to, data)

	actual := GetRange(from, to)

	if actual[0x00] != 0x66 {
		t.Errorf("TestSetRange() failed: at 0x00, expected 0x66, got 0x%04x", actual[0x50])
	}
}