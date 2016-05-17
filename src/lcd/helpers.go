package lcd

import . "common"

// Return a 16-byte tile from the given string.
//
// example:
//
//"33333333" <- line 0
//"31111113"
//"31111113"
//"31111113"
//"31111113"
//"31111113"
//"31111113"
//"33333333"
func MakeTile(input string) []byte {
	result := make([]byte, 16, 16)

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			h, l := helperGetPixel(input[y*8+x])
			i := y * 2
			b := uint8(7 - x)
			result[i] = SetBit(result[i], b, h)
			result[i+1] = SetBit(result[i+1], b, l)
		}
	}

	return result
}

func helperGetPixel(c byte) (uint8, uint8) {
	var h uint8
	var l uint8

	switch c {
	case '.':
		l = 0
		h = 0
	case '1':
		l = 1
		h = 0
	case '2':
		l = 0
		h = 1
	case '3':
		l = 1
		h = 1
	}

	return h, l
}
