package cpu

func translate_r8(dst uint8) uint8 {
	// Translate 3-bit register code to actual register reference
	switch dst {
		case 0:
			return gbz80.B()
		case 1:
			return gbz80.C()
		case 2:
			return gbz80.D()
		case 3:
			return gbz80.E()
		case 4:
			return gbz80.H()
		case 5:
			return gbz80.L()
		case 6:
			return gbz80.HL()
		case 7:
			return gbz80.A()
		default:
			return 0
		}
	}
}