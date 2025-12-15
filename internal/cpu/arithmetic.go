package cpu

// TODO: Implement after memory is finished
func (c *Cpu) adc() {

}

// TODO: Implement after memory is finished
func (c *Cpu) sbc() {

}

// TODO: Implement after memory is finished
func (c *Cpu) dec() {

}

// inx increments the X register by one
// flags:
//   - Zero: set if resulting X = 0.
//   - Negative: set if bit 7 of resulting X is set.
func (c *Cpu) inx() {
	c.X += 1
	c.updateNegativeFlag(c.X)
	c.updateZeroFlag(c.X)
}

// iny increments the Y register by one
// flags:
//   - Zero: set if resulting Y = 0.
//   - Negative: set if bit 7 of rsulting Y is set.
func (c *Cpu) iny() {
	c.Y += 1
	c.updateNegativeFlag(c.Y)
	c.updateZeroFlag(c.Y)
}

// dex decrements the X register by one
// flags:
//   - Zero: set if resulting X = 0.
//   - Negative: set if bit 7 of resulting X is set.
func (c *Cpu) dex() {
	c.X -= 1
	c.updateNegativeFlag(c.X)
	c.updateZeroFlag(c.X)
}

// dey decrements the Y register by one
// flags:
//   - Zero: set if resulting Y = 0.
//   - Negative: set if bit 7 of resulting Y is set.
func (c *Cpu) dey() {
	c.Y -= 1
	c.updateNegativeFlag(c.Y)
	c.updateZeroFlag(c.Y)
}
