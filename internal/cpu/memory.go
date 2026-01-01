package cpu

import "errors"

const ZERO_PAGE uint16 = 0x0

type AddressingMode uint8

const (
	Implicit AddressingMode = iota
	Accumulator
	Immediate
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

func (mode AddressingMode) absoluteAddress(cpu *Cpu, base uint16) *Address {
	switch mode {
	case ZeroPage:
		return newAddress(false, ZERO_PAGE+base)

	case ZeroPageX:
		address := uint16(byte(base + uint16(cpu.X)))
		return newAddress(false, address)

	case ZeroPageY:
		address := uint16(byte(base + uint16(cpu.Y)))
		return newAddress(false, address)

	case Absolute:
		return newAddress(false, base)

	case AbsoluteX, AbsoluteXCross:
		address := base + uint16(cpu.X)
		return newAddress(isPageCrossed(base, address), address)

	case AbsoluteY, AbsoluteYCross:
		address := base + uint16(cpu.Y)
		return newAddress(isPageCrossed(base, address), address)

	case IndirectX:
		ptr := byte(base + uint16(cpu.X))
		low := cpu.Read(uint16(ptr))
		high := cpu.Read(uint16(ptr + 1))
		address := uint16(high)<<8 | uint16(low)
		return newAddress(false, address)

	case IndirectY, IndirectYCross:
		ptr := uint8(base)
		low := cpu.Read(uint16(ptr))
		high := cpu.Read(uint16(ptr + 1))
		derefBase := uint16(high)<<8 | uint16(low)
		deref := derefBase + uint16(cpu.Y)
		return newAddress(isPageCrossed(derefBase, deref), deref)

	case Relative:
		offset := int8(cpu.Read(cpu.ProgramCounter))
		cpu.ProgramCounter += 1
		jumpAddr := uint16(int32(cpu.ProgramCounter) + int32(offset))
		return newAddress(isPageCrossed(cpu.ProgramCounter, jumpAddr), jumpAddr)
	}

	panic(errors.New("addressing mode not implemented"))
}

func (mode AddressingMode) Write(cpu *Cpu, data byte) {
	if mode == Accumulator {
		cpu.Accumulator = data
		return
	}

	address := mode.getAddressAndAdvanceProgramCounter(cpu)
	cpu.Write(address.Value, data)
}

func (mode AddressingMode) Read(cpu *Cpu) byte {
	if mode == Accumulator {
		return cpu.Accumulator
	}

	if mode == Immediate {
		return cpu.Fetch()
	}

	address := mode.getAddressAndAdvanceProgramCounter(cpu)
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

func (mode AddressingMode) getAddressAndAdvanceProgramCounter(cpu *Cpu) *Address {
	var address *Address
	if mode == Absolute || mode == AbsoluteX || mode == AbsoluteY {
		arg := cpu.Read16(cpu.ProgramCounter)
		address = mode.absoluteAddress(cpu, arg)
		cpu.ProgramCounter += 2
	} else {
		arg := cpu.Read(cpu.ProgramCounter)
		address = mode.absoluteAddress(cpu, uint16(arg))
		cpu.ProgramCounter += 1
	}
	return address
}
