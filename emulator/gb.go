package emulator

import (
	"fmt"
	"os"
)

// MakeDMG Create a new instance of the DMG (Game Boy)
func MakeDMG() *DMG {
	var mem [MemorySize]uint8
	return &DMG{
		Gbz80:  MakeGbz80(),
		Memory: mem,
	}
}

// PrintMemory prints the Full Memory
func (dmg *DMG) PrintMemory() {
	for i := 0; i < MemorySize/64; i++ {
		fmt.Printf("#%04x : ", i*64)
		for j := 0; j < 64; j++ {
			fmt.Printf("%x", dmg.Memory[i*64+j])
		}
		fmt.Print("\n")
	}
}

// PrintMemorySection prints a part of the Memory
func (dmg *DMG) PrintMemorySection(start int, end int) {
	if start < end {
		for i := start; i < end; i++ {
			if i%64 == 0 {
				fmt.Printf("\n#%04x : ", i)
			}
			fmt.Printf("%x", dmg.Memory[i])
		}
	}
}

// Print the DMG state
func (dmg *DMG) Print() {
	fmt.Println("DMG Model :")
	fmt.Println("-----------")
	fmt.Println("GB Z80 CPU Registers :")
	dmg.Gbz80.Print()

	// --------------------------------------
	dmg.PrintMemorySection(VRAMStart, VRAMEnd)

	// --------------------------------------

	// --------------------------------------

}

func (dmg *DMG) LoadROM(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if len(data) > MemorySize {
		return fmt.Errorf("ROM too large: %d bytes", len(data))
	}

	copy(dmg.Memory[:], data)

	return nil
}

func (dmg *DMG) Step() {
	dmg.ExecuteCurrentInstruction()
}

func (dmg *DMG) ExecuteCurrentInstruction() {

	// Implementation based on opcode decoding methods recommended at :
	// https://archive.gbdev.io/salvage/decoding_gbz80_opcodes/Decoding%20Gamboy%20Z80%20Opcodes.html

	// First check for prefix & read opcode
	isCBPrefixed := false
	opcode := dmg.Memory[dmg.Gbz80.Pc]
	var d uint8
	print(d)
	if opcode == 0xCB {
		isCBPrefixed = true
		d = dmg.Memory[dmg.Gbz80.Pc+1]
		opcode = dmg.Memory[dmg.Gbz80.Pc+2]
		dmg.Gbz80.Pc += 2
	}
	dmg.Gbz80.Pc += 1

	// Compute cpu matrix path values
	x := opcode & 0b11000000 >> 6
	y := opcode & 0b00111000 >> 3
	z := opcode & 0b00000111
	p := opcode & 0b00010000 >> 4
	q := opcode & 0b00001000 >> 3

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
					var nn uint16
					nn = uint16(dmg.Memory[dmg.Gbz80.Pc])<<8 + uint16(dmg.Memory[dmg.Gbz80.Pc+1])
					dmg.Gbz80.Pc += 2
					LDn16SP(dmg, nn)
				case 2:
					Stop(dmg)
				case 3:
					db := int8(dmg.Memory[dmg.Gbz80.Pc])
					dmg.Gbz80.Pc += 1
					JRd(dmg, db)
				default:
					db := int8(dmg.Memory[dmg.Gbz80.Pc])
					cc := CC[y-4]
					dmg.Gbz80.Pc += 1
					JRCCd(dmg, cc, db)
				}
			case 0x1:
				// 16 bit load ops
				if q == 0 {
					var nn uint16
					nn = uint16(dmg.Memory[dmg.Gbz80.Pc])<<8 + uint16(dmg.Memory[dmg.Gbz80.Pc+1])
					LDr16n16(dmg, RP[p], nn)
					dmg.Gbz80.Pc += 2
				} else {
					// TODO: Add HL, rp[p]
				}
			case 0x2:
				// Indirect Load ops
				if q == 0 {
					switch p {
					case 0:
						LDr16A(dmg, dmg.Gbz80.Bc)
					case 1:
						LDr16A(dmg, dmg.Gbz80.De)
					case 2:
						LDr16A(dmg, dmg.Gbz80.Hl)
						dmg.Gbz80.Hl++
					case 3:
						LDr16A(dmg, dmg.Gbz80.Hl)
						dmg.Gbz80.Hl--
					}
				} else {
					switch p {
					case 0:
						LDAr16(dmg, dmg.Gbz80.Bc)
					case 1:
						LDAr16(dmg, dmg.Gbz80.De)
					case 2:
						LDAr16(dmg, dmg.Gbz80.Hl)
						dmg.Gbz80.Hl++
					case 3:
						LDAr16(dmg, dmg.Gbz80.Hl)
						dmg.Gbz80.Hl--
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
				n := dmg.Memory[dmg.Gbz80.Pc]
				LDr8n8(dmg, R[y], n)
				dmg.Gbz80.Pc++
			case 0x7:
				// Assorted ops on accumulator flags
				switch y {
				case 0:
					Rlca(dmg)
				case 1:
					Rrca(dmg)
				case 2:
					Rla(dmg)
				case 3:
					Rra(dmg)
				case 4:
					Daa(dmg)
				case 5:
					Cpl(dmg)
				case 6:
					Scf(dmg)
				case 7:
					Ccf(dmg)
				}
			}
		case 1:
			// 8 Bit Loading
			if z == 6 && y == 6 {
				Halt(dmg)
			} else {
				LDr8r8(dmg, R[y], R[z])
			}
		case 2:
			// ALU operations on acc & registers
			switch y {
			case 0:
			// ADD A, r[z]
			case 1:
			// ADC A, r[z]
			case 2:
			// SUB A, r[z]
			case 3:
			// SBC A, r[z]
			case 4:
			// AND A, r[z]
			case 5:
			// XOR A, r[z]
			case 6:
			// OR A, r[z]
			case 7:
				// CP A, r[z]
			}
		case 3:
			switch z {
			case 0:
				// Conditional return, mem-mapped register loads and stack operations
				if y <= 3 {
					fmt.Printf("RET %s", DIS_CC[y])
					RetCc(dmg, CC[y])
				} else if y == 4 {
					n := dmg.Memory[dmg.Gbz80.Pc]
					dmg.Gbz80.Pc += 1
					fmt.Printf("LDH %d, A", n)
					// TODO IMPLEMENT
				} else if y == 5 {
					db := int8(dmg.Memory[dmg.Gbz80.Pc])
					dmg.Gbz80.Pc += 1
					fmt.Printf("ADD SP, %d", db)
					// TODO IMPLEMENT
				} else if y == 6 {
					n := dmg.Memory[dmg.Gbz80.Pc]
					dmg.Gbz80.Pc += 1
					fmt.Printf("LDH A, %d", n)
					// TODO IMPLEMENT
				} else if y == 7 {
					db := int8(dmg.Memory[dmg.Gbz80.Pc])
					dmg.Gbz80.Pc += 1
					fmt.Printf("LD HL, SP+ %d", db)
					// TODO IMPLEMENT
				}
			case 1:
				// POP & various ops
				if q == 0 {
					fmt.Printf("POP %s", DIS_RP2[p])
					Popr16(dmg, RP2[p])
				} else {
					if p == 0 {
						fmt.Printf("RET")
						Ret(dmg)
					} else if p == 1 {
						fmt.Printf("RETI")
						// TODO IMPLEMENT
					} else if p == 2 {
						fmt.Printf("JP HL")
						Jphl(dmg)
					} else if p == 3 {
						fmt.Printf("LD SP, HL")
						// TODO IMPLEMENT
					}
				}
			case 2:
				// Conditional jumps
				if y <= 3 {
					var nn uint16
					nn = uint16(dmg.Memory[dmg.Gbz80.Pc])<<8 + uint16(dmg.Memory[dmg.Gbz80.Pc+1])
					dmg.Gbz80.Pc += 2
					fmt.Printf("JP %s, %d", DIS_CC[y], nn)
					Jpccn16(dmg, CC[y], nn)
				} else if y == 4 {
					fmt.Printf("LDH C, A")
				} else if y == 5 {
					var nn uint16
					nn = uint16(dmg.Memory[dmg.Gbz80.Pc])<<8 + uint16(dmg.Memory[dmg.Gbz80.Pc+1])
					dmg.Gbz80.Pc += 2
					fmt.Printf("LD (%d), A", nn)
					// TODO: Implement
				} else if y == 6 {
					fmt.Printf("LDH A, C")
				} else if y == 7 {
					var nn uint16
					nn = uint16(dmg.Memory[dmg.Gbz80.Pc])<<8 + uint16(dmg.Memory[dmg.Gbz80.Pc+1])
					dmg.Gbz80.Pc += 2
					fmt.Printf("LD A, (%d)", nn)
					// TODO: Implement
				}
			case 3:
				if y == 0 {
					var nn uint16
					nn = uint16(dmg.Memory[dmg.Gbz80.Pc])<<8 + uint16(dmg.Memory[dmg.Gbz80.Pc+1])
					dmg.Gbz80.Pc += 2
					fmt.Printf("JP %d", nn)
					Jpn16(dmg, nn)
				} else if y == 6 {
					fmt.Printf("DI")
					// TODO: Implement
				} else if y == 7 {
					fmt.Printf("EI")
					// TODO: Implement
				} else {
					fmt.Printf("WARNING : INVALID INSTRUCTIONS")
				}
			case 4:
				if y <= 3 {
					var nn uint16
					nn = uint16(dmg.Memory[dmg.Gbz80.Pc])<<8 + uint16(dmg.Memory[dmg.Gbz80.Pc+1])
					dmg.Gbz80.Pc += 2
					fmt.Printf("CALL %s, %d", DIS_CC[y], nn)
					// TODO: Implement
				} else {
					fmt.Printf("WARNING : INVALID INSTRUCTIONS")
				}
			case 5:
				if q == 0 {
					fmt.Printf("PUSH %s", DIS_RP2[p])
					Pushr16(dmg, RP2[p])
				} else {
					if p == 0 {
						var nn uint16
						nn = uint16(dmg.Memory[dmg.Gbz80.Pc])<<8 + uint16(dmg.Memory[dmg.Gbz80.Pc+1])
						dmg.Gbz80.Pc += 2
						fmt.Printf("CALL %d", nn)
						Calln16(dmg, nn)
					} else {
						fmt.Printf("WARNING : INVALID INSTRUCTIONS")
					}
				}
			case 6:
				n := dmg.Memory[dmg.Gbz80.Pc]
				dmg.Gbz80.Pc += 1
				fmt.Printf("%s %d", DIS_ALU[y], n)
				// TODO: Implement
			case 7:
				fmt.Printf("RST %d", y*8)
				// TODO: Implement
			}
		}

	} else {
		// CB prefixed operations
		switch x {
		case 0x0:
			// rotation/Shift instructions

		case 0x1:
			// Bit test instruction

		case 0x2:
			// Reset bit
			dmg.Gbz80.ResetBitInR8Register(R[z], y)
		case 0x3:
			// Set Bit instructions

		}

	}

}
