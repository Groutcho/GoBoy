package common

// set the 8 high bits of <target> with <value>.
func SetHighBits(value uint8, target uint16) uint16 {
	return (target & 0x00FF) | (uint16(value) << 8)
}

// set the 8 low bits of <target> with <value>.
func SetLowBits(value uint8, target uint16) uint16 {
	return (target & 0xFF00) | uint16(value)
}

// return the 8 high bits of <value>.
func GetHighBits(value uint16) uint8 {
	return uint8((value & 0xFF00) >> 8)
}

// return the 8 low bits of <value>.
func GetLowBits(value uint16) uint8 {
	return uint8(value & 0x00FF)
}

// return the 4 low bits of <value>.
func GetLowNibble(value uint8) uint8 {
	return uint8(value & 0x0F)
}

// return the 4 high bits of <value>.
func GetHighNibble(value uint8) uint8 {
	return uint8((value & 0xF0) >> 4)
}

func SetBit(value uint8, bit uint8, set uint8) uint8 {
	if set != 0 {
		return value | (1 << bit)
	} else {
		return value & ^(1 << bit)
	}
}

func GetBit(value uint8, bit uint8) uint8 {
	if value & (1 << bit) == 0 {
		return 0
	} else {
		return 1
	}
}