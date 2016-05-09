letters = "BCDEHL"
seconds = "BCDEHL*"
start = 0x40

for l in letters:
	for s in seconds:
		if s == "*":
			s = "(HL)"
		txt = "{l}, {s}".format(l=l, s=s)
		print("ld   {txt:7}         {x:02X}              {cycles}    ----     // {l}={s}".format(txt=txt, l=l, s=s, x=start, cycles=8 if s == "(HL)" else 4))

		if s == "(HL)":
			start += 2
		else:
			start += 1