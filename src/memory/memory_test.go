package memory

import "testing"

func TestGet(t *testing.T) {
	RAM[0xFF21] = 0x28

	if b := Get(0xFF21); b != 0x28 {
		t.Errorf("TestGet() failed: expected 0x28, got 0x%04x", b)
	}
}

func TestSet(t *testing.T) {
	Write(0x5247, 0x99)

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

	SetRange(from, data)

	actual := GetRange(from, uint16(len(data)))

	if actual[0x00] != 0x66 {
		t.Errorf("TestSetRange() failed: at 0x00, expected 0x66, got 0x%04x", actual[0x50])
	}
}

func TestDMATransfer(t *testing.T) {
	for i := 0; i < 40; i++ {
		RAM[0x100+i] = 0x88
	}

	Write(0xFF46, 0x01)

	for i := 0; i < 40; i++ {
		if RAM[0xFE00+i] != 0x88 {
			t.Error("DMA transfer failed.")
		}
	}
}
