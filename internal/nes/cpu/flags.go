package cpu

/// 7  bit  0
/// ---- ----
/// NV1B DIZC
/// |||| ||||
/// |||| |||+- Carry
/// |||| ||+-- Zero
/// |||| |+--- Interrupt Disable
/// |||| +---- Decimal
/// |||+------ (No CPU effect; see: the B flag)
/// ||+------- (No CPU effect; always pushed as 1)
/// |+-------- Overflow
/// +--------- Negative

// updateZeroFlag sets the Z status flag to 1 if the result is zero.
func (c *Cpu) updateZeroFlag(result byte) {
	if result == 0 {
		// result is 0, set zero bit to 1
		c.Status |= 0b0000_0010
	} else {
		// result is not 0, set zero bit to 0
		c.Status &= 0b1111_1101
	}
}

// updateNegativeFlag sets the N status flag to 1 if result is negative.
func (c *Cpu) updateNegativeFlag(result byte) {
	if result&0b1000_0000 != 0 {
		// result negative is 1, set negative to 1
		c.Status |= 0b1000_0000
	} else {
		// result negative is 0, set negative to 0
		c.Status &= 0b0111_1111
	}
}

// clc clears (assings 0) the carry flag.
func (c *Cpu) clc() {
	c.Status &= 0b1111_1110
}

// sec sets (assings 1) the carry flag.
func (c *Cpu) sec() {
	c.Status |= 0b0000_0001
}

// cli clears the interrupt disable flag.
func (c *Cpu) cli() {
	c.Status &= 0b1111_1011
}

// sei sets the interrupt disable flag.
func (c *Cpu) sei() {
	c.Status |= 0b0000_0100
}

// cld clears the decimal flag.
func (c *Cpu) cld() {
	c.Status &= 0b1111_0111
}

// sed sets the decimal flag.
func (c *Cpu) sed() {
	c.Status |= 0b0000_1000
}

// clv clears the overflow flag.
func (c *Cpu) clv() {
	c.Status &= 0b1011_1111
}
