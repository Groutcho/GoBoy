package cpu

func jumpToOperand16(addr uint16) int {
	PC = addr
	return 4
}

func jumpToOffset(offset int) int {
	SetPC(uint16(int(GetPC()) + offset))
	return 3
}

func xC3_jp() int {
	addr := FetchOperand16()
	return jumpToOperand16(addr)
}

func xE9_jp() int {
	SetPC(GetHL())
	return 1
}

func xDA_jp() int {
	addr := FetchOperand16()
	if GetFlagCy() {
		return jumpToOperand16(addr)
	} else {
		return 3
	}
}

func xD2_jp() int {
	addr := FetchOperand16()
	if !GetFlagCy() {
		return jumpToOperand16(addr)
	} else {
		return 3
	}
}

func xC2_jp() int {
	addr := FetchOperand16()
	if !GetFlagZf() {
		return jumpToOperand16(addr)
	} else {
		return 3
	}
}

func xCA_jp() int {
	addr := FetchOperand16()
	if GetFlagZf() {
		return jumpToOperand16(addr)
	} else {
		return 3
	}
}

func x18_jr() int {
	offset := FetchOperand8s()
	return jumpToOffset(offset)
}

func x38_jr() int {
	offset := FetchOperand8s()
	if GetFlagCy() {
		return jumpToOffset(offset)
	} else {
		return 2
	}
}

func x30_jr() int {
	offset := FetchOperand8s()
	if !GetFlagCy() {
		return jumpToOffset(offset)
	} else {
		return 2
	}
}

func x20_jr() int {
	offset := FetchOperand8s()
	if !GetFlagZf() {
		return jumpToOffset(offset)
	} else {
		return 2
	}
}

func x28_jr() int {
	offset := FetchOperand8s()
	if GetFlagZf() {
		return jumpToOffset(offset)
	} else {
		return 2
	}
}
