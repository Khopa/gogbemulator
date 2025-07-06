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

func MakeGbz80() *Gbz80 {
	return &Gbz80{
		af: 0,
		bc: 0,
		de: 0,
		hl: 0,
		sp: 0,
		pc: 0,
	}
}

func (gbz80 *Gbz80) Print() {
	fmt.Printf("A|F %016b\n", gbz80.af)
	fmt.Printf("B|C %016b\n", gbz80.bc)
	fmt.Printf("D|E %016b\n", gbz80.de)
	fmt.Printf("H|L %016b\n", gbz80.hl)
	fmt.Printf("SP: %016b\n", gbz80.sp)
	fmt.Printf("PC: %016b\n", gbz80.pc)
}
