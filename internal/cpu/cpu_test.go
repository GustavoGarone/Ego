package cpu

import (
	"testing"

	"github.com/GustavoGarone/ego/internal/bus"
)

func TestLda(t *testing.T) {
	var accumulatorParam uint8 = 0xc0 // -128, 11000000
	program := []uint8{0xa9, accumulatorParam, 0x00}
	bus := bus.NewBus([]uint8{0}, program) // Empty dummy ram, program as rom
	cpu := NewCpu(bus)
	cpu.Run()
	if cpu.Accumulator != accumulatorParam {
		t.Errorf("LDA failed to load to accumulator. Got %x want %x", cpu.Accumulator, accumulatorParam)
	}
	var expected uint8 = 0b1000_0000
	if cpu.Status != expected {
		t.Errorf("LDA failed to update Status register. Got %b want %b", cpu.Status, expected)
	}
}

func TestRun(t *testing.T) {
	var accumulatorParam uint8 = 0xc0
	program := []uint8{0xa9, 0xc0, 0xaa, 0x00}
	bus := bus.NewBus([]uint8{0}, program)
	cpu := NewCpu(bus)
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
	bus := bus.NewBus([]uint8{0}, program)
	cpu := NewCpu(bus)
	cpu.updateZeroFlag(0)
	var expect uint8 = 0b0000_0010
	if cpu.Status != expect {
		t.Errorf("Failed to update flag. Got %b want %b", cpu.Status, expect)
	}
	cpu.updateZeroFlag(1)
	expect = 0b000_0000
	if cpu.Status != expect {
		t.Errorf("Failed to update flag. Got %b want %b", cpu.Status, expect)
	}
}

func TestUpdateNegativeFLag(t *testing.T) {
	program := []uint8{0x00}
	bus := bus.NewBus([]uint8{0}, program)
	cpu := NewCpu(bus)
	cpu.updateNegativeFlag(0b1001_1110)
	var expect uint8 = 0b1000_0000
	if cpu.Status != expect {
		t.Errorf("Failed to update flag. Got %b want %b", cpu.Status, expect)
	}
	cpu.updateNegativeFlag(1)
	expect = 0b000_0000
	if cpu.Status != expect {
		t.Errorf("Failed to update flag. Got %b want %b", cpu.Status, expect)
	}
}
