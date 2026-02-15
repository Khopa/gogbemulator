package emulator

// Miscellaneous instructions for the GBZ80 CPU

// NOP No operation, does nothing
func NOP(dmg *DMG) {
	// do nothing !
}

// Halt Enter low power mode
func Halt(dmg *DMG) {
	if dmg.Gbz80.Ime {
		dmg.Gbz80.Halt()
	} else {

	}
}

// Stop Enter CPU very low power mode.
// (Also used to switch between GBC double speed and normal speed CPU modes.)
func Stop(dmg *DMG) {
	dmg.Gbz80.Stop()
}
