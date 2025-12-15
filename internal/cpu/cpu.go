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

func (c *Cpu) Fetch() uint8 {
	return c.Program[c.ProgramCounter]
}

func (c *Cpu) Execute(opcode uint8) {
	c.ProgramCounter += 1
	switch opcode {
	case 0x00:
		return
	case 0xa9:
		c.ProgramCounter += 1 // a9 comes with an argument
		arg := c.Fetch()
		c.lda(arg)
	}
}

func (c *Cpu) lda(value uint8) {
	c.Accumulator = value
	c.updateZeroAndNegativeFlags(c.Accumulator)
}

func (c *Cpu) updateZeroAndNegativeFlags(result uint8) {
	if result == 0 {
		// result is 0, set zero bit to 1
		c.Status |= 0b0000_0010
	} else {
		// result is not 0, set zero bit to 0
		c.Status &= 0b1111_1101
	}

	if result&0b1000_0000 != 0 {
		// result negative is 1, set negative to 1
		c.Status |= 0b1000_0000
	} else {
		// result negative is 0, set negative to 0
		c.Status &= 0b0111_1111
	}
}
