package emulator

const MemorySize = 0x10000 // 64KB Memory

// Memory Map consts for ROM, RAM & I/O
const (
	ROMBank0Start      = 0x0000 // 16KB Cartridge ROM
	ROMBank0End        = 0x3FFF
	ROMBank1Start      = 0x4000 // From cartridge, switchable bank
	ROMBank1End        = 0x7FFF
	VRAMStart          = 0x8000 // 8KB Video RAM
	VRAMEnd            = 0x9FFF
	ExternalRAMStart   = 0xA000 // 8KB External RAM
	ExternalRAMEnd     = 0xBFFF
	WRAMStart          = 0xC000 // 8KB Work RAM
	WRAMEnd            = 0xDFFF
	EchoRAMStart       = 0xE000 // Echo of 8KB Work RAM
	EchoRAMEnd         = 0xFDFF
	OAMStart           = 0xFE00 // Sprite Attribute Table (OAM)
	OAMEnd             = 0xFE9F
	IOPortStart        = 0xFF00 // I/O Ports
	IOPortEnd          = 0xFF7F
	HRAMStart          = 0xFF80 // High RAM
	HRAMEnd            = 0xFFFF
	InterruptEnableReg = 0xFFFF // Interrupt Enable Register
)
