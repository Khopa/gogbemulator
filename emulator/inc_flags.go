package emulator

// Flags related instructions for the GBZ80 CPU

// Daa Decimal Adjust Accumulator
// Designed to be used after performing an arithmetic instruction (ADD, ADC, SUB, SBC) whose inputs were in Binary-Coded Decimal (BCD),
// adjusting the result to likewise be in BCD.
func Daa(dmg *DMG) {
	var adjust int16
	var result int16
	if dmg.Gbz80.SubtractionFlag() {
		if dmg.Gbz80.HalfCarryFlag() {
			adjust += 0x6
		}
		if dmg.Gbz80.CarryFlag() {
			adjust += 0x60
		}
		result = int16(dmg.Gbz80.A()) - adjust
	} else {
		if dmg.Gbz80.HalfCarryFlag() || dmg.Gbz80.A()&0xF > 0x9 {
			adjust += 0x6
		}
		if dmg.Gbz80.CarryFlag() || dmg.Gbz80.A() > 0x99 {
			adjust += 0x60
		}
		result = int16(dmg.Gbz80.A()) + adjust
	}

	// Update accumulator
	dmg.Gbz80.SetR8Register(R8_A, uint8(result&0xFF))

	// Set flags
	if result == 0 {
		dmg.Gbz80.setFlag(FLAG_Z, true)
	}
	dmg.Gbz80.setFlag(FLAG_H, false)
	dmg.Gbz80.setFlag(FLAG_C, result > 0xFF || result < 0)

}
