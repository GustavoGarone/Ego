package cpu

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

	// The ROM the CPU should read from
	Program []uint8
}

func NewCpu(program []uint8) *Cpu {
	return &Cpu{
		Accumulator:    0,
		X:              0,
		Y:              0,
		Status:         0,
		Stack:          0,
		ProgramCounter: 0,
		Program:        program,
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
	return c.Program[c.ProgramCounter]
}

// Execute will handle a program instruction. Returns true if execution is done.
func (c *Cpu) Execute(opcode uint8) bool {
	switch opcode {
	case 0x00:
		return true
	case 0xea:
		return false // NOP
	case 0xa9:
		c.lda()
	case 0xa2:
		c.ldx()
	case 0xa0:
		c.ldy()
	case 0xaa:
		c.tax()
	case 0xa8:
		c.tay()
	case 0x8a:
		c.txa()
	case 0x98:
		c.tya()
	case 0x69:
		c.adc()
	case 0xe9:
		c.sbc()
	case 0xc6:
		c.dec()
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
