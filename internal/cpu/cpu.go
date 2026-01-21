package cpu

import (
	"github.com/GustavoGarone/ego/internal/bus"
)

type Cpu struct {
	// Accumulator, alongside with the ALU, supports the
	// status register for carrying, overflow etc.
	Accumulator byte

	// X and Y are used for several addressing modes.
	X, Y byte

	// Stack can be accessed using interrupts, pulls, pushes
	// and transfers.
	Stack byte

	// Status is used by the ALU. PHP, PLP, arithmetic,
	// testing, and branch instructions can access this register.
	Status byte

	// ProgramCounter can be accessed either by allowing the CPU's
	// fetch logic increment the address bus, an interrupt and using
	// the RTS/JMP/JSR/Branch instructions.
	ProgramCounter uint16

	// Bus is the system bus the CPU is connected to. It allows the CPU to read and write
	// to memory and peripherals.
	bus *bus.Bus
}

func New(bus *bus.Bus) *Cpu {
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

// Resets the CPU, setting registers to zero and defaulting the stack, status
// and program counter registers to the default `reset` states.
func (c *Cpu) Reset() {
	c.Accumulator = 0
	c.X = 0
	c.Y = 0
	c.Stack = 0xFD
	c.Status = 0x24
	c.ProgramCounter = c.bus.Read16(0xFFFC)
}

// Runs the Fetch/Decode/Execute loop.
func (c *Cpu) Run() {
	for {
		opcode := c.Fetch()
		if c.Execute(opcode) {
			break
		}
	}
}

// Fetch gets the current opcode
func (c *Cpu) Fetch() byte {
	data := c.bus.Read(c.ProgramCounter)
	c.ProgramCounter += 1
	return data
}

// Writes data to the address in the bus
func (c *Cpu) Write(address uint16, data byte) {
	c.bus.Write(address, data)
}

// Write16 writes 2 bytes of data to the address in the bus
func (c *Cpu) Write16(address uint16, data uint16) {
	c.bus.Write16(address, data)
}

// Reads from the address in the bus
func (c *Cpu) Read(address uint16) byte {
	return c.bus.Read(address)
}

// Read16 reads 2 bytes of data from the address in the bus
func (c *Cpu) Read16(address uint16) uint16 {
	return c.bus.Read16(address)
}

// pushes the value to the stack
func (c *Cpu) push(data byte) {
	c.Write(0x0100+uint16(c.Stack), data)
	c.Stack -= 1
}

// push16 pushes 2 bytes of data to the stack
func (c *Cpu) push16(data uint16) {
	low := byte(data >> 8)
	high := byte(data & 0xFF)
	c.push(high)
	c.push(low)
}

// pops the value from the stack
func (c *Cpu) pop() byte {
	c.Stack += 1
	return c.Read(0x0100 + uint16(c.Stack))
}

// pops 2 bytes of data from the stack
func (c *Cpu) pop16() uint16 {
	low := uint16(c.pop())
	high := uint16(c.pop())
	return (high << 8) | low
}

// Execute will handle a program instruction. Returns true if execution is done.
func (c *Cpu) Execute(opcode byte) bool {
	switch opcode {
	case 0x00:
		c.brk()
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
	case 0x4c:
		c.jmp(Absolute)
	case 0x6c:
		c.jmp(None)
	case 0x20:
		c.jsr()
	case 0x60:
		c.rts()
	case 0x40:
		c.rti()
	case 0x0a:
		c.asl(Accumulator)
	case 0x06:
		c.asl(ZeroPage)
	case 0x16:
		c.asl(ZeroPageX)
	case 0x0e:
		c.asl(Absolute)
	case 0x1e:
		c.asl(AbsoluteX)
	case 0x4a:
		c.asl(Accumulator)
	case 0x46:
		c.asl(ZeroPage)
	case 0x56:
		c.asl(ZeroPageX)
	case 0x4e:
		c.asl(Absolute)
	case 0x5e:
		c.asl(AbsoluteX)
	case 0x2a:
		c.rol(Accumulator)
	case 0x26:
		c.rol(ZeroPage)
	case 0x36:
		c.rol(ZeroPageX)
	case 0x2e:
		c.rol(Absolute)
	case 0x3e:
		c.rol(AbsoluteX)
	case 0x6a:
		c.ror(Accumulator)
	case 0x66:
		c.ror(ZeroPage)
	case 0x76:
		c.ror(ZeroPageX)
	case 0x6e:
		c.ror(Absolute)
	case 0x7e:
		c.ror(AbsoluteX)
	case 0x29:
		c.and(Immediate)
	case 0x25:
		c.and(ZeroPage)
	case 0x35:
		c.and(ZeroPageX)
	case 0x2d:
		c.and(Absolute)
	case 0x3d:
		c.and(AbsoluteX)
	case 0x39:
		c.and(AbsoluteY)
	case 0x21:
		c.and(IndirectX)
	case 0x31:
		c.and(IndirectY)
	case 0x09:
		c.ora(Immediate)
	case 0x05:
		c.ora(ZeroPage)
	case 0x15:
		c.ora(ZeroPageX)
	case 0x0d:
		c.ora(Absolute)
	case 0x1d:
		c.ora(AbsoluteX)
	case 0x19:
		c.ora(AbsoluteY)
	case 0x01:
		c.ora(IndirectX)
	case 0x11:
		c.ora(IndirectY)
	case 0x49:
		c.eor(Immediate)
	case 0x45:
		c.eor(ZeroPage)
	case 0x55:
		c.eor(ZeroPageX)
	case 0x4d:
		c.eor(Absolute)
	case 0x5d:
		c.eor(AbsoluteX)
	case 0x59:
		c.eor(AbsoluteY)
	case 0x41:
		c.eor(IndirectX)
	case 0x51:
		c.eor(IndirectY)
	case 0x24:
		c.bit(ZeroPage)
	case 0x2c:
		c.bit(Absolute)
	case 0xc9:
		c.cmp(Immediate)
	case 0xc5:
		c.cmp(ZeroPage)
	case 0xd5:
		c.cmp(ZeroPageX)
	case 0xcd:
		c.cmp(Absolute)
	case 0xdd:
		c.cmp(AbsoluteX)
	case 0xd9:
		c.cmp(AbsoluteY)
	case 0xc1:
		c.cmp(IndirectX)
	case 0xd1:
		c.cmp(IndirectY)
	case 0xe0:
		c.cpx(Immediate)
	case 0xe4:
		c.cpx(ZeroPage)
	case 0xec:
		c.cpx(Absolute)
	case 0xc0:
		c.cpy(Immediate)
	case 0xc4:
		c.cpy(ZeroPage)
	case 0xcc:
		c.cpy(Absolute)
	case 0x48:
		c.pha()
	case 0x08:
		c.php()
	case 0x68:
		c.pla()
	case 0x28:
		c.plp()
	case 0xba:
		c.tsx()
	case 0x9a:
		c.txs()
	}
	return false
}
