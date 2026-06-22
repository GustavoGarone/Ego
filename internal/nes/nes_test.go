package nes

import "testing"

func TestRun(t *testing.T) {
	var accumulatorParam uint8 = 0xc0 // -128, 0b1100_0000

	program := []uint8{0xa9, accumulatorParam, 0xaa, 0x00}
	rom := createRom(t, program)

	nes, err := New(rom)
	if err != nil {
		t.Errorf("Failed to load program: %v", err)
	}

	nes.cpu.ProgramCounter = 0x8000
	nes.Run()

	if nes.cpu.Accumulator != accumulatorParam {
		t.Errorf("LDA failed to load to accumulator. Got %x want %x", nes.cpu.Accumulator, accumulatorParam)
	}
	if nes.cpu.Accumulator != nes.cpu.X {
		t.Errorf("Values between accumulator and X differ. Got Acummulator = %x and X = %x", nes.cpu.Accumulator, nes.cpu.X)
	}
}

func createRom(t *testing.T, data []byte) []byte {
	t.Helper()
	rom := make([]byte, 16+16384+8192) // Header + 16KB program + 8KB graphics
	rom[0] = 'N'
	rom[1] = 'E'
	rom[2] = 'S'
	rom[3] = 0x1a
	rom[4] = 1 // chunks of program
	rom[5] = 1 // chunks of graphics

	copy(rom[16:], data)
	return rom
}
