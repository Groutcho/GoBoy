package cpu

import "testing"
import . "memory"

func TestOr(t* testing.T) {
	SetA(0x0F)
	SetB(0xF0)
	xB0_or()
	testRegister(t, REG_A, 0xFF)

	SetA(0x80)
	SetH(0x0F)
	xB4_or()
	testRegister(t, REG_A, 0x8F)

	SetA(0x80)
	Set(0x0000, 0x98)
	SetHL(0x0000)
	xB6_or()
	testRegister(t, REG_A, 0x98)

	SetA(0x80)
	Set(0x0000, 0x52)
	SetPC(0x0000)
	xF6_or()
	testRegister(t, REG_A, 0xD2)
}

func TestXor(t* testing.T) {
	SetA(0x0F)
	SetB(0xF0)
	xA8_xor()
	testRegister(t, REG_A, 0xFF)

	SetA(0x80)
	SetH(0x0F)
	xAC_xor()
	testRegister(t, REG_A, 0x8F)

	SetA(0x80)
	Set(0x0000, 0x98)
	SetHL(0x0000)
	xAE_xor()
	testRegister(t, REG_A, 0x18)

	SetA(0x80)
	Set(0x0000, 0x52)
	SetPC(0x0000)
	xEE_xor()
	testRegister(t, REG_A, 0xD2)
}

func TestAnd(t* testing.T) {
	SetA(0x0F)
	SetB(0xF0)
	xA0_and()
	testRegister(t, REG_A, 0x00)

	SetA(0x80)
	SetH(0x0F)
	xA4_and()
	testRegister(t, REG_A, 0x00)

	SetA(0x80)
	Set(0x0000, 0x98)
	SetHL(0x0000)
	xA6_and()
	testRegister(t, REG_A, 0x80)

	SetA(0x80)
	Set(0x0000, 0x52)
	SetPC(0x0000)
	xE6_and()
	testRegister(t, REG_A, 0x00)
}