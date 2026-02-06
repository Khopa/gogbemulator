package emulator

// Miscellaneous instructions for the GBZ80 CPU

// NOP No operation, does nothing
func NOP(dmg *DMG) {
	// do nothing !
}

// Enter low power mode
func Halt(dmg *DMG) {
	if dmg.Gbz80.Ime {
		dmg.Gbz80.Halt()
	} else {

	}
}
