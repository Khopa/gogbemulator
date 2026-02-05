package emulator

// Increment/Decrement instructions for the GBZ80 CPU

// IncR8 increment a R8 register
func IncR8(dmg *DMG, r8 R8Register) {
	dmg.gbz80.IncrementR8Register(r8)
}

// DecR8 decrement a R8 register
func DecR8(dmg *DMG, r8 R8Register) {
	dmg.gbz80.DecrementR8Register(r8)
}

// IncR16 increment a R16 register
func IncR16(dmg *DMG, r16 R16Register) {
	dmg.gbz80.IncrementR16Register(r16)
}

// DecR16 decrement a R16 register
func DecR16(dmg *DMG, r16 R16Register) {
	dmg.gbz80.DecrementR16Register(r16)
}
