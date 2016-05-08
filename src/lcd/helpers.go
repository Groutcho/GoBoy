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

	for i := 0; i < 64; i++ {
		B := int(i/4)
		b := uint8(7 - ((i % 4) * 2))
		h, l := getPixel(input[i])
		result[B] = SetBit(result[B], b, h)
		result[B] = SetBit(result[B], b - 1, l)
	}

	return result
}

func getPixel(c byte) (uint8, uint8) {
	var h uint8
	var l uint8

	switch (c) {
		case '.': l = 0; h = 0
		case '1': l = 1; h = 0
		case '2': l = 0; h = 1
		case '3': l = 1; h = 1
	}

	return h, l
}