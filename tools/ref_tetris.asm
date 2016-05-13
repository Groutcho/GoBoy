0x00           NOP
0xC3 0x50 0x01 JP $0150
0xC3 0x8B 0x02 JP $028B
0xAF           XOR A
0x21 0xFF 0xDF LD HL,$DFFF
0x0E 0x10      LD C,$10
0x06 0x00      LD B,$00
0x32           LD [HLD],A
0x05           DEC B
0x20 0xFC      JR NZ,$FC ; 0x293
0x0D           DEC C
0x20 0xF9      JR NZ,$F9 ; 0x293
0x3E 0x01      LD A,$01
0xF3           DI
0xE0 0x0F      LDH [$0F],A ; IF
0xE0 0xFF      LDH [$FF],A ; IE
0xE0 0x42      LDH [$42],A ; SCY
0xE0 0x43      LDH [$43],A ; SCX
0xE0 0xA4      LDH [$A4],A ; HIMEM
0xE0 0x41      LDH [$41],A ; STAT
0xE0 0x01      LDH [$01],A ; SB
0xE0 0x02      LDH [$02],A ; SC
0x3E 0x80      LD A,$80
0xE0 0x40      LDH [$40],A ; LCDC
0xF0 0x44      LDH A,[$44] ; LY
0xFE 0x94      CP $94
0x20 0xFA      JR NZ,$FA ; 0x2B2
0x3E 0x03      LD A,$03
0x3E 0xE4      LD A,$E4
0xE0 0x47      LDH [$47],A ; BGP
0xE0 0x48      LDH [$48],A ; OBP0
0x3E 0xC4      LD A,$C4
0xE0 0x49      LDH [$49],A ; OBP1
0x21 0x26 0xFF LD HL,$FF26
0x3E 0xFF      LD A,$FF
0x36 0x77      LD [HL],$77
0xEA 0x00 0x20 LD [$2000],A
0x31 0xFF 0xCF LD SP,$CFFF
0x20 0xFC      JR NZ,$FC ; 0x2DF
0x21 0xFF 0xCF LD HL,$CFFF
0x20 0xFC      JR NZ,$FC ; 0x2EA
0x20 0xF9      JR NZ,$F9 ; 0x2EA
0x21 0xFF 0x9F LD HL,$9FFF
0x0E 0x20      LD C,$20
0x20 0xFC      JR NZ,$FC ; 0x2F9
0x20 0xF9      JR NZ,$F9 ; 0x2F9
0x21 0xFF 0xFE LD HL,$FEFF
0x20 0xFC      JR NZ,$FC ; 0x305
0x21 0xFE 0xFF LD HL,$FFFE
0x06 0x80      LD B,$80
0x20 0xFC      JR NZ,$FC ; 0x30E
0x0E 0xB6      LD C,$B6
0x06 0x0C      LD B,$0C
0x21 0xC7 0x2A LD HL,$2AC7
0x2A           LD A,[HLI]
0xE2           LD [C],A
0x0C           INC C
0x20 0xFA      JR NZ,$FA ; 0x319
0xCD 0xE9 0x27 CALL $27E9
0xCD 0xF3 0x7F CALL $7FF3
0x3E 0x09      LD A,$09
0x3E 0x37      LD A,$37
0xE0 0xC0      LDH [$C0],A ; HIMEM
0x3E 0x1C      LD A,$1C
0xE0 0xC1      LDH [$C1],A ; HIMEM
0x3E 0x24      LD A,$24
0xE0 0xE1      LDH [$E1],A ; HIMEM
0xFB           EI
0xE0 0x4A      LDH [$4A],A ; WY
0xE0 0x4B      LDH [$4B],A ; WX
0xE0 0x06      LDH [$06],A ; TMA
0xCD 0xFA 0x29 CALL $29FA
0xCD 0x77 0x03 CALL $0377
0xCD 0xF0 0x7F CALL $7FF0
0xF0 0x80      LDH A,[$80] ; HIMEM
0xE6 0x0F      AND $0F
0xFE 0x0F      CP $0F
0xCA 0x9A 0x02 JP Z,$029A
0x21 0xA6 0xFF LD HL,$FFA6
0x06 0x02      LD B,$02
0x7E           LD A,[HL]
0xA7           AND A
0x28 0x01      JR Z,$01 ; 0x35F
0x35           DEC [HL]
0x2C           INC L
0x20 0xF7      JR NZ,$F7 ; 0x35A
0xF0 0xC5      LDH A,[$C5] ; HIMEM
0x28 0x04      JR Z,$04 ; 0x36C
0xF0 0x85      LDH A,[$85] ; HIMEM
0x28 0xFB      JR Z,$FB ; 0x36C
0xE0 0x85      LDH [$85],A ; HIMEM
0xC3 0x43 0x03 JP $0343
0xF0 0xE1      LDH A,[$E1] ; HIMEM
0xEF           RST $28
0x29           ADD HL,HL
0x1C           INC E
0x3D           DEC A
0x1D           DEC E
0xA8           XOR B
0x12           LD [DE],A
0xDF           RST $18
0x61           LD H,C
0x81           ADD A,C
0x19           ADD HL,DE
0x04           INC B
0xE6 0x04      AND $04
0x14           INC D
0xF0 0x14      LDH A,[$14] ; NR14
0x6B           LD L,E
0x1A           LD A,[DE]
0x1B           DEC DE
0x1E 0x71      LD E,$71
0x1F           RRA
0x7A           LD A,D
0x89           ADC A,C
0x15           DEC D
0x23           INC HL
0x16 0x8D      LD D,$8D
0x16 0xDE      LD D,$DE
0x16 0x4F      LD D,$4F
0x17           RLA
0x77           LD [HL],A
0xE4           -
0x21 0xFF 0x9B LD HL,$9BFF
0x01 0x00 0x04 LD BC,$0400
0x3E 0x2F      LD A,$2F
0x0B           DEC BC
0x78           LD A,B
0xB1           OR C
0x20 0xF8      JR NZ,$F8 ; 0x27EF
0xC9           RET
0x3E 0x20      LD A,$20
0xE0 0x00      LDH [$00],A ; P1/JOYP
0xF0 0x00      LDH A,[$00] ; P1/JOYP
0x2F           CPL
0xCB 0x37      SWAP A
0x47           LD B,A
0x3E 0x10      LD A,$10
0xB0           OR B
0x4F           LD C,A
0xA9           XOR C
0xA1           AND C
0xE0 0x81      LDH [$81],A ; HIMEM
0x79           LD A,C
0xE0 0x80      LDH [$80],A ; HIMEM
0x3E 0x30      LD A,$30
0xF5           PUSH AF
0xC5           PUSH BC
0xD5           PUSH DE
0xE5           PUSH HL
0xFA 0x7F 0xDF LD A,[$DF7F]
0xFE 0x01      CP $01
0x28 0x46      JR Z,$46 ; 0x65A4
0xFE 0x02      CP $02
0x28 0x7B      JR Z,$7B ; 0x65DD
0xFA 0x7E 0xDF LD A,[$DF7E]
0x20 0x7B      JR NZ,$7B ; 0x65E3
0xF0 0xE4      LDH A,[$E4] ; HIMEM
0x28 0x0D      JR Z,$0D ; 0x657A
0xEA 0xE0 0xDF LD [$DFE0],A
0xEA 0xE8 0xDF LD [$DFE8],A
0xEA 0xF0 0xDF LD [$DFF0],A
0xEA 0xF8 0xDF LD [$DFF8],A
0xCD 0x52 0x65 CALL $6552
0xCD 0x0E 0x6A CALL $6A0E
0xCD 0x2E 0x6A CALL $6A2E
0xCD 0x79 0x68 CALL $6879
0xCD 0x52 0x6A CALL $6A52
0xCD 0x75 0x6C CALL $6C75
0xCD 0x96 0x6A CALL $6A96
0xEA 0x7F 0xDF LD [$DF7F],A
0xE1           POP HL
0xD1           POP DE
0xC1           POP BC
0xF1           POP AF
0xCD 0xF8 0x69 CALL $69F8
0xEA 0xE1 0xDF LD [$DFE1],A
0xEA 0xF1 0xDF LD [$DFF1],A
0xEA 0xF9 0xDF LD [$DFF9],A
0x21 0xBF 0xDF LD HL,$DFBF
0xCB 0xBE      RES 7,[HL]
0x21 0x9F 0xDF LD HL,$DF9F
0x21 0xAF 0xDF LD HL,$DFAF
0x21 0xCF 0xDF LD HL,$DFCF
0x21 0x1A 0x6F LD HL,$6F1A
0xCD 0xC9 0x69 CALL $69C9
0xEA 0x7E 0xDF LD [$DF7E],A
0x21 0xFB 0x65 LD HL,$65FB
0xCD 0x8E 0x69 CALL $698E
0x18 0xB7      JR $B7 ; 0x658F
0x21 0xFF 0x65 LD HL,$65FF
0x18 0xF6      JR $F6 ; 0x65D3
0x18 0x85      JR $85 ; 0x6568
0x21 0x7E 0xDF LD HL,$DF7E
0xFE 0x28      CP $28
0x28 0xEC      JR Z,$EC ; 0x65D8
0xFE 0x20      CP $20
0x28 0xE0      JR Z,$E0 ; 0x65D0
0xFE 0x18      CP $18
0x28 0xE4      JR Z,$E4 ; 0x65D8
0xFE 0x10      CP $10
0x20 0x97      JR NZ,$97 ; 0x658F
0x34           INC [HL]
0x18 0x94      JR $94 ; 0x658F
0x21 0x0A 0x6F LD HL,$6F0A
0xCD 0x3E 0x69 CALL $693E
0xF0 0x04      LDH A,[$04] ; DIV
0xE6 0x1F      AND $1F
0x3E 0xD0      LD A,$D0
0x80           ADD A,B
0xEA 0xF5 0xDF LD [$DFF5],A
0x21 0x38 0x68 LD HL,$6838
0xC3 0x95 0x69 JP $6995
0x21 0xF4 0xDF LD HL,$DFF4
0x21 0xF5 0xDF LD HL,$DFF5
0xFE 0x0E      CP $0E
0x30 0x0A      JR NC,$0A ; 0x686F
0xE6 0xF0      AND $F0
0x0E 0x1D      LD C,$1D
0xFE 0x1E      CP $1E
0xCA 0x1F 0x69 JP Z,$691F
0x18 0xEE      JR $EE ; 0x6867
0xFA 0xF0 0xDF LD A,[$DFF0]
0xCA 0xA8 0x68 JP Z,$68A8
0xCA 0x3D 0x68 JP Z,$683D
0xFA 0xF1 0xDF LD A,[$DFF1]
0xCA 0xF3 0x68 JP Z,$68F3
0xCA 0x54 0x68 JP Z,$6854
0x21 0xDA 0x6E LD HL,$6EDA
0x21 0x97 0x68 LD HL,$6897
0xEA 0xF6 0xDF LD [$DFF6],A
0x21 0x94 0x68 LD HL,$6894
0x3E 0x00      LD A,$00
0x21 0x9C 0x68 LD HL,$689C
0x21 0x99 0x68 LD HL,$6899
0x18 0xEC      JR $EC ; 0x68BD
0x21 0xA1 0x68 LD HL,$68A1
0x21 0x9E 0x68 LD HL,$689E
0x18 0xDB      JR $DB ; 0x68BD
0x3E 0x02      LD A,$02
0x21 0xA6 0x68 LD HL,$68A6
0x21 0xA3 0x68 LD HL,$68A3
0x18 0xCA      JR $CA ; 0x68BD
0xFE 0x09      CP $09
0x28 0xC4      JR Z,$C4 ; 0x68C0
0xFE 0x13      CP $13
0x28 0xD1      JR Z,$D1 ; 0x68D1
0xFE 0x17      CP $17
0x28 0xDE      JR Z,$DE ; 0x68E2
0x28 0x17      JR Z,$17 ; 0x691F
0xFE 0x00      CP $00
0xC8           RET Z
0x28 0x05      JR Z,$05 ; 0x6915
0x28 0x05      JR Z,$05 ; 0x6919
0x18 0x02      JR $02 ; 0x691B
0xE0 0x1D      LDH [$1D],A ; NR33
0xE0 0x1A      LDH [$1A],A ; NR30
0x18 0x25      JR $25 ; 0x6963
0xCB 0xFE      SET 7,[HL]
0xEA 0xF4 0xDF LD [$DFF4],A
0x0E 0x16      LD C,$16
0x06 0x04      LD B,$04
0x18 0x0C      JR $0C ; 0x69A1
0x0E 0x1A      LD C,$1A
0x06 0x05      LD B,$05
0x18 0x05      JR $05 ; 0x69A1
0x20 0xFA      JR NZ,$FA ; 0x69A1
0xEA 0x71 0xDF LD [$DF71],A
0xCB 0x27      SLA A
0x09           ADD HL,BC
0x4E           LD C,[HL]
0x46           LD B,[HL]
0x69           LD L,C
0x60           LD H,B
0x7C           LD A,H
0x0E 0x30      LD C,$30
0xFE 0x40      CP $40
0x20 0xF8      JR NZ,$F8 ; 0x69CC
0xEA 0xE9 0xDF LD [$DFE9],A
0xEA 0x9F 0xDF LD [$DF9F],A
0xEA 0xAF 0xDF LD [$DFAF],A
0xEA 0xBF 0xDF LD [$DFBF],A
0xEA 0xCF 0xDF LD [$DFCF],A
0xE0 0x25      LDH [$25],A ; NR51
0xEA 0x78 0xDF LD [$DF78],A
0x3E 0x08      LD A,$08
0xE0 0x12      LDH [$12],A ; NR12
0xE0 0x17      LDH [$17],A ; NR22
0xE0 0x21      LDH [$21],A ; NR42
0xE0 0x14      LDH [$14],A ; NR14
0xE0 0x19      LDH [$19],A ; NR24
0xE0 0x23      LDH [$23],A ; NR44
0xE0 0x10      LDH [$10],A ; NR10
0x11 0xE0 0xDF LD DE,$DFE0
0x28 0x0C      JR Z,$0C ; 0x6A21
0x21 0x00 0x65 LD HL,$6500
0xCD 0xA9 0x69 CALL $69A9
0xE9           JP [HL]
0x28 0x07      JR Z,$07 ; 0x6A2D
0x21 0x10 0x65 LD HL,$6510
0xCD 0xAD 0x69 CALL $69AD
0x11 0xF8 0xDF LD DE,$DFF8
0x28 0x0C      JR Z,$0C ; 0x6A41
0x21 0x20 0x65 LD HL,$6520
0x28 0x07      JR Z,$07 ; 0x6A4D
0x21 0x28 0x65 LD HL,$6528
0xCD 0xD6 0x69 CALL $69D6
0x21 0xE8 0xDF LD HL,$DFE8
0xFE 0xFF      CP $FF
0x28 0xF2      JR Z,$F2 ; 0x6A4E
0x21 0x30 0x65 LD HL,$6530
0xCD 0x44 0x6B CALL $6B44
0xCD 0x6D 0x6A CALL $6A6D
0xFA 0xE9 0xDF LD A,[$DFE9]
0x21 0xEF 0x6A LD HL,$6AEF
0x28 0x06      JR Z,$06 ; 0x6A7E
0x18 0xF7      JR $F7 ; 0x6A75
0xEA 0x76 0xDF LD [$DF76],A
0xEA 0x79 0xDF LD [$DF79],A
0xEA 0x7A 0xDF LD [$DF7A],A
0xEA 0x75 0xDF LD [$DF75],A
0xEA 0x77 0xDF LD [$DF77],A
0x28 0x3D      JR Z,$3D ; 0x6AD9
0x21 0x75 0xDF LD HL,$DF75
0xFA 0x78 0xDF LD A,[$DF78]
0x28 0x37      JR Z,$37 ; 0x6ADD
0xFE 0x03      CP $03
0x28 0x2F      JR Z,$2F ; 0x6AD9
0xBE           CP [HL]
0x20 0x33      JR NZ,$33 ; 0x6AE2
0x2D           DEC L
0x36 0x00      LD [HL],$00
0xFA 0x79 0xDF LD A,[$DF79]
0xCB 0x46      BIT 0,[HL]
0xCA 0xC0 0x6A JP Z,$6AC0
0xFA 0x7A 0xDF LD A,[$DF7A]
0x28 0x04      JR Z,$04 ; 0x6ACB
0xCB 0xD0      SET 2,B
0xCB 0xF0      SET 6,B
0xFA 0xF9 0xDF LD A,[$DFF9]
0x28 0x04      JR Z,$04 ; 0x6AD5
0xCB 0xD8      SET 3,B
0xCB 0xF8      SET 7,B
0x18 0xF9      JR $F9 ; 0x6AD6
0x18 0xDE      JR $DE ; 0x6AC0
0x20 0xF1      JR NZ,$F1 ; 0x6AD9
0x20 0xEB      JR NZ,$EB ; 0x6AD9
0x0A           LD A,[BC]
0x03           INC BC
0x11 0x80 0xDF LD DE,$DF80
0xCD 0x3E 0x6B CALL $6B3E
0x11 0x90 0xDF LD DE,$DF90
0x11 0xA0 0xDF LD DE,$DFA0
0x11 0xB0 0xDF LD DE,$DFB0
0x11 0xC0 0xDF LD DE,$DFC0
0x21 0x90 0xDF LD HL,$DF90
0x11 0x94 0xDF LD DE,$DF94
0xCD 0x33 0x6B CALL $6B33
0x21 0xA0 0xDF LD HL,$DFA0
0x11 0xA4 0xDF LD DE,$DFA4
0x21 0xB0 0xDF LD HL,$DFB0
0x11 0xB4 0xDF LD DE,$DFB4
0x21 0xC0 0xDF LD HL,$DFC0
0x11 0xC4 0xDF LD DE,$DFC4
0x01 0x10 0x04 LD BC,$0410
0x21 0x92 0xDF LD HL,$DF92
0x36 0x01      LD [HL],$01
0x85           ADD A,L
0x6F           LD L,A
0x20 0xF8      JR NZ,$F8 ; 0x6B9B
0xEA 0x9E 0xDF LD [$DF9E],A
0xEA 0xAE 0xDF LD [$DFAE],A
0xEA 0xBE 0xDF LD [$DFBE],A
0x62           LD H,D
0x18 0x2A      JR $2A ; 0x6BE4
0xCD 0xEA 0x6B CALL $6BEA
0xCD 0xFF 0x6B CALL $6BFF
0x5F           LD E,A
0x57           LD D,A
0x73           LD [HL],E
0x72           LD [HL],D
0x71           LD [HL],C
0x21 0x70 0xDF LD HL,$DF70
0x28 0xCA      JR Z,$CA ; 0x6BAE
0xC3 0x8F 0x6C JP $6C8F
0x3A           LD A,[HLD]
0x13           INC DE
0x7B           LD A,E
0x22           LD [HLI],A
0x18 0xF1      JR $F1 ; 0x6BF0
0x18 0x2C      JR $2C ; 0x6C35
0xFA 0x70 0xDF LD A,[$DF70]
0x20 0x10      JR NZ,$10 ; 0x6C20
0xFA 0xB8 0xDF LD A,[$DFB8]
0xCB 0x7F      BIT 7,A
0x28 0x09      JR Z,$09 ; 0x6C20
0xFE 0x06      CP $06
0x20 0x04      JR NZ,$04 ; 0x6C20
0x3E 0x40      LD A,$40
0xE0 0x1C      LDH [$1C],A ; NR32
0x7D           LD A,L
0xC6 0x09      ADD A,$09
0x20 0xDD      JR NZ,$DD ; 0x6C06
0xC6 0x04      ADD A,$04
0xCB 0x7E      BIT 7,[HL]
0x20 0xD5      JR NZ,$D5 ; 0x6C06
0xCD 0x98 0x6D CALL $6D98
0xC3 0x6A 0x6D JP $6D6A
0xCD 0xF6 0x6B CALL $6BF6
0x54           LD D,H
0x28 0x1F      JR Z,$1F ; 0x6C6C
0x28 0x04      JR Z,$04 ; 0x6C55
0xC3 0x8D 0x6C JP $6C8D
0x18 0xD5      JR $D5 ; 0x6C41
0x21 0xE9 0xDF LD HL,$DFE9
0xEA 0x70 0xDF LD [$DF70],A
0xCA 0x35 0x6C JP Z,$6C35
0xC2 0x09 0x6C JP NZ,$6C09
0xCA 0x3A 0x6C JP Z,$6C3A
0xFE 0x9D      CP $9D
0xCA 0xBA 0x6B JP Z,$6BBA
0xFE 0xA0      CP $A0
0x20 0x1A      JR NZ,$1A ; 0x6CBC
0x11 0x81 0xDF LD DE,$DF81
0x67           LD H,A
0xFE 0x04      CP $04
0xCA 0xED 0x6C JP Z,$6CED
0xC6 0x05      ADD A,$05
0x5D           LD E,L
0x28 0x0F      JR Z,$0F ; 0x6CE8
0x21 0x33 0x6E LD HL,$6E33
0xC3 0x04 0x6D JP $6D04
0x18 0x17      JR $17 ; 0x6D04
0x11 0xC6 0xDF LD DE,$DFC6
0x21 0xC5 0x6E LD HL,$6EC5
0xFE 0xCB      CP $CB
0x20 0xF8      JR NZ,$F8 ; 0x6CF5
0x21 0xC4 0xDF LD HL,$DFC4
0x18 0x2E      JR $2E ; 0x6D32
0x28 0x21      JR Z,$21 ; 0x6D2D
0x28 0x19      JR Z,$19 ; 0x6D29
0xFA 0xBF 0xDF LD A,[$DFBF]
0x20 0x05      JR NZ,$05 ; 0x6D1E
0x16 0x00      LD D,$00
0x18 0x15      JR $15 ; 0x6D3E
0x18 0x05      JR $05 ; 0x6D32
0x20 0x4F      JR NZ,$4F ; 0x6D88
0x28 0x02      JR Z,$02 ; 0x6D47
0x1E 0x01      LD E,$01
0x20 0x13      JR NZ,$13 ; 0x6D65
0xF6 0x80      OR $80
0xF6 0x05      OR $05
0xCB 0x86      RES 0,[HL]
0x11 0x70 0xDF LD DE,$DF70
0x28 0x09      JR Z,$09 ; 0x6D7B
0x3C           INC A
0x11 0x10 0x00 LD DE,$0010
0xC3 0x83 0x6C JP $6C83
0x21 0x9E 0xDF LD HL,$DF9E
0x21 0xAE 0xDF LD HL,$DFAE
0x21 0xBE 0xDF LD HL,$DFBE
0x18 0xAC      JR $AC ; 0x6D3B
0xCB 0x3F      SRL A
0x26 0x00      LD H,$00
0x5E           LD E,[HL]
0xC6 0x06      ADD A,$06
0x28 0x18      JR Z,$18 ; 0x6DBA
0x0E 0x13      LD C,$13
0x28 0x0E      JR Z,$0E ; 0x6DBC
0x0E 0x18      LD C,$18
0x28 0x08      JR Z,$08 ; 0x6DBC
0x28 0x02      JR Z,$02 ; 0x6DBC
0xFA 0x71 0xDF LD A,[$DF71]
0x18 0x09      JR $09 ; 0x6DD7
0x11 0xFC 0x6D LD DE,$6DFC
0xCD 0x8F 0x6D CALL $6D8F
0xCB 0x40      BIT 0,B
0x20 0x02      JR NZ,$02 ; 0x6DE3
0xCB 0x33      SWAP E
0xCB 0x5F      BIT 3,A
0x28 0x06      JR Z,$06 ; 0x6DF0
0x26 0xFF      LD H,$FF
0xF6 0xF0      OR $F0
0x18 0x02      JR $02 ; 0x6DF2
0x18 0xBE      JR $BE ; 0x6DBA
0xC3 0x53 0x65 JP $6553
0xC3 0xD6 0x69 JP $69D6