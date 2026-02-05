package emulator

// Load instructions for the GBZ80 CPU
// See : https://rgbds.gbdev.io/docs/v0.9.4/gbz80.7

// LDr8r8 Load the value from src registers into dst register
func LDr8r8(dmg *DMG, dst R8Register, src R8Register) {
	dmg.gbz80.SetR8Register(dst, dmg.gbz80.GetR8Register(src))
}

// LDr8n8 Copy immediate 8-bit value into dst register
func LDr8n8(dmg *DMG, dst R8Register, value uint8) {
	dmg.gbz80.SetR8Register(dst, value)
}

// LDr16n16 Load the value from src registers into dst register
func LDr16n16(dmg *DMG, dst R16Register, value uint16) {
	dmg.gbz80.SetR16Register(dst, value)
}

// LDHLr8 Copy the value in register r8 into the byte pointed to by HL.
func LDHLr8(dmg *DMG, src R8Register) {
	dmg.SetMemoryU8(dmg.gbz80.HL(), dmg.gbz80.GetR8Register(src))
}

// LDHLn8 Copy the value n8 into the byte pointed to by HL.
func LDHLn8(dmg *DMG, n8 uint8) {
	dmg.SetMemoryU8(dmg.gbz80.HL(), uint8(n8))
}

// LDr8HL Copy the value pointed to by HL into register r8.
func LDr8HL(dmg *DMG, dst R8Register) {
	dmg.gbz80.SetR8Register(dst, dmg.GetMemoryU8(dmg.gbz80.HL()))
}

// LDr16A Copy the value in register A into the byte pointed to by r16.
func LDr16A(dmg *DMG, r16 uint16) {
	dmg.SetMemoryU8(r16, dmg.gbz80.A())
}

// LDHn16A Copy the value in register A into the byte at address n16.
func LDHn16A(dmg *DMG, n16 uint16) {
	dmg.SetMemoryU8(n16, dmg.gbz80.A())
}

// LDHCA Copy the value in register A into the byte at address $FF00+C.
func LDHCA(dmg *DMG) {
	dmg.SetMemoryU8(0xFF00+uint16(dmg.gbz80.C()), dmg.gbz80.A())
}
