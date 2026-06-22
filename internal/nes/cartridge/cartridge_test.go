package cartridge

import (
	"bytes"
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	program := []byte{0x01, 0x02, 0x03}
	graphics := []byte{0x04, 0x05}

	cart := New(program, graphics)

	if !bytes.Equal(cart.ProgramData, program) {
		t.Errorf("New() ProgramData mismatch. Got %v, want %v", cart.ProgramData, program)
	}
	if !bytes.Equal(cart.GraphicsData, graphics) {
		t.Errorf("New() GraphicsData mismatch. Got %v, want %v", cart.GraphicsData, graphics)
	}
}

func TestNoGraphics(t *testing.T) {
	program := []byte{0xAA, 0xBB, 0xCC}
	cart := NoGraphics(program)

	if !bytes.Equal(cart.ProgramData, program) {
		t.Errorf("NoGraphics() ProgramData mismatch")
	}
	if len(cart.GraphicsData) != 0 {
		t.Errorf("NoGraphics() GraphicsData should be empty, got size %d", len(cart.GraphicsData))
	}
}

func TestRead(t *testing.T) {
	tests := []struct {
		name        string
		programData []byte
		address     uint16
		expectedVal byte
	}{
		{
			name:        "Exact match inside 16KB ROM (lower bank)",
			programData: makeFakeProgram(16384, map[uint16]byte{0x0000: 0x42}), // index 0
			address:     0x8000,
			expectedVal: 0x42,
		},
		{
			name:        "Mirrored access in 16KB ROM (upper bank mirrors lower bank)",
			programData: makeFakeProgram(16384, map[uint16]byte{0x0001: 0x99}), // index 1
			address:     0xC001,                                                // (0xC001 - 0x8000) % 16384 = 1
			expectedVal: 0x99,
		},
		{
			name:        "Normal access inside 32KB ROM (upper bank)",
			programData: makeFakeProgram(32768, map[uint16]byte{0x4000: 0x77}), // index 16384
			address:     0xC000,                                                // (0xC000 - 0x8000) % 32768 = 16384
			expectedVal: 0x77,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cart := NoGraphics(tt.programData)
			got := cart.Read(tt.address)
			if got != tt.expectedVal {
				t.Errorf("Read(0x%04x) failed. Got %x, want %x", tt.address, got, tt.expectedVal)
			}
		})
	}
}

func TestParse(t *testing.T) {
	buildROM := func(programUnits, graphicsUnits byte, flag6, flag7 byte, includeTrainer bool, totalSize int) []byte {
		rom := make([]byte, totalSize)
		copy(rom[0:4], []byte{'N', 'E', 'S', 0x1A})
		rom[4] = programUnits
		rom[5] = graphicsUnits
		rom[6] = flag6
		rom[7] = flag7
		if includeTrainer {
			rom[6] |= 0b0000_0100
		}
		return rom
	}

	t.Run("Valid standard iNES parsing without Trainer", func(t *testing.T) {
		totalSize := 24592 // 16 bytes header + 16384 (1 unit PRG) + 8192 (1 unit CHR)
		rom := buildROM(1, 1, 0x21, 0x40, false, totalSize)

		// Fill edge markers to check indexing bounds
		rom[16] = 0xAA
		rom[16+16384-1] = 0xBB
		rom[16+16384] = 0xCC
		rom[24591] = 0xDD

		cart, err := Parse(rom)
		if err != nil {
			t.Fatalf("Unexpected Parse error: %v", err)
		}

		// Mapper: 0x40 (high) | 0x02 (low) = 0x42 (66)
		if cart.MapperID != 0x42 {
			t.Errorf("Expected Mapper ID 66, got %d", cart.MapperID)
		}
		if cart.Mirroring != 1 {
			t.Errorf("Expected Mirroring 1, got %d", cart.Mirroring)
		}
		if len(cart.ProgramData) != 16384 || cart.ProgramData[0] != 0xAA || cart.ProgramData[16383] != 0xBB {
			t.Errorf("ProgramData slice bounds malformed")
		}
		if len(cart.GraphicsData) != 8192 || cart.GraphicsData[0] != 0xCC || cart.GraphicsData[8191] != 0xDD {
			t.Errorf("GraphicsData slice bounds malformed")
		}
	})

	t.Run("Valid iNES parsing with Trainer offset", func(t *testing.T) {
		totalSize := 16912 // 16 bytes header + 512 trainer + 16384 PRG + 0 CHR
		rom := buildROM(1, 0, 0x00, 0x00, true, totalSize)
		rom[16+512] = 0x99 // First byte of ProgramData after trainer

		cart, err := Parse(rom)
		if err != nil {
			t.Fatalf("Unexpected Parse error: %v", err)
		}

		if len(cart.ProgramData) != 16384 || cart.ProgramData[0] != 0x99 {
			t.Errorf("Failed to offset ProgramData extraction window over trainer data")
		}
	})

	t.Run("Invalid Header Magic Numbers", func(t *testing.T) {
		badHeader := []byte("NOTANESROMxxxxxx")
		_, err := Parse(badHeader)
		if err == nil || !errors.Is(err, ErrInvalidNesFile) {
			t.Errorf("Expected invalid file error, got: %v", err)
		}
	})

	t.Run("Truncated File Error Check", func(t *testing.T) {
		// Header specifies 1 unit ProgramData (16KB) but slice size is only 100 bytes
		rom := buildROM(1, 0, 0x00, 0x00, false, 100)
		_, err := Parse(rom)
		if err == nil || !errors.Is(err, ErrSmallFile) {
			t.Errorf("Expected truncation error, got: %v", err)
		}
	})
}

// Helper utility to make testing buffers cleanly
func makeFakeProgram(size int, values map[uint16]byte) []byte {
	b := make([]byte, size)
	for addr, val := range values {
		if int(addr) < size {
			b[addr] = val
		}
	}
	return b
}
