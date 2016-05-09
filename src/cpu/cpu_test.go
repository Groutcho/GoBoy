package cpu

import "testing"
import . "memory"

func TestFetch(t* testing.T) {
	arr := make([]byte, 4)
	arr[0] = 0x00
	arr[1] = 0x58
	arr[2] = 0xCB
	arr[3] = 0x98

	LoadProgram(arr)
	SetPC(0x0000)

	startingPC := GetPC()
	currentPC := startingPC

	if opcode, _ := Fetch(); opcode != 0x00 {
		t.Errorf("TestFetch(): at first Fetch(), expected 00, got %02x", opcode)
	}

	if currentPC = GetPC(); currentPC != (startingPC + 1) {
		t.Errorf("TestFetch(): at first Fetch(), expected PC @ 0x%04x, got 0x%04x", (startingPC + 1), currentPC)
	}

	if opcode, _ := Fetch(); opcode != 0x58 {
		t.Errorf("TestFetch(): at second Fetch(), expected 58, got %02x", opcode)
	}

	if currentPC = GetPC(); currentPC != (startingPC + 2) {
		t.Errorf("TestFetch(): at second Fetch(), expected PC @ 0x%04x, got 0x%04x", (startingPC + 2), currentPC)
	}

	if opcode, _ := Fetch(); opcode != 0x197 {
		t.Errorf("TestFetch(): at last Fetch(), expected extended opcode conversion CB 98 -> 197, got %02x", opcode)
	}

	if currentPC = GetPC(); currentPC != (startingPC + 4) {
		t.Errorf("TestFetch(): at last Fetch(), expected PC @ 0x%04x, got 0x%04x", (startingPC + 4), currentPC)
	}
}

func op0() int {
	SetA(0x27)
	return 1
}

func op1() int {
	SetA(0x44)
	return 1
}

func op2() int {
	SetA(0x38)
	return 1
}

func op3() int {
	SetA(0xFF)
	return 1
}

func TestExecuteNext(t* testing.T) {
	SetPC(0x0000)

	arr := make([]byte, 4)
	arr[0] = 0xF8
	arr[1] = 0xF9
	arr[2] = 0xFA
	arr[3] = 0xFB

	LoadProgram(arr)
	SetPC(0x0000)

	dispatch_table[0xF8] = op0
	dispatch_table[0xF9] = op1
	dispatch_table[0xFA] = op2
	dispatch_table[0xFB] = op3

	ExecuteNext()
	if a := GetA(); a != 0x27 {
		t.Errorf("TestExecuteNext(): expected value 0x27 in register A, got 0x%02x", a)
	}

	ExecuteNext()
	if a := GetA(); a != 0x44 {
		t.Errorf("TestExecuteNext(): expected value 0x44 in register A, got 0x%02x", a)
	}

	ExecuteNext()
	if a := GetA(); a != 0x38 {
		t.Errorf("TestExecuteNext(): expected value 0x38 in register A, got 0x%02x", a)
	}

	ExecuteNext()
	if a := GetA(); a != 0xFF {
		t.Errorf("TestExecuteNext(): expected value 0xFF in register A, got 0x%02x", a)
	}
}

// Test a conditional jump that should return to 0x0000 (0x0001 after incrementing PC)
func TestSimpleProgram01(t* testing.T) {
	program := []byte {
		0x10, 0x00, 	 // stop
		0x3E, 0xF0, 	 // ld A, 0xF0
		0x06, 0x02, 	 // ld B, 0x02
		0x80, 			 // add A, B
		0x11, 0x89, 0xF8,// ld DE 0xF889
		0xD6, 0xF2,		 // sub A, 0xF2 (should give zero)
		0xCA, 0x00, 0x00,// jump to 0x0000 if zero flag
		0x10, 0x00,
	}

	LoadProgram(program)
	SetFlagZf(false)
	SetPC(0x0002) // at instruction ld A, 0xF0
	StartTest(100)

	if de := GetDE(); de != 0xF889 {
		t.Errorf("TestSimpleProgram01() failed: expected DE @ 0xF889, got 0x%04X", de)
	}

	if pc := GetPC(); pc != 0x0001 { // incremented after fetch, so not 0x0000
		t.Errorf("TestSimpleProgram01() failed: expected PC @ 0x0001, got 0x%04X", pc)
	}
}

// Test an overflowing increment
func TestSimpleProgram02(t* testing.T) {
	program := []byte {
		0x3C,		 	// inc A
		0x10, 0x00,		// stop
	}

	SetPC(0x0000)
	SetFlagH(false)
	SetA(0x0F)
	LoadProgram(program)
	StartTest(100)

	if a := GetA(); a!= 0x10 {
		t.Errorf("TestSimpleProgram02() failed: expected A @ 0x10, got 0x%02X", a)
	}
	if !GetFlagH() {
		t.Error("TestSimpleProgram02() failed: half carry should be set.")
	}
}

// measure the duration of a ld instruction
func BenchmarkLdRegistersInstructions(b *testing.B) {
	Set(0x0000, 0x78) // ld A, B

	for i := 0; i < b.N; i++ {
		SetPC(0x0000)
		ExecuteNext()
	}
}

func TestInterruptsEnables(t *testing.T) {
	var vblank 	bool
	var timer 	bool
	var lcdStat bool
	var joypad 	bool
	var serial	bool

	ResetMemory()

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

	ResetMemory()

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
