package ir

// --- [ Binary instructions ] -------------------------------------------------

// ~~~ [ add ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewAdd appends a new add instruction to the basic block based on the given
// operands.
func (block *Block) NewAdd(x, y Value) *InstAdd {
	inst := NewAdd(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fadd ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFAdd appends a new fadd instruction to the basic block based on the given
// operands.
func (block *Block) NewFAdd(x, y Value) *InstFAdd {
	inst := NewFAdd(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewSub appends a new sub instruction to the basic block based on the given
// operands.
func (block *Block) NewSub(x, y Value) *InstSub {
	inst := NewSub(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fsub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFSub appends a new fsub instruction to the basic block based on the given
// operands.
func (block *Block) NewFSub(x, y Value) *InstFSub {
	inst := NewFSub(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewMul appends a new mul instruction to the basic block based on the given
// operands.
func (block *Block) NewMul(x, y Value) *InstMul {
	inst := NewMul(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fmul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFMul appends a new fmul instruction to the basic block based on the given
// operands.
func (block *Block) NewFMul(x, y Value) *InstFMul {
	inst := NewFMul(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ udiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewUDiv appends a new udiv instruction to the basic block based on the given
// operands.
func (block *Block) NewUDiv(x, y Value) *InstUDiv {
	inst := NewUDiv(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ sdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewSDiv appends a new sdiv instruction to the basic block based on the given
// operands.
func (block *Block) NewSDiv(x, y Value) *InstSDiv {
	inst := NewSDiv(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFDiv appends a new fdiv instruction to the basic block based on the given
// operands.
func (block *Block) NewFDiv(x, y Value) *InstFDiv {
	inst := NewFDiv(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ urem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewURem appends a new urem instruction to the basic block based on the given
// operands.
func (block *Block) NewURem(x, y Value) *InstURem {
	inst := NewURem(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ srem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewSRem appends a new srem instruction to the basic block based on the given
// operands.
func (block *Block) NewSRem(x, y Value) *InstSRem {
	inst := NewSRem(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ frem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFRem appends a new frem instruction to the basic block based on the given
// operands.
func (block *Block) NewFRem(x, y Value) *InstFRem {
	inst := NewFRem(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}
