package cpu

// ldx loads a byte of memory into the X register.
// flags:
//   - Zero: set if X = 0.
//   - Negative: set if bit 7 of X is set.
func (c *Cpu) ldx() {
	c.ProgramCounter += 1
	arg := c.Fetch()
	c.X = arg
	c.updateNegativeFlag(c.X)
	c.updateZeroFlag(c.X)
}

// ldy loads a byte of memory into the Y register.
// flags:
//   - Zero: set if Y = 0.
//   - Negative: set if bit 7 of Y is set.
func (c *Cpu) ldy() {
	c.ProgramCounter += 1
	arg := c.Fetch()
	c.Y = arg
	c.updateNegativeFlag(c.Y)
	c.updateZeroFlag(c.Y)
}

// lda loads a byte of memory into the accumulator
// flags:
//   - Zero: set if A = 0.
//   - Negative: set if bit 7 of A is set.
func (c *Cpu) lda(mode AddressingMode) {
	c.ProgramCounter += 1
	arg := mode.Read(c, 0)
	c.Accumulator = arg
	c.updateNegativeFlag(c.Accumulator)
	c.updateZeroFlag(c.Accumulator)
}
