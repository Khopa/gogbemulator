package cpu

var OPCODES = map[uint8]func(*Gbz80){
	0x00: NOP,
}
