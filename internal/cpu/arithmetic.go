package cpu

// adc adds the carry flag and a memory value to the accumulator. The carry flag
// is then set to the carry value coming out of bit 7.
// flags:
//   - Carry: set if result > $FF (unsigned overflow occurred)
//   - Zero: set if result == 0
//   - Overflow: set if result sign is different from both the Accumulator and memory's.
//     (result ^ Accumulator) & (result ^ memory) & $80
//     Therefore, the overflow flag is set if result < -128 or result > 127
//   - Negative: set if bit 7 of result is set.
func (c *Cpu) adc(mode AddressingMode) {
	memory := mode.Read(c, 0)
	carryFlag := (c.Status & 1)
	result := c.Accumulator + mode.Read(c, 0) + carryFlag
	if result > 0xff {
		c.Status |= 0b0000_0001
	} else {
		c.Status &= 0b1111_1110
	}
	// see http://www.6502.org/tutorials/vflag.html
	if ((result ^ c.Accumulator) & (result ^ memory) & 0x80) != 0 {
		c.Status |= 0b0100_0000
	} else {
		c.Status &= 0b1011_1111
	}
	c.updateNegativeFlag(result)
	c.updateZeroFlag(result)
}

// sbc subtracts a memory value and the NOT of the carry flag from the accumulator.
// The carry flag is then set to the NOT value  of bit 7.
// A = A - memory - !Carry
// flags:
//   - Carry: set if result >= $00 (unsigned underflow, !(result < $00))
//   - Zero: set if result == 0
//   - Overflow: set if result sign is different from both the Accumulator and the same as memory's.
//     (result ^ Accumulator) & (result ^ !memory) & $80
//     Therefore, the overflow flag is set if result < -128 or result > 127
//   - Negative: set if bit 7 of result is set.
func (c *Cpu) sbc(mode AddressingMode) {
	memory := mode.Read(c, 0)
	carryFlag := (c.Status & 1)
	result := c.Accumulator - mode.Read(c, 0) - ^carryFlag
	if ^result < 0xff {
		c.Status |= 0b0000_0001
	} else {
		c.Status &= 0b1111_1110
	}
	if ((result ^ c.Accumulator) & (result ^ ^memory) & 0x80) != 0 {
		c.Status |= 0b0100_0000
	} else {
		c.Status &= 0b1011_1111
	}
	c.updateNegativeFlag(result)
	c.updateZeroFlag(result)
}

// dec substracts 1 from a memory location
// This is a read-modify-write instruction, meaning that it first writes the
// original value back to memory before the modified value. (?)
// This extra write can matter if targeting a hardware register.
// As we are not emulating the hardware on that level, we don't  have to worry about the extra write.
// flags:
//   - Zero: set if resulting X = 0.
//   - Negative: set if bit 7 of resulting X is set.
func (c *Cpu) dec(mode AddressingMode) {
	value := mode.Read(c, 0)
	value -= 1
	mode.Write(c, 0, value)
	c.updateNegativeFlag(value)
	c.updateZeroFlag(value)
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
