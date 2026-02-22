package emulator

// JRd Relative Jump to address d.
// The target address is encoded as a signed 8-bit offset from the address immediately following the JR instruction
// so it must be between -128 and 127 bytes away.
func JRd(dmg *DMG, d int8) {
	tmp := int32(d) + int32(dmg.Gbz80.PC())
	if tmp < 0 || tmp > 65535 {
		// Should not happen, gb compiler shouldn't allow this !
		panic("overflow")
	}
	dmg.Gbz80.SetR16Register(R16_PC, uint16(tmp))
}

// JRCCd Relative Jump to address d if condition "CC" is met.
// The target address n16 is encoded as a signed 8-bit offset from the address immediately following the JR instruction
// so it must be between -128 and 127 bytes away.
func JRCCd(dmg *DMG, CC uint8, d int8) {
	conditionMet := false
	switch CC {
	case CC_NZ:
		conditionMet = !dmg.Gbz80.ZeroFlag()
	case CC_Z:
		conditionMet = dmg.Gbz80.ZeroFlag()
	case CC_NC:
		conditionMet = !dmg.Gbz80.CarryFlag()
	case CC_C:
		conditionMet = dmg.Gbz80.CarryFlag()
	}
	if conditionMet {
		JRd(dmg, d)
	}
}

// Jpn16 Jump to address n16; effectively, copy n16 into PC.
func Jpn16(dmg *DMG, n16 uint16) {
	dmg.Gbz80.SetR16Register(R16_PC, n16)
}

// Jpccn16 Jump to address n16 if condition cc is met.
func Jpccn16(dmg *DMG, CC uint8, n16 uint16) {
	conditionMet := false
	switch CC {
	case CC_NZ:
		conditionMet = !dmg.Gbz80.ZeroFlag()
	case CC_Z:
		conditionMet = dmg.Gbz80.ZeroFlag()
	case CC_NC:
		conditionMet = !dmg.Gbz80.CarryFlag()
	case CC_C:
		conditionMet = dmg.Gbz80.CarryFlag()
	}
	if conditionMet {
		dmg.Gbz80.SetR16Register(R16_PC, n16)
	}
}

// Jphl Jump to address in HL; effectively, copy the value in register HL into PC.
func Jphl(dmg *DMG) {
	dmg.Gbz80.SetR16Register(R16_PC, dmg.Gbz80.GetR16Register(R16_HL))
}
