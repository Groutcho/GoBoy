package cpu

func jumpToOperand16() int {
	addr := FetchOperand16()
	SetPC(addr)
	return 4
}

func jumpToOffset() int {
	offset := FetchOperand8s()
	SetPC(uint16(int(GetPC()) + offset))
	return 3
}

func xC3_jp() int {
	return jumpToOperand16()
}

func xE9_jp() int {
	SetPC(GetHL())
	return 1
}

func xDA_jp() int {
	if GetFlagCy() {
		return jumpToOperand16()
	} else {
		return 3
	}
}

func xD2_jp() int {
	if !GetFlagCy() {
		return jumpToOperand16()
	} else {
		return 3
	}
}

func xC2_jp() int {
	if !GetFlagZf() {
		return jumpToOperand16()
	} else {
		return 3
	}
}

func xCA_jp() int {
	if GetFlagZf() {
		return jumpToOperand16()
	} else {
		return 3
	}
}

func x18_jr() int {
	return jumpToOffset()
}

func x38_jr() int {
	if GetFlagCy() {
		return jumpToOffset()
	} else {
		return 2
	}
}

func x30_jr() int {
	if !GetFlagCy() {
		return jumpToOffset()
	} else {
		return 2
	}
}

func x20_jr() int {
	if !GetFlagZf() {
		return jumpToOffset()
	} else {
		return 2
	}
}

func x28_jr() int {
	if GetFlagZf() {
		return jumpToOffset()
	} else {
		return 2
	}
}