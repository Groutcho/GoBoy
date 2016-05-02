#!/usr/bin/python3
#encoding=utf-8

import re
import sys
from textwrap import dedent, indent
from binascii import unhexlify

skip_re = re.compile(r'[ ,]*')

mnemonic_re = re.compile(r'[a-z]{2,4}')
short_reg_re = re.compile(r'\b(A|B|C|D|E|H|L)\b')
long_ref_reg_re = re.compile(r'\((AF|BC|DE|HL|SP|PC)\)')
long_reg_re = re.compile(r'\b(AF|BC|DE|HL|SP|PC)\b')
flag_re = re.compile(r'\b(CY|Z)\b')
opcode_re = re.compile(r'[0-9A-F]{2}')
ex_opcode_re = re.compile(r'CB [0-9A-F]{2}')
operand_16bit_re = re.compile(r'%2')
operand_ref_16bit_re = re.compile(r'\(%2\)')
operand_8bit_re = re.compile(r'%1')
operand_uns8bit_re = re.compile(r'%s')
not_re = re.compile(r'!')
comment_re = re.compile(r'//.*')
constant_re = re.compile(r'\d+')
cond_cycle_re = re.compile(r';')
equal_re = re.compile(r'=')
flags_re = re.compile(r'[z01-][01-][h01-][c01-]')

MNEMONIC = 'mnemonic'
REF_REG = 'register pointer'
REG16 = '16bit register'
OPERAND16 = '16bit operand'
OPERAND16REF = '16bit reference operand'
OPERAND8 = '8bit operand'
EX_OPCODE = 'extended opcode'
OPCODE = 'opcode'
REG8 = '8bit register'
CONSTANT = 'constant'
FLAGS = 'flags'
COMMENT = 'comment'
OPERANDU8 = '8bit unsigned operand'
FLAG = 'flag'
BOOL_NOT = 'not'
CYCLE_COUNT = 'cycles'
CYCLE_SEP = 'conditional cycles separator'
ERROR = 'error'

FLAG_MAY_CHANGE = 0
FLAG_SET_TO_ONE = 1
FLAG_SET_TO_ZERO = 2
FLAG_UNCHANGED = 3

class Instruction(object):
    def __init__(self, mnemonic, description, opcode, cycles, flags, comment, line):
        self.mnemonic = mnemonic
        self.opcode = opcode
        self.description = description
        self.cycles = cycles
        self.flags = flags
        self.comment = comment
        self.line = line

    @staticmethod
    def parse(line):
        mnemonic = line[:4].strip()
        description = tokenize(line[4:22])
        opcode = parse_opcode(line[21:26].strip())
        cycles = parse_cycles(line[33:38])
        flags = parse_flags(line[42:46])
        comment = line[51:]

        return Instruction(mnemonic, description, opcode, cycles, flags, comment, line)

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


def parse_flags(s):
    flags = []
    for flag in s:
        if flag == '-':
            flags.append(FLAG_UNCHANGED)
        elif flag == '0':
            flags.append(FLAG_SET_TO_ZERO)
        elif flag == '1':
            flags.append(FLAG_SET_TO_ONE)
        else:
            flags.append(FLAG_MAY_CHANGE)

    return flags


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
        t, i, p = apply(s, flag_re, i)
        if t:
            error = False
            tokens.append(Token(FLAG, p, t))
            continue
        t, i, p = apply(s, operand_ref_16bit_re, i)
        if t:
            error = False
            tokens.append(Token(OPERAND16REF, p, t.strip('()')))
            continue
        t, i, p = apply(s, not_re, i)
        if t:
            error = False
            tokens.append(Token(BOOL_NOT, p, t))
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


def get(tokens, code):
    for t in tokens:
        if t.code == code:
            return t

    return None


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


def parse_ld(tokens, s):
    opA = tokens[0]
    opB = tokens[1]

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
        dest = 'Set(FetchOperand16(), '
    elif opA.code == OPERANDU8:
        dest = 'int8(FetchOperand8())'
    elif opA.code in (REG8, REG16):
        dest = 'Set%s(' % opA.value
    elif opA.code == REF_REG:
        dest = 'Set(Get%s(), ' % opA.value

    if dest is None or src is None:
        return '// parsing error'

    return '{}{})'.format(dest, src)


def parse_rst(tokens, s):
    addr = unhexlify(tokens[0].value)[0]

    return  "Call(0x{:04x})".format(addr)


def parse_set(tokens, s):
    assert tokens[1].code == CONSTANT
    assert tokens[2].code in (REG8, REF_REG)
    n_bit = int(tokens[1].value)

    if tokens[0].value == 'res':
        value = '0'
    else:
        value = '1'

    if tokens[2].code == REG8:
        code = indent(dedent("""
        value := SetBit(Get{0}(), {1}, {2})
        Set{0}(value)""".format(tokens[2].value, n_bit, value)), '\t')
    else:
        code = indent(dedent("""
        addr := Get{0}()
        value := SetBit(Get(addr), {1}, {2})
        Set(addr, value)""".format(tokens[2].value, n_bit, value)), '\t')

    return code.strip()


def parse_swap(tokens, s):
    target = tokens[1]

    if target.code == REG8:
        operation = 'Set{0}(Swap(Get{0}()))'.format(target.value)
    elif target.code == REF_REG:
        operation = 'addr := Get%s()\n' % target.value
        operation += 'Set(addr, Swap(Get(addr)))'.format(target.value)

    return operation


def custom_impl_F8():
    """
    ld   HL, SP+%s
    """

    code =  "operand := int32(FetchOperand8())\n" +\
            "\tSetHL(uint16(int32(GetSP()) + operand))"


    return code


def parse_add(tokens, s):
    opA = tokens[1]
    opB = tokens[2]

    assert opA.code in (REG8, REG16), "first operand of add should be 8bit or 16bit register"
    assert opB.code in (REG8, REG16, OPERANDU8, OPERAND8, REF_REG), "second operand of add invalid %s" % opB.code

    opAstr = 'Get{}()'.format(opA.value)

    opBstr = '<error>'
    if opB.code in (REG8, REG16):
        opBstr = 'Get{}()'.format(opB.value)
    elif opB.code == OPERAND8:
        opBstr = 'FetchOperand8()'
    elif opB.code == OPERANDU8:
        opBstr = 'int8(FetchOperand8())'
    elif opB.code == OPERAND16REF:
        opBstr = 'Get(FetchOperand16())'
    elif opB.code == REF_REG:
        opBstr = 'Get(Get{}())'.format(opB.value)

    dest = 'Set{}'.format(opA.value)

    code = 'opA := {}\n'.format(opAstr)
    code += '\topB := {}\n'.format(opBstr)

    if tokens[0].value == 'add':
        code += '\t{}(opA + opB)'.format(dest)
    elif tokens[0].value == 'adc':
        code += '\n\tif GetFlagCy() {\n'
        code += '\t\t{}(opA + opB + 1)\n'.format(dest)
        code += '\t} else {\n'
        code += '\t\t{}(opA + opB)\n'.format(dest)
        code += '\t}'
    return code


def parse_inc(tokens, s):
    target = tokens[1]

    assert target.code in (REG8, REG16, REF_REG)

    if target.code in (REG16, REG8):
        code = "Inc{0}()".format(target.value)
    elif target.code == REF_REG:
        code = 'addr := Get%s()\n' % target.value
        code += "\tSet(addr, Get(addr) + 1)"

    return code


def parse_dec(tokens, s):
    target = tokens[1]

    assert target.code in (REG8, REG16, REF_REG)

    if target.code in (REG16, REG8):
        code = "Dec{0}()".format(target.value)
    elif target.code == REF_REG:
        code = 'addr := Get%s()\n' % target.value
        code += "\tSet(addr, Get(addr) - 1)"

    return code


def parse_push(tokens, s):
    reg = tokens[1]
    assert reg.code == REG16

    code = 'Push(Get{}())'.format(reg.value)
    return code


def parse_nop(tokens, s):
    return None


def get_flag(flag):
    if flag == 'CY':
        return 'GetFlagCy'
    elif flag == 'Z':
        return 'GetFlagZf'


def parse_call(tokens, s):
    if tokens[1].code == OPERAND16:
        return "Call(FetchOperand16())"
    elif tokens[1].code == BOOL_NOT:
        if tokens[2].code == FLAG:
            fl = get_flag(tokens[2].value)
            code = 'jmp := false\n\n'
            code += '\tif !{}() {{\n'.format(fl)
            code += '\t\tjmp = true\n'
            code += '\t\tCall(FetchOperand16())\n'
            code += '\t}'
            return code

    elif tokens[1].code == FLAG:
        fl = get_flag(tokens[1].value)
        code = 'jmp := false\n\n'
        code += '\tif {}() {{\n'.format(fl)
        code += '\t\tjmp = true\n'
        code += '\t\tCall(FetchOperand16())\n'
        code += '\t}'
        return code


def parse_bit(tokens, s):
    n_bit = tokens[1].value
    if tokens[2].code == REG8:
        target = 'Get{}()'.format(tokens[2].value)
    else:
        target = 'Get(Get{}())'.format(tokens[2].value)

    code = 'if set := GetBit({}, {}); set {{\n'.format(n_bit, target)
    code += '\t\tSetFlagZf(true)\n'
    code += '\t} else {\n'
    code += '\t\tSetFlagZf(false)\n'
    code += '\t}\n\n'

    code += '\tSetFlagN(false)\n'
    code += '\tSetFlagHc(false)\n'

    return code


dispatch = {\
    'ld': parse_ld,
    'rst': parse_rst,
    #'swap': parse_swap,
    #'inc': parse_inc,
    #'dec': parse_dec,
    #'push': parse_push,
    #'nop': parse_nop,
    #'res': parse_set,
    #'set': parse_set,
    #'add': parse_add,
    #'adc': parse_add,
    #'call': parse_call,
    #'bit': parse_bit,
}

custom_impl_table = {'F8': custom_impl_F8}


unsupp = set()

def make_func(instr, file):

    if instr.opcode[0] in custom_impl_table:
        result = custom_impl_table[instr.opcode[0]]()
    elif instr.mnemonic in dispatch:
        result = dispatch[instr.mnemonic](instr.description, instr.line)
    else:
        unsupp.add('unsupported mnemonic: %s' % instr.mnemonic)
        return

    comment = '// {} - {}'.format(instr.line[:21].strip(), instr.comment.strip('/ \n'))

    print(comment, file=file)
    print('func ' + instr.func_name() + '() uint8 {', file=file)

    print(instr.opcode[0])
    if instr.opcode[1]:
        oc = unhexlify(instr.opcode[0])[0] + 0xFF
    else:
        oc = unhexlify(instr.opcode[0])[0]

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


def make_dispatch_table(instructions):
    with open('../src/cpu/table.go', 'wt') as f:
        print("// this package is automatically generated.", file=f)
        print('package cpu\n', file=f)

        print('func init() {', file=f)

        for opcode, name in instructions:
            print('\tdispatch_table[0x{:03X}] = {}'.format(opcode, name), file=f)

        print('}', file=f)

header = \
"""
// this package is automatically generated by instruction_generator.py
package cpu

import . "memory"

"""


def main():
    with open('instructions.CPU') as cpu:
        with open('../src/cpu/instructions.go', 'wt') as f:

            print(header, file=f)

            func_names = []

            for line in cpu.readlines():
                if line.strip():
                    instruction = Instruction.parse(line)
                    n = make_func(instruction, f)
                    if n:
                        func_names.append(n)
            make_dispatch_table(func_names)
            for un in unsupp:
                print('// %s' % un, file=f)

if __name__ == "__main__":
    main()