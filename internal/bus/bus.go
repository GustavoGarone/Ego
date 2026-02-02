package bus

import (
	"github.com/GustavoGarone/ego/internal/cartridge"
)

type Bus struct {
	Ram  []byte
	Cart *cartridge.Cartridge
}

func New(cartridge *cartridge.Cartridge) *Bus {
	return &Bus{
		Ram:  make([]byte, 2048),
		Cart: cartridge,
	}
}

const (
	memoryRangeCPU       uint16 = 0x1FFF
	memoryRangePPU       uint16 = 0x3FFF
	memoryRangeAPU       uint16 = 0x4017
	memoryRangeCartridge uint16 = 0x4020
)

const (
	memoryAPU     uint16 = 0x4015
	memoryJoypad1 uint16 = 0x4016
	memoryJoypad2 uint16 = 0x4017
)

const (
	// mirrorCPU should be used to mirror every two kilobytes
	mirrorCPU uint16 = 0x07FF
)

func (b *Bus) Read(address uint16) byte {
	switch {
	case address <= memoryRangeCPU:
		return b.Ram[address&mirrorCPU]
	case address <= memoryRangePPU:
		// TODO: add PPU handling
		return 0
	case address <= memoryRangeAPU:
		switch address {
		case memoryAPU:
			// TODO: add APU handling
			return 0
		case memoryJoypad1:
			// TODO: add joypad handling
			return 0
		case memoryJoypad2:
			// TODO: add joypad handling
			return 0
		default:
			return 0
		}
	case address >= memoryRangeCartridge:
		return b.Cart.Read(address)
	default:
		return 0
	}
}

func (b *Bus) Read16(address uint16) uint16 {
	low := uint16(b.Read(address))
	high := uint16(b.Read(address + 1))
	return (high << 8) | low
}

func (b *Bus) Write(address uint16, data byte) {
	switch {
	case address <= memoryRangeCPU:
		b.Ram[address&mirrorCPU] = data
	case address <= memoryRangePPU:
		// TODO: add PPU handling
		return
	case address >= memoryRangeCartridge:
		b.Cart.Write(address, data)
	default:
		return
	}
}

func (b *Bus) Write16(address uint16, data uint16) {
	low := byte(data & 0x00FF)
	high := byte((data & 0xFF00) >> 8)
	b.Write(address, low)
	b.Write(address+1, high)
}
