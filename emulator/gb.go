package emulator

import (
	"fmt"

	"khopa.github.io/gogbemulator/emulator/cpu"
)

type DMG struct {
	gbz80 *cpu.Gbz80
}

func MakeDMG() *DMG {

	// Create CPU
	var gbz80 *cpu.Gbz80
	gbz80 = cpu.MakeGbz80()

	// Create DMG Game boy model
	return &DMG{
		gbz80,
	}

}

func (dmg *DMG) Print() {
	fmt.Println("DMG Model :")
	fmt.Println("-----------\n")
	fmt.Println("GB Z80 CPU Registers :\n")
	dmg.gbz80.Print()
}
