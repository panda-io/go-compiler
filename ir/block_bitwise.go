package ir

// --- [ Bitwise instructions ] ------------------------------------------------

// ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewShl appends a new shl instruction to the basic block based on the given
// operands.
func (block *Block) NewShl(x, y Value) *InstShl {
	inst := NewShl(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewLShr appends a new lshr instruction to the basic block based on the given
// operands.
func (block *Block) NewLShr(x, y Value) *InstLShr {
	inst := NewLShr(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewAShr appends a new ashr instruction to the basic block based on the given
// operands.
func (block *Block) NewAShr(x, y Value) *InstAShr {
	inst := NewAShr(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewAnd appends a new and instruction to the basic block based on the given
// operands.
func (block *Block) NewAnd(x, y Value) *InstAnd {
	inst := NewAnd(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewOr appends a new or instruction to the basic block based on the given
// operands.
func (block *Block) NewOr(x, y Value) *InstOr {
	inst := NewOr(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewXor appends a new xor instruction to the basic block based on the given
// operands.
func (block *Block) NewXor(x, y Value) *InstXor {
	inst := NewXor(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}
