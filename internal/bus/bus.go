package bus

type Bus struct {
	Ram []uint8
	Rom []uint8
}

func NewBus(ram []uint8, rom []uint8) *Bus {
	return &Bus{
		Ram: ram,
		Rom: rom,
	}
}

func (b *Bus) Read(address uint16) uint8 {
	// & 0x07FF mods every two kilobytes to account for mirroring
	return b.Ram[address&0x07FF]
}

func (b *Bus) Read16(address uint16) uint16 {
	low := uint16(b.Read(address))
	high := uint16(b.Read(address + 1))
	return (high << 8) | low
}

func (b *Bus) Write(address uint16, data uint8) {
	b.Ram[address&0x07FF] = data
}

func (b *Bus) Write16(address uint16, data uint16) {
	low := uint8(data & 0x00FF)
	high := uint8((data & 0xFF00) >> 8)
	b.Write(address, low)
	b.Write(address+1, high)
}
