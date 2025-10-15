package emulator

import "fmt"

/**
 * Create a new instance of the DMG (Game Boy)
 */
func MakeDMG() *DMG {
	var mem [MemorySize]uint8
	return &DMG{
		gbz80:  MakeGbz80(),
		memory: mem,
	}
}

/**
 * Print the DMG state
 */
func (dmg *DMG) Print() {
	fmt.Println("DMG Model :")
	fmt.Println("-----------")
	fmt.Println("GB Z80 CPU Registers :")
	dmg.gbz80.Print()
}
