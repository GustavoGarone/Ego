package cpu

// cmp compares the accumulator to a memory value, setting flags as appropriate but not
// modifying any registers. The comparison is implemented as a subtraction,
// setting carry if there is no borrow.
// flags:
//   - carry: A >= memory
//   - Zero: A == memory
//   - Negative: result bit 7
func (c *Cpu) cmp(mode AddressingMode) {
	memory := mode.Read(c)
	result := c.Accumulator - memory
	if c.Accumulator >= memory {
		c.Status |= 0b0000_0001
	}

	c.updateNegativeFlag(result)
	c.updateZeroFlag(result)
}

// cpx compares X to a memory value, setting flags as appropriate but not
// modifying any registers. The comparison is implemented as a subtraction,
// setting carry if there is no borrow.
// flags:
//   - carry: X >= memory
//   - Zero: X == memory
//   - Negative: result bit 7
func (c *Cpu) cpx(mode AddressingMode) {
	memory := mode.Read(c)
	result := c.X - memory
	if c.X >= memory {
		c.Status |= 0b0000_0001
	}

	c.updateNegativeFlag(result)
	c.updateZeroFlag(result)
}

// cpy compares Y to a memory value, setting flags as appropriate but not
// modifying any registers. The comparison is implemented as a subtraction,
// setting carry if there is no borrow.
// flags:
//   - carry: Y >= memory
//   - Zero: Y == memory
//   - Negative: result bit 7
func (c *Cpu) cpy(mode AddressingMode) {
	memory := mode.Read(c)
	result := c.Y - memory
	if c.Y >= memory {
		c.Status |= 0b0000_0001
	}

	c.updateNegativeFlag(result)
	c.updateZeroFlag(result)
}
