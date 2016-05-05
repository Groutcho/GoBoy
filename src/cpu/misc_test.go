package cpu

import "testing"

func Test37_scf(t* testing.T) {
	SetFlagCy(false)
	x37_scf()

	if !GetFlagCy() {
		t.Error("Test37_scf() failed: expected CY set.")
	}
}