package emulator

import "testing"

func TestDecR8(t *testing.T) {
	dmg := MakeDMG()
	dmg.gbz80.SetR8Register(R8_B, 25)
	if dmg.gbz80.B() != 25 {
		t.Error("expected B to be set to 25")
	}
	dmg.gbz80.Print()
	dmg.gbz80.DecrementR8Register(R8_B)
	if dmg.gbz80.B() != 24 {
		t.Error("expected B to be decremented to 24")
	}
}

func TestIncR8(t *testing.T) {
	dmg := MakeDMG()
	dmg.gbz80.SetR8Register(R8_B, 25)
	if dmg.gbz80.B() != 25 {
		t.Error("expected B to be set to 25")
	}
	dmg.gbz80.Print()
	dmg.gbz80.IncrementR8Register(R8_B)
	if dmg.gbz80.B() != 26 {
		t.Error("expected B to be incremented to 26")
	}
}

func TestDecR16(t *testing.T) {
	dmg := MakeDMG()
	dmg.gbz80.SetR16Register(R16_BC, 25)
	if dmg.gbz80.bc != 25 {
		t.Error("expected BC to be set to 25")
	}
	dmg.gbz80.Print()
	dmg.gbz80.DecrementR16Register(R16_BC)
	if dmg.gbz80.bc != 24 {
		t.Error("expected BC to be decremented to 24")
	}
}

func TestIncR16(t *testing.T) {
	dmg := MakeDMG()
	dmg.gbz80.SetR16Register(R16_BC, 25)
	if dmg.gbz80.bc != 25 {
		t.Error("expected BC to be set to 25")
	}
	dmg.gbz80.Print()
	dmg.gbz80.IncrementR16Register(R16_BC)
	if dmg.gbz80.bc != 26 {
		t.Error("expected BC to be incremented to 26")
	}
}
