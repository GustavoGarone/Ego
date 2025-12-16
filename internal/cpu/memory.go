package cpu

const ZERO_PAGE uint16 = 0x0

type AddressingMode uint8

const (
	Implicit AddressingMode = iota
	Accumulator
	Imeddiate
	ZeroPage
	ZeroPageX
	ZeroPageY
	Relative
	Absolute
	AbsoluteX
	AbsoluteXCross
	AbsoluteY
	AbsoluteYCross
	IndirectX
	IndirectY
	IndirectYCross
	None
)

func (mode AddressingMode) absoluteAddress(cpu *cpu.Cpu, base uint16) *Address {
	switch mode {
	case ZeroPage:
		return newAddress(false, ZERO_PAGE+base)

	case ZeroPageX:
		position := uint8(ZERO_PAGE + base)
		var address = uint16(position + cpu.X)
		return newAddress(uint8(address) < position, address)

	case ZeroPageY:
		position := uint8(ZERO_PAGE + base)
		var address = uint16(position + cpu.Y)
		return newAddress(uint8(address) < position, address)

	case Absolute:
		return newAddress(false, base)

	case AbsoluteX, AbsoluteXCross:
		address := base + uint16(cpu.X)
		return newAddress(isPageCrossed(base, address), address)

	case AbsoluteY, AbsoluteYCross:
		address := base + uint16(cpu.Y)
		return newAddress(isPageCrossed(base, address), address)

	case IndirectX:
		pointer := uint8(base) + cpu.X
		low := cpu.Read(uint16(pointer))
		high := cpu.Read(uint16((pointer + 1)))
		address := uint16(high)<<8 | uint16(low)
		return newAddress(false, address)

	case IndirectY, IndirectYCross:
		low := cpu.Read(uint16(base))
		high := cpu.Read(uint16(uint8(base) + 1))
		derefBase := uint16(high)<<8 | uint16(low)
		deref := derefBase + uint16(cpu.Y)
		return newAddress(isPageCrossed(derefBase, deref), deref)
	}

	panic(errors.New("addressing mode not implemented"))
}

func (mode AddressingMode) Write(cpu *cpu.Cpu, position uint16, data uint8) {
	if mode == Accumulator {
		cpu.Accumulator = data
		return
	}

	var address *Address
	if mode == Absolute || mode == AbsoluteX || mode == AbsoluteY {
		arg := cpu.Read16(cpu.ProgramCounter)
		address = mode.absoluteAddress(cpu, arg)
	} else {
		arg := cpu.Read(cpu.ProgramCounter)
		address = mode.absoluteAddress(cpu, uint16(arg))
	}

	cpu.Write(address.Value, data)
}

func (mode AddressingMode) Read(cpu *cpu.Cpu, position uint16) uint8 {
	if mode == Accumulator {
		return cpu.Accumulator
	}

	if mode == Imeddiate {
		return cpu.Read(cpu.ProgramCounter)
	}

	var address *Address
	if mode == Absolute || mode == AbsoluteX || mode == AbsoluteY {
		arg := cpu.Read16(cpu.ProgramCounter)
		address = mode.absoluteAddress(cpu, arg)
	} else {
		arg := cpu.Read(cpu.ProgramCounter)
		address = mode.absoluteAddress(cpu, uint16(arg))
	}

	if address.PageCrossed && mode.isPageCrossMode() {
		// Simulate extra cycle for page crossing
	}

	return cpu.Read(address.Value)
}

func newAddress(pageCrossed bool, value uint16) *Address {
	return &Address{
		PageCrossed: pageCrossed,
		Value:       value,
	}
}

// isPageCrossed returns true if the two addresses are on different memory pages
func isPageCrossed(addr1, addr2 uint16) bool {
	return addr1&0xFF00 != addr2&0xFF00
}

func (mode AddressingMode) isPageCrossMode() bool {
	return mode == AbsoluteXCross || mode == AbsoluteYCross || mode == IndirectYCross
}

type Address struct {
	PageCrossed bool
	Value       uint16
}
