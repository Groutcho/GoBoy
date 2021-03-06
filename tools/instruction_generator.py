#!/usr/bin/python3
#encoding=utf-8

import re
import sys
from textwrap import dedent, indent
from binascii import unhexlify

skip_re = re.compile(r'[ ,]*')

short_reg_re = re.compile(r'\b(A|B|C|D|E|H|L)\b')
long_ref_reg_re = re.compile(r'\((AF|BC|DE|HL|SP|PC)\)')
long_reg_re = re.compile(r'\b(AF|BC|DE|HL|SP|PC)\b')
operand_16bit_re = re.compile(r'%2')
operand_ref_16bit_re = re.compile(r'\(%2\)')
operand_8bit_re = re.compile(r'%1')
operand_uns8bit_re = re.compile(r'%s')
constant_re = re.compile(r'\d+')

REF_REG = 'register pointer'
REG16 = '16bit register'
OPERAND16 = '16bit operand'
OPERAND16REF = '16bit reference operand'
OPERANDU8 = '8bit unsigned operand'
OPERAND8 = '8bit operand'
REG8 = '8bit register'
CONSTANT = 'constant'
ERROR = 'error'

write_packages = (\
	'adc',
	'add',
	'and',
	'bit',
	'cp',
	'dec',
	'inc',
	'ld',
	'or',
	'res',
	'rst',
	'sbc',
	'set',
	'sla',
	'sra',
	'srl',
	'sub',
	'swap',
	'xor',
	'rlca',
	'rla',
	'rrca',
	'rra',
	'rlc',
	'rl',
	'rrc',
	'rr'
)
write_table = ( 'ld',
	'adc',
	'add',
	'and',
	'bit',
	'call',
	'cp',
	'cpl',
	'dec',
	'inc',
	'jp',
	'jr',
	'ldd',
	'ldi',
	'or',
	'pop',
	'push',
	'res',
	'ret',
	'reti',
	'rst',
	'sbc',
	'set',
	'sla',
	'sra',
	'srl',
	'sub',
	'swap',
	'xor',
	'rlca',
	'rla',
	'rrca',
	'rra',
	'rlc',
	'rl',
	'rrc',
	'rr'
)

dont_require_memory = ('rst', 'rlca', 'rrca', 'rla', 'rra')

class Instruction(object):
	def __init__(self, mnemonic, pseudocode, opcode, cycles, comment, line):
		self.mnemonic = mnemonic
		self.opcode = opcode
		self.pseudocode = pseudocode
		self.cycles = cycles
		self.comment = comment
		self.line = line

	@staticmethod
	def parse(line):
		mnemonic = line[:4].strip()
		pseudocode = tokenize(line[4:22])
		opcode = parse_opcode(line[21:26].strip())
		cycles = parse_cycles(line[33:38])
		comment = line[51:]

		return Instruction(mnemonic, pseudocode, opcode, cycles, comment, line)

	def full_opcode(self):
		oc = unhexlify(self.opcode[0])[0]
		if self.opcode[1]:
			return 0x100 + oc
		else:
			return oc

	def func_name(self):
		if self.opcode[1]:
			return 'xCB_{}_{}'.format(self.opcode[0], self.mnemonic)
		else:
			return 'x{}_{}'.format(self.opcode[0], self.mnemonic)


def parse_opcode(s):
	if len(s) == 2:
		return (s, False)

	if len(s) == 5:
		print(s[-2:])
		return (s[-2:], True)


class Token(object):
	def __init__(self, code, pos, value=None):
		self.code = code
		self.pos = pos
		self.value = value

	def __repr__(self):
		if self.value:
			return "{} ({})".format(self.code, self.value, self.pos)
		else:
			return "{}".format(self.code, self.pos)


def apply(s, regexp, i):
	m = skip_re.match(s[i:])
	if m:
		i += m.end(0)

	m = regexp.match(s[i:])
	if m:
		return m.group(0), i + m.end(0), i + m.start(0)
	else:
		return None, i, i


def tokenize(text):
	tokens = []
	i = 0
	t = 0

	limit = 40

	s = text[:22]

	while limit and i < len(s) -1:
		error = True
		t, i, p = apply(s, long_ref_reg_re, i)
		if t:
			error = False
			tokens.append(Token(REF_REG, p, t.strip('()')))
			continue
		t, i, p = apply(s, long_reg_re, i)
		if t:
			error = False
			tokens.append(Token(REG16, p, t))
			continue
		t, i, p = apply(s, short_reg_re, i)
		if t:
			error = False
			tokens.append(Token(REG8, p, t))
			continue
		t, i, p = apply(s, operand_16bit_re, i)
		if t:
			error = False
			tokens.append(Token(OPERAND16, p, t))
			continue
		t, i, p = apply(s, operand_uns8bit_re, i)
		if t:
			error = False
			tokens.append(Token(OPERANDU8, p, t))
			continue
		t, i, p = apply(s, operand_8bit_re, i)
		if t:
			error = False
			tokens.append(Token(OPERAND8, p, t))
			continue
		t, i, p = apply(s, operand_ref_16bit_re, i)
		if t:
			error = False
			tokens.append(Token(OPERAND16REF, p, t.strip('()')))
			continue
		t, i, p = apply(s, constant_re, i)
		if t:
			tokens.append(Token(CONSTANT, p, t))
			error = False

		limit -= 1
		if error:
			tokens.append(Token(ERROR, p, None))
			i+=1

	return tokens


def parse_cycles(s):
	cycles = s.strip()
	try:
		if ';' in cycles:
			a, b = cycles.split(';')
			return (int(int(a)/4), int(int(b)/4))
		return int(int(cycles)/4), 0
	except ValueError as err:
		print('{}: {}'.format(s, err))
		sys.exit(1)
		return 0, 0


def parse_ld(instruction):
	opA = instruction.pseudocode[0]
	opB = instruction.pseudocode[1]

	src = None
	if opB.code == OPERAND8:
		src = 'FetchOperand8()'
	elif opB.code == OPERAND16:
		src = 'FetchOperand16()'
	elif opB.code == OPERAND16REF:
		src = 'Get(FetchOperand16())'
	elif opB.code == OPERANDU8:
		src = 'int8(FetchOperand8())'
	elif opB.code in (REG8, REG16):
		src = 'Get%s()' % opB.value
	elif opB.code == REF_REG:
		src = 'Get(Get%s())' % opB.value

	dest = None
	if opA.code == OPERAND16REF:
		dest = 'Write(FetchOperand16(), '
	elif opA.code == OPERANDU8:
		dest = 'int8(FetchOperand8())'
	elif opA.code in (REG8, REG16):
		dest = 'Set%s(' % opA.value
	elif opA.code == REF_REG:
		dest = 'Write(Get%s(), ' % opA.value

	if dest is None or src is None:
		return '// parsing error'

	return '{}{})'.format(dest, src)


def parse_rst(instruction):
	addr = unhexlify(instruction.pseudocode[0].value)[0]

	return  "Call(0x{:04x})".format(addr)


def parse_set(instr):
	assert instr.pseudocode[0].code == CONSTANT
	assert instr.pseudocode[1].code in (REG8, REF_REG)
	n_bit = int(instr.pseudocode[0].value)

	if instr.mnemonic == 'res':
		value = '0'
	else:
		value = '1'

	if instr.pseudocode[1].code == REG8:
		code = indent(dedent("""
		value := SetBit(Get{0}(), {1}, {2})
		Set{0}(value)""".format(instr.pseudocode[1].value, n_bit, value)), '\t')
	else:
		code = indent(dedent("""
		addr := Get{0}()
		value := SetBit(Get(addr), {1}, {2})
		Write(addr, value)""".format(instr.pseudocode[1].value, n_bit, value)), '\t')

	return code.strip()


def parse_swap(instruction):
	target = instruction.pseudocode[0]

	if target.code == REG8:
		operation = '{0} = Swap(Get{0}())'.format(target.value)
	elif target.code == REF_REG:
		operation = 'addr := Get%s()\n' % target.value
		operation += '\tWrite(addr, Swap(Get(addr)))'.format(target.value)

	return operation


def custom_impl_F8():
	"""
	ld   HL, SP+%s
	"""

	code = \
	"""
	operand := int(FetchOperand8s())
	result := int(GetSP()) + operand
	SetHL(uint16(result))

	hc := 0
	if operand > 0 {
		if GetLowNibble(uint8(operand)) + GetLowNibble(GetLowBits(GetSP())) > 0xF {
			hc = 1
		} else {
			hc = 0
		}
	}

	SetFlags(result, F_SET_0, F_SET_0, hc, F_SET_IF, F_16bit)
	"""

	return code.strip()


def custom_impl_08():
		"""
		ld   %2, SP
		"""

		code = \
		"""
		Write16(FetchOperand16(), GetSP())
		"""

		return code.strip()


def custom_impl_E0():
	"""
	read from io-port %1 (memory FF00+%1)
	"""

	code = \
	"""
	offset := int(FetchOperand8())
	addr := 0xFF00 + uint16(offset)
	value := GetA()
	Write(addr, value)
	"""

	return code.strip()


def custom_impl_F0():
	"""
	read from io-port %1 (memory FF00+%1)
	"""

	code = \
	"""
	offset := int(FetchOperand8())
	addr := 0xFF00 + uint16(offset)
	value := Get(addr)
	SetA(value)
	"""

	return code.strip()


def custom_impl_E2():
	"""
	write to io-port C (memory FF00+C
	"""

	code = \
	"""
	offset := GetC()
	addr := 0xFF00 + uint16(offset)
	value := GetA()
	Write(addr, value)
	"""

	return code.strip()


def custom_impl_F2():
	"""
	read from io-port C (memory FF00+C)
	"""

	code = \
	"""
	offset := GetC()
	addr := 0xFF00 + uint16(offset)
	value := Get(addr)
	SetA(value)
	"""

	return code.strip()


def parse_inc(instr):
	target = instr.pseudocode[0]

	pattern = \
	"""
	original := {dest}
	value := original {sign} 1
	{assign}

	hc := Is{hctype}HalfCarry({nibble}, uint8(1))

	SetFlags(int(value), F_SET_IF, F_SET_{n}, hc, F_IGNORE, F_{width}bit)
	"""

	pattern16 = \
	"""
	Set{reg}(Get{reg}(){sign}1)
	"""

	if instr.mnemonic == 'inc':
		sign = '+'
		n = 0
		hctype = 'Add'
	else:
		sign = '-'
		n = 1
		hctype = 'Sub'

	width = 8
	nibble = "original"
	if target.code in (REG8, REG16):
		dest = "Get{}()".format(target.value)
		assign = "Set{}(value)".format(target.value)
		if target.code == REG16:
			width = 16
			nibble = "GetHighBits(original)"
	elif target.code == REF_REG:
		dest = "Get(Get{}())".format(target.value)
		assign = "Write(Get{}(), value)".format(target.value)

	if target.code == REG16:
		code = pattern16.format(
			reg=target.value, sign=sign)
	else:
		code = pattern.format(
			dest=dest, sign=sign, 
			assign=assign, width=width,
			nibble=nibble, n=n, hctype=hctype)

	return code.strip()


def parse_rotate(instr):
	template_rlc = \
	"""
	value := {value}
	msb := GetBit(value, 7)

	value = value << 1
	value = SetBit(value, 0, {edge})

	SetFlagCy(msb == 1)
	SetFlagZf({flagz})
	SetFlagN(false)
	SetFlagH(false)

	{target}
	"""

	template_rrc = \
	"""
	value := {value}
	lsb := GetBit(value, 0)

	value = value >> 1
	value = SetBit(value, 7, {edge})

	SetFlagCy(lsb == 1)
	SetFlagZf({flagz})
	SetFlagN(false)
	SetFlagH(false)

	{target}
	"""

	m = instr.mnemonic
	dest = instr.pseudocode[0]
	flagz = "value == 0"

	if m in ('rlca', 'rla', 'rrca', 'rra'):
		value = "GetA()"
		target = "SetA(value)"
		flagz = "false"
	elif dest.code == REG8:
		value = "Get%s()" % dest.value
		target = "Set%s(value)" % dest.value
	else:
		value = "Get(GetHL())"
		target = "Write(GetHL(), value)"

	if m in ('rlc', 'rl', 'rlca', 'rla'):
		if m == 'rlc':
			edge = 'msb'
		else:
			edge = 'uint8(GetFlagCyInt())'
		code = template_rlc.format(value=value, target=target, edge=edge, flagz=flagz)
	elif m in ('rrc', 'rr', 'rra', 'rrca'):
		if m == 'rrc':
			edge = 'lsb'
		else:
			edge = 'uint8(GetFlagCyInt())'
		code = template_rrc.format(value=value, target=target, edge=edge, flagz=flagz)

	return code.strip()


def parse_bitwise(instr):
	template_sla = \
	"""
	value := {value}

	SetFlagCy(value & 0x80 == 0x80)

	value = value << 1

	SetFlagZf(value == 0)
	SetFlagN(false)
	SetFlagH(false)

	{target}
	"""

	template_sra = \
	"""
	value := {value}
	msb := GetBit(value, 7)

	SetFlagCy(GetBit(value, 0) == 1)

	value = value >> 1
	value = SetBit(value, 7, msb)

	SetFlagZf(value == 0)
	SetFlagN(false)
	SetFlagH(false)

	{target}
	"""

	template_srl = \
	"""
	value := {value}

	SetFlagCy(value & 0x01 == 1)

	value = value >> 1

	SetFlagZf(value == 0)
	SetFlagN(false)
	SetFlagH(false)

	{target}
	"""

	m = instr.mnemonic
	dest = instr.pseudocode[0]
	if m == 'sra':
		bit1 = 'value |= 0x10'
		msb = 'msb := value & 0x80'

	if dest.code == REG8:
		value = "Get%s()" % dest.value
		target = "Set%s(value)" % dest.value
	else:
		value = "Get(GetHL())"
		target = "Write(GetHL(), value)"

	if m == 'sra':
		code = template_sra.format(value=value, target=target)
	elif m == 'sla':
		code = template_sla.format(value=value, target=target)
	elif m == 'srl':
		code = template_srl.format(value=value, target=target)

	return code.strip()


def parse_cp(instr):
	template = \
	"""
	left := int(GetA())
	right := int({right})
	result := left - right

	hc := IsSubHalfCarry(uint8(left), uint8(right))

	SetFlags(result, F_SET_IF, F_SET_1, hc, F_SET_IF, F_8bit)
	"""

	op = instr.pseudocode[0]

	if op.code == REG8:
		right = 'Get%s()' % op.value
	elif op.code == OPERAND8:
		right = 'FetchOperand8()'
	elif op.code == REF_REG:
		right = 'Get(GetHL())'

	code = template.format(right=right)
	return code.strip()


def parse_logic(instr):
	template = \
	"""
	result := GetA() {op} {right}
	SetA(result)

	SetFlags(int(result), F_SET_IF, F_SET_0, {hcflag}, F_SET_0, F_8bit)
	"""

	right = instr.pseudocode[0]

	if right.code == REG8:
		right = 'Get%s()' % right.value
	elif right.code == OPERAND8:
		right = 'FetchOperand8()'
	elif right.code == REF_REG:
		right = 'Get(GetHL())'

	m = instr.mnemonic
	if m == 'and':
		op = '&'
	elif m == 'xor':
		op = '^'
	elif m == 'or':
		op = '|'

	hcflag = 'F_SET_0'
	if m == 'and':
		hcflag = 'F_SET_1'

	code = template.format(op=op, right=right, hcflag=hcflag)
	return code.strip()


def parse_add(instr):
	opA = instr.pseudocode[0]
	opB = instr.pseudocode[1]

	assert opA.code in (REG8, REG16), "first operand of add should be 8bit or 16bit register"
	assert opB.code in (REG8, REG16, OPERANDU8, OPERAND8, REF_REG), "second operand of add invalid %s" % opB.code

	opAstr = 'Get{}()'.format(opA.value)

	pattern = \
	"""
	left := {left}
	right := {right}
	result := int(left) {sign} int(right){opt}

	Set{dest}(uint{width}(result))
	hcarry := {hc}
	SetFlags(result, {zFlag}, F_SET_{optype}, hcarry, F_SET_IF, F_{width}bit)
	"""

	m = instr.mnemonic
	left = ""
	right = ""
	sign = "+" if m in ('add', 'adc') else "-"
	optype = "0" if m in ('add', 'adc') else "1"

	opt = ""
	if m == 'adc':
		opt = " + GetFlagCyInt()"
	elif m == 'sbc':
		opt = " - GetFlagCyInt()"

	dest = ""
	hc = "Is{}HalfCarry(left, right)"
	width = 8
	zFlag = "F_SET_IF"

	left = 'Get{}()'.format(opA.value)
	dest = opA.value
	if opA.code == REG16:
		width = 16
		zFlag = "F_IGNORE"
		if opB.code == OPERANDU8:
			hc = "Is{}HalfCarry(GetHighBits(left), right)"
		else:
			hc = "Is{}HalfCarry(GetHighBits(left), GetHighBits(right))"

	if m in ('adc', 'add'):
		hc = hc.format('Add')
	else:
		hc = hc.format('Sub')

	b = opB.code
	if b in (REG8, REG16):
		right = 'Get{}()'.format(opB.value)
	elif b == OPERAND8:
		right = 'FetchOperand8()'
	elif b == OPERANDU8:
		right = 'FetchOperand8()'
	elif b == OPERAND16REF:
		right = 'Get(FetchOperand16())'
	elif b == REF_REG:
		right = 'Get(Get{}())'.format(opB.value)

	code = pattern.format(left=left, right=right, opt=opt, hc=hc,
						  dest=dest, width=width, zFlag=zFlag, sign=sign, optype=optype)

	return code.strip()


def parse_bit(instruction):
	operand = instruction.pseudocode[1]
	n_bit = instruction.pseudocode[0].value
	if operand.code == REG8:
		target = 'Get{}()'.format(operand.value)
	else:
		target = 'Get(Get{}())'.format(operand.value)

	code = 'set := GetBit({}, {});\n'.format(target, n_bit)
	code += '\tSetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)'

	return code


dispatch = {\
	'ld': parse_ld,
	'rst': parse_rst,
	'swap': parse_swap,
	'inc': parse_inc,
	'dec': parse_inc,
	'res': parse_set,
	'set': parse_set,
	'add': parse_add,
	'adc': parse_add,
	'sub': parse_add,
	'sbc': parse_add,
	'bit': parse_bit,

	'and': parse_logic,
	'or': parse_logic,
	'xor': parse_logic,

	'cp': parse_cp,

	'sra': parse_bitwise,
	'sla': parse_bitwise,
	'srl': parse_bitwise,

	'rlca': parse_rotate,
	'rla': parse_rotate,
	'rra': parse_rotate,
	'rrca': parse_rotate,

	'rlc': parse_rotate,
	'rl': parse_rotate,
	'rrc': parse_rotate,
	'rr': parse_rotate,
}

custom_impl_table = {\
	0xF0: custom_impl_F0,
	0xF8: custom_impl_F8,
	0xE0: custom_impl_E0,
	0x08: custom_impl_08,
	0xE2: custom_impl_E2,
	0xF2: custom_impl_F2,
	0xF0: custom_impl_F0,
	}


unsupp = set()

def make_func(instr, file):
	if instr.opcode[1]:
		oc = unhexlify(instr.opcode[0])[0] + 256
	else:
		oc = unhexlify(instr.opcode[0])[0]

	if oc in custom_impl_table:
		result = custom_impl_table[oc]()
	elif instr.mnemonic in dispatch:
		result = dispatch[instr.mnemonic](instr)
	else:
		unsupp.add('unsupported mnemonic: %s' % instr.mnemonic)
		return

	comment = '// {} - {}'.format(instr.line[:21].strip(), instr.comment.strip('/ \n'))

	print(comment, file=file)
	print('func ' + instr.func_name() + '() int {', file=file)



	ret = (oc, instr.func_name())

	if result:
		print('\t%s\n' % result, file=file)

	if instr.cycles[1] == 0:
		print('\treturn %d' % instr.cycles[0], file=file)
	else:
		print(indent(dedent("""\
		if jmp {{
		\treturn {}
		}} else {{
		\treturn {}
		}}""".format(instr.cycles[0], instr.cycles[1])), '\t'), file=file)

	print('}\n', file=file)
	return ret


def make_disasm_table(instructions:dict):
	with open('../src/cpu/disasm_table.go', 'wt') as f:
		print("// this file is automatically generated by instruction_generator.py", file=f)
		print('package cpu\n', file=f)

		print('func init() {', file=f)

		for instr in sorted(instructions, key=lambda x: x.full_opcode()):
			print("{:02X}".format(instr.full_opcode()))
			line = instr.line[:21]
			arg_type = "NO_ARG"
			if "%1" in line:
				arg_type = "U8"
			elif "%s" in line:
				arg_type = "S8"
			elif "%2" in line:
				arg_type = "U16"

			line = line.replace("%1", "0x%02X").replace("%2", "0x%04X").replace("%s", "%d")

			print('\tdisasm_table[0x{code:03X}].text = "{line}"'.format(
				code=instr.full_opcode(),
				line=line.strip()),
				file=f)
			print('\tdisasm_table[0x{code:03X}].argument = {arg}'.format(
				code=instr.full_opcode(),
				arg=arg_type),
				file=f)

		print('}', file=f)


def make_dispatch_table(instructions:dict):
	with open('../src/cpu/table.go', 'wt') as f:
		print("// this file is automatically generated by instruction_generator.py", file=f)
		print('package cpu\n', file=f)

		print('func init() {', file=f)

		for instr in sorted(instructions, key=lambda x: x.full_opcode()):
			if instr.mnemonic in write_table:
				print("{:02X}".format(instr.full_opcode()))
				print('\tdispatch_table[0x{:03X}] = {}'.format(instr.full_opcode(), instr.func_name()), file=f)

		print('}', file=f)

header = \
"""// this file is automatically generated by instruction_generator.py
package cpu
{imports}
"""

def requirements(mnemonic):
	imports = []

	if mnemonic in ('add', 'bit', 'rla', 'ld', 'res',
		'rl', 'rlc', 'rlca', 'rr', 'rrc', 'rra', 'rrca', 'set', 'sra'):
		imports.append('import . "common"')
	if mnemonic in ('add', 'bit', 'adc', 'and', 'dec', 'inc', 'ld', 'swap', 'sub', 'xor',
		'cp', 'or', 'res', 'rlc', 'rl', 'rr', 'rrc', 'sbc', 'set', 'sla', 'sra', 'srl'):
		imports.append('import . "memory"')

	if imports:
		return '\n' + '\n'.join(imports)
	else:
		return ''

def parse_file(filename):
	result = dict()
	lst = []
	with open(filename) as cpu:
		for line in cpu.readlines():
			if line.strip() and line[0] != '#':
				instruction = Instruction.parse(line)
				if instruction.mnemonic not in result:
					result[instruction.mnemonic] = []
				result[instruction.mnemonic].append(instruction)
				lst.append(instruction)
	return result, lst


def main():
	instr_dict, instr_list = parse_file('instructions.CPU')

	for mnemonic in instr_dict:
		if mnemonic in write_packages:
			with open('../src/cpu/{}.go'.format(mnemonic), 'wt') as f:
				print(header.format(imports=requirements(mnemonic)), file=f)

				for instruction in instr_dict[mnemonic]:
					n = make_func(instruction, f)

	make_dispatch_table(instr_list)
	make_disasm_table(instr_list)
	for un in unsupp:
		print(un)

if __name__ == "__main__":
	main()