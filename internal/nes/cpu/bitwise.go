package cpu

// and executes an AND bitwise operation between a value and the accumulator.
// A = A & memory
// flags:
//   - Zero: set if result is 0
//   - Negative: set if the result bit 7 is set
func (c *Cpu) and(mode AddressingMode) {
	result := c.Accumulator & mode.Read(c)
	c.Accumulator = result

	c.updateNegativeFlag(result)
	c.updateZeroFlag(result)
}

// ora inclusive ORs a memory value and the accumulator.
// A = A | memory
// flags:
//   - Zero: set if result is 0
//   - Negative: set if the result bit 7 is set
func (c *Cpu) ora(mode AddressingMode) {
	result := c.Accumulator | mode.Read(c)
	c.Accumulator = result

	c.updateNegativeFlag(result)
	c.updateZeroFlag(result)
}

// eor exclusive ORs a memory value and the accumulator.
// A = A ^ memory
// flags:
//   - Zero: set if result is 0
//   - Negative: set if the result bit 7 is set
func (c *Cpu) eor(mode AddressingMode) {
	result := c.Accumulator ^ mode.Read(c)
	c.Accumulator = result

	c.updateNegativeFlag(result)
	c.updateZeroFlag(result)
}

// BIT modifies flags according to a given memory value. The zero flag
// is set depending on the result of A & memory. Bit's 7 and 6 of the memory
// value are set directly into the negative and overflow flags.
// flags:
//   - Zero: set if A & value == 0
//   - Overflow: value's 6th bit
//   - Negative: value's 7th bit
func (c *Cpu) bit(mode AddressingMode) {
	value := mode.Read(c)
	result := c.Accumulator & value

	c.updateZeroFlag(result)
	c.updateNegativeFlag(result)
	// Updating Overflow flag
	mask := value & 0b0100_0000 // Creates a mask based on the memory value
	c.Status &= 0b1011_1111     // Clears the overflow flag
	c.Status |= mask            // If mask's 6th bit is set, so will the overflow.
}
