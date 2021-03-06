// this file is automatically generated by instruction_generator.py
package cpu

import . "common"
import . "memory"

// bit  0, B - test bit 0 of register B
func xCB_40_bit() int {
	set := GetBit(GetB(), 0);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  0, C - test bit 0 of register C
func xCB_41_bit() int {
	set := GetBit(GetC(), 0);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  0, D - test bit 0 of register D
func xCB_42_bit() int {
	set := GetBit(GetD(), 0);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  0, E - test bit 0 of register E
func xCB_43_bit() int {
	set := GetBit(GetE(), 0);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  0, H - test bit 0 of register H
func xCB_44_bit() int {
	set := GetBit(GetH(), 0);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  0, L - test bit 0 of register L
func xCB_45_bit() int {
	set := GetBit(GetL(), 0);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  0, (HL) - [HL] & {2^0}
func xCB_46_bit() int {
	set := GetBit(Get(GetHL()), 0);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 3
}

// bit  0, A - test bit 0 of register A
func xCB_47_bit() int {
	set := GetBit(GetA(), 0);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  1, B - test bit 1 of register B
func xCB_48_bit() int {
	set := GetBit(GetB(), 1);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  1, C - test bit 1 of register C
func xCB_49_bit() int {
	set := GetBit(GetC(), 1);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  1, D - test bit 1 of register D
func xCB_4A_bit() int {
	set := GetBit(GetD(), 1);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  1, E - test bit 1 of register E
func xCB_4B_bit() int {
	set := GetBit(GetE(), 1);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  1, H - test bit 1 of register H
func xCB_4C_bit() int {
	set := GetBit(GetH(), 1);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  1, L - test bit 1 of register L
func xCB_4D_bit() int {
	set := GetBit(GetL(), 1);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  1, (HL) - [HL] & {2^1}
func xCB_4E_bit() int {
	set := GetBit(Get(GetHL()), 1);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 3
}

// bit  1, A - test bit 1 of register A
func xCB_4F_bit() int {
	set := GetBit(GetA(), 1);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  2, B - test bit 2 of register B
func xCB_50_bit() int {
	set := GetBit(GetB(), 2);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  2, C - test bit 2 of register C
func xCB_51_bit() int {
	set := GetBit(GetC(), 2);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  2, D - test bit 2 of register D
func xCB_52_bit() int {
	set := GetBit(GetD(), 2);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  2, E - test bit 2 of register E
func xCB_53_bit() int {
	set := GetBit(GetE(), 2);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  2, H - test bit 2 of register H
func xCB_54_bit() int {
	set := GetBit(GetH(), 2);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  2, L - test bit 2 of register L
func xCB_55_bit() int {
	set := GetBit(GetL(), 2);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  2, (HL) - [HL] & {2^2}
func xCB_56_bit() int {
	set := GetBit(Get(GetHL()), 2);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 3
}

// bit  2, A - test bit 2 of register A
func xCB_57_bit() int {
	set := GetBit(GetA(), 2);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  3, B - test bit 3 of register B
func xCB_58_bit() int {
	set := GetBit(GetB(), 3);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  3, C - test bit 3 of register C
func xCB_59_bit() int {
	set := GetBit(GetC(), 3);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  3, D - test bit 3 of register D
func xCB_5A_bit() int {
	set := GetBit(GetD(), 3);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  3, E - test bit 3 of register E
func xCB_5B_bit() int {
	set := GetBit(GetE(), 3);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  3, H - test bit 3 of register H
func xCB_5C_bit() int {
	set := GetBit(GetH(), 3);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  3, L - test bit 3 of register L
func xCB_5D_bit() int {
	set := GetBit(GetL(), 3);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  3, (HL) - [HL] & {2^3}
func xCB_5E_bit() int {
	set := GetBit(Get(GetHL()), 3);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 3
}

// bit  3, A - test bit 3 of register A
func xCB_5F_bit() int {
	set := GetBit(GetA(), 3);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  4, B - test bit 4 of register B
func xCB_60_bit() int {
	set := GetBit(GetB(), 4);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  4, C - test bit 4 of register C
func xCB_61_bit() int {
	set := GetBit(GetC(), 4);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  4, D - test bit 4 of register D
func xCB_62_bit() int {
	set := GetBit(GetD(), 4);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  4, E - test bit 4 of register E
func xCB_63_bit() int {
	set := GetBit(GetE(), 4);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  4, H - test bit 4 of register H
func xCB_64_bit() int {
	set := GetBit(GetH(), 4);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  4, L - test bit 4 of register L
func xCB_65_bit() int {
	set := GetBit(GetL(), 4);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  4, (HL) - [HL] & {2^4}
func xCB_66_bit() int {
	set := GetBit(Get(GetHL()), 4);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 3
}

// bit  4, A - test bit 4 of register A
func xCB_67_bit() int {
	set := GetBit(GetA(), 4);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  5, B - test bit 5 of register B
func xCB_68_bit() int {
	set := GetBit(GetB(), 5);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  5, C - test bit 5 of register C
func xCB_69_bit() int {
	set := GetBit(GetC(), 5);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  5, D - test bit 5 of register D
func xCB_6A_bit() int {
	set := GetBit(GetD(), 5);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  5, E - test bit 5 of register E
func xCB_6B_bit() int {
	set := GetBit(GetE(), 5);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  5, H - test bit 5 of register H
func xCB_6C_bit() int {
	set := GetBit(GetH(), 5);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  5, L - test bit 5 of register L
func xCB_6D_bit() int {
	set := GetBit(GetL(), 5);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  5, (HL) - [HL] & {2^5}
func xCB_6E_bit() int {
	set := GetBit(Get(GetHL()), 5);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 3
}

// bit  5, A - test bit 5 of register A
func xCB_6F_bit() int {
	set := GetBit(GetA(), 5);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  6, B - test bit 6 of register B
func xCB_70_bit() int {
	set := GetBit(GetB(), 6);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  6, C - test bit 6 of register C
func xCB_71_bit() int {
	set := GetBit(GetC(), 6);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  6, D - test bit 6 of register D
func xCB_72_bit() int {
	set := GetBit(GetD(), 6);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  6, E - test bit 6 of register E
func xCB_73_bit() int {
	set := GetBit(GetE(), 6);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  6, H - test bit 6 of register H
func xCB_74_bit() int {
	set := GetBit(GetH(), 6);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  6, L - test bit 6 of register L
func xCB_75_bit() int {
	set := GetBit(GetL(), 6);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  6, (HL) - [HL] & {2^6}
func xCB_76_bit() int {
	set := GetBit(Get(GetHL()), 6);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 3
}

// bit  6, A - test bit 6 of register A
func xCB_77_bit() int {
	set := GetBit(GetA(), 6);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  7, B - test bit 7 of register B
func xCB_78_bit() int {
	set := GetBit(GetB(), 7);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  7, C - test bit 7 of register C
func xCB_79_bit() int {
	set := GetBit(GetC(), 7);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  7, D - test bit 7 of register D
func xCB_7A_bit() int {
	set := GetBit(GetD(), 7);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  7, E - test bit 7 of register E
func xCB_7B_bit() int {
	set := GetBit(GetE(), 7);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  7, H - test bit 7 of register H
func xCB_7C_bit() int {
	set := GetBit(GetH(), 7);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  7, L - test bit 7 of register L
func xCB_7D_bit() int {
	set := GetBit(GetL(), 7);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

// bit  7, (HL) - [HL] & {2^7}
func xCB_7E_bit() int {
	set := GetBit(Get(GetHL()), 7);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 3
}

// bit  7, A - test bit 7 of register A
func xCB_7F_bit() int {
	set := GetBit(GetA(), 7);
	SetFlags(int(set), F_SET_IF, F_SET_0, F_SET_0, F_IGNORE, F_8bit)

	return 2
}

