package emulator

import "testing"

func TestDecR8(t *testing.T) {
	dmg := MakeDMG()
	dmg.Gbz80.SetR8Register(R8_B, 25)
	if dmg.Gbz80.B() != 25 {
		t.Error("expected B to be set to 25")
	}
	dmg.Gbz80.Print()
	DecR8(dmg, R8_B)
	if dmg.Gbz80.B() != 24 {
		t.Error("expected B to be decremented to 24")
	}
}

func TestIncR8(t *testing.T) {
	dmg := MakeDMG()
	dmg.Gbz80.SetR8Register(R8_B, 25)
	if dmg.Gbz80.B() != 25 {
		t.Error("expected B to be set to 25")
	}
	dmg.Gbz80.Print()
	IncR8(dmg, R8_B)
	if dmg.Gbz80.B() != 26 {
		t.Error("expected B to be incremented to 26")
	}
}

func TestDecR16(t *testing.T) {
	dmg := MakeDMG()
	dmg.Gbz80.SetR16Register(R16_BC, 25)
	if dmg.Gbz80.Bc != 25 {
		t.Error("expected BC to be set to 25")
	}
	dmg.Gbz80.Print()
	DecR16(dmg, R16_BC)
	if dmg.Gbz80.Bc != 24 {
		t.Error("expected BC to be decremented to 24")
	}
}

func TestIncR16(t *testing.T) {
	dmg := MakeDMG()
	dmg.Gbz80.SetR16Register(R16_BC, 25)
	if dmg.Gbz80.Bc != 25 {
		t.Error("expected BC to be set to 25")
	}
	IncR16(dmg, R16_BC)
	if dmg.Gbz80.Bc != 26 {
		t.Error("expected BC to be incremented to 26")
	}
}
