package emulator

// Stack related instructions
// See : https://rgbds.gbdev.io/docs/v0.9.4/gbz80.7

// Pushr16 Push register r16 into the stack.
func Pushr16(dmg *DMG, r16 R16Register) {
	value := dmg.Gbz80.GetR16Register(r16)
	DecR16(dmg, R16_SP)
	dmg.SetMemoryU8(dmg.Gbz80.SP(), uint8(value&0xff00>>8))
	DecR16(dmg, R16_SP)
	dmg.SetMemoryU8(dmg.Gbz80.SP(), uint8(value&0x00ff))
}

// PushAf Push register AF into the stack.
func PushAf(dmg *DMG) {
	DecR16(dmg, R16_SP)
	dmg.SetMemoryU8(dmg.Gbz80.SP(), dmg.Gbz80.GetR8Register(R8_A))
	DecR16(dmg, R16_SP)
	dmg.SetMemoryU8(dmg.Gbz80.SP(), dmg.Gbz80.GetR8Register(R8_F))
}

// Popr16 Pop register from the stack.
func Popr16(dmg *DMG, r16 R16Register) {
	var val uint16
	val = uint16(dmg.GetMemoryU8(dmg.Gbz80.SP()))
	IncR16(dmg, R16_SP)
	val = val<<8 + uint16(dmg.GetMemoryU8(dmg.Gbz80.SP()))
	IncR16(dmg, R16_SP)
	dmg.Gbz80.SetR16Register(r16, val)
}

// PopAf Pop register AF from the stack.
func PopAf(dmg *DMG) {
	dmg.Gbz80.SetR8Register(R8_F, dmg.GetMemoryU8(dmg.Gbz80.SP()))
	IncR16(dmg, R16_SP)
	dmg.Gbz80.SetR8Register(R8_A, dmg.GetMemoryU8(dmg.Gbz80.SP()))
	IncR16(dmg, R16_SP)
}

// PopAf Pop register AF from the stack.
func Calln16(dmg *DMG, n16 uint16) {
	Pushr16(dmg, R16_PC)
	// JP 16
}

// Ret Return from subroutine. This is basically a POP PC (if such an instruction existed).
// See POP r16 for an explanation of how POP works.
func Ret(dmg *DMG) {
	Popr16(dmg, R16_PC)
}

// RetCc Return from subroutine if condition cc is met.
func RetCc(dmg *DMG, CC uint8) {
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
		Popr16(dmg, R16_PC)
	}
}
