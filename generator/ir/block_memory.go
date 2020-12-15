package ir

// --- [ Memory instructions ] -------------------------------------------------

// ~~~ [ alloca ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewAlloca appends a new alloca instruction to the basic block based on the
// given element type.
func (block *Block) NewAlloca(elemType Type) *InstAlloca {
	inst := NewAlloca(elemType)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ load ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewLoad appends a new load instruction to the basic block based on the given
// element type and source address.
func (block *Block) NewLoad(elemType Type, src Value) *InstLoad {
	inst := NewLoad(elemType, src)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ store ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewStore appends a new store instruction to the basic block based on the
// given source value and destination address.
func (block *Block) NewStore(src, dst Value) *InstStore {
	inst := NewStore(src, dst)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fence ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFence appends a new fence instruction to the basic block based on the
// given atomic ordering.
func (block *Block) NewFence(ordering AtomicOrdering) *InstFence {
	inst := NewFence(ordering)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ cmpxchg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewCmpXchg appends a new cmpxchg instruction to the basic block based on the
// given address, value to compare against, new value to store, and atomic
// orderings for success and failure.
func (block *Block) NewCmpXchg(ptr, cmp, new Value, successOrdering, failureOrdering AtomicOrdering) *InstCmpXchg {
	inst := NewCmpXchg(ptr, cmp, new, successOrdering, failureOrdering)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ atomicrmw ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewAtomicRMW appends a new atomicrmw instruction to the basic block based on
// the given atomic operation, destination address, operand and atomic ordering.
func (block *Block) NewAtomicRMW(op AtomicOp, dst, x Value, ordering AtomicOrdering) *InstAtomicRMW {
	inst := NewAtomicRMW(op, dst, x, ordering)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ getelementptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewGetElementPtr appends a new getelementptr instruction to the basic block
// based on the given element type, source address and element indices.
func (block *Block) NewGetElementPtr(elemType Type, src Value, indices ...Value) *InstGetElementPtr {
	inst := NewGetElementPtr(elemType, src, indices...)
	block.Insts = append(block.Insts, inst)
	return inst
}
