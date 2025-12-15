package cpu

// lda loads a byte of memory into the X register.
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

// lda loads a byte of memory into the Y register.
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

// tya copies the current contents of the Y register into the Accumulator.
// flags:
//   - Zero: set if accumulator = 0.
//   - Negative: set if bit 7 of accumulator is set.
func (c *Cpu) tya() {
	c.Accumulator = c.Y
	c.updateNegativeFlag(c.Accumulator)
	c.updateZeroFlag(c.Accumulator)
}

// tya copies the current contents of the X register into the Accumulator.
// flags:
//   - Zero: set if accumulator = 0.
//   - Negative: set if bit 7 of accumulator is set.
func (c *Cpu) txa() {
	c.Accumulator = c.Y
	c.updateNegativeFlag(c.Accumulator)
	c.updateZeroFlag(c.Accumulator)
}

// tay copies the current contents of the accumulator into the Y register.
// flags:
//   - Zero: set if Y = 0.
//   - Negative: set if bit 7 of Y is set.
func (c *Cpu) tay() {
	c.Y = c.Accumulator
	c.updateNegativeFlag(c.Y)
	c.updateZeroFlag(c.Y)
}

// tax copies the current contents of the accumulator into the X register.
// flags:
//   - Zero: set if X = 0.
//   - Negative: set if bit 7 of X is set.
func (c *Cpu) tax() {
	c.X = c.Accumulator
	c.updateNegativeFlag(c.X)
	c.updateZeroFlag(c.X)
}

// lda loads a byte of memory into the accumulator
// flags:
//   - Zero: set if A = 0.
//   - Negative: set if bit 7 of A is set.
func (c *Cpu) lda() {
	c.ProgramCounter += 1
	arg := c.Fetch()
	c.Accumulator = arg
	c.updateNegativeFlag(c.Accumulator)
	c.updateZeroFlag(c.Accumulator)
}

func (c *Cpu) updateZeroFlag(result uint8) {
	if result == 0 {
		// result is 0, set zero bit to 1
		c.Status |= 0b0000_0010
	} else {
		// result is not 0, set zero bit to 0
		c.Status &= 0b1111_1101
	}
}

func (c *Cpu) updateNegativeFlag(result uint8) {
	if result&0b1000_0000 != 0 {
		// result negative is 1, set negative to 1
		c.Status |= 0b1000_0000
	} else {
		// result negative is 0, set negative to 0
		c.Status &= 0b0111_1111
	}
}
