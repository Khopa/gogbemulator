package emulator

import "fmt"

type DMG struct {
	gbz80  *Gbz80
	memory [MemorySize]uint8 // 64KB Memory
}

// SetMemoryU8 sets memory at address to value
func (dmg *DMG) SetMemoryU8(address uint16, value uint8) {
	dmg.memory[address] = value
}

// GetMemoryU8 gets memory value at address
func (dmg *DMG) GetMemoryU8(address uint16) uint8 {
	return dmg.memory[address]
}

type Gbz80 struct {
	af uint16 // Accumulator & Flags
	bc uint16 // B&C registers
	de uint16 // D&E registers
	hl uint16 // H&L registers
	sp uint16 // Stack pointer
	pc uint16 // Program Counter
}

type R8Register uint8

const (
	R8_B  R8Register = 0 // B register
	R8_C  R8Register = 1 // C register
	R8_D  R8Register = 2 // D register
	R8_E  R8Register = 3 // E register
	R8_H  R8Register = 4 // H register
	R8_L  R8Register = 5 // L register
	R8_HL R8Register = 6 // HL register
	R8_A  R8Register = 7 // Accumulator
)

type R16Register uint8

const (
	R16_BC R16Register = 0 // B&C registers
	R16_DE R16Register = 1 // D&E registers
	R16_HL R16Register = 2 // H&L registers
	R16_SP R16Register = 3 // Stack Pointer
	R16_AF R16Register = 4 // A&F Registers
)

// Conditions Table (For CPU Matrix)
const (
	CC_NZ = 0
	CC_Z  = 1
	CC_NC = 2
	CC_C  = 3
)

// ALU Table (For CPU Matrix)
const (
	ALU_ADD_A = 0
	ALU_ADC_A = 1
	ALU_SUB   = 2
	ALU_SBC_A = 3
	ALU_AND   = 4
	ALU_XOR   = 5
	ALU_OR    = 6
	ALU_CP    = 7
)

// ROT Table (For CPU Matrix)
const (
	ROT_RLC  = 0
	ROT_RRC  = 0
	ROT_RL   = 0
	ROT_RR   = 0
	ROT_SLA  = 0
	ROT_SRA  = 0
	ROT_SWAP = 0
	ROT_SRL  = 0
)

// CPU Matrix Tables
var R = []R8Register{R8_B, R8_C, R8_D, R8_E, R8_H, R8_L, R8_HL, R8_A}
var RP = []R16Register{R16_BC, R16_DE, R16_HL, R16_SP}
var RP2 = []R16Register{R16_BC, R16_DE, R16_HL, R16_AF}
var CC = []uint8{CC_NZ, CC_Z, CC_NC, CC_C}
var ALU = []uint8{ALU_ADD_A, ALU_ADC_A, ALU_SUB, ALU_SBC_A, ALU_AND, ALU_XOR, ALU_OR, ALU_CP}
var ROT = []uint8{ROT_RLC, ROT_RRC, ROT_RL, ROT_RR, ROT_SLA, ROT_SRA, ROT_SWAP, ROT_SRL}

// MakeGbz80 Create a new instance of the Z80 Registers
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

// Print the CPU registers in binary format.
func (gbz80 *Gbz80) Print() {
	fmt.Printf("A|F %016b\n", gbz80.af)
	fmt.Printf("B|C %016b\n", gbz80.bc)
	fmt.Printf("D|E %016b\n", gbz80.de)
	fmt.Printf("H|L %016b\n", gbz80.hl)
	fmt.Printf("SP: %016b\n", gbz80.sp)
	fmt.Printf("PC: %016b\n", gbz80.pc)
}

// Reset the CPU registers to their initial state.
func (gbz80 *Gbz80) Reset() {
	gbz80.af = 0b0000000000000000
	gbz80.bc = 0b0000000000000000
	gbz80.de = 0b0000000000000000
	gbz80.hl = 0b0000000000000000
	gbz80.sp = 0b0000000000000000
	gbz80.pc = 0b0000000000000000
}

// A Get Accumulator (A) and Flag Registers access (F)
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

func (gbz80 *Gbz80) HL() uint16 {
	return gbz80.hl
}

// Stack Pointer (SP) and Program Counter (PC) access

func (gbz80 *Gbz80) SP() uint16 {
	return gbz80.sp
}

func (gbz80 *Gbz80) PC() uint16 {
	return gbz80.pc
}

// SetR8Register Setter for 8-bit registers
func (gbz80 *Gbz80) SetR8Register(reg R8Register, value uint8) {
	switch reg {
	case R8_A:
		gbz80.af = (gbz80.af & 0x00FF) | (uint16(value) << 8)
	case R8_B:
		gbz80.bc = (gbz80.bc & 0x00FF) | (uint16(value) << 8)
	case R8_C:
		gbz80.bc = (gbz80.bc & 0xFF00) | uint16(value)
	case R8_D:
		gbz80.de = (gbz80.de & 0x00FF) | (uint16(value) << 8)
	case R8_E:
		gbz80.de = (gbz80.de & 0xFF00) | uint16(value)
	case R8_H:
		gbz80.hl = (gbz80.hl & 0x00FF) | (uint16(value) << 8)
	case R8_L:
		gbz80.hl = (gbz80.hl & 0xFF00) | uint16(value)
	}
}

func (gbz80 *Gbz80) ResetBitInR8Register(reg R8Register, bit uint8) {
	switch reg {
	case R8_A:
		gbz80.af = (gbz80.af & 0x00FF) | ((gbz80.af & 0xFF00) ^ (0x0100 << bit))
	case R8_B:
		gbz80.bc = (gbz80.bc & 0x00FF) | ((gbz80.bc & 0xFF00) ^ (0x0100 << bit))
	case R8_C:
		gbz80.bc = (gbz80.bc & 0xFF00) | ((gbz80.bc & 0x00FF) ^ (0x01 << bit))
	case R8_D:
		gbz80.de = (gbz80.de & 0x00FF) | ((gbz80.de & 0xFF00) ^ (0x0100 << bit))
	case R8_E:
		gbz80.de = (gbz80.de & 0xFF00) | ((gbz80.de & 0x00FF) ^ (0x01 << bit))
	case R8_H:
		gbz80.hl = (gbz80.hl & 0x00FF) | ((gbz80.hl & 0xFF00) ^ (0x0100 << bit))
	case R8_L:
		gbz80.hl = (gbz80.hl & 0xFF00) | ((gbz80.hl & 0x00FF) ^ (0x01 << bit))
	}
}

// GetR8Register Getter for 8-bit registers
func (gbz80 *Gbz80) GetR8Register(reg R8Register) uint8 {
	switch reg {
	case R8_A:
		return uint8(gbz80.bc >> 8)
	case R8_B:
		return uint8(gbz80.bc >> 8)
	case R8_C:
		return uint8(gbz80.bc & 0x00FF)
	case R8_D:
		return uint8(gbz80.de >> 8)
	case R8_E:
		return uint8(gbz80.de & 0x00FF)
	case R8_H:
		return uint8(gbz80.hl >> 8)
	case R8_L:
		return uint8(gbz80.hl & 0x00FF)
	}
	return 0
}

// SetR16Register Setter for 16-bit registers
func (gbz80 *Gbz80) SetR16Register(reg R16Register, value uint16) {
	switch reg {
	case R16_BC:
		gbz80.bc = value
	case R16_DE:
		gbz80.de = value
	case R16_HL:
		gbz80.hl = value
	case R16_SP:
		gbz80.sp = value
	}
}

// GetR16Register Getter for 16-bit registers
func (gbz80 *Gbz80) GetR16Register(reg R16Register) uint16 {
	switch reg {
	case R16_BC:
		return gbz80.bc
	case R16_DE:
		return gbz80.de
	case R16_HL:
		return gbz80.hl
	case R16_SP:
		return gbz80.sp
	}
	return 0
}

// IncrementR8Register Increment Value in R8 register
func (gbz80 *Gbz80) IncrementR8Register(reg R8Register) {
	gbz80.SetR8Register(reg, gbz80.GetR8Register(reg)+1)
}

// DecrementR8Register Decrement Value in R8 register
func (gbz80 *Gbz80) DecrementR8Register(reg R8Register) {
	gbz80.SetR8Register(reg, gbz80.GetR8Register(reg)-1)
}

// IncrementR16Register Increment Value in R16 register
func (gbz80 *Gbz80) IncrementR16Register(reg R16Register) {
	gbz80.SetR16Register(reg, gbz80.GetR16Register(reg)+1)
}

// DecrementR16Register Decrement Value in R16 register
func (gbz80 *Gbz80) DecrementR16Register(reg R16Register) {
	gbz80.SetR16Register(reg, gbz80.GetR16Register(reg)-1)
}
