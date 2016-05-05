package cpu

// call to %2, SP=SP-2, (SP)=PC, PC=%2
func xCD_call() int {
	addr := FetchOperand16()
	Call(addr)

	return 6
}

// conditional call if CY
func xDC_call() int {
	addr := FetchOperand16()
	if GetFlagCy() {
		Call(addr)
		return 6
	} else {
		return 3
	}
}

// conditional call if !CY
func xD4_call() int {
	addr := FetchOperand16()
	if !GetFlagCy() {
		Call(addr)
		return 6
	} else {
		return 3
	}
}

// conditional call if Z
func xC4_call() int {
	addr := FetchOperand16()
	if GetFlagZf() {
		Call(addr)
		return 6
	} else {
		return 3
	}
}

// conditional call if !Z
func xCC_call() int {
	addr := FetchOperand16()
	if !GetFlagZf() {
		Call(addr)
		return 6
	} else {
		return 3
	}
}