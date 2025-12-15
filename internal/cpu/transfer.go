package cpu

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
