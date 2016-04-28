#! /usr/bin/python3
#encoding=utf-8

import os
import sys
import argparse
from binascii import unhexlify

def parse_options():
    parser = argparse.ArgumentParser(
        description='disasm - disassemble GB binaries')
    parser.add_argument('filename',
        metavar="FILE",
        help='the file to disassemble')

    args = parser.parse_args()

    return args

def die(msg):
    print("FATAL: %s" % msg)
    sys.exit(1)

class Instruction(object):
    def __init__(mnemonic, opcode, operands):
        self.mnemonic = mnemonic
        self.opcode = opcode
        self.operands = operands

    def __str__(self):
        return "{:4}".format(self.mnemonic)

def parse_opcode(_bytes):
    elements = _bytes.split()
    opcode = None
    opcode_ex = None
    operandA = False
    operandB = False

    try:
        opcode = unhexlify(elements[0].strip())[0]
    except:
        pass
    if opcode == 0xCB:
        try:
            opcode_ex = unhexlify(elements[1].strip())[0]
        except:
            pass
    
    if not opcode_ex:
        if len(elements) > 1:
            operandA = True

        if len(elements) == 3:
            operandB = True

    return opcode, opcode_ex, operandA, operandB

func_template = """/*
{}
*/
func x{:X}{}_{}({}) {{
    // TODO: implement
}}"""

def print_go_func(comment, mnemonic, signature, opcode, cycles, flags):
    args = ''
    if opcode[2]:
        args += "opA uint8"
    if opcode[3]:
        args += ", opB uint8"
    comment = comment + ' - ' + signature if signature else comment
    print(func_template.format(comment, opcode[0], hex(opcode[1]) if opcode[1] else '', mnemonic, args))
    print('')


def load_instructions():
    instr_map = dict()
    lineno = 1
    with open('instructions') as fd:
        for line in fd.readlines():
            line = line.strip()
            if not line or line[0] == '#':
                lineno += 1
                continue

            spl = line.split(':')

            try:
                mnemonic = spl[0].strip()
                signature = spl[1].strip()
                opcode = parse_opcode(spl[2].strip())
                cycles = spl[3].strip()
                flags = spl[4].strip()
                comment = spl[5].strip()

                # print(mnemonic, signature, opcode, cycles, flags, comment)
                if opcode[0]:
                    print_go_func(comment, mnemonic, signature, opcode, cycles, flags)
            except IndexError as ex:
                print('{:3d}: {}'.format(lineno, spl))

            lineno += 1

            # extended = None
            # signature = line[5:15].strip()
            # try:
            #     opcode = unhexlify([15:17])[0]
            #     if opcode == 0xCB:
            #         extended = unhexlify(line[18:20])[0]
            #     else:
            #         operands = line[24:]

            #     instr_map[opcode] = (mnemonic, signature, opcode, extended)
            # except:
            #     pass

    return instr_map


def main():
    args = parse_options()
    instructions = load_instructions()
    return

    if not os.path.exists(args.filename):
        die("no such file: %s" % args.filename)

    with open(args.filename, 'rb') as fd:
        while True:
            byte = fd.read(1)
            if (len(byte) == 0):
                sys.exit(0)

            opcode = byte[0]
            if opcode in instructions:
                mnemonic, sig, opcode, extended = instructions[opcode]
                print('{:6}'.format(mnemonic), end='')
                if sig:
                    print(' ', end='')
                    print('{:10}'.format(sig), end='')
                print('')
            else:
                pass
                # print(hex(opcode))
            # input()



if __name__ == "__main__":
    main()