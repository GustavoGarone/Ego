package bus

type Bus struct {
	Ram [0x800]uint8
	Rom [0x8000]uint8
}

func (b *Bus) Read(address uint16) uint8 {
	return b.Ram[address]
}

func (b *Bus) Read16(address uint16) uint16 {
	low := uint16(b.Read(address))
	high := uint16(b.Read(address + 1))
	return (high << 8) | low
}

func (b *Bus) Write(address uint16, data uint8) {
	b.Ram[address] = data
}

func (b *Bus) Write16(address uint16, data uint16) {
	low := uint8(data & 0x00FF)
	high := uint8((data & 0xFF00) >> 8)
	b.Write(address, low)
	b.Write(address+1, high)
}
