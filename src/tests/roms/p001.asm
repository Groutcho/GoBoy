[0x0100] 3E FF 		ld A, FF 	; put FF in A
[0x0102] 06 10		ld B, 10 	; put 16D in B
[0x0104] 21 00 88	ld HL, 8800 ; put 0x8800 in HL
[0x0107] 22			ldi			; ld [HL], A, inc HL
[0x0108] 05			dec B		; decrement B
[0x0109] C2 07 01	jp !Z 0107	; if Z not zero, jump to 7000
[0x010C] 3E 0		ld A, 0		; put 0h in A
[0x010E] 06 10		ld B, 10	; put 16D in B
[0x0110] 22			ldi			; ld [HL], A, inc HL
[0x0111] 05			dec B		; decrement B
[0x0112] C2 10 01	jp !Z 0107	; if Z not zero, jump to 7000
[0x0115] 3E 01		ld A, 01	; put 01 in A
[0x0117] EA	00 98	ld (9800), A; put A at (9800)
[0x011A] EA 02 98   ld (9802), A; put A at (9802)
[0x011D] C3 1D 01	jp 011D		; jump to 11A
