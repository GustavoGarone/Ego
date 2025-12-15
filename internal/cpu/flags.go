package cpu

// 7  bit  0
// ---- ----
// NV1B DIZC
// |||| ||||
// |||| |||+- Carry
// |||| ||+-- Zero
// |||| |+--- Interrupt Disable
// |||| +---- Decimal
// |||+------ (No CPU effect; see: the B flag)
// ||+------- (No CPU effect; always pushed as 1)
// |+-------- Overflow
// +--------- Negative
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

func (c *Cpu) clc() {
	c.Status &= 0b1111_1110
}

func (c *Cpu) sec() {
	c.Status |= 0b0000_0001
}

func (c *Cpu) cli() {
	c.Status &= 0b1111_1011
}

func (c *Cpu) sei() {
	c.Status |= 0b0000_0100
}

func (c *Cpu) cld() {
	c.Status &= 0b1111_0111
}

func (c *Cpu) sed() {
	c.Status |= 0b0000_1000
}

func (c *Cpu) clv() {
	c.Status &= 0b1011_1111
}
