package emulator

import "testing"

func TestRlca(t *testing.T) {
	dmg := MakeDMG()
	dmg.Gbz80.SetR8Register(R8_A, 153) // 153 = 0b10011001
	if dmg.Gbz80.A() != 153 {
		t.Error("expected A to be set to 153")
	}
	dmg.Gbz80.Print()
	Rlca(dmg) //  0b10011001 << 1 = 0b00110011 = 51
	if dmg.Gbz80.A() != 51 {
		t.Errorf("expected A to be shifted to 51, not %d", dmg.Gbz80.A())
	}
	if !dmg.Gbz80.CarryFlag() {
		t.Error("Expected Carry Flag set")
	}
	dmg.Gbz80.Print()
}

func TestRrca(t *testing.T) {
	dmg := MakeDMG()
	dmg.Gbz80.SetR8Register(R8_A, 153) // 153 = 0b10011001
	if dmg.Gbz80.A() != 153 {
		t.Error("expected A to be set to 153")
	}
	dmg.Gbz80.Print()
	Rrca(dmg) //  0b10011001 >> 1 = 0b11001100 = 204
	if dmg.Gbz80.A() != 204 {
		t.Errorf("expected A to be shifted to 204, not %d", dmg.Gbz80.A())
	}
	if !dmg.Gbz80.CarryFlag() {
		t.Error("Expected Carry Flag set")
	}
	dmg.Gbz80.Print()
}

func TestRlcaNoCarry(t *testing.T) {
	dmg := MakeDMG()
	dmg.Gbz80.SetR8Register(R8_A, 16) // 16 = 0b00010000
	if dmg.Gbz80.A() != 16 {
		t.Error("expected A to be set to 16")
	}
	dmg.Gbz80.Print()
	Rlca(dmg) //  0b00010000 << 1 = 0b00100000 = 32
	if dmg.Gbz80.A() != 32 {
		t.Errorf("expected A to be shifted to 32, not %d", dmg.Gbz80.A())
	}
	if dmg.Gbz80.CarryFlag() {
		t.Error("Expected Carry Flag not set")
	}
	dmg.Gbz80.Print()
}

func TestRrcaNoCarry(t *testing.T) {
	dmg := MakeDMG()
	dmg.Gbz80.SetR8Register(R8_A, 16) // 16 = 0b00010000
	if dmg.Gbz80.A() != 16 {
		t.Error("expected A to be set to 16")
	}
	dmg.Gbz80.Print()
	Rrca(dmg) //  0b00010000 >> 1 = 0b00001000 = 8
	if dmg.Gbz80.A() != 8 {
		t.Errorf("expected A to be shifted to 8, not %d", dmg.Gbz80.A())
	}
	if dmg.Gbz80.CarryFlag() {
		t.Error("Expected Carry Flag not set")
	}
	dmg.Gbz80.Print()
}

func TestRlaWithCarrySetBefore(t *testing.T) {
	dmg := MakeDMG()
	dmg.Gbz80.SetR8Register(R8_A, 0b10001000)
	dmg.Gbz80.setFlag(FLAG_C, true)
	if dmg.Gbz80.A() != 0b10001000 {
		t.Error("expected A to be set to 0b10001000")
	}
	dmg.Gbz80.Print()
	Rla(dmg) //  0b10001000 << 1 (with carry) = 0b00010001
	if dmg.Gbz80.A() != 0b00010001 {
		t.Errorf("expected A to be shifted to 0b00010001, not %b", dmg.Gbz80.A())
	}
	if !dmg.Gbz80.CarryFlag() {
		t.Error("Expected Carry Flag set because bit 7 was set before rotation")
	}
	dmg.Gbz80.Print()
}

func TestRlaWithNoCarrySetBefore(t *testing.T) {
	dmg := MakeDMG()
	dmg.Gbz80.SetR8Register(R8_A, 0b10001000)
	dmg.Gbz80.setFlag(FLAG_C, false)
	if dmg.Gbz80.A() != 0b10001000 {
		t.Error("expected A to be set to 0b10001000")
	}
	dmg.Gbz80.Print()
	Rla(dmg) //  0b10001000 << 1 (with no carry) = 0b00010000
	if dmg.Gbz80.A() != 0b00010000 {
		t.Errorf("expected A to be shifted to 0b00010001, not %b", dmg.Gbz80.A())
	}
	if !dmg.Gbz80.CarryFlag() {
		t.Error("Expected Carry Flag set because bit 7 was set before rotation")
	}
	dmg.Gbz80.Print()
}

func TestRlaWithoutAnyCarry(t *testing.T) {
	dmg := MakeDMG()
	dmg.Gbz80.SetR8Register(R8_A, 0b00001000)
	dmg.Gbz80.setFlag(FLAG_C, false)
	if dmg.Gbz80.A() != 0b00001000 {
		t.Error("expected A to be set to 0b00001000")
	}
	dmg.Gbz80.Print()
	Rla(dmg) //  0b00001000 << 1 (with no carry) = 0b00010000
	if dmg.Gbz80.A() != 0b00010000 {
		t.Errorf("expected A to be shifted to 0b00010000, not %b", dmg.Gbz80.A())
	}
	if dmg.Gbz80.CarryFlag() {
		t.Error("Expected Carry Flag not set because bit 7 was not to 1 before rotation")
	}
	dmg.Gbz80.Print()
}

func TestRraWithCarrySetBefore(t *testing.T) {
	dmg := MakeDMG()
	dmg.Gbz80.SetR8Register(R8_A, 0b00001001)
	dmg.Gbz80.setFlag(FLAG_C, true)
	if dmg.Gbz80.A() != 0b00001001 {
		t.Error("expected A to be set to 0b00001001")
	}
	dmg.Gbz80.Print()
	Rra(dmg) //  0b00001001 >> 1 (with carry) = 0b10000100
	if dmg.Gbz80.A() != 0b10000100 {
		t.Errorf("expected A to be shifted to 0b10000100, not %b", dmg.Gbz80.A())
	}
	if !dmg.Gbz80.CarryFlag() {
		t.Error("Expected Carry Flag set because bit 1 was set before rotation")
	}
	dmg.Gbz80.Print()
}

func TestRraWithNoCarrySetBefore(t *testing.T) {
	dmg := MakeDMG()
	dmg.Gbz80.SetR8Register(R8_A, 0b00001001)
	dmg.Gbz80.setFlag(FLAG_C, false)
	if dmg.Gbz80.A() != 0b00001001 {
		t.Error("expected A to be set to 0b00001001")
	}
	dmg.Gbz80.Print()
	Rra(dmg) //  0b00001001 >> 1 (with no carry) = 0b00000100
	if dmg.Gbz80.A() != 0b00000100 {
		t.Errorf("expected A to be shifted to 0b00000100, not %b", dmg.Gbz80.A())
	}
	if !dmg.Gbz80.CarryFlag() {
		t.Error("Expected Carry Flag set because bit 1 was set before rotation")
	}
	dmg.Gbz80.Print()
}

func TestRraWithoutAnyCarry(t *testing.T) {
	dmg := MakeDMG()
	dmg.Gbz80.SetR8Register(R8_A, 0b00001000)
	dmg.Gbz80.setFlag(FLAG_C, false)
	if dmg.Gbz80.A() != 0b00001000 {
		t.Error("expected A to be set to 0b00001000")
	}
	dmg.Gbz80.Print()
	Rra(dmg) //  0b00001000 >> 1 (with no carry) = 0b00000100
	if dmg.Gbz80.A() != 0b00000100 {
		t.Errorf("expected A to be shifted to 0b00000100, not %b", dmg.Gbz80.A())
	}
	if dmg.Gbz80.CarryFlag() {
		t.Error("Expected Carry Flag not set because bit 1 was not set before rotation")
	}
	dmg.Gbz80.Print()
}
