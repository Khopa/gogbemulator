package emulator

import (
	"fmt"
)

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

// PrintMemorySection prints a part of the memory
func (dmg *DMG) PrintMemorySection(start int, end int) {
	if start < end {
		for i := start; i < end; i++ {
			if i%64 == 0 {
				fmt.Printf("\n#%04x : ", i)
			}
			fmt.Printf("%x", dmg.memory[i])
		}
	}
}

// Print the DMG state
func (dmg *DMG) Print() {
	fmt.Println("DMG Model :")
	fmt.Println("-----------")
	fmt.Println("GB Z80 CPU Registers :")
	dmg.gbz80.Print()

	// --------------------------------------
	dmg.PrintMemorySection(VRAMStart, VRAMEnd)

	// --------------------------------------

	// --------------------------------------

}

func (dmg *DMG) Run() {

	dmg.gbz80.pc = ProgramStart

	i := 0
	for {
		i++

		if i == 1000 {
			break
		}
	}

}

func (dmg *DMG) ExecuteCurrentInstruction() {

	// Implementation based on opcode decoding methods recommended at :
	// https://archive.gbdev.io/salvage/decoding_gbz80_opcodes/Decoding%20Gamboy%20Z80%20Opcodes.html

	// First check for prefix & read opcode
	isCBPrefixed := false
	opcode := dmg.memory[dmg.gbz80.pc]
	if opcode == 0xCB {
		isCBPrefixed = true
		opcode = dmg.memory[dmg.gbz80.pc+1]
		dmg.gbz80.pc += 1
	}
	dmg.gbz80.pc += 1

	// Compute cpu matrix path values
	x := opcode & 0xC0
	y := opcode & 0x38
	z := opcode & 0x07
	p := opcode & 0x30
	q := opcode & 0x08

	if !isCBPrefixed {

		switch x {
		case 0:
			switch z {
			case 0x0:
				// Jump ops
				switch y {
				case 0:
					NOP(dmg)
				case 1:
					// TODO: LD (nn), SP
				case 2:
					// TODO: STOP
				case 3:
					// TODO: JR d
				default:
					// TODO: JR cc[y-4], d
				}
			case 0x1:
				// 16 bit load ops
				if q == 0 {
					var nn uint16
					nn = uint16(dmg.memory[dmg.gbz80.pc])<<8 + uint16(dmg.memory[dmg.gbz80.pc+1])
					LDr16n16(dmg, RP[p], nn)
					dmg.gbz80.pc += 2
				} else {
					// TODO: Add HL, rp[p]
				}
			case 0x2:
				// Indirect Load ops
				if q == 0 {
					switch p {
					case 0:
						LDr16A(dmg, dmg.gbz80.bc)
					case 1:
						LDr16A(dmg, dmg.gbz80.de)
					case 2:
						LDr16A(dmg, dmg.gbz80.hl)
						dmg.gbz80.hl++
					case 3:
						LDr16A(dmg, dmg.gbz80.hl)
						dmg.gbz80.hl--
					}
				} else {
					switch p {
					case 0:
						LDAr16(dmg, dmg.gbz80.bc)
					case 1:
						LDAr16(dmg, dmg.gbz80.de)
					case 2:
						LDAr16(dmg, dmg.gbz80.hl)
						dmg.gbz80.hl++
					case 3:
						LDAr16(dmg, dmg.gbz80.hl)
						dmg.gbz80.hl--
					}
				}
			case 0x3:
				// 16 Bit Inc Dec
				if q == 0 {
					IncR16(dmg, RP[p])
				} else {
					DecR16(dmg, RP[p])
				}
			case 0x4:
				IncR8(dmg, R[y])
			case 0x5:
				DecR8(dmg, R[y])
			case 0x6:
				n := dmg.memory[dmg.gbz80.pc]
				LDr8n8(dmg, R[y], n)
				dmg.gbz80.pc++
			case 0x7:
				// Assorted ops on accumulator flags
				switch y {

				}
			}
		case 1:
			if z == 6 && y == 6 {
				// HALT
			} else {
				LDr8r8(dmg, R[y], R[z])
			}
		case 2:
			// ALU operations

		case 3:
			// 16 Bit Inc Dec
		case 4:
			// 16 Bit Inc Dec
		}

	} else {
		switch x {
		case 0x0:
			// rotation/Shift instructions

		case 0x1:
			// Bit test instruction

		case 0x2:
			// Reset bit
			dmg.gbz80.ResetBitInR8Register(R[z], y)
		case 0x3:
			// Set Bit instructions

		}

	}

}
