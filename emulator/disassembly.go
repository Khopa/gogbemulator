package emulator

import "fmt"

// CPU Matrix Tables
var DIS_R = []string{"B", "C", "D", "E", "H", "L", "(HL)", "A"}
var DIS_RP = []string{"BC", "DE", "HL", "SP"}
var DIS_RP2 = []string{"BC", "DE", "HL", "AF"}
var DIS_CC = []string{"NZ", "Z", "NC", "C"}
var DIS_ALU = []string{"ADD A,", "ADC A,", "SUB", "SBC A,", "AND", "XOR", "OR", "CP"}
var DIS_ROT = []string{"RLC", "RRC", "RL", "RR", "SLA", "SRA", "SWAP", "SRL"}

// Disassembly Dissasemble given ROM
func Disassembly(path string, startPos uint16, lines int) string {
	dmg := MakeDMG()
	err := dmg.LoadROM(path)
	if err != nil {
		panic(err)
	}
	var dissasembly string
	dissasembly = "CURRENT INSTRUCTIONS ===> "

	i := 0
	dmg.Gbz80.Pc = startPos

	for {
		ins := DisassembleCurrentInstructions(dmg)
		// fmt.Printf("%d # %s\n", startPos+uint16(i), ins)
		dissasembly += fmt.Sprintf("%d # %s\n", startPos+uint16(i), ins)
		if dmg.Gbz80.PC() >= 0xFFFF {
			break
		}
		i++
		if i > lines && lines != -1 {
			break
		}
	}
	return dissasembly
}

func DisassembleCurrentInstructions(dmg *DMG) string {

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
					return "NOP"
				case 1:
					var nn uint16
					nn = uint16(dmg.Memory[dmg.Gbz80.Pc])<<8 + uint16(dmg.Memory[dmg.Gbz80.Pc+1])
					dmg.Gbz80.Pc += 2
					return fmt.Sprintf("LD %d, SP", nn)
				case 2:
					return "STOP"
				case 3:
					db := int8(dmg.Memory[dmg.Gbz80.Pc])
					dmg.Gbz80.Pc += 1
					return fmt.Sprintf("JR %d", db)
				default:
					db := int8(dmg.Memory[dmg.Gbz80.Pc])
					dmg.Gbz80.Pc += 1
					return fmt.Sprintf("JR %s, %d", DIS_CC[y-4], db)
				}
			case 0x1:
				// 16 bit load ops
				if q == 0 {
					var nn uint16
					nn = uint16(dmg.Memory[dmg.Gbz80.Pc])<<8 + uint16(dmg.Memory[dmg.Gbz80.Pc+1])
					dmg.Gbz80.Pc += 2
					return fmt.Sprintf("LD %s, %d", DIS_RP[p], nn)
				} else {
					return fmt.Sprintf("ADD HL, %s", DIS_RP[p])
				}
			case 0x2:
				// Indirect Load ops
				if q == 0 {
					switch p {
					case 0:
						return "LD (BC), A"
					case 1:
						return "LD (DE), A"
					case 2:
						return "LD (HL+), A"
					case 3:
						return "LD (HL-), A"
					}
				} else {
					switch p {
					case 0:
						return "LD A, (BC)"
					case 1:
						return "LD A, (DE)"
					case 2:
						return "LD A, (HL+)"
					case 3:
						return "LD A, (HL-)"
					}
				}
			case 0x3:
				// 16 Bit Inc Dec
				if q == 0 {
					return fmt.Sprintf("INC %s", DIS_RP[p])
				} else {
					return fmt.Sprintf("DEC %s", DIS_RP[p])
				}
			case 0x4:
				return fmt.Sprintf("INC %s", DIS_R[y])
			case 0x5:
				return fmt.Sprintf("DEC %s", DIS_R[y])
			case 0x6:
				n := dmg.Memory[dmg.Gbz80.Pc]
				dmg.Gbz80.Pc++
				return fmt.Sprintf("LD %s, %d", DIS_R[y], n)
			case 0x7:
				// Assorted ops on accumulator flags
				switch y {
				case 0:
					return "RLCA"
				case 1:
					return "RRCA"
				case 2:
					return "RLA"
				case 3:
					return "RRA"
				case 4:
					return "DAA"
				case 5:
					return "CPL"
				case 6:
					return "SCF"
				case 7:
					return "CCF"
				}
			}
		case 1:
			// 8 Bit Loading
			if z == 6 && y == 6 {
				return "HALT"
			} else {
				return fmt.Sprintf("LD %s, %s", DIS_R[y], DIS_R[z])
			}
		case 2:
			// ALU operations on acc & registers
			return fmt.Sprintf("%s %s", DIS_ALU[y], DIS_R[z])
		case 3:
			switch z {
			case 0:
				// Conditional return, mem-mapped register loads and stack operations
				if y <= 3 {
					return fmt.Sprintf("RET %s", DIS_CC[y])
				} else if y == 4 {
					n := dmg.Memory[dmg.Gbz80.Pc]
					dmg.Gbz80.Pc += 1
					return fmt.Sprintf("LDH %d, A", n)
				} else if y == 5 {
					db := int8(dmg.Memory[dmg.Gbz80.Pc])
					dmg.Gbz80.Pc += 1
					return fmt.Sprintf("ADD SP, %d", db)
				} else if y == 6 {
					n := dmg.Memory[dmg.Gbz80.Pc]
					dmg.Gbz80.Pc += 1
					return fmt.Sprintf("LDH A, %d", n)
				} else if y == 7 {
					db := int8(dmg.Memory[dmg.Gbz80.Pc])
					dmg.Gbz80.Pc += 1
					return fmt.Sprintf("LD HL, SP+ %d", db)
				}
			case 1:
				// POP & various ops
				if q == 0 {
					return fmt.Sprintf("POP %s", DIS_RP2[p])
				} else {
					if p == 0 {
						return "RET"
					} else if p == 1 {
						return "RETI"
					} else if p == 2 {
						return "JP HL"
					} else if p == 3 {
						return "LD SP, HL"
					}
				}
			case 2:
				// Conditional jumps
				if y <= 3 {
					var nn uint16
					nn = uint16(dmg.Memory[dmg.Gbz80.Pc])<<8 + uint16(dmg.Memory[dmg.Gbz80.Pc+1])
					dmg.Gbz80.Pc += 2
					return fmt.Sprintf("JP %s, %d", DIS_CC[y], nn)
				} else if y == 4 {
					return "LDH C, A"
				} else if y == 5 {
					var nn uint16
					nn = uint16(dmg.Memory[dmg.Gbz80.Pc])<<8 + uint16(dmg.Memory[dmg.Gbz80.Pc+1])
					dmg.Gbz80.Pc += 2
					return fmt.Sprintf("LD (%d), A", nn)
				} else if y == 6 {
					return "LDH A, C"
				} else if y == 7 {
					var nn uint16
					nn = uint16(dmg.Memory[dmg.Gbz80.Pc])<<8 + uint16(dmg.Memory[dmg.Gbz80.Pc+1])
					dmg.Gbz80.Pc += 2
					return fmt.Sprintf("LD A, (%d)", nn)
				}
			case 3:
				if y == 0 {
					var nn uint16
					nn = uint16(dmg.Memory[dmg.Gbz80.Pc])<<8 + uint16(dmg.Memory[dmg.Gbz80.Pc+1])
					dmg.Gbz80.Pc += 2
					return fmt.Sprintf("JP %d", nn)
				} else if y == 6 {
					return "DI"
				} else if y == 7 {
					return "EI"
				} else {
					return "INVALID INSTRUCTIONS"
				}
			case 4:
				if y <= 3 {
					var nn uint16
					nn = uint16(dmg.Memory[dmg.Gbz80.Pc])<<8 + uint16(dmg.Memory[dmg.Gbz80.Pc+1])
					dmg.Gbz80.Pc += 2
					return fmt.Sprintf("CALL %s, %d", DIS_CC[y], nn)
				} else {
					return "INVALID INSTRUCTIONS"
				}
			case 5:
				if q == 0 {
					return fmt.Sprintf("PUSH %s", DIS_RP2[p])
				} else {
					if p == 0 {
						var nn uint16
						nn = uint16(dmg.Memory[dmg.Gbz80.Pc])<<8 + uint16(dmg.Memory[dmg.Gbz80.Pc+1])
						dmg.Gbz80.Pc += 2
						return fmt.Sprintf("CALL %d", nn)
					} else {
						return "INVALID INSTRUCTIONS"
					}
				}
			case 6:
				n := dmg.Memory[dmg.Gbz80.Pc]
				dmg.Gbz80.Pc += 1
				return fmt.Sprintf("%s %d", DIS_ALU[y], n)
			case 7:
				return fmt.Sprintf("RST %d", y*8)
			}

		}

	} else {
		// CB prefixed operations
		switch x {
		case 0:
			// rotation/Shift instructions
			return fmt.Sprintf("%s %s", DIS_ROT[y], DIS_R[z])
		case 1:
			// Bit test instruction
			return fmt.Sprintf("BIT %d %s", y, DIS_R[z])
		case 2:
			// Reset bit
			return fmt.Sprintf("RES %d %s", y, DIS_R[z])
		case 3:
			// Set bit
			return fmt.Sprintf("SET %d %s", y, DIS_R[z])

		}
	}

	return "?????"

}
