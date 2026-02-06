package emulator

// Rlca Rotate the contents of register A to the left. Bit 7 copied to C, and to bit 0
func Rlca(dmg *DMG) {
	dmg.Gbz80.SetR8Register(R8_C, dmg.Gbz80.A()&0x80>>7)
	dmg.Gbz80.SetR8Register(R8_A, dmg.Gbz80.A()<<1|dmg.Gbz80.C())
}

// Rrca Rotate the contents of register A to the right. Bit 1 copied to C, and to bit 7
func Rrca(dmg *DMG) {
	dmg.Gbz80.SetR8Register(R8_C, dmg.Gbz80.A()&0x01)
	dmg.Gbz80.SetR8Register(R8_A, dmg.Gbz80.A()>>1|dmg.Gbz80.C()<<7)
}
