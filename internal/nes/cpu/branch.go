package cpu

// bcc branches to a nearby location by adding the relative offset to the program
// counter if the carry flag is clear.
// The offset is signed and has a range of [-128, 127] relative to the
// first byte *after* the branch instruction
func (c *Cpu) bcc(mode AddressingMode) {
	if (c.Status & 1) == 0 {
		// Base irrelevant to "Relative" Addressing mode
		address := mode.absoluteAddress(c, 0)
		c.ProgramCounter = address.Value
		// TODO: add aditional handling if it's page crossed
	}
}

// bcs branches to a nearby location by adding the relative offset to the program
// counter if the carry flag is set.
// The offset is signed and has a range of [-128, 127] relative to the
// first byte *after* the branch instruction
func (c *Cpu) bcs(mode AddressingMode) {
	if (c.Status & 1) == 1 {
		address := mode.absoluteAddress(c, 0)
		c.ProgramCounter = address.Value
		// TODO: add aditional handling if it's page crossed
	}
}

// beq branches to a nearby location by adding the relative offset to the program
// counter if the zero flag is set.
// The offset is signed and has a range of [-128, 127] relative to the
// first byte *after* the branch instruction
func (c *Cpu) beq(mode AddressingMode) {
	if (c.Status & 0b0000_0010) != 1 {
		address := mode.absoluteAddress(c, 0)
		c.ProgramCounter = address.Value
		// TODO: add aditional handling if it's page crossed
	}
}

// bne branches to a nearby location by adding the relative offset to the program
// counter if the zero flag is clear.
// The offset is signed and has a range of [-128, 127] relative to the
// first byte *after* the branch instruction
func (c *Cpu) bne(mode AddressingMode) {
	if (c.Status & 0b0000_0010) == 0 {
		address := mode.absoluteAddress(c, 0)
		c.ProgramCounter = address.Value
		// TODO: add aditional handling if it's page crossed
	}
}

// bpl branches to a nearby location by adding the relative offset to the program
// counter if the negative flag is clear.
// The offset is signed and has a range of [-128, 127] relative to the
// first byte *after* the branch instruction
func (c *Cpu) bpl(mode AddressingMode) {
	if (c.Status & 0b1000_0000) == 0 {
		address := mode.absoluteAddress(c, 0)
		c.ProgramCounter = address.Value
		// TODO: add aditional handling if it's page crossed
	}
}

// bmi branches to a nearby location by adding the relative offset to the program
// counter if the negative flag is set.
// The offset is signed and has a range of [-128, 127] relative to the
// first byte *after* the branch instruction
func (c *Cpu) bmi(mode AddressingMode) {
	if (c.Status & 0b1000_0000) != 0 {
		address := mode.absoluteAddress(c, 0)
		c.ProgramCounter = address.Value
		// TODO: add aditional handling if it's page crossed
	}
}

// bvc branches to a nearby location by adding the relative offset to the program
// counter if the overflow flag is clear.
// The offset is signed and has a range of [-128, 127] relative to the
// first byte *after* the branch instruction
func (c *Cpu) bvc(mode AddressingMode) {
	if (c.Status & 0b0100_0000) == 0 {
		address := mode.absoluteAddress(c, 0)
		c.ProgramCounter = address.Value
		// TODO: add aditional handling if it's page crossed
	}
}

// bvs branches to a nearby location by adding the relative offset to the program
// counter if the overflow flag is set.
// The offset is signed and has a range of [-128, 127] relative to the
// first byte *after* the branch instruction
func (c *Cpu) bvs(mode AddressingMode) {
	if (c.Status & 0b0100_0000) != 0 {
		address := mode.absoluteAddress(c, 0)
		c.ProgramCounter = address.Value
		// TODO: add aditional handling if it's page crossed
	}
}
