package lcd

import "testing"

func CompareTiles(t* testing.T, actual *[16]byte, expected *[16]byte) {
	for i := 0; i < 16; i++ {
		if actual[i] != expected[i] {
			t.Error("tiles don't match")
			t.Errorf("at %d: expected: 0x%02X, got: 0x%02X", i, expected[i], actual[i])
			return
		}
	}
}

func printTile(t* testing.T, tile [16]byte) {
	for i := 0; i < 16; i+=2 {
		t.Logf("%02X%02X", tile[i], tile[i+1])
	}
}

func TestMakeTile(t* testing.T) {

	tile := ""  +
	 "........" +
	 "........" +
	 "........" +
	 "........" +
	 "........" +
	 "........" +
	 "........" +
	 "........"

	result := MakeTile(tile)
	expected := new([16]byte)

	CompareTiles(t, result, expected)

	tile = ""  +
	 "...1...." +
	 "........" +
	 "........" +
	 "........" +
	 "........" +
	 "........" +
	 "........" +
	 "........"

	result = MakeTile(tile)
	expected = new([16]byte)
	expected[0] = 0x1

	printTile(t, *result)
	CompareTiles(t, result, expected)
}