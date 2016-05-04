package cpu

import "testing"

func testFlags(t* testing.T, Z bool, N bool, H bool, C bool) {
	if GetFlagZf() != Z {
		if Z {
			t.Error("TestFlags() failed: expected Z set")
		} else {
			t.Error("TestFlags() failed: expected Z unset")			
		}
	}

	if GetFlagN() != N {
		if N {
			t.Error("TestFlags() failed: expected N set")
		} else {
			t.Error("TestFlags() failed: expected N unset")			
		}
	}

	if GetFlagH() != H {
		if H {
			t.Error("TestFlags() failed: expected H set")
		} else {
			t.Error("TestFlags() failed: expected H unset")			
		}
	}

	if GetFlagCy() != C {
		if C {
			t.Error("TestFlags() failed: expected C set")
		} else {
			t.Error("TestFlags() failed: expected C unset")			
		}
	}
}