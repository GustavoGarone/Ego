package cpu

import (
	"testing"

	"github.com/GustavoGarone/ego/internal/bus"
	"github.com/GustavoGarone/ego/internal/cartridge"
)

func TestLda(t *testing.T) {
	var accumulatorParam uint8 = 0xc0 // -128, 0b1100_0000
	program := []uint8{0xa9, accumulatorParam, 0x00}
	cart := cartridge.NoGraphics(program)
	bus := bus.New(cart)
	cpu := New(bus)
	cpu.ProgramCounter = 0x8000

	cpu.Run()
	if cpu.Accumulator != accumulatorParam {
		t.Errorf("LDA failed to load to accumulator. Got %x want %x", cpu.Accumulator, accumulatorParam)
	}
	var want uint8 = 0b1000_0000
	got := cpu.Status & want
	if got != want {
		t.Errorf("LDA failed to update Status register. Got %b want %b", got, want)
	}
}

func TestRun(t *testing.T) {
	var accumulatorParam uint8 = 0xc0 // -128, 0b1100_0000
	program := []uint8{0xa9, accumulatorParam, 0xaa, 0x00}
	cart := cartridge.NoGraphics(program)
	bus := bus.New(cart)
	cpu := New(bus)
	cpu.ProgramCounter = 0x8000

	cpu.Run()
	if cpu.Accumulator != accumulatorParam {
		t.Errorf("LDA failed to load to accumulator. Got %x want %x", cpu.Accumulator, accumulatorParam)
	}
	if cpu.Accumulator != cpu.X {
		t.Errorf("Values between accumulator and X differ. Got Acummulator = %x and X = %x", cpu.Accumulator, cpu.X)
	}
}

func TestUpdateZeroFlag(t *testing.T) {
	program := []uint8{0x00}
	cart := cartridge.NoGraphics(program)
	bus := bus.New(cart)
	cpu := New(bus)
	cpu.ProgramCounter = 0x8000

	cpu.updateZeroFlag(0)
	var want uint8 = 0b0000_0010
	got := cpu.Status & want
	if got != want {
		t.Errorf("Failed to update flag. Got %b want %b", got, want)
	}
	cpu.updateZeroFlag(1)
	want = 0b000_0000
	if cpu.Status != want {
		t.Errorf("Failed to update flag. Got %b want %b", got, want)
	}
}

func TestUpdateNegativeFLag(t *testing.T) {
	program := []uint8{0x00}
	cart := cartridge.NoGraphics(program)
	bus := bus.New(cart)
	cpu := New(bus)
	cpu.ProgramCounter = 0x8000

	cpu.updateNegativeFlag(0b1001_1110)
	var want uint8 = 0b1000_0000
	got := cpu.Status & want
	if got != want {
		t.Errorf("Failed to update flag. Got %b want %b", got, want)
	}
	cpu.updateNegativeFlag(1)
	want = 0b000_0000
	if cpu.Status != want {
		t.Errorf("Failed to update flag. Got %b want %b", got, want)
	}
}
