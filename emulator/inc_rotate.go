package emulator

// Rotate & Shift instructions for the GBZ80 CPU

// Rlca Rotate the contents of register A to the left. Bit 7 copied to C Flag, and to bit 0
func Rlca(dmg *DMG) {
	dmg.Gbz80.setFlag(FLAG_C, dmg.Gbz80.A()&0x80 == 0x80)
	bit7 := dmg.Gbz80.A() & 0x80 >> 7
	dmg.Gbz80.SetR8Register(R8_A, dmg.Gbz80.A()<<1|bit7)
}

// Rrca Rotate the contents of register A to the right. Bit 1 copied to C Flag, and to bit 7
func Rrca(dmg *DMG) {
	dmg.Gbz80.setFlag(FLAG_C, dmg.Gbz80.A()&0x01 == 0x01)
	bit1 := dmg.Gbz80.A() & 0x01 << 7
	dmg.Gbz80.SetR8Register(R8_A, dmg.Gbz80.A()>>1|bit1)
}

// Rla Rotate register A left, through the C flag
func Rla(dmg *DMG) {
	carry := dmg.Gbz80.A()&0x80 == 0x80
	if dmg.Gbz80.CarryFlag() {
		dmg.Gbz80.SetR8Register(R8_A, dmg.Gbz80.A()<<1|0x01)
	} else {
		dmg.Gbz80.SetR8Register(R8_A, dmg.Gbz80.A()<<1)
	}
	dmg.Gbz80.setFlag(FLAG_C, carry)
}

// Rra Rotate register A right, through the C flag
func Rra(dmg *DMG) {
	carry := dmg.Gbz80.A()&0x01 == 0x01
	if dmg.Gbz80.CarryFlag() {
		dmg.Gbz80.SetR8Register(R8_A, dmg.Gbz80.A()>>1|0x80)
	} else {
		dmg.Gbz80.SetR8Register(R8_A, dmg.Gbz80.A()>>1)
	}
	dmg.Gbz80.setFlag(FLAG_C, carry)
}
