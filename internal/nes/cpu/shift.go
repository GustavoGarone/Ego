package cpu

// setBitAsCarry sets the carry flag to equal value's n'th bit.
// this considers a byte's rightmost bit as the 0th bit.
func (c *Cpu) setBitAsCarry(value byte, n int8) {
	c.Status &= 0b1111_1110 // Clears carry flag
	c.Status |= (value & (1 << (n)))
}

// asl shifts all of the bits of a memory value or the accumulator one position to the left.
// bit 7 is shifted into the carry flag.
// value = value << 1
// flags:
//   - Carry: set if value's bit 7 is set
//   - Zero: set if result == 0
//   - Negative: set if result's bit 7 is set
func (c *Cpu) asl(mode AddressingMode) {
	value := mode.Read(c)
	c.setBitAsCarry(value, 7)
	result := value << 1
	mode.Write(c, result)

	c.updateNegativeFlag(result)
	c.updateZeroFlag(result)
}

// lsr shifts all the bits of a memory value or the accumulator one position to the right
// 0 is shifted into the bit 7, bit 0 is shifted into the carry flag.
// value = value >> 1
//   - Carry: set to value's bit 0
//   - Zero: set if result == 0
//   - Negative: set to 0.
func (c *Cpu) lsr(mode AddressingMode) {
	value := mode.Read(c)
	c.setBitAsCarry(value, 0)
	result := value >> 1
	mode.Write(c, result)

	c.updateNegativeFlag(result)
	c.updateZeroFlag(result)
}

// rol shifts a memory value or the accumulator to the left, treating the carry
// flag as though it is both above bit 7 and below bit 0. That is, carry -> 0 and
// 7 -> carry.
// flags:
//   - Carry: set if value's bit 7 is set
//   - Zero: set if result == 0
//   - Negative: set if result's bit 7 is set
func (c *Cpu) rol(mode AddressingMode) {
	value := mode.Read(c)
	result := value << 1
	result |= c.Status & 1
	c.setBitAsCarry(value, 7)

	c.updateNegativeFlag(result)
	c.updateZeroFlag(result)
}

// ror shifts a memory value or the accumulator to the right, trating the carry
// flag as though it is both above bit 7 and below bit 0. That is, carry -> 7 and
// 0 -> carry.
//   - Carry: set if value's bit 7 is set
//   - Zero: set if result == 0
//   - Negative: set if result's bit 7 is set
func (c *Cpu) ror(mode AddressingMode) {
	value := mode.Read(c)
	result := (value >> 1)
	result |= c.Status << 7
	c.setBitAsCarry(value, 0)

	c.updateNegativeFlag(result)
	c.updateZeroFlag(result)
}
