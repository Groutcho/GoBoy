package cpu

import "testing"
import . "memory"

func TestCD_call(t* testing.T) {

	program := []byte{
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0xCD, 0x15, 0xFF, // call 0xFF15
		0x00, // garbage <- PC before call
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0xFF, // garbage
	}

	SetPC(0x000A)
	SetSP(0x0014)
	LoadProgram(program)
	ExecuteNext()

	CheckRegister(t, REG_PC, 0xFF15)
	CheckRegister(t, REG_SP, 0x0012)

	t.Logf("SP: %04X", GetSP())
	t.Logf("addr: %04X", Get16(GetSP()))
	return_address := Get16(GetSP())
	t.Log(DumpRange(0x0000, 0x0014))
	if return_address != 0x000D {
		t.Errorf("TestxCD_call() failed: return address invalid. Expected 0x000D, got 0x%04x", return_address)
	}
}

func TestDC_call(t* testing.T) {

	program := []byte{
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0xDC, 0x15, 0xFF, // call 0xFF15
		0x00, // garbage <- PC before call
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0xFF, // garbage
	}

	LoadProgram(program)

	SetPC(0x000A)
	SetSP(0x0014)
	SetFlagCy(true)
	ExecuteNext()

	CheckRegister(t, REG_PC, 0xFF15)
	CheckRegister(t, REG_SP, 0x0012)

	t.Logf("SP: %04X", GetSP())
	t.Logf("addr: %04X", Get16(GetSP()))
	return_address := Get16(GetSP())
	t.Log(DumpRange(0x0000, 0x0014))
	if return_address != 0x000D {
		t.Errorf("TestxCD_call() failed: return address invalid. Expected 0x000D, got 0x%04x", return_address)
	}

	SetPC(0x000A)
	SetSP(0x0014)
	SetFlagCy(false)
	ExecuteNext()

	CheckRegister(t, REG_PC, 0x000D)
	CheckRegister(t, REG_SP, 0x0014)
}

func TestD4_call(t* testing.T) {

	program := []byte{
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0xD4, 0x15, 0xFF, // call 0xFF15
		0x00, // garbage <- PC before call
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0xFF, // garbage
	}

	LoadProgram(program)

	SetPC(0x000A)
	SetSP(0x0014)
	SetFlagCy(false)
	ExecuteNext()

	CheckRegister(t, REG_PC, 0xFF15)
	CheckRegister(t, REG_SP, 0x0012)

	t.Logf("SP: %04X", GetSP())
	t.Logf("addr: %04X", Get16(GetSP()))
	return_address := Get16(GetSP())
	t.Log(DumpRange(0x0000, 0x0014))
	if return_address != 0x000D {
		t.Errorf("TestxCD_call() failed: return address invalid. Expected 0x000D, got 0x%04x", return_address)
	}

	SetPC(0x000A)
	SetSP(0x0014)
	SetFlagCy(true)
	ExecuteNext()

	CheckRegister(t, REG_PC, 0x000D)
	CheckRegister(t, REG_SP, 0x0014)
}

func TestC4_call(t* testing.T) {

	program := []byte{
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0xC4, 0x15, 0xFF, // call 0xFF15
		0x00, // garbage <- PC before call
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0xFF, // garbage
	}

	LoadProgram(program)

	SetPC(0x000A)
	SetSP(0x0014)
	SetFlagZf(true)
	ExecuteNext()

	CheckRegister(t, REG_PC, 0xFF15)
	CheckRegister(t, REG_SP, 0x0012)

	t.Logf("SP: %04X", GetSP())
	t.Logf("addr: %04X", Get16(GetSP()))
	return_address := Get16(GetSP())
	t.Log(DumpRange(0x0000, 0x0014))
	if return_address != 0x000D {
		t.Errorf("TestxCD_call() failed: return address invalid. Expected 0x000D, got 0x%04x", return_address)
	}

	SetPC(0x000A)
	SetSP(0x0014)
	SetFlagZf(false)
	ExecuteNext()

	CheckRegister(t, REG_PC, 0x000D)
	CheckRegister(t, REG_SP, 0x0014)
}

func TestC4CC_call(t* testing.T) {

	program := []byte{
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0xCC, 0x15, 0xFF, // call 0xFF15
		0x00, // garbage <- PC before call
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0x00, // garbage
		0xFF, // garbage
	}

	LoadProgram(program)

	SetPC(0x000A)
	SetSP(0x0014)
	SetFlagZf(false)
	ExecuteNext()

	CheckRegister(t, REG_PC, 0xFF15)
	CheckRegister(t, REG_SP, 0x0012)

	t.Logf("SP: %04X", GetSP())
	t.Logf("addr: %04X", Get16(GetSP()))
	return_address := Get16(GetSP())
	t.Log(DumpRange(0x0000, 0x0014))
	if return_address != 0x000D {
		t.Errorf("TestxCD_call() failed: return address invalid. Expected 0x000D, got 0x%04x", return_address)
	}

	SetPC(0x000A)
	SetSP(0x0014)
	SetFlagZf(true)
	ExecuteNext()

	CheckRegister(t, REG_PC, 0x000D)
	CheckRegister(t, REG_SP, 0x0014)
}
