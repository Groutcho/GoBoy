// this file is automatically generated by instruction_generator.py
package cpu

func init() {
	dispatch_table[0x001] = x01_ld
	dispatch_table[0x002] = x02_ld
	dispatch_table[0x003] = x03_inc
	dispatch_table[0x004] = x04_inc
	dispatch_table[0x005] = x05_dec
	dispatch_table[0x006] = x06_ld
	dispatch_table[0x007] = x07_rlca
	dispatch_table[0x009] = x09_add
	dispatch_table[0x00A] = x0A_ld
	dispatch_table[0x00B] = x0B_dec
	dispatch_table[0x00C] = x0C_inc
	dispatch_table[0x00D] = x0D_dec
	dispatch_table[0x00E] = x0E_ld
	dispatch_table[0x00F] = x0F_rrca
	dispatch_table[0x011] = x11_ld
	dispatch_table[0x012] = x12_ld
	dispatch_table[0x013] = x13_inc
	dispatch_table[0x014] = x14_inc
	dispatch_table[0x015] = x15_dec
	dispatch_table[0x016] = x16_ld
	dispatch_table[0x017] = x17_rla
	dispatch_table[0x018] = x18_jr
	dispatch_table[0x019] = x19_add
	dispatch_table[0x01A] = x1A_ld
	dispatch_table[0x01B] = x1B_dec
	dispatch_table[0x01C] = x1C_inc
	dispatch_table[0x01D] = x1D_dec
	dispatch_table[0x01E] = x1E_ld
	dispatch_table[0x01F] = x1F_rra
	dispatch_table[0x020] = x20_jr
	dispatch_table[0x021] = x21_ld
	dispatch_table[0x022] = x22_ldi
	dispatch_table[0x023] = x23_inc
	dispatch_table[0x024] = x24_inc
	dispatch_table[0x025] = x25_dec
	dispatch_table[0x026] = x26_ld
	dispatch_table[0x028] = x28_jr
	dispatch_table[0x029] = x29_add
	dispatch_table[0x02A] = x2A_ldi
	dispatch_table[0x02B] = x2B_dec
	dispatch_table[0x02C] = x2C_inc
	dispatch_table[0x02D] = x2D_dec
	dispatch_table[0x02E] = x2E_ld
	dispatch_table[0x02F] = x2F_cpl
	dispatch_table[0x030] = x30_jr
	dispatch_table[0x031] = x31_ld
	dispatch_table[0x032] = x32_ldd
	dispatch_table[0x033] = x33_inc
	dispatch_table[0x034] = x34_inc
	dispatch_table[0x035] = x35_dec
	dispatch_table[0x036] = x36_ld
	dispatch_table[0x038] = x38_jr
	dispatch_table[0x039] = x39_add
	dispatch_table[0x03A] = x3A_ldd
	dispatch_table[0x03B] = x3B_dec
	dispatch_table[0x03C] = x3C_inc
	dispatch_table[0x03D] = x3D_dec
	dispatch_table[0x03E] = x3E_ld
	dispatch_table[0x040] = x40_ld
	dispatch_table[0x041] = x41_ld
	dispatch_table[0x042] = x42_ld
	dispatch_table[0x043] = x43_ld
	dispatch_table[0x044] = x44_ld
	dispatch_table[0x045] = x45_ld
	dispatch_table[0x046] = x46_ld
	dispatch_table[0x047] = x47_ld
	dispatch_table[0x048] = x48_ld
	dispatch_table[0x049] = x49_ld
	dispatch_table[0x04A] = x4A_ld
	dispatch_table[0x04B] = x4B_ld
	dispatch_table[0x04C] = x4C_ld
	dispatch_table[0x04D] = x4D_ld
	dispatch_table[0x04E] = x4E_ld
	dispatch_table[0x04F] = x4F_ld
	dispatch_table[0x050] = x50_ld
	dispatch_table[0x051] = x51_ld
	dispatch_table[0x052] = x52_ld
	dispatch_table[0x053] = x53_ld
	dispatch_table[0x054] = x54_ld
	dispatch_table[0x055] = x55_ld
	dispatch_table[0x056] = x56_ld
	dispatch_table[0x057] = x57_ld
	dispatch_table[0x058] = x58_ld
	dispatch_table[0x059] = x59_ld
	dispatch_table[0x05A] = x5A_ld
	dispatch_table[0x05B] = x5B_ld
	dispatch_table[0x05C] = x5C_ld
	dispatch_table[0x05D] = x5D_ld
	dispatch_table[0x05E] = x5E_ld
	dispatch_table[0x05F] = x5F_ld
	dispatch_table[0x060] = x60_ld
	dispatch_table[0x061] = x61_ld
	dispatch_table[0x062] = x62_ld
	dispatch_table[0x063] = x63_ld
	dispatch_table[0x064] = x64_ld
	dispatch_table[0x065] = x65_ld
	dispatch_table[0x066] = x66_ld
	dispatch_table[0x067] = x67_ld
	dispatch_table[0x068] = x68_ld
	dispatch_table[0x069] = x69_ld
	dispatch_table[0x06A] = x6A_ld
	dispatch_table[0x06B] = x6B_ld
	dispatch_table[0x06C] = x6C_ld
	dispatch_table[0x06D] = x6D_ld
	dispatch_table[0x06E] = x6E_ld
	dispatch_table[0x06F] = x6F_ld
	dispatch_table[0x070] = x70_ld
	dispatch_table[0x071] = x71_ld
	dispatch_table[0x072] = x72_ld
	dispatch_table[0x073] = x73_ld
	dispatch_table[0x074] = x74_ld
	dispatch_table[0x075] = x75_ld
	dispatch_table[0x077] = x77_ld
	dispatch_table[0x078] = x78_ld
	dispatch_table[0x079] = x79_ld
	dispatch_table[0x07A] = x7A_ld
	dispatch_table[0x07B] = x7B_ld
	dispatch_table[0x07C] = x7C_ld
	dispatch_table[0x07D] = x7D_ld
	dispatch_table[0x07E] = x7E_ld
	dispatch_table[0x07F] = x7F_ld
	dispatch_table[0x080] = x80_add
	dispatch_table[0x081] = x81_add
	dispatch_table[0x082] = x82_add
	dispatch_table[0x083] = x83_add
	dispatch_table[0x084] = x84_add
	dispatch_table[0x085] = x85_add
	dispatch_table[0x086] = x86_add
	dispatch_table[0x087] = x87_add
	dispatch_table[0x088] = x88_adc
	dispatch_table[0x089] = x89_adc
	dispatch_table[0x08A] = x8A_adc
	dispatch_table[0x08B] = x8B_adc
	dispatch_table[0x08C] = x8C_adc
	dispatch_table[0x08D] = x8D_adc
	dispatch_table[0x08E] = x8E_adc
	dispatch_table[0x08F] = x8F_adc
	dispatch_table[0x090] = x90_sub
	dispatch_table[0x091] = x91_sub
	dispatch_table[0x092] = x92_sub
	dispatch_table[0x093] = x93_sub
	dispatch_table[0x094] = x94_sub
	dispatch_table[0x095] = x95_sub
	dispatch_table[0x096] = x96_sub
	dispatch_table[0x097] = x97_sub
	dispatch_table[0x098] = x98_sbc
	dispatch_table[0x099] = x99_sbc
	dispatch_table[0x09A] = x9A_sbc
	dispatch_table[0x09B] = x9B_sbc
	dispatch_table[0x09C] = x9C_sbc
	dispatch_table[0x09D] = x9D_sbc
	dispatch_table[0x09E] = x9E_sbc
	dispatch_table[0x09F] = x9F_sbc
	dispatch_table[0x0A0] = xA0_and
	dispatch_table[0x0A1] = xA1_and
	dispatch_table[0x0A2] = xA2_and
	dispatch_table[0x0A3] = xA3_and
	dispatch_table[0x0A4] = xA4_and
	dispatch_table[0x0A5] = xA5_and
	dispatch_table[0x0A6] = xA6_and
	dispatch_table[0x0A7] = xA7_and
	dispatch_table[0x0A8] = xA8_xor
	dispatch_table[0x0A9] = xA9_xor
	dispatch_table[0x0AA] = xAA_xor
	dispatch_table[0x0AB] = xAB_xor
	dispatch_table[0x0AC] = xAC_xor
	dispatch_table[0x0AD] = xAD_xor
	dispatch_table[0x0AE] = xAE_xor
	dispatch_table[0x0AF] = xAF_xor
	dispatch_table[0x0B0] = xB0_or
	dispatch_table[0x0B1] = xB1_or
	dispatch_table[0x0B2] = xB2_or
	dispatch_table[0x0B3] = xB3_or
	dispatch_table[0x0B4] = xB4_or
	dispatch_table[0x0B5] = xB5_or
	dispatch_table[0x0B6] = xB6_or
	dispatch_table[0x0B7] = xB7_or
	dispatch_table[0x0B8] = xB8_cp
	dispatch_table[0x0B9] = xB9_cp
	dispatch_table[0x0BA] = xBA_cp
	dispatch_table[0x0BB] = xBB_cp
	dispatch_table[0x0BC] = xBC_cp
	dispatch_table[0x0BD] = xBD_cp
	dispatch_table[0x0BE] = xBE_cp
	dispatch_table[0x0BF] = xBF_cp
	dispatch_table[0x0C0] = xC0_ret
	dispatch_table[0x0C1] = xC1_pop
	dispatch_table[0x0C2] = xC2_jp
	dispatch_table[0x0C3] = xC3_jp
	dispatch_table[0x0C4] = xC4_call
	dispatch_table[0x0C5] = xC5_push
	dispatch_table[0x0C6] = xC6_add
	dispatch_table[0x0C7] = xC7_rst
	dispatch_table[0x0C8] = xC8_ret
	dispatch_table[0x0C9] = xC9_ret
	dispatch_table[0x0CA] = xCA_jp
	dispatch_table[0x0CC] = xCC_call
	dispatch_table[0x0CD] = xCD_call
	dispatch_table[0x0CE] = xCE_adc
	dispatch_table[0x0CF] = xCF_rst
	dispatch_table[0x0D0] = xD0_ret
	dispatch_table[0x0D1] = xD1_pop
	dispatch_table[0x0D2] = xD2_jp
	dispatch_table[0x0D4] = xD4_call
	dispatch_table[0x0D5] = xD5_push
	dispatch_table[0x0D6] = xD6_sub
	dispatch_table[0x0D7] = xD7_rst
	dispatch_table[0x0D8] = xD8_ret
	dispatch_table[0x0D9] = xD9_reti
	dispatch_table[0x0DA] = xDA_jp
	dispatch_table[0x0DC] = xDC_call
	dispatch_table[0x0DE] = xDE_sbc
	dispatch_table[0x0DF] = xDF_rst
	dispatch_table[0x0E0] = xE0_ld
	dispatch_table[0x0E1] = xE1_pop
	dispatch_table[0x0E2] = xE2_ld
	dispatch_table[0x0E5] = xE5_push
	dispatch_table[0x0E6] = xE6_and
	dispatch_table[0x0E7] = xE7_rst
	dispatch_table[0x0E8] = xE8_add
	dispatch_table[0x0E9] = xE9_jp
	dispatch_table[0x0EA] = xEA_ld
	dispatch_table[0x0EE] = xEE_xor
	dispatch_table[0x0EF] = xEF_rst
	dispatch_table[0x0F0] = xF0_ld
	dispatch_table[0x0F1] = xF1_pop
	dispatch_table[0x0F2] = xF2_ld
	dispatch_table[0x0F5] = xF5_push
	dispatch_table[0x0F6] = xF6_or
	dispatch_table[0x0F7] = xF7_rst
	dispatch_table[0x0F8] = xF8_ld
	dispatch_table[0x0F9] = xF9_ld
	dispatch_table[0x0FA] = xFA_ld
	dispatch_table[0x0FE] = xFE_cp
	dispatch_table[0x0FF] = xFF_rst
	dispatch_table[0x100] = xCB_00_rlc
	dispatch_table[0x101] = xCB_01_rlc
	dispatch_table[0x102] = xCB_02_rlc
	dispatch_table[0x103] = xCB_03_rlc
	dispatch_table[0x104] = xCB_04_rlc
	dispatch_table[0x105] = xCB_05_rlc
	dispatch_table[0x106] = xCB_06_rlc
	dispatch_table[0x107] = xCB_07_rlc
	dispatch_table[0x108] = xCB_08_rrc
	dispatch_table[0x109] = xCB_09_rrc
	dispatch_table[0x10A] = xCB_0A_rrc
	dispatch_table[0x10B] = xCB_0B_rrc
	dispatch_table[0x10C] = xCB_0C_rrc
	dispatch_table[0x10D] = xCB_0D_rrc
	dispatch_table[0x10E] = xCB_0E_rrc
	dispatch_table[0x10F] = xCB_0F_rrc
	dispatch_table[0x110] = xCB_10_rl
	dispatch_table[0x111] = xCB_11_rl
	dispatch_table[0x112] = xCB_12_rl
	dispatch_table[0x113] = xCB_13_rl
	dispatch_table[0x114] = xCB_14_rl
	dispatch_table[0x115] = xCB_15_rl
	dispatch_table[0x116] = xCB_16_rl
	dispatch_table[0x117] = xCB_17_rl
	dispatch_table[0x118] = xCB_18_rr
	dispatch_table[0x119] = xCB_19_rr
	dispatch_table[0x11A] = xCB_1A_rr
	dispatch_table[0x11B] = xCB_1B_rr
	dispatch_table[0x11C] = xCB_1C_rr
	dispatch_table[0x11D] = xCB_1D_rr
	dispatch_table[0x11E] = xCB_1E_rr
	dispatch_table[0x11F] = xCB_1F_rr
	dispatch_table[0x120] = xCB_20_sla
	dispatch_table[0x121] = xCB_21_sla
	dispatch_table[0x122] = xCB_22_sla
	dispatch_table[0x123] = xCB_23_sla
	dispatch_table[0x124] = xCB_24_sla
	dispatch_table[0x125] = xCB_25_sla
	dispatch_table[0x126] = xCB_26_sla
	dispatch_table[0x127] = xCB_27_sla
	dispatch_table[0x128] = xCB_28_sra
	dispatch_table[0x129] = xCB_29_sra
	dispatch_table[0x12A] = xCB_2A_sra
	dispatch_table[0x12B] = xCB_2B_sra
	dispatch_table[0x12C] = xCB_2C_sra
	dispatch_table[0x12D] = xCB_2D_sra
	dispatch_table[0x12E] = xCB_2E_sra
	dispatch_table[0x12F] = xCB_2F_sra
	dispatch_table[0x130] = xCB_30_swap
	dispatch_table[0x131] = xCB_31_swap
	dispatch_table[0x132] = xCB_32_swap
	dispatch_table[0x133] = xCB_33_swap
	dispatch_table[0x134] = xCB_34_swap
	dispatch_table[0x135] = xCB_35_swap
	dispatch_table[0x136] = xCB_36_swap
	dispatch_table[0x137] = xCB_37_swap
	dispatch_table[0x138] = xCB_38_srl
	dispatch_table[0x139] = xCB_39_srl
	dispatch_table[0x13A] = xCB_3A_srl
	dispatch_table[0x13B] = xCB_3B_srl
	dispatch_table[0x13C] = xCB_3C_srl
	dispatch_table[0x13D] = xCB_3D_srl
	dispatch_table[0x13E] = xCB_3E_srl
	dispatch_table[0x13F] = xCB_3F_srl
	dispatch_table[0x140] = xCB_40_bit
	dispatch_table[0x141] = xCB_41_bit
	dispatch_table[0x142] = xCB_42_bit
	dispatch_table[0x143] = xCB_43_bit
	dispatch_table[0x144] = xCB_44_bit
	dispatch_table[0x145] = xCB_45_bit
	dispatch_table[0x146] = xCB_46_bit
	dispatch_table[0x147] = xCB_47_bit
	dispatch_table[0x148] = xCB_48_bit
	dispatch_table[0x149] = xCB_49_bit
	dispatch_table[0x14A] = xCB_4A_bit
	dispatch_table[0x14B] = xCB_4B_bit
	dispatch_table[0x14C] = xCB_4C_bit
	dispatch_table[0x14D] = xCB_4D_bit
	dispatch_table[0x14E] = xCB_4E_bit
	dispatch_table[0x14F] = xCB_4F_bit
	dispatch_table[0x150] = xCB_50_bit
	dispatch_table[0x151] = xCB_51_bit
	dispatch_table[0x152] = xCB_52_bit
	dispatch_table[0x153] = xCB_53_bit
	dispatch_table[0x154] = xCB_54_bit
	dispatch_table[0x155] = xCB_55_bit
	dispatch_table[0x156] = xCB_56_bit
	dispatch_table[0x157] = xCB_57_bit
	dispatch_table[0x158] = xCB_58_bit
	dispatch_table[0x159] = xCB_59_bit
	dispatch_table[0x15A] = xCB_5A_bit
	dispatch_table[0x15B] = xCB_5B_bit
	dispatch_table[0x15C] = xCB_5C_bit
	dispatch_table[0x15D] = xCB_5D_bit
	dispatch_table[0x15E] = xCB_5E_bit
	dispatch_table[0x15F] = xCB_5F_bit
	dispatch_table[0x160] = xCB_60_bit
	dispatch_table[0x161] = xCB_61_bit
	dispatch_table[0x162] = xCB_62_bit
	dispatch_table[0x163] = xCB_63_bit
	dispatch_table[0x164] = xCB_64_bit
	dispatch_table[0x165] = xCB_65_bit
	dispatch_table[0x166] = xCB_66_bit
	dispatch_table[0x167] = xCB_67_bit
	dispatch_table[0x168] = xCB_68_bit
	dispatch_table[0x169] = xCB_69_bit
	dispatch_table[0x16A] = xCB_6A_bit
	dispatch_table[0x16B] = xCB_6B_bit
	dispatch_table[0x16C] = xCB_6C_bit
	dispatch_table[0x16D] = xCB_6D_bit
	dispatch_table[0x16E] = xCB_6E_bit
	dispatch_table[0x16F] = xCB_6F_bit
	dispatch_table[0x170] = xCB_70_bit
	dispatch_table[0x171] = xCB_71_bit
	dispatch_table[0x172] = xCB_72_bit
	dispatch_table[0x173] = xCB_73_bit
	dispatch_table[0x174] = xCB_74_bit
	dispatch_table[0x175] = xCB_75_bit
	dispatch_table[0x176] = xCB_76_bit
	dispatch_table[0x177] = xCB_77_bit
	dispatch_table[0x178] = xCB_78_bit
	dispatch_table[0x179] = xCB_79_bit
	dispatch_table[0x17A] = xCB_7A_bit
	dispatch_table[0x17B] = xCB_7B_bit
	dispatch_table[0x17C] = xCB_7C_bit
	dispatch_table[0x17D] = xCB_7D_bit
	dispatch_table[0x17E] = xCB_7E_bit
	dispatch_table[0x17F] = xCB_7F_bit
	dispatch_table[0x180] = xCB_80_res
	dispatch_table[0x181] = xCB_81_res
	dispatch_table[0x182] = xCB_82_res
	dispatch_table[0x183] = xCB_83_res
	dispatch_table[0x184] = xCB_84_res
	dispatch_table[0x185] = xCB_85_res
	dispatch_table[0x186] = xCB_86_res
	dispatch_table[0x187] = xCB_87_res
	dispatch_table[0x188] = xCB_88_res
	dispatch_table[0x189] = xCB_89_res
	dispatch_table[0x18A] = xCB_8A_res
	dispatch_table[0x18B] = xCB_8B_res
	dispatch_table[0x18C] = xCB_8C_res
	dispatch_table[0x18D] = xCB_8D_res
	dispatch_table[0x18E] = xCB_8E_res
	dispatch_table[0x18F] = xCB_8F_res
	dispatch_table[0x190] = xCB_90_res
	dispatch_table[0x191] = xCB_91_res
	dispatch_table[0x192] = xCB_92_res
	dispatch_table[0x193] = xCB_93_res
	dispatch_table[0x194] = xCB_94_res
	dispatch_table[0x195] = xCB_95_res
	dispatch_table[0x196] = xCB_96_res
	dispatch_table[0x197] = xCB_97_res
	dispatch_table[0x198] = xCB_98_res
	dispatch_table[0x199] = xCB_99_res
	dispatch_table[0x19A] = xCB_9A_res
	dispatch_table[0x19B] = xCB_9B_res
	dispatch_table[0x19C] = xCB_9C_res
	dispatch_table[0x19D] = xCB_9D_res
	dispatch_table[0x19E] = xCB_9E_res
	dispatch_table[0x19F] = xCB_9F_res
	dispatch_table[0x1A0] = xCB_A0_res
	dispatch_table[0x1A1] = xCB_A1_res
	dispatch_table[0x1A2] = xCB_A2_res
	dispatch_table[0x1A3] = xCB_A3_res
	dispatch_table[0x1A4] = xCB_A4_res
	dispatch_table[0x1A5] = xCB_A5_res
	dispatch_table[0x1A6] = xCB_A6_res
	dispatch_table[0x1A7] = xCB_A7_res
	dispatch_table[0x1A8] = xCB_A8_res
	dispatch_table[0x1A9] = xCB_A9_res
	dispatch_table[0x1AA] = xCB_AA_res
	dispatch_table[0x1AB] = xCB_AB_res
	dispatch_table[0x1AC] = xCB_AC_res
	dispatch_table[0x1AD] = xCB_AD_res
	dispatch_table[0x1AE] = xCB_AE_res
	dispatch_table[0x1AF] = xCB_AF_res
	dispatch_table[0x1B0] = xCB_B0_res
	dispatch_table[0x1B1] = xCB_B1_res
	dispatch_table[0x1B2] = xCB_B2_res
	dispatch_table[0x1B3] = xCB_B3_res
	dispatch_table[0x1B4] = xCB_B4_res
	dispatch_table[0x1B5] = xCB_B5_res
	dispatch_table[0x1B6] = xCB_B6_res
	dispatch_table[0x1B7] = xCB_B7_res
	dispatch_table[0x1B8] = xCB_B8_res
	dispatch_table[0x1B9] = xCB_B9_res
	dispatch_table[0x1BA] = xCB_BA_res
	dispatch_table[0x1BB] = xCB_BB_res
	dispatch_table[0x1BC] = xCB_BC_res
	dispatch_table[0x1BD] = xCB_BD_res
	dispatch_table[0x1BE] = xCB_BE_res
	dispatch_table[0x1BF] = xCB_BF_res
	dispatch_table[0x1C0] = xCB_C0_set
	dispatch_table[0x1C1] = xCB_C1_set
	dispatch_table[0x1C2] = xCB_C2_set
	dispatch_table[0x1C3] = xCB_C3_set
	dispatch_table[0x1C4] = xCB_C4_set
	dispatch_table[0x1C5] = xCB_C5_set
	dispatch_table[0x1C6] = xCB_C6_set
	dispatch_table[0x1C7] = xCB_C7_set
	dispatch_table[0x1C8] = xCB_C8_set
	dispatch_table[0x1C9] = xCB_C9_set
	dispatch_table[0x1CA] = xCB_CA_set
	dispatch_table[0x1CB] = xCB_CB_set
	dispatch_table[0x1CC] = xCB_CC_set
	dispatch_table[0x1CD] = xCB_CD_set
	dispatch_table[0x1CE] = xCB_CE_set
	dispatch_table[0x1CF] = xCB_CF_set
	dispatch_table[0x1D0] = xCB_D0_set
	dispatch_table[0x1D1] = xCB_D1_set
	dispatch_table[0x1D2] = xCB_D2_set
	dispatch_table[0x1D3] = xCB_D3_set
	dispatch_table[0x1D4] = xCB_D4_set
	dispatch_table[0x1D5] = xCB_D5_set
	dispatch_table[0x1D6] = xCB_D6_set
	dispatch_table[0x1D7] = xCB_D7_set
	dispatch_table[0x1D8] = xCB_D8_set
	dispatch_table[0x1D9] = xCB_D9_set
	dispatch_table[0x1DA] = xCB_DA_set
	dispatch_table[0x1DB] = xCB_DB_set
	dispatch_table[0x1DC] = xCB_DC_set
	dispatch_table[0x1DD] = xCB_DD_set
	dispatch_table[0x1DE] = xCB_DE_set
	dispatch_table[0x1DF] = xCB_DF_set
	dispatch_table[0x1E0] = xCB_E0_set
	dispatch_table[0x1E1] = xCB_E1_set
	dispatch_table[0x1E2] = xCB_E2_set
	dispatch_table[0x1E3] = xCB_E3_set
	dispatch_table[0x1E4] = xCB_E4_set
	dispatch_table[0x1E5] = xCB_E5_set
	dispatch_table[0x1E6] = xCB_E6_set
	dispatch_table[0x1E7] = xCB_E7_set
	dispatch_table[0x1E8] = xCB_E8_set
	dispatch_table[0x1E9] = xCB_E9_set
	dispatch_table[0x1EA] = xCB_EA_set
	dispatch_table[0x1EB] = xCB_EB_set
	dispatch_table[0x1EC] = xCB_EC_set
	dispatch_table[0x1ED] = xCB_ED_set
	dispatch_table[0x1EE] = xCB_EE_set
	dispatch_table[0x1EF] = xCB_EF_set
	dispatch_table[0x1F0] = xCB_F0_set
	dispatch_table[0x1F1] = xCB_F1_set
	dispatch_table[0x1F2] = xCB_F2_set
	dispatch_table[0x1F3] = xCB_F3_set
	dispatch_table[0x1F4] = xCB_F4_set
	dispatch_table[0x1F5] = xCB_F5_set
	dispatch_table[0x1F6] = xCB_F6_set
	dispatch_table[0x1F7] = xCB_F7_set
	dispatch_table[0x1F8] = xCB_F8_set
	dispatch_table[0x1F9] = xCB_F9_set
	dispatch_table[0x1FA] = xCB_FA_set
	dispatch_table[0x1FB] = xCB_FB_set
	dispatch_table[0x1FC] = xCB_FC_set
	dispatch_table[0x1FD] = xCB_FD_set
	dispatch_table[0x1FE] = xCB_FE_set
	dispatch_table[0x1FF] = xCB_FF_set
}
