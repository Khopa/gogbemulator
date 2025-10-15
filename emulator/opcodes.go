package emulator

var OPCODES = map[uint8]func(*DMG){
	0x00: NOP,
}
