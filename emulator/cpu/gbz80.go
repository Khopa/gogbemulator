package cpu

import "fmt"

type Gbz80 struct {
	af uint16 // Accumulator & Flags
	bc uint16 // B&C registers
	de uint16 // D&E registers
	hl uint16 // H&L registers
	sp uint16 // Stack pointer
	pc uint16 // Program Counter
}

/**
 * Create a new instance of the Z80 Registers
 */
func MakeGbz80() *Gbz80 {
	return &Gbz80{
		af: 0b0000000010000000,
		bc: 0b0000000000000000,
		de: 0b0000000000000000,
		hl: 0b0000000000000000,
		sp: 0b0000000000000000,
		pc: 0b0000000000000000,
	}
}

/**
 * Print the CPU registers in binary format.
 */
func (gbz80 *Gbz80) Print() {
	fmt.Printf("A|F %016b\n", gbz80.af)
	fmt.Printf("B|C %016b\n", gbz80.bc)
	fmt.Printf("D|E %016b\n", gbz80.de)
	fmt.Printf("H|L %016b\n", gbz80.hl)
	fmt.Printf("SP: %016b\n", gbz80.sp)
	fmt.Printf("PC: %016b\n", gbz80.pc)
}

/**
 * Reset the CPU registers to their initial state.
 */
func (gbz80 *Gbz80) Reset() {
	gbz80.af = 0b0000000000000000
	gbz80.bc = 0b0000000000000000
	gbz80.de = 0b0000000000000000
	gbz80.hl = 0b0000000000000000
	gbz80.sp = 0b0000000000000000
	gbz80.pc = 0b0000000000000000
}

// Accumulator (A) and Flag Registers access (F)

func (gbz80 *Gbz80) A() uint8 {
	return uint8(gbz80.af >> 8)
}

// Flag Register (F) bits:
// +-----+-----+--------------------------+
// | Bit | Name| Explanation              |
// +-----+-----+--------------------------+
// |  7  |  z  | Zero flag                |
// |  6  |  n  | Subtraction flag (BCD)   |
// |  5  |  h  | Half Carry flag (BCD)    |
// |  4  |  c  | Carry flag               |
// +-----+-----+--------------------------+

func (gbz80 *Gbz80) ZeroFlag() bool {
	return (gbz80.af & 0b0001) != 0
}

func (gbz80 *Gbz80) SubtractionFlag() bool {
	return (gbz80.af & 0b0010) != 0
}

func (gbz80 *Gbz80) HalfCarryFlag() bool {
	return (gbz80.af & 0b0100) != 0
}

func (gbz80 *Gbz80) CarryFlag() bool {
	return (gbz80.af & 0b1000) != 0
}

// BC Registers access

func (gbz80 *Gbz80) B() uint8 {
	return uint8(gbz80.bc >> 8)
}

func (gbz80 *Gbz80) C() uint8 {
	return uint8(gbz80.bc & 0x00FF)
}

// DE Registers access

func (gbz80 *Gbz80) D() uint8 {
	return uint8(gbz80.de >> 8)
}

func (gbz80 *Gbz80) E() uint8 {
	return uint8(gbz80.de & 0x00FF)
}

// HL Registers access

func (gbz80 *Gbz80) H() uint8 {
	return uint8(gbz80.hl >> 8)
}

func (gbz80 *Gbz80) L() uint8 {
	return uint8(gbz80.hl & 0x00FF)
}

// Stack Pointer (SP) and Program Counter (PC) access

func (gbz80 *Gbz80) SP() uint16 {
	return gbz80.sp
}

func (gbz80 *Gbz80) PC() uint16 {
	return gbz80.pc
}
