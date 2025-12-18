package cpu

import "github.com/GustavoGarone/ego/internal/bus"

type Cpu struct {
	// Accumulator, alongside with the ALU, supports the
	// status register for carrying, overflow etc.
	Accumulator uint8

	// X and Y are used for several addressing modes.
	X, Y uint8

	// Stack can be accessed using interrupts, pulls, pushes
	// and transfers.
	Stack uint8

	// Status is used by the ALU. PHP, PLP, arithmetic,
	// testing, and branch instructions can access this register.
	Status uint8

	// ProgramCounter can be accessed either by allowing the CPU's
	// fetch logic increment the address bus, an interrupt and using
	// the RTS/JMP/JSR/Branch instructions.
	ProgramCounter uint16

	// Bus is the system bus the CPU is connected to.  It allows the CPU to read and write
	// to memory and peripherals.
	bus *bus.Bus
}

func NewCpu(bus *bus.Bus) *Cpu {
	return &Cpu{
		Accumulator:    0,
		X:              0,
		Y:              0,
		Status:         0,
		Stack:          0,
		ProgramCounter: 0,
		bus:            bus,
	}
}

func (c *Cpu) Run() {
	for {
		opcode := c.Fetch()
		if c.Execute(opcode) {
			break
		}
		c.ProgramCounter += 1
	}
}

// Fetch gets the current opcode
func (c *Cpu) Fetch() uint8 {
	return c.bus.Rom[c.ProgramCounter]
}

// Execute will handle a program instruction. Returns true if execution is done.
func (c *Cpu) Execute(opcode uint8) bool {
	switch opcode {
	case 0x00:
		return true
	case 0xea:
		return false // NOP
	case 0xa9:
		c.lda(Immediate)
	case 0xa5:
		c.lda(ZeroPage)
	case 0xb5:
		c.lda(ZeroPageX)
	case 0xad:
		c.lda(Absolute)
	case 0xbd:
		c.lda(AbsoluteX)
	case 0xb9:
		c.lda(AbsoluteY)
	case 0xa1:
		c.lda(IndirectX)
	case 0xb1:
		c.lda(IndirectY)
	case 0xa2:
		c.ldx(Immediate)
	case 0xa6:
		c.ldx(ZeroPage)
	case 0xb6:
		c.ldx(ZeroPageY)
	case 0xae:
		c.ldx(Absolute)
	case 0xbe:
		c.ldx(AbsoluteY)
	case 0xa0:
		c.ldy(Immediate)
	case 0xa4:
		c.ldy(ZeroPage)
	case 0xb4:
		c.ldy(ZeroPageX)
	case 0xac:
		c.ldy(Absolute)
	case 0xbc:
		c.ldy(AbsoluteX)
	case 0x85:
		c.sta(ZeroPage)
	case 0x95:
		c.sta(ZeroPageX)
	case 0x8d:
		c.sta(Absolute)
	case 0x9d:
		c.sta(AbsoluteX)
	case 0x99:
		c.sta(AbsoluteY)
	case 0x81:
		c.sta(IndirectX)
	case 0x91:
		c.sta(IndirectY)
	case 0x86:
		c.stx(ZeroPage)
	case 0x96:
		c.stx(ZeroPageY)
	case 0x8e:
		c.stx(Absolute)
	case 0x84:
		c.sty(ZeroPage)
	case 0x94:
		c.sty(ZeroPageX)
	case 0x8c:
		c.sty(Absolute)
	case 0xaa:
		c.tax()
	case 0xa8:
		c.tay()
	case 0x8a:
		c.txa()
	case 0x98:
		c.tya()
	case 0x69:
		c.adc(Immediate)
	case 0x65:
		c.adc(ZeroPage)
	case 0x75:
		c.adc(ZeroPageX)
	case 0x6d:
		c.adc(Absolute)
	case 0x7d:
		c.adc(AbsoluteX)
	case 0x79:
		c.adc(AbsoluteY)
	case 0x61:
		c.adc(IndirectX)
	case 0x71:
		c.adc(IndirectY)
	case 0xe9:
		c.sbc(Immediate)
	case 0xe5:
		c.sbc(ZeroPage)
	case 0xf5:
		c.sbc(ZeroPageX)
	case 0xed:
		c.sbc(Absolute)
	case 0xfd:
		c.sbc(AbsoluteX)
	case 0xf9:
		c.sbc(AbsoluteY)
	case 0xe1:
		c.sbc(IndirectX)
	case 0xf1:
		c.sbc(IndirectY)
	case 0xc6:
		c.dec(ZeroPage)
	case 0xd6:
		c.dec(ZeroPageX)
	case 0xce:
		c.dec(Absolute)
	case 0xde:
		c.dec(AbsoluteX)
	case 0xe8:
		c.inx()
	case 0xc8:
		c.iny()
	case 0xca:
		c.dex()
	case 0x88:
		c.dey()
	case 0x18:
		c.clc()
	case 0x38:
		c.sec()
	case 0x58:
		c.cli()
	case 0x78:
		c.sei()
	case 0xd8:
		c.cld()
	case 0xf8:
		c.sed()
	case 0xb8:
		c.clv()
	}
	return false
}

func (c *Cpu) Write(address uint16, data uint8) {
	c.bus.Write(address, data)
}

func (c *Cpu) Write16(address uint16, data uint16) {
	c.bus.Write16(address, data)
}

func (c *Cpu) Read(address uint16) uint8 {
	return c.bus.Read(address)
}

func (c *Cpu) Read16(address uint16) uint16 {
	return c.bus.Read16(address)
}
