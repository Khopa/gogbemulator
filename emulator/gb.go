package emulator

import "fmt"

// MakeDMG Create a new instance of the DMG (Game Boy)
func MakeDMG() *DMG {
	var mem [MemorySize]uint8
	return &DMG{
		gbz80:  MakeGbz80(),
		memory: mem,
	}
}

// PrintMemory prints the Full Memory
func (dmg *DMG) PrintMemory() {
	for i := 0; i < MemorySize/64; i++ {
		fmt.Printf("#%04x : ", i*64)
		for j := 0; j < 64; j++ {
			fmt.Printf("%x", dmg.memory[i*64+j])
		}
		fmt.Print("\n")
	}
}

// Print the DMG state
func (dmg *DMG) Print() {
	fmt.Println("DMG Model :")
	fmt.Println("-----------")
	fmt.Println("GB Z80 CPU Registers :")
	dmg.gbz80.Print()
	dmg.PrintMemory()
}
