package cpu

// Load instructions for the GBZ80 CPU

// Load the value from src registers into dst register
func LDr8r8(gbz80 *Gbz80, dst R8Register, src R8Register) {
	gbz80.SetR8Register(dst, gbz80.GetR8Register(src))
}

// Copy immediate 8-bit value into dst register
func LDr8n8(gbz80 *Gbz80, dst R8Register, value uint8) {
	gbz80.SetR8Register(dst, value)
}
