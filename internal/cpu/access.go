package cpu

// ldx loads a byte of memory into the X register.
// flags:
//   - Zero: set if X = 0.
//   - Negative: set if bit 7 of X is set.
func (c *Cpu) ldx(mode AddressingMode) {
	c.ProgramCounter += 1
	c.X = mode.Read(c, 0)
	c.updateNegativeFlag(c.X)
	c.updateZeroFlag(c.X)
}

// ldy loads a byte of memory into the Y register.
// flags:
//   - Zero: set if Y = 0.
//   - Negative: set if bit 7 of Y is set.
func (c *Cpu) ldy(mode AddressingMode) {
	c.ProgramCounter += 1
	c.Y = mode.Read(c, 0)
	c.updateNegativeFlag(c.Y)
	c.updateZeroFlag(c.Y)
}

// lda loads a byte of memory into the accumulator
// flags:
//   - Zero: set if A = 0.
//   - Negative: set if bit 7 of A is set.
func (c *Cpu) lda(mode AddressingMode) {
	c.ProgramCounter += 1
	c.Accumulator = mode.Read(c, 0)
	c.updateNegativeFlag(c.Accumulator)
	c.updateZeroFlag(c.Accumulator)
}

// sta stores the Accumulator value into memory
func (c *Cpu) sta(mode AddressingMode) {
	mode.Write(c, 0, c.Accumulator)
}

// stx Stores the X register into memory
func (c *Cpu) stx(mode AddressingMode) {
	mode.Write(c, 0, c.X)
}

// sty Stores the Y register into memory
func (c *Cpu) sty(mode AddressingMode) {
	mode.Write(c, 0, c.Y)
}
