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

write_packages = ('ld', 'add', 'adc', 'bit', 'set', 'res', 'sub', 'sbc', 'rst')
write_table = ('ld', 'call', 'add', 'adc', 'bit', 'set', 'res', 'sub', 'sbc', 'rst', 'pop')

dont_require_memory = ('rst',)

class Instruction(object):
    def __init__(self, mnemonic, pseudocode, opcode, cycles, flags, comment, line):
        self.mnemonic = mnemonic
        self.opcode = opcode
        self.pseudocode = pseudocode
        self.cycles = cycles
        self.flags = flags
        self.comment = comment
        self.line = line

    @staticmethod
    def parse(line):
        mnemonic = line[:4].strip()
        pseudocode = tokenize(line[4:22])
        opcode = parse_opcode(line[21:26].strip())
        cycles = parse_cycles(line[33:38])
        flags = parse_flags(line[42:46])
        comment = line[51:]

        return Instruction(mnemonic, pseudocode, opcode, cycles, flags, comment, line)

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
        Set(addr, value)""".format(instr.pseudocode[1].value, n_bit, value)), '\t')

    return code.strip()


def parse_swap(instruction):
    target = instruction.pseudocode[1]

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

    code = \
    """
    operand := int(FetchOperand8s())
    result := int(GetSP()) + operand
    SetHL(uint16(result))

    hc := 0
    if operand > 0 {
        if getLowNibble(uint8(operand)) + getLowNibble(getLowBits(GetSP())) > 0xF {
            hc = 1
        } else {
            hc = 0
        }
    }

    SetFlags(result, F_SET_0, F_SET_0, hc, F_SET_IF, F_16bit)
    """

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
    hc = "IsHalfCarry(left, right)"
    width = 8
    zFlag = "F_SET_IF"

    left = 'Get{}()'.format(opA.value)
    dest = opA.value
    if opA.code == REG16:
        width = 16
        zFlag = "F_IGNORE"
        if opB.code == OPERANDU8:
            hc = "IsHalfCarry(getHighBits(left), right)"
        else:
            hc = "IsHalfCarry(getHighBits(left), getHighBits(right))"

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


def parse_inc(instruction):
    target = instruction.pseudocode[0]

    assert target.code in (REG8, REG16, REF_REG)

    if target.code in (REG16, REG8):
        code = "Inc{0}()".format(target.value)
    elif target.code == REF_REG:
        code = 'addr := Get%s()\n' % target.value
        code += "\tSet(addr, Get(addr) + 1)"

    return code


def parse_dec(instruction):
    target = instruction.pseudocode[0]

    assert target.code in (REG8, REG16, REF_REG)

    if target.code in (REG16, REG8):
        code = "Dec{0}()".format(target.value)
    elif target.code == REF_REG:
        code = 'addr := Get%s()\n' % target.value
        code += "\tSet(addr, Get(addr) - 1)"

    return code


def parse_push(instruction):
    reg = instruction.pseudocode[1]
    assert reg.code == REG16

    code = 'Push(Get{}())'.format(reg.value)
    return code


def parse_nop(instruction):
    return None


def get_flag(flag):
    if flag == 'CY':
        return 'GetFlagCy'
    elif flag == 'Z':
        return 'GetFlagZf'


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
    'dec': parse_dec,
    #'push': parse_push,
    'nop': parse_nop,
    'res': parse_set,
    'set': parse_set,
    'add': parse_add,
    'adc': parse_add,
    'sub': parse_add,
    'sbc': parse_add,
    'bit': parse_bit,
}

custom_impl_table = {0xF8: custom_impl_F8}


unsupp = set()

def make_func(instr, file):
    if instr.opcode[1]:
        oc = unhexlify(instr.opcode[0])[0] + 0xFF
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


def make_dispatch_table(instructions:dict):
    with open('../src/cpu/table.go', 'wt') as f:
        print("// this package is automatically generated.", file=f)
        print('package cpu\n', file=f)

        print('func init() {', file=f)

        for k, v in instructions.items():
            if k in write_table:
                for instr in v:
                    if instr.opcode[1]:
                        oc = 0xFF + unhexlify(instr.opcode[0])[0]
                    else:
                        oc = unhexlify(instr.opcode[0])[0]
                    print('\tdispatch_table[0x{:03X}] = {}'.format(oc, instr.func_name()), file=f)

        print('}', file=f)

header = \
"""// this package is automatically generated by instruction_generator.py
package cpu

import . "memory"

"""

header_nomem = \
"""// this package is automatically generated by instruction_generator.py
package cpu

"""

def parse_file(filename):
    result = dict()
    with open(filename) as cpu:
        for line in cpu.readlines():
            if line.strip():
                instruction = Instruction.parse(line)
                if instruction.mnemonic not in result:
                    result[instruction.mnemonic] = []
                result[instruction.mnemonic].append(instruction)
    return result

def main():
    instructions = parse_file('instructions.CPU')

    for mnemonic in instructions:
        if mnemonic in write_packages:
            with open('../src/cpu/{}.go'.format(mnemonic), 'wt') as f:
                if mnemonic not in dont_require_memory:
                    print(header, file=f)
                else:
                    print(header_nomem, file=f)


                for instruction in instructions[mnemonic]:
                    n = make_func(instruction, f)

    make_dispatch_table(instructions)
    for un in unsupp:
        print(un)

if __name__ == "__main__":
    main()