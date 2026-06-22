package cpu

// pha pushes the accumulator to the current stack position and decrements the
// stack pointer
func (c *Cpu) pha() {
	c.push(c.Accumulator) // SP decrease implemented in push
}

// pla increments the stack pointer and then loads the value at that stack
// position into the accumulator
func (c *Cpu) pla() {
	c.Accumulator = c.pop()
}

// php stores a byte to the stack containing the 6 status flags and B flag and
// then decrements the stack pointer. The B flag and extra bit are both pushed as
// 1. The bit order is NV1BDIZC (high to low)
// flag:
//   - B: Break - Pushed as 1. This flag only exists as pushed to the stack.
func (c *Cpu) php() {
	flags := c.Status | 0b0011_0000
	c.push(flags)
}

// plp increments the stack pointer and then loads the value at that stack
// position into the status flags. The B and extra bit are ignored.
func (c *Cpu) plp() {
	c.Status = c.pop()
}

// txs transfers x to the stack pointer
func (c *Cpu) txs() {
	c.Stack = c.X
}

// tsx trasnfers the stack pointer to x
func (c *Cpu) tsx() {
	c.X = c.Stack
}
