package cpu

func (c *Cpu) jmp(mode AddressingMode) {
	var address uint16

	if mode == Absolute {
		address = c.Read16(c.ProgramCounter)
	} else {
		pointer := c.Read16(c.ProgramCounter)
		low := c.Read(pointer)

		// JMP Indirect mode bug -- increments incorrectly the high byte if
		// it ends in 0xFF (like 0x02FF).
		highAddress := (pointer & 0xFF00) | uint16(uint8(pointer)+1)
		high := c.Read(highAddress)
		address = uint16(high<<8) | uint16(low)
	}

	c.ProgramCounter = address
}

func (c *Cpu) jsr() {
	target := c.Read16(c.ProgramCounter)
	ret := c.ProgramCounter + 1
	c.push16(ret)
	c.ProgramCounter = target
}

func (c *Cpu) rts() {
	address := c.pop16()
	c.ProgramCounter = address + 1
}

func (c *Cpu) brk() {
	c.push16(c.ProgramCounter + 1)
	c.push(c.Status | 0b0011_0000)
	c.Status |= 0b0000_0100
	c.ProgramCounter = c.Read16(0xFFFE)
}

func (c *Cpu) rti() {
	c.Status = (c.pop() & 0b1100_1111) | (c.Status & 0b0011_0000)
	c.ProgramCounter = c.pop16()
}
