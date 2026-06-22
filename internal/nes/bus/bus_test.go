package bus

import (
	"testing"

	"github.com/GustavoGarone/ego/internal/nes/cartridge"
)

func TestNew(t *testing.T) {
	cart := cartridge.NoGraphics([]byte{0x00})
	b := New(cart)

	if len(b.Ram) != 2048 {
		t.Errorf("Expected RAM size to be 2048 bytes, got %d", len(b.Ram))
	}
	if b.Cart != cart {
		t.Errorf("Cartridge reference not set correctly on Bus")
	}
}

func TestRAMMirroring(t *testing.T) {
	cart := cartridge.NoGraphics([]byte{0x00})
	b := New(cart)
	b.Write(0x0005, 0xA5)

	tests := []struct {
		name    string
		address uint16
		want    byte
	}{
		{"Real RAM read", 0x0005, 0xA5},
		{"First mirror (0x0800)", 0x0805, 0xA5},  // 0x0805 & 0x07FF = 0x0005
		{"Second mirror (0x1000)", 0x1005, 0xA5}, // 0x1005 & 0x07FF = 0x0005
		{"Third mirror (0x1800)", 0x1805, 0xA5},  // 0x1805 & 0x07FF = 0x0005
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := b.Read(tt.address)
			if got != tt.want {
				t.Errorf("Read(0x%04x) failed. Got %x, want %x", tt.address, got, tt.want)
			}
		})
	}
}

func TestReadWrite16(t *testing.T) {
	cart := cartridge.NoGraphics([]byte{0x00})
	b := New(cart)

	var address uint16 = 0x0100
	var value uint16 = 0x1234
	b.Write16(address, value)

	if gotLow := b.Read(address); gotLow != 0x34 {
		t.Errorf("Expected low byte at 0x%04x to be 0x34, got 0x%x", address, gotLow)
	}

	if gotHigh := b.Read(address + 1); gotHigh != 0x12 {
		t.Errorf("Expected high byte at 0x%04x to be 0x12, got 0x%x", address+1, gotHigh)
	}

	if got16 := b.Read16(address); got16 != value {
		t.Errorf("Read16(0x%04x) failed. Got 0x%04x, want 0x%04x", address, got16, value)
	}
}

func TestCartridgeDelegation(t *testing.T) {
	programData := make([]byte, 16384)
	programData[0] = 0xEE

	cart := cartridge.NoGraphics(programData)
	b := New(cart)

	got := b.Read(0x8000)
	if got != 0xEE {
		t.Errorf("Expected cartridge delegation at 0x8000 to return 0xEE, got 0x%02X", got)
	}
}

func TestStubRanges(t *testing.T) {
	cart := cartridge.NoGraphics([]byte{0x00})
	b := New(cart)

	stubs := []struct {
		name    string
		address uint16
	}{
		{"PPU Space", 0x2000},
		{"APU Memory", 0x4015},
		{"Joypad 1", 0x4016},
		{"Joypad 2", 0x4017},
		{"Unmapped Guard Region", 0x401F},
	}

	for _, tt := range stubs {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.Read(tt.address); got != 0 {
				t.Errorf("Stub address 0x%04X should return 0x00, got 0x%02X", tt.address, got)
			}
			// Ensure writing doesn't panic
			b.Write(tt.address, 0xFF)
		})
	}
}
