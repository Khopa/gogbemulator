package emulator

import "testing"

func TestDaaNoFlags(t *testing.T) {
	dmg := MakeDMG()
	dmg.Gbz80.SetR8Register(R8_A, 0)
	dmg.Gbz80.setFlag(FLAG_Z, false)
	dmg.Gbz80.setFlag(FLAG_C, false)
	dmg.Gbz80.setFlag(FLAG_N, false)
	dmg.Gbz80.setFlag(FLAG_H, false)

	Daa(dmg)
	if dmg.Gbz80.A() != 0 {
		t.Errorf("expected A to still be zero, but was %d", dmg.Gbz80.A())
	}
	if !dmg.Gbz80.ZeroFlag() {
		t.Error("Zero Flag should be set")
	}
}

func TestDaaWithSubstractionAndCarryFlag(t *testing.T) {
	dmg := MakeDMG()
	dmg.Gbz80.SetR8Register(R8_A, 128)
	dmg.Gbz80.setFlag(FLAG_Z, false)
	dmg.Gbz80.setFlag(FLAG_C, true)
	dmg.Gbz80.setFlag(FLAG_N, true)
	dmg.Gbz80.setFlag(FLAG_H, false)

	Daa(dmg) // -96
	if dmg.Gbz80.A() != 32 {
		t.Errorf("expected A to be 32, but was %x", dmg.Gbz80.A())
	}
	if dmg.Gbz80.ZeroFlag() {
		t.Error("Zero Flag should not be set")
	}
	if dmg.Gbz80.HalfCarryFlag() {
		t.Error("Half Carry should be cleared")
	}
}

func TestDaaWithSubstractionAndCarryAndHalfCarryFlag(t *testing.T) {
	dmg := MakeDMG()
	dmg.Gbz80.SetR8Register(R8_A, 128)
	dmg.Gbz80.setFlag(FLAG_Z, false)
	dmg.Gbz80.setFlag(FLAG_C, true)
	dmg.Gbz80.setFlag(FLAG_N, true)
	dmg.Gbz80.setFlag(FLAG_H, true)

	Daa(dmg) // -96 - 6
	if dmg.Gbz80.A() != 26 {
		t.Errorf("expected A to be 26, but was %x", dmg.Gbz80.A())
	}
	if dmg.Gbz80.ZeroFlag() {
		t.Error("Zero Flag should not be set")
	}
	if dmg.Gbz80.HalfCarryFlag() {
		t.Error("Half Carry should be cleared")
	}
}

func TestZeroingDaaWithSubstractionAndCarryAndHalfCarryFlag(t *testing.T) {
	dmg := MakeDMG()
	dmg.Gbz80.SetR8Register(R8_A, 102)
	dmg.Gbz80.setFlag(FLAG_Z, false)
	dmg.Gbz80.setFlag(FLAG_C, true)
	dmg.Gbz80.setFlag(FLAG_N, true)
	dmg.Gbz80.setFlag(FLAG_H, true)

	Daa(dmg) // -96 - 6
	if dmg.Gbz80.A() != 0 {
		t.Errorf("expected A to be 0, but was %x", dmg.Gbz80.A())
	}
	if !dmg.Gbz80.ZeroFlag() {
		t.Error("Zero Flag should be set")
	}
	if dmg.Gbz80.HalfCarryFlag() {
		t.Error("Half Carry should be cleared")
	}
}

func TestOverflowingDaaWithSubstractionAndCarryAndHalfCarryFlag(t *testing.T) {
	dmg := MakeDMG()
	dmg.Gbz80.SetR8Register(R8_A, 0x32)
	dmg.Gbz80.setFlag(FLAG_Z, false)
	dmg.Gbz80.setFlag(FLAG_C, true)
	dmg.Gbz80.setFlag(FLAG_N, true)
	dmg.Gbz80.setFlag(FLAG_H, true)

	Daa(dmg) // -96 - 6
	if dmg.Gbz80.A() != 0xCC {
		t.Errorf("expected A to be 0xCC, but was %x", dmg.Gbz80.A())
	}
	if dmg.Gbz80.ZeroFlag() {
		t.Error("Zero Flag should not be set")
	}
	if !dmg.Gbz80.CarryFlag() {
		t.Error("Carry should be unchanged")
	}
	if dmg.Gbz80.HalfCarryFlag() {
		t.Error("Half Carry should be cleared")
	}
}
