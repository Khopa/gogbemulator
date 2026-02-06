package emulator

import "testing"

func TestRlca(t *testing.T) {
	dmg := MakeDMG()
	dmg.gbz80.SetR8Register(R8_A, 153) // 153 = 0b10011001
	if dmg.gbz80.A() != 153 {
		t.Error("expected A to be set to 153")
	}
	dmg.gbz80.Print()
	Rlca(dmg) //  0b10011001 << 1 = 0b00110011 = 51
	if dmg.gbz80.A() != 51 {
		t.Errorf("expected A to be shifted to 51, not %d", dmg.gbz80.A())
	}
	if dmg.gbz80.C() != 1 {
		t.Error("Expected Carry Flag set")
	}
	dmg.gbz80.Print()
}

func TestRrca(t *testing.T) {
	dmg := MakeDMG()
	dmg.gbz80.SetR8Register(R8_A, 153) // 153 = 0b10011001
	if dmg.gbz80.A() != 153 {
		t.Error("expected A to be set to 153")
	}
	dmg.gbz80.Print()
	Rrca(dmg) //  0b10011001 >> 1 = 0b11001100 = 204
	if dmg.gbz80.A() != 204 {
		t.Errorf("expected A to be shifted to 204, not %d", dmg.gbz80.A())
	}
	if dmg.gbz80.C() != 1 {
		t.Error("Expected Carry Flag set")
	}
	dmg.gbz80.Print()
}

func TestRlcaNoCarry(t *testing.T) {
	dmg := MakeDMG()
	dmg.gbz80.SetR8Register(R8_A, 16) // 16 = 0b00010000
	if dmg.gbz80.A() != 16 {
		t.Error("expected A to be set to 16")
	}
	dmg.gbz80.Print()
	Rlca(dmg) //  0b00010000 << 1 = 0b00100000 = 32
	if dmg.gbz80.A() != 32 {
		t.Errorf("expected A to be shifted to 32, not %d", dmg.gbz80.A())
	}
	if dmg.gbz80.C() != 0 {
		t.Error("Expected Carry Flag not set")
	}
	dmg.gbz80.Print()
}

func TestRrcaNoCarry(t *testing.T) {
	dmg := MakeDMG()
	dmg.gbz80.SetR8Register(R8_A, 16) // 16 = 0b00010000
	if dmg.gbz80.A() != 16 {
		t.Error("expected A to be set to 16")
	}
	dmg.gbz80.Print()
	Rrca(dmg) //  0b00010000 >> 1 = 0b00001000 = 8
	if dmg.gbz80.A() != 8 {
		t.Errorf("expected A to be shifted to 8, not %d", dmg.gbz80.A())
	}
	if dmg.gbz80.C() != 0 {
		t.Error("Expected Carry Flag not set")
	}
	dmg.gbz80.Print()
}
