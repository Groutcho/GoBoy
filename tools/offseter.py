pairs = [\
	('B', 0),
	('C', 1),
	('D', 2),
	('E', 3),
	('H', 4),
	('L', 5),
	('A', 7)\
	]

template = "{:4} {}         {:2x}              {:2d}    {}     // {}"

# LD (HL),r      1  70+rb            7     ------  [HL]=r
def offset(s):
	for letter, offs in pairs:
		for bit in range(0, 8):
			x = 0x80 + (8 * bit) + offs
			print(s.replace('<b>', str(bit)).replace('$$', hex(x)[-2:].upper()).replace('<r>', letter))

s = "res  <b>,<r>             CB $$           8    ----     // reset bit <b> of register <r>"
offset(s)

# addr = ('CY ', '!CY', '!Z ', 'Z  ')
# ops = ('DC', 'D4', 'C4', 'CC')

# for i in range(len(addr)):
# 	print(s.replace('<f>', addr[i]).replace('xx', ops[i]))