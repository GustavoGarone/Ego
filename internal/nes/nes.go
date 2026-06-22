package nes

import (
	"github.com/GustavoGarone/ego/internal/emulator"
	"github.com/GustavoGarone/ego/internal/nes/bus"
	"github.com/GustavoGarone/ego/internal/nes/cartridge"
	"github.com/GustavoGarone/ego/internal/nes/cpu"
)

type Nes struct {
	cpu *cpu.Cpu
}

var (
	_ emulator.Emulator = &Nes{}
)

// New creates a new emulator with the provided ROM data.
func New(romData []byte) (*Nes, error) {
	cartridge, err := cartridge.Parse(romData)
	if err != nil {
		return nil, err
	}

	bus := bus.New(cartridge)
	cpu := cpu.New(bus)
	return &Nes{cpu: cpu}, nil
}

func (n *Nes) Run() {
	for {
		shouldStop := n.cpu.Step()
		if shouldStop {
			break
		}
	}
}
