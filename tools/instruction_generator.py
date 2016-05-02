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


def tokenize(s):
    tokens = []
    i = 0
    t = 0

    limit = 40

    while limit and i < len(s) -1:
        error = True
        t, i, p = apply(s, mnemonic_re, i)
        if t:
            error = False
            tokens.append(Token(MNEMONIC, p, t))
            continue
        t, i, p = apply(s, cond_cycle_re, i)
        if t:
            error = False
            tokens.append(Token(CYCLE_SEP, p, t))
            continue
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
        t, i, p = apply(s, ex_opcode_re, i)
        if t:
            error = False
            tokens.append(Token(EX_OPCODE, p, t))
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
        t, i, p = apply(s, opcode_re, i)
        if t:
            error = False
            if 34 < p < 38:
                tokens.append(Token(CYCLE_COUNT, p, t))
            elif i > 21:
                tokens.append(Token(OPCODE, p, t))
            else:
                tokens.append(Token(CONSTANT, p, t))
            continue
        t, i, p = apply(s, not_re, i)
        if t:
            error = False
            tokens.append(Token(BOOL_NOT, p, t))
            continue
        t, i, p = apply(s, constant_re, i)
        if t:
            error = False
            if 34 < p < 38:
                tokens.append(Token(CYCLE_COUNT, p, t))
            else:
                tokens.append(Token(CONSTANT, p, t))
            continue
        t, i, p = apply(s, flags_re, i)
        if t:
            error = False
            tokens.append(Token(FLAGS, p, t))
            continue
        t, i, p = apply(s, comment_re, i)
        if t:
            error = False
            tokens.append(Token(COMMENT, p, t.lstrip('/ ')))
            continue

        limit -= 1
        if error:
            tokens.append(Token(ERROR, p, None))
            i+=1

    if limit == 0:
        print('could not parse "%s"' % s[i:])

    return tokens


def get(tokens, code):
    for t in tokens:
        if t.code == code:
            return t

    return None


def parse_cycles(s):
    cycles = s[34:38].strip()
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
    assert tokens[0].code == MNEMONIC
    opA = tokens[1]
    opB = tokens[2]

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
    addr = unhexlify(tokens[1].value)[0]
    
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
    #'rst': parse_rst,
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


unsupp = set()

def make_func(tokens, s, file):
    s = s.strip()
    desc = s[:21].strip()
    cycle_a, cycle_b = parse_cycles(s)
    extended = False

    opcode = get(tokens, OPCODE)
    if not opcode:
        extended = True
        opcode = get(tokens, EX_OPCODE)

    mnemonic = tokens[0]

    if mnemonic.value in dispatch:
        result = dispatch[mnemonic.value](tokens, s)
    else:
        unsupp.add('unsupported mnemonic: %s' % mnemonic.value)
        return

    if extended:
        func_name = "x{}_{}".format(opcode.value.replace(' ', '_'), mnemonic.value)
    else:
        func_name = "x{}_{}".format(opcode.value, mnemonic.value)
    comment = '// {} - {}'.format(desc, get(tokens, COMMENT).value)

    print(comment, file=file)
    print('func ' + func_name + '() uint8 {', file=file)

    if opcode.code == EX_OPCODE:
        oc = unhexlify(opcode.value[-2:])[0] + 0xFF
        ret =(oc, func_name)
    else:
        oc = unhexlify(opcode.value)[0]
        ret = (oc, func_name)

    if result:
        print('\t%s\n' % result, file=file)

    if cycle_b == 0:
        print('\treturn %d' % cycle_a, file=file)  

    else:
        print(indent(dedent("""\
        if jmp {{
        \treturn {}
        }} else {{
        \treturn {}
        }}""".format(cycle_a, cycle_b)), '\t'), file=file)

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


def main(use_file=True):
    if use_file:
        with open('instructions.CPU') as cpu:
            with open('../src/cpu/instructions.go', 'wt') as f:
                print("// this package is automatically generated.", file=f)
                print('package cpu\n', file=f)
                print('import . "memory"\n', file=f)

                func_names = []

                for line in cpu.readlines():
                    if line.strip():
                        tokens = tokenize(line)
                        n = make_func(tokens, line, f)
                        if n:
                            func_names.append(n)
                make_dispatch_table(func_names)
                for un in unsupp:
                    print('// %s' % un, file=f)

    else:
        txt = """
bit  0,H             CB 44           8    z01-     // test bit 0 of register H
bit  1,H             CB 4C           8    z01-     // test bit 1 of register H
bit  2,H             CB 54           8    z01-     // test bit 2 of register H
bit  3,H             CB 5C           8    z01-     // test bit 3 of register H
bit  4,H             CB 64           8    z01-     // test bit 4 of register H
        """
        for line in txt.split('\n'):
            if line.strip():
                tokens = tokenize(line)
                try:
                    make_func(tokens, line)
                except AssertionError as err:
                    print('{}: {}'.format(line, err))

if __name__ == "__main__":
    main()