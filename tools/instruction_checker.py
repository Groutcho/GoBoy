#!/usr/bin/python3
#encoding=utf-8

import sys
from binascii import unhexlify

class Instruction(object):
	def __init__(self, mnem, opcode, extended=False, comment='// MISSING comment'):
		self.mnemonic = mnem,
		self.opcode = opcode
		self.extended = extended
		self.comment = comment

	def full_opcode(self):
		if self.extended:
			return 'CB_{}'.format(self.opcode)
		return self.opcode

def check():
	with open('instructions.CPU') as cpu:
		opcodes = set()
		lineno = 1
		error = False
		infos = []
		for line in cpu.readlines():
			line = line.strip()
			if line == '' or line[0] == '#':
				lineno += 1
				continue

			info = [0, 0, 0, 0, 0]

			opcode = line[21:32].replace('nn', '').strip()
			if opcode in opcodes:
				pass
				# error = True
				# print('ERROR:{}: {}'.format(lineno, opcode))
				# sys.exit(1)
			else:
				extended = False
				if opcode[:2] == 'CB':
					opcode = unhexlify(opcode[-2:])
					extended = True
				else:
					opcode = unhexlify(opcode[:2])

				print(hex(opcode[0]))

				opcodes.add(opcode)
				mnemonic = line[0:5].strip() #mnemonic
				comment = line[51:]
				infos.append(Instruction(mnemonic, opcode, extended, comment))
			lineno += 1

		# if not error:
		# 	print('{} instructions checked. No error found.'.format(len(opcodes)))

		return infos

func_template = '{}\nfunc x{}_{}() {{\n{}\n}}\n'

def make_templates(instructions):
	print("package cpu\n")
	for ix in instructions:
		print(func_template.format(ix.comment, ix.full_opcode(), ix.mnemonic, '    // TODO: implement this instruction'))

# package main

# import "fmt"

# func hello() {
# 	fmt.Println("hello, world!")
# }

# type voidFunc func()

# func main() {
# 	arr := make([]voidFunc, 40, 40)
# 	for i := 0; i < len(arr); i++ {
# 		arr[i] = hello
# 	}
# 	for i := 0; i < len(arr); i++ {
# 		arr[i]()
# 	}
# }
def make_dispatch_table(instructions):
	# print('dispatch := make([]voidFunc, 512, 512)')
	for ix in instructions:
		opcode = ix.opcode
		print(opcode)
		return
		if ix.extended:
			opcode += 0xCB
		# print('dispatch[0x0{}] = x{}_{}'.format(ix.opcode, ix.full_opcode().upper(), ix.mnemonic))


lines = check()
# make_templates(lines)
make_dispatch_table(lines)