package cpu

// asl shifts all of the bits of a memory value or the accumulator one position to the left
// value = value << 1
// flags:
//   - Carry: set if value's bit 7 is set
//   - Zero: set if result == 0
//   - Negative: set if result's bit 7 is set
func (c *Cpu) asl(mode AddressingMode) {
	value := mode.Read(c, 0)
	result := value << 1
	mode.Write(c, 0, result)

	c.updateNegativeFlag(result)
	c.updateZeroFlag(result)
}
